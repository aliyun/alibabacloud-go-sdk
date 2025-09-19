// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserVpcRequest interface {
	dara.Model
	String() string
	GoString() string
	SetK8sRegionId(v string) *ListUserVpcRequest
	GetK8sRegionId() *string
}

type ListUserVpcRequest struct {
	// Region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	K8sRegionId *string `json:"K8sRegionId,omitempty" xml:"K8sRegionId,omitempty"`
}

func (s ListUserVpcRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserVpcRequest) GoString() string {
	return s.String()
}

func (s *ListUserVpcRequest) GetK8sRegionId() *string {
	return s.K8sRegionId
}

func (s *ListUserVpcRequest) SetK8sRegionId(v string) *ListUserVpcRequest {
	s.K8sRegionId = &v
	return s
}

func (s *ListUserVpcRequest) Validate() error {
	return dara.Validate(s)
}
