// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallAppPluginRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPluginId(v string) *UninstallAppPluginRequest
	GetPluginId() *string
}

type UninstallAppPluginRequest struct {
	// The gateway plugin ID.
	//
	// example:
	//
	// 1bae9ceaceea432d91c7069fab0dfc02
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
}

func (s UninstallAppPluginRequest) String() string {
	return dara.Prettify(s)
}

func (s UninstallAppPluginRequest) GoString() string {
	return s.String()
}

func (s *UninstallAppPluginRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *UninstallAppPluginRequest) SetPluginId(v string) *UninstallAppPluginRequest {
	s.PluginId = &v
	return s
}

func (s *UninstallAppPluginRequest) Validate() error {
	return dara.Validate(s)
}
