// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceLinkedRoleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *CreateServiceLinkedRoleShrinkRequest
	GetRegionId() *string
	SetSdkRequestShrink(v string) *CreateServiceLinkedRoleShrinkRequest
	GetSdkRequestShrink() *string
}

type CreateServiceLinkedRoleShrinkRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SdkRequestShrink *string `json:"SdkRequest,omitempty" xml:"SdkRequest,omitempty"`
}

func (s CreateServiceLinkedRoleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceLinkedRoleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateServiceLinkedRoleShrinkRequest) GetSdkRequestShrink() *string {
	return s.SdkRequestShrink
}

func (s *CreateServiceLinkedRoleShrinkRequest) SetRegionId(v string) *CreateServiceLinkedRoleShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServiceLinkedRoleShrinkRequest) SetSdkRequestShrink(v string) *CreateServiceLinkedRoleShrinkRequest {
	s.SdkRequestShrink = &v
	return s
}

func (s *CreateServiceLinkedRoleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
