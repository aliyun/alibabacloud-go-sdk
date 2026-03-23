// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRepairApRadioRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApMac(v string) *RepairApRadioRequest
	GetApMac() *string
	SetAppCode(v string) *RepairApRadioRequest
	GetAppCode() *string
	SetAppName(v string) *RepairApRadioRequest
	GetAppName() *string
	SetRadioIndex(v string) *RepairApRadioRequest
	GetRadioIndex() *string
}

type RepairApRadioRequest struct {
	// This parameter is required.
	ApMac *string `json:"ApMac,omitempty" xml:"ApMac,omitempty"`
	// This parameter is required.
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	RadioIndex *string `json:"RadioIndex,omitempty" xml:"RadioIndex,omitempty"`
}

func (s RepairApRadioRequest) String() string {
	return dara.Prettify(s)
}

func (s RepairApRadioRequest) GoString() string {
	return s.String()
}

func (s *RepairApRadioRequest) GetApMac() *string {
	return s.ApMac
}

func (s *RepairApRadioRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *RepairApRadioRequest) GetAppName() *string {
	return s.AppName
}

func (s *RepairApRadioRequest) GetRadioIndex() *string {
	return s.RadioIndex
}

func (s *RepairApRadioRequest) SetApMac(v string) *RepairApRadioRequest {
	s.ApMac = &v
	return s
}

func (s *RepairApRadioRequest) SetAppCode(v string) *RepairApRadioRequest {
	s.AppCode = &v
	return s
}

func (s *RepairApRadioRequest) SetAppName(v string) *RepairApRadioRequest {
	s.AppName = &v
	return s
}

func (s *RepairApRadioRequest) SetRadioIndex(v string) *RepairApRadioRequest {
	s.RadioIndex = &v
	return s
}

func (s *RepairApRadioRequest) Validate() error {
	return dara.Validate(s)
}
