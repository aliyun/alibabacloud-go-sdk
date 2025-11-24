// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSwimLaneGroupListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServiceMeshId(v string) *GetSwimLaneGroupListRequest
	GetServiceMeshId() *string
}

type GetSwimLaneGroupListRequest struct {
	// The ASM instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s GetSwimLaneGroupListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSwimLaneGroupListRequest) GoString() string {
	return s.String()
}

func (s *GetSwimLaneGroupListRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *GetSwimLaneGroupListRequest) SetServiceMeshId(v string) *GetSwimLaneGroupListRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *GetSwimLaneGroupListRequest) Validate() error {
	return dara.Validate(s)
}
