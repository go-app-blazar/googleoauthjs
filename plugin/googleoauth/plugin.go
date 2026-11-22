package googleoauth

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/go-app-blazar/blazar/blazarplugin"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

// Config is the configuration for the Google OAuth plugin.
type Config struct {
}

// plugin is the implementation of the Google OAuth plugin.
type plugin struct {
	config Config
}

var _ blazarplugin.Plugin = (*plugin)(nil)

// NewPlugin creates a new Google OAuth plugin.
func NewPlugin(config Config) blazarplugin.Plugin {
	return &plugin{
		config: config,
	}
}

// Register registers the plugin against the go-app handler and the HTTP mux.
func (p *plugin) Register(handler *app.Handler, mux *http.ServeMux) {
	slog.InfoContext(context.TODO(), "Registering Google OAuth plugin", "config", p.config)

	handler.Scripts = append(handler.Scripts, "https://accounts.google.com/gsi/client")

	handler.CacheableResources = append(handler.CacheableResources, "https://accounts.google.com/gsi/client")
}
