// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleClusterNodePoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *ScaleClusterNodePoolRequestBody) *ScaleClusterNodePoolRequest
	GetBody() *ScaleClusterNodePoolRequestBody
	SetClusterId(v string) *ScaleClusterNodePoolRequest
	GetClusterId() *string
	SetNodepoolId(v string) *ScaleClusterNodePoolRequest
	GetNodepoolId() *string
}

type ScaleClusterNodePoolRequest struct {
	// The HTTP request body.
	Body *ScaleClusterNodePoolRequestBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// eck-xxxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The node pool ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// np2201da356f5245cf8faa522a8a4c8224
	NodepoolId *string `json:"NodepoolId,omitempty" xml:"NodepoolId,omitempty"`
}

func (s ScaleClusterNodePoolRequest) String() string {
	return dara.Prettify(s)
}

func (s ScaleClusterNodePoolRequest) GoString() string {
	return s.String()
}

func (s *ScaleClusterNodePoolRequest) GetBody() *ScaleClusterNodePoolRequestBody {
	return s.Body
}

func (s *ScaleClusterNodePoolRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ScaleClusterNodePoolRequest) GetNodepoolId() *string {
	return s.NodepoolId
}

func (s *ScaleClusterNodePoolRequest) SetBody(v *ScaleClusterNodePoolRequestBody) *ScaleClusterNodePoolRequest {
	s.Body = v
	return s
}

func (s *ScaleClusterNodePoolRequest) SetClusterId(v string) *ScaleClusterNodePoolRequest {
	s.ClusterId = &v
	return s
}

func (s *ScaleClusterNodePoolRequest) SetNodepoolId(v string) *ScaleClusterNodePoolRequest {
	s.NodepoolId = &v
	return s
}

func (s *ScaleClusterNodePoolRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ScaleClusterNodePoolRequestBody struct {
	// The expected number of instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s ScaleClusterNodePoolRequestBody) String() string {
	return dara.Prettify(s)
}

func (s ScaleClusterNodePoolRequestBody) GoString() string {
	return s.String()
}

func (s *ScaleClusterNodePoolRequestBody) GetCount() *int32 {
	return s.Count
}

func (s *ScaleClusterNodePoolRequestBody) SetCount(v int32) *ScaleClusterNodePoolRequestBody {
	s.Count = &v
	return s
}

func (s *ScaleClusterNodePoolRequestBody) Validate() error {
	return dara.Validate(s)
}
