// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApSsidConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApMac(v string) *DeleteApSsidConfigRequest
	GetApMac() *string
	SetAppCode(v string) *DeleteApSsidConfigRequest
	GetAppCode() *string
	SetAppName(v string) *DeleteApSsidConfigRequest
	GetAppName() *string
	SetRadioIndex(v string) *DeleteApSsidConfigRequest
	GetRadioIndex() *string
	SetSsid(v string) *DeleteApSsidConfigRequest
	GetSsid() *string
}

type DeleteApSsidConfigRequest struct {
	// This parameter is required.
	ApMac *string `json:"ApMac,omitempty" xml:"ApMac,omitempty"`
	// This parameter is required.
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	RadioIndex *string `json:"RadioIndex,omitempty" xml:"RadioIndex,omitempty"`
	// This parameter is required.
	Ssid *string `json:"Ssid,omitempty" xml:"Ssid,omitempty"`
}

func (s DeleteApSsidConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteApSsidConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteApSsidConfigRequest) GetApMac() *string {
	return s.ApMac
}

func (s *DeleteApSsidConfigRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *DeleteApSsidConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *DeleteApSsidConfigRequest) GetRadioIndex() *string {
	return s.RadioIndex
}

func (s *DeleteApSsidConfigRequest) GetSsid() *string {
	return s.Ssid
}

func (s *DeleteApSsidConfigRequest) SetApMac(v string) *DeleteApSsidConfigRequest {
	s.ApMac = &v
	return s
}

func (s *DeleteApSsidConfigRequest) SetAppCode(v string) *DeleteApSsidConfigRequest {
	s.AppCode = &v
	return s
}

func (s *DeleteApSsidConfigRequest) SetAppName(v string) *DeleteApSsidConfigRequest {
	s.AppName = &v
	return s
}

func (s *DeleteApSsidConfigRequest) SetRadioIndex(v string) *DeleteApSsidConfigRequest {
	s.RadioIndex = &v
	return s
}

func (s *DeleteApSsidConfigRequest) SetSsid(v string) *DeleteApSsidConfigRequest {
	s.Ssid = &v
	return s
}

func (s *DeleteApSsidConfigRequest) Validate() error {
	return dara.Validate(s)
}
