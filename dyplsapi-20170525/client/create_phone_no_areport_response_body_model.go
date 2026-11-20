// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePhoneNoAReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreatePhoneNoAReportResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CreatePhoneNoAReportResponseBody
	GetCode() *string
	SetData(v *CreatePhoneNoAReportResponseBodyData) *CreatePhoneNoAReportResponseBody
	GetData() *CreatePhoneNoAReportResponseBodyData
	SetMessage(v string) *CreatePhoneNoAReportResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreatePhoneNoAReportResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreatePhoneNoAReportResponseBody
	GetSuccess() *bool
}

type CreatePhoneNoAReportResponseBody struct {
	// Details of denied access.
	//
	// example:
	//
	// 0
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Request status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The struct of the A-number filing result.
	Data *CreatePhoneNoAReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Error message for failure.
	//
	// example:
	//
	// 手机号码***已存在
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The returned ID.
	//
	// example:
	//
	// 1D73E648-0978-18A5-B089-3BB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded. Valid values:
	//
	// - true: Succeeded.
	//
	// - false: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreatePhoneNoAReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePhoneNoAReportResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePhoneNoAReportResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreatePhoneNoAReportResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePhoneNoAReportResponseBody) GetData() *CreatePhoneNoAReportResponseBodyData {
	return s.Data
}

func (s *CreatePhoneNoAReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePhoneNoAReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePhoneNoAReportResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreatePhoneNoAReportResponseBody) SetAccessDeniedDetail(v string) *CreatePhoneNoAReportResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreatePhoneNoAReportResponseBody) SetCode(v string) *CreatePhoneNoAReportResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePhoneNoAReportResponseBody) SetData(v *CreatePhoneNoAReportResponseBodyData) *CreatePhoneNoAReportResponseBody {
	s.Data = v
	return s
}

func (s *CreatePhoneNoAReportResponseBody) SetMessage(v string) *CreatePhoneNoAReportResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePhoneNoAReportResponseBody) SetRequestId(v string) *CreatePhoneNoAReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePhoneNoAReportResponseBody) SetSuccess(v bool) *CreatePhoneNoAReportResponseBody {
	s.Success = &v
	return s
}

func (s *CreatePhoneNoAReportResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePhoneNoAReportResponseBodyData struct {
	// Creation result. Valid values:
	//
	// - true: Succeeded.
	//
	// - false: Failed.
	//
	// example:
	//
	// true
	CreateResult *bool `json:"CreateResult,omitempty" xml:"CreateResult,omitempty"`
	// Enumeration of creation types. A value of 1 indicates success, and negative values indicate creation failure. Valid values:
	//
	// - 1: Succeeded.
	//
	// - -1: The mobile phone number has already been submitted for filing; no need to submit again.
	//
	// - -2: The mobile phone number has already been successfully filed; no need to file again.
	//
	// example:
	//
	// 1
	FailType *int64 `json:"FailType,omitempty" xml:"FailType,omitempty"`
}

func (s CreatePhoneNoAReportResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreatePhoneNoAReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreatePhoneNoAReportResponseBodyData) GetCreateResult() *bool {
	return s.CreateResult
}

func (s *CreatePhoneNoAReportResponseBodyData) GetFailType() *int64 {
	return s.FailType
}

func (s *CreatePhoneNoAReportResponseBodyData) SetCreateResult(v bool) *CreatePhoneNoAReportResponseBodyData {
	s.CreateResult = &v
	return s
}

func (s *CreatePhoneNoAReportResponseBodyData) SetFailType(v int64) *CreatePhoneNoAReportResponseBodyData {
	s.FailType = &v
	return s
}

func (s *CreatePhoneNoAReportResponseBodyData) Validate() error {
	return dara.Validate(s)
}
