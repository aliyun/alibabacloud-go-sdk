// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScheduleConfSettingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateScheduleConfSettingsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateScheduleConfSettingsResponseBody
	GetSuccess() *bool
	SetVendorRequestId(v string) *UpdateScheduleConfSettingsResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *UpdateScheduleConfSettingsResponseBody
	GetVendorType() *string
}

type UpdateScheduleConfSettingsResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s UpdateScheduleConfSettingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduleConfSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateScheduleConfSettingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateScheduleConfSettingsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateScheduleConfSettingsResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *UpdateScheduleConfSettingsResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *UpdateScheduleConfSettingsResponseBody) SetRequestId(v string) *UpdateScheduleConfSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateScheduleConfSettingsResponseBody) SetSuccess(v bool) *UpdateScheduleConfSettingsResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateScheduleConfSettingsResponseBody) SetVendorRequestId(v string) *UpdateScheduleConfSettingsResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *UpdateScheduleConfSettingsResponseBody) SetVendorType(v string) *UpdateScheduleConfSettingsResponseBody {
	s.VendorType = &v
	return s
}

func (s *UpdateScheduleConfSettingsResponseBody) Validate() error {
	return dara.Validate(s)
}
