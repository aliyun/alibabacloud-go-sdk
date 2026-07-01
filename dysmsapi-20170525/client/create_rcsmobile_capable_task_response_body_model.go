// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRCSMobileCapableTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateRCSMobileCapableTaskResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CreateRCSMobileCapableTaskResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *CreateRCSMobileCapableTaskResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *CreateRCSMobileCapableTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateRCSMobileCapableTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateRCSMobileCapableTaskResponseBody
	GetSuccess() *bool
}

type CreateRCSMobileCapableTaskResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值示例值
	Code *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateRCSMobileCapableTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRCSMobileCapableTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRCSMobileCapableTaskResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateRCSMobileCapableTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateRCSMobileCapableTaskResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *CreateRCSMobileCapableTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateRCSMobileCapableTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRCSMobileCapableTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateRCSMobileCapableTaskResponseBody) SetAccessDeniedDetail(v string) *CreateRCSMobileCapableTaskResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateRCSMobileCapableTaskResponseBody) SetCode(v string) *CreateRCSMobileCapableTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRCSMobileCapableTaskResponseBody) SetData(v map[string]interface{}) *CreateRCSMobileCapableTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateRCSMobileCapableTaskResponseBody) SetMessage(v string) *CreateRCSMobileCapableTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateRCSMobileCapableTaskResponseBody) SetRequestId(v string) *CreateRCSMobileCapableTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRCSMobileCapableTaskResponseBody) SetSuccess(v bool) *CreateRCSMobileCapableTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateRCSMobileCapableTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
