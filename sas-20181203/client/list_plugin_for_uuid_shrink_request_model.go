// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPluginForUuidShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTypesShrink(v string) *ListPluginForUuidShrinkRequest
	GetTypesShrink() *string
	SetUuid(v string) *ListPluginForUuidShrinkRequest
	GetUuid() *string
}

type ListPluginForUuidShrinkRequest struct {
	// The plug-in types.
	TypesShrink *string `json:"Types,omitempty" xml:"Types,omitempty"`
	// The UUID of the server.
	//
	// >  You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to query the UUIDs of servers.
	//
	// This parameter is required.
	//
	// example:
	//
	// bdb7071f-129d-4ceb-af80-4cf70c4571c6
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListPluginForUuidShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPluginForUuidShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListPluginForUuidShrinkRequest) GetTypesShrink() *string {
	return s.TypesShrink
}

func (s *ListPluginForUuidShrinkRequest) GetUuid() *string {
	return s.Uuid
}

func (s *ListPluginForUuidShrinkRequest) SetTypesShrink(v string) *ListPluginForUuidShrinkRequest {
	s.TypesShrink = &v
	return s
}

func (s *ListPluginForUuidShrinkRequest) SetUuid(v string) *ListPluginForUuidShrinkRequest {
	s.Uuid = &v
	return s
}

func (s *ListPluginForUuidShrinkRequest) Validate() error {
	return dara.Validate(s)
}
