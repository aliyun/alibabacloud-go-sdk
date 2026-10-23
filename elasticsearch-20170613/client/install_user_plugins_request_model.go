// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallUserPluginsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *InstallUserPluginsRequest
	GetBody() *string
	SetForce(v bool) *InstallUserPluginsRequest
	GetForce() *bool
}

type InstallUserPluginsRequest struct {
	// The list of custom plugins.
	//
	// example:
	//
	// [
	//
	//   {
	//
	//     "elasticsearchVersion": "8.17.0",
	//
	//     "pluginType": "CUSTOM_PLUGIN",
	//
	//     "name": "analysis-stempel",
	//
	//     "description": "The Stempel Analysis plugin is ***",
	//
	//     "version": "8.17.0",
	//
	//     "fileVersion": "CAEQnwIYgYDAzYfNj4oaI***"
	//
	//   }
	//
	// ]
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// Specifies whether to forcibly install custom user plugins. Default value: false. Instances with the legacy architecture do not support forced installation.
	Force *bool `json:"force,omitempty" xml:"force,omitempty"`
}

func (s InstallUserPluginsRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallUserPluginsRequest) GoString() string {
	return s.String()
}

func (s *InstallUserPluginsRequest) GetBody() *string {
	return s.Body
}

func (s *InstallUserPluginsRequest) GetForce() *bool {
	return s.Force
}

func (s *InstallUserPluginsRequest) SetBody(v string) *InstallUserPluginsRequest {
	s.Body = &v
	return s
}

func (s *InstallUserPluginsRequest) SetForce(v bool) *InstallUserPluginsRequest {
	s.Force = &v
	return s
}

func (s *InstallUserPluginsRequest) Validate() error {
	return dara.Validate(s)
}
