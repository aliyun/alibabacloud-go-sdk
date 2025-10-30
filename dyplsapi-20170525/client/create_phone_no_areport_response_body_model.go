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
	// example:
	//
	// 0
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// 请求状态码
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// A号码报备结果结构体
	Data *CreatePhoneNoAReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 失败错误提示
	//
	// example:
	//
	// 手机号码***已存在
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 返回id
	//
	// example:
	//
	// 1D73E648-0978-18A5-B089-3BB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 请求是否成功
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
	// 创建结果
	//
	// example:
	//
	// true
	CreateResult *bool `json:"CreateResult,omitempty" xml:"CreateResult,omitempty"`
	// 创建类型枚举，1为成功，负数为创建失败
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
