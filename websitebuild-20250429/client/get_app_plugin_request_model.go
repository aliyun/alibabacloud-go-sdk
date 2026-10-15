// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppPluginRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPluginId(v string) *GetAppPluginRequest
	GetPluginId() *string
}

type GetAppPluginRequest struct {
	// The gateway plug-in ID.
	//
	// example:
	//
	// 1bae9ceaceea432d91c7069fab0dfc02
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
}

func (s GetAppPluginRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppPluginRequest) GoString() string {
	return s.String()
}

func (s *GetAppPluginRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *GetAppPluginRequest) SetPluginId(v string) *GetAppPluginRequest {
	s.PluginId = &v
	return s
}

func (s *GetAppPluginRequest) Validate() error {
	return dara.Validate(s)
}
