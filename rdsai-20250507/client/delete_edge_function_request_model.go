// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEdgeFunctionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteEdgeFunctionRequest
	GetClientToken() *string
	SetEdgeFunctionName(v string) *DeleteEdgeFunctionRequest
	GetEdgeFunctionName() *string
	SetInstanceName(v string) *DeleteEdgeFunctionRequest
	GetInstanceName() *string
	SetRegionId(v string) *DeleteEdgeFunctionRequest
	GetRegionId() *string
}

type DeleteEdgeFunctionRequest struct {
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// ef-****
	EdgeFunctionName *string `json:"EdgeFunctionName,omitempty" xml:"EdgeFunctionName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteEdgeFunctionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEdgeFunctionRequest) GoString() string {
	return s.String()
}

func (s *DeleteEdgeFunctionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteEdgeFunctionRequest) GetEdgeFunctionName() *string {
	return s.EdgeFunctionName
}

func (s *DeleteEdgeFunctionRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DeleteEdgeFunctionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteEdgeFunctionRequest) SetClientToken(v string) *DeleteEdgeFunctionRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteEdgeFunctionRequest) SetEdgeFunctionName(v string) *DeleteEdgeFunctionRequest {
	s.EdgeFunctionName = &v
	return s
}

func (s *DeleteEdgeFunctionRequest) SetInstanceName(v string) *DeleteEdgeFunctionRequest {
	s.InstanceName = &v
	return s
}

func (s *DeleteEdgeFunctionRequest) SetRegionId(v string) *DeleteEdgeFunctionRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteEdgeFunctionRequest) Validate() error {
	return dara.Validate(s)
}
