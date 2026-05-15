// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAiCallTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateAiCallTaskResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CreateAiCallTaskResponseBody
	GetCode() *string
	SetData(v int64) *CreateAiCallTaskResponseBody
	GetData() *int64
	SetMessage(v string) *CreateAiCallTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAiCallTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAiCallTaskResponseBody
	GetSuccess() *bool
}

type CreateAiCallTaskResponseBody struct {
	// example:
	//
	// Access Denied
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1213123123*****
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 8C93B345-F702-5449-BA7E-7D110D4BF798
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAiCallTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAiCallTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAiCallTaskResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateAiCallTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateAiCallTaskResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateAiCallTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAiCallTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAiCallTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAiCallTaskResponseBody) SetAccessDeniedDetail(v string) *CreateAiCallTaskResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateAiCallTaskResponseBody) SetCode(v string) *CreateAiCallTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAiCallTaskResponseBody) SetData(v int64) *CreateAiCallTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreateAiCallTaskResponseBody) SetMessage(v string) *CreateAiCallTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAiCallTaskResponseBody) SetRequestId(v string) *CreateAiCallTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAiCallTaskResponseBody) SetSuccess(v bool) *CreateAiCallTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAiCallTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
