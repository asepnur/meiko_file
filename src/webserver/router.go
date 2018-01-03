package webserver

import (
	"net/http"

	"github.com/asepnur/meiko_course/src/util/auth"
	"github.com/asepnur/meiko_course/src/webserver/handler/file"

	"github.com/julienschmidt/httprouter"
)

// Load returns all routing of this server
func loadRouter(r *httprouter.Router) {
	// ========================== File Handler ==========================
	r.GET("/api/v1/filerouter", auth.OptionalAuthorize(file.RouterFileHandler))
	r.GET("/api/v1/file/:payload/:filename", file.GetFileHandler)
	r.GET("/api/v1/admin/file/types/available", auth.MustAuthorize(file.AvailableTypes))
	r.GET("/api/v1/image/:payload", auth.MustAuthorize(file.GetProfileHandler))
	r.POST("/api/v1/image/profile", auth.MustAuthorize(file.UploadProfileImageHandler))
	r.POST("/api/admin/v1/image/information", auth.MustAuthorize(file.UploadInformationImageHandler))
	// r.POST("/api/v1/file/assignment", auth.MustAuthorize(file.UploadAssignmentHandler))
	r.POST("/api/v1/file", auth.MustAuthorize(file.UploadFileHandler))
	r.GET("/static/*filepath", file.StaticHandler)
	// ======================== End File Handler ========================

	// Catch
	r.NotFound = http.HandlerFunc(file.IndexHandler)
	// r.MethodNotAllowed = http.RedirectHandler("/", http.StatusPermanentRedirect)
}
