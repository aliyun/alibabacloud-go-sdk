// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFixedNoAReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateFixedNoAReportResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CreateFixedNoAReportResponseBody
	GetCode() *string
	SetData(v *CreateFixedNoAReportResponseBodyData) *CreateFixedNoAReportResponseBody
	GetData() *CreateFixedNoAReportResponseBodyData
	SetMessage(v string) *CreateFixedNoAReportResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateFixedNoAReportResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateFixedNoAReportResponseBody
	GetSuccess() *bool
}

type CreateFixedNoAReportResponseBody struct {
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
	// The struct of the reporting result for number A.
	Data *CreateFixedNoAReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Error message for failure.
	//
	// example:
	//
	// 固话号码***已存在
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

func (s CreateFixedNoAReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFixedNoAReportResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFixedNoAReportResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateFixedNoAReportResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateFixedNoAReportResponseBody) GetData() *CreateFixedNoAReportResponseBodyData {
	return s.Data
}

func (s *CreateFixedNoAReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateFixedNoAReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFixedNoAReportResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateFixedNoAReportResponseBody) SetAccessDeniedDetail(v string) *CreateFixedNoAReportResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateFixedNoAReportResponseBody) SetCode(v string) *CreateFixedNoAReportResponseBody {
	s.Code = &v
	return s
}

func (s *CreateFixedNoAReportResponseBody) SetData(v *CreateFixedNoAReportResponseBodyData) *CreateFixedNoAReportResponseBody {
	s.Data = v
	return s
}

func (s *CreateFixedNoAReportResponseBody) SetMessage(v string) *CreateFixedNoAReportResponseBody {
	s.Message = &v
	return s
}

func (s *CreateFixedNoAReportResponseBody) SetRequestId(v string) *CreateFixedNoAReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFixedNoAReportResponseBody) SetSuccess(v bool) *CreateFixedNoAReportResponseBody {
	s.Success = &v
	return s
}

func (s *CreateFixedNoAReportResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateFixedNoAReportResponseBodyData struct {
	// Creation result.
	//
	// example:
	//
	// true
	CreateResult *bool `json:"CreateResult,omitempty" xml:"CreateResult,omitempty"`
	// Enumeration of creation types. The value 1 indicates success, and negative numbers indicate creation failure. Valid values:
	//
	// - 1: Succeeded.
	//
	// - -1: The phone number has already been submitted for reporting; no need to submit again.
	//
	// - -2: The phone number has already been successfully reported; no need to report again.
	//
	// example:
	//
	// 1
	FailType *int64 `json:"FailType,omitempty" xml:"FailType,omitempty"`
}

func (s CreateFixedNoAReportResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateFixedNoAReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateFixedNoAReportResponseBodyData) GetCreateResult() *bool {
	return s.CreateResult
}

func (s *CreateFixedNoAReportResponseBodyData) GetFailType() *int64 {
	return s.FailType
}

func (s *CreateFixedNoAReportResponseBodyData) SetCreateResult(v bool) *CreateFixedNoAReportResponseBodyData {
	s.CreateResult = &v
	return s
}

func (s *CreateFixedNoAReportResponseBodyData) SetFailType(v int64) *CreateFixedNoAReportResponseBodyData {
	s.FailType = &v
	return s
}

func (s *CreateFixedNoAReportResponseBodyData) Validate() error {
	return dara.Validate(s)
}
