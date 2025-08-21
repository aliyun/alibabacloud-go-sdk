// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvalidateThirdPartyAppLoginStateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *InvalidateThirdPartyAppLoginStateShrinkRequest
	GetDeviceInfoShrink() *string
	SetThirdPartyAppId(v string) *InvalidateThirdPartyAppLoginStateShrinkRequest
	GetThirdPartyAppId() *string
}

type InvalidateThirdPartyAppLoginStateShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// com.*.*.*
	ThirdPartyAppId *string `json:"ThirdPartyAppId,omitempty" xml:"ThirdPartyAppId,omitempty"`
}

func (s InvalidateThirdPartyAppLoginStateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s InvalidateThirdPartyAppLoginStateShrinkRequest) GoString() string {
	return s.String()
}

func (s *InvalidateThirdPartyAppLoginStateShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *InvalidateThirdPartyAppLoginStateShrinkRequest) GetThirdPartyAppId() *string {
	return s.ThirdPartyAppId
}

func (s *InvalidateThirdPartyAppLoginStateShrinkRequest) SetDeviceInfoShrink(v string) *InvalidateThirdPartyAppLoginStateShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *InvalidateThirdPartyAppLoginStateShrinkRequest) SetThirdPartyAppId(v string) *InvalidateThirdPartyAppLoginStateShrinkRequest {
	s.ThirdPartyAppId = &v
	return s
}

func (s *InvalidateThirdPartyAppLoginStateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
