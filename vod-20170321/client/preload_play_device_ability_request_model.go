// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreloadPlayDeviceAbilityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBrand(v string) *PreloadPlayDeviceAbilityRequest
	GetBrand() *string
	SetResourceRealOwnerId(v int64) *PreloadPlayDeviceAbilityRequest
	GetResourceRealOwnerId() *int64
}

type PreloadPlayDeviceAbilityRequest struct {
	Brand               *string `json:"Brand,omitempty" xml:"Brand,omitempty"`
	ResourceRealOwnerId *int64  `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
}

func (s PreloadPlayDeviceAbilityRequest) String() string {
	return dara.Prettify(s)
}

func (s PreloadPlayDeviceAbilityRequest) GoString() string {
	return s.String()
}

func (s *PreloadPlayDeviceAbilityRequest) GetBrand() *string {
	return s.Brand
}

func (s *PreloadPlayDeviceAbilityRequest) GetResourceRealOwnerId() *int64 {
	return s.ResourceRealOwnerId
}

func (s *PreloadPlayDeviceAbilityRequest) SetBrand(v string) *PreloadPlayDeviceAbilityRequest {
	s.Brand = &v
	return s
}

func (s *PreloadPlayDeviceAbilityRequest) SetResourceRealOwnerId(v int64) *PreloadPlayDeviceAbilityRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *PreloadPlayDeviceAbilityRequest) Validate() error {
	return dara.Validate(s)
}
