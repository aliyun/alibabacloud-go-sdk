// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPluginForUuidRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTypes(v []*string) *ListPluginForUuidRequest
	GetTypes() []*string
	SetUuid(v string) *ListPluginForUuidRequest
	GetUuid() *string
}

type ListPluginForUuidRequest struct {
	// The plug-in types.
	Types []*string `json:"Types,omitempty" xml:"Types,omitempty" type:"Repeated"`
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

func (s ListPluginForUuidRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPluginForUuidRequest) GoString() string {
	return s.String()
}

func (s *ListPluginForUuidRequest) GetTypes() []*string {
	return s.Types
}

func (s *ListPluginForUuidRequest) GetUuid() *string {
	return s.Uuid
}

func (s *ListPluginForUuidRequest) SetTypes(v []*string) *ListPluginForUuidRequest {
	s.Types = v
	return s
}

func (s *ListPluginForUuidRequest) SetUuid(v string) *ListPluginForUuidRequest {
	s.Uuid = &v
	return s
}

func (s *ListPluginForUuidRequest) Validate() error {
	return dara.Validate(s)
}
