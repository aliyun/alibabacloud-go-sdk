// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppPluginFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPluginId(v string) *ListAppPluginFilesRequest
	GetPluginId() *string
	SetPluginVersion(v string) *ListAppPluginFilesRequest
	GetPluginVersion() *string
}

type ListAppPluginFilesRequest struct {
	// The ID of the gateway plugin.
	//
	// example:
	//
	// aliplayer-react
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// The version number. If you do not specify this parameter, the currently active version is used.
	//
	// example:
	//
	// 1.1.3
	PluginVersion *string `json:"PluginVersion,omitempty" xml:"PluginVersion,omitempty"`
}

func (s ListAppPluginFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppPluginFilesRequest) GoString() string {
	return s.String()
}

func (s *ListAppPluginFilesRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *ListAppPluginFilesRequest) GetPluginVersion() *string {
	return s.PluginVersion
}

func (s *ListAppPluginFilesRequest) SetPluginId(v string) *ListAppPluginFilesRequest {
	s.PluginId = &v
	return s
}

func (s *ListAppPluginFilesRequest) SetPluginVersion(v string) *ListAppPluginFilesRequest {
	s.PluginVersion = &v
	return s
}

func (s *ListAppPluginFilesRequest) Validate() error {
	return dara.Validate(s)
}
