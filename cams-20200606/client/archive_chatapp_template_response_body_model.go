// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iArchiveChatappTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ArchiveChatappTemplateResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ArchiveChatappTemplateResponseBody
	GetCode() *string
	SetData(v *ArchiveChatappTemplateResponseBodyData) *ArchiveChatappTemplateResponseBody
	GetData() *ArchiveChatappTemplateResponseBodyData
	SetMessage(v string) *ArchiveChatappTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *ArchiveChatappTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ArchiveChatappTemplateResponseBody
	GetSuccess() *bool
}

type ArchiveChatappTemplateResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The request status code.
	//
	// - OK indicates that the request was successful.
	//
	// - For other error codes, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ArchiveChatappTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - **true**: The call was successful.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ArchiveChatappTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ArchiveChatappTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ArchiveChatappTemplateResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ArchiveChatappTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *ArchiveChatappTemplateResponseBody) GetData() *ArchiveChatappTemplateResponseBodyData {
	return s.Data
}

func (s *ArchiveChatappTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ArchiveChatappTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ArchiveChatappTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ArchiveChatappTemplateResponseBody) SetAccessDeniedDetail(v string) *ArchiveChatappTemplateResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ArchiveChatappTemplateResponseBody) SetCode(v string) *ArchiveChatappTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *ArchiveChatappTemplateResponseBody) SetData(v *ArchiveChatappTemplateResponseBodyData) *ArchiveChatappTemplateResponseBody {
	s.Data = v
	return s
}

func (s *ArchiveChatappTemplateResponseBody) SetMessage(v string) *ArchiveChatappTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *ArchiveChatappTemplateResponseBody) SetRequestId(v string) *ArchiveChatappTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ArchiveChatappTemplateResponseBody) SetSuccess(v bool) *ArchiveChatappTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *ArchiveChatappTemplateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ArchiveChatappTemplateResponseBodyData struct {
	// The request status code.
	//
	// - OK indicates that the request was successful.
	//
	// - For other error codes, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// 示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 示例值示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - **true**: The call was successful.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ArchiveChatappTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ArchiveChatappTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *ArchiveChatappTemplateResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *ArchiveChatappTemplateResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *ArchiveChatappTemplateResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *ArchiveChatappTemplateResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *ArchiveChatappTemplateResponseBodyData) SetCode(v string) *ArchiveChatappTemplateResponseBodyData {
	s.Code = &v
	return s
}

func (s *ArchiveChatappTemplateResponseBodyData) SetMessage(v string) *ArchiveChatappTemplateResponseBodyData {
	s.Message = &v
	return s
}

func (s *ArchiveChatappTemplateResponseBodyData) SetRequestId(v string) *ArchiveChatappTemplateResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *ArchiveChatappTemplateResponseBodyData) SetSuccess(v bool) *ArchiveChatappTemplateResponseBodyData {
	s.Success = &v
	return s
}

func (s *ArchiveChatappTemplateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
