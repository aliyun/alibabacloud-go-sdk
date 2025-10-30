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
	Data *CreateFixedNoAReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 失败错误提示
	//
	// example:
	//
	// 示例值示例值
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
