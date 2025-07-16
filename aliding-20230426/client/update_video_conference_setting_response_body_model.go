// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVideoConferenceSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCaseResult(v string) *UpdateVideoConferenceSettingResponseBody
	GetCaseResult() *string
	SetCode(v string) *UpdateVideoConferenceSettingResponseBody
	GetCode() *string
	SetRequestId(v string) *UpdateVideoConferenceSettingResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *UpdateVideoConferenceSettingResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *UpdateVideoConferenceSettingResponseBody
	GetVendorType() *string
}

type UpdateVideoConferenceSettingResponseBody struct {
	// example:
	//
	// success
	CaseResult *string `json:"caseResult,omitempty" xml:"caseResult,omitempty"`
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s UpdateVideoConferenceSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateVideoConferenceSettingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVideoConferenceSettingResponseBody) GetCaseResult() *string {
	return s.CaseResult
}

func (s *UpdateVideoConferenceSettingResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateVideoConferenceSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateVideoConferenceSettingResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *UpdateVideoConferenceSettingResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *UpdateVideoConferenceSettingResponseBody) SetCaseResult(v string) *UpdateVideoConferenceSettingResponseBody {
	s.CaseResult = &v
	return s
}

func (s *UpdateVideoConferenceSettingResponseBody) SetCode(v string) *UpdateVideoConferenceSettingResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateVideoConferenceSettingResponseBody) SetRequestId(v string) *UpdateVideoConferenceSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVideoConferenceSettingResponseBody) SetVendorRequestId(v string) *UpdateVideoConferenceSettingResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *UpdateVideoConferenceSettingResponseBody) SetVendorType(v string) *UpdateVideoConferenceSettingResponseBody {
	s.VendorType = &v
	return s
}

func (s *UpdateVideoConferenceSettingResponseBody) Validate() error {
	return dara.Validate(s)
}
