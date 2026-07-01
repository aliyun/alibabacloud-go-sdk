// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRCSMobileCapableTaskResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryRCSMobileCapableTaskResultResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QueryRCSMobileCapableTaskResultResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *QueryRCSMobileCapableTaskResultResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *QueryRCSMobileCapableTaskResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryRCSMobileCapableTaskResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryRCSMobileCapableTaskResultResponseBody
	GetSuccess() *bool
}

type QueryRCSMobileCapableTaskResultResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值
	Code *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryRCSMobileCapableTaskResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryRCSMobileCapableTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRCSMobileCapableTaskResultResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryRCSMobileCapableTaskResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryRCSMobileCapableTaskResultResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *QueryRCSMobileCapableTaskResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryRCSMobileCapableTaskResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryRCSMobileCapableTaskResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryRCSMobileCapableTaskResultResponseBody) SetAccessDeniedDetail(v string) *QueryRCSMobileCapableTaskResultResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryRCSMobileCapableTaskResultResponseBody) SetCode(v string) *QueryRCSMobileCapableTaskResultResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRCSMobileCapableTaskResultResponseBody) SetData(v map[string]interface{}) *QueryRCSMobileCapableTaskResultResponseBody {
	s.Data = v
	return s
}

func (s *QueryRCSMobileCapableTaskResultResponseBody) SetMessage(v string) *QueryRCSMobileCapableTaskResultResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRCSMobileCapableTaskResultResponseBody) SetRequestId(v string) *QueryRCSMobileCapableTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRCSMobileCapableTaskResultResponseBody) SetSuccess(v bool) *QueryRCSMobileCapableTaskResultResponseBody {
	s.Success = &v
	return s
}

func (s *QueryRCSMobileCapableTaskResultResponseBody) Validate() error {
	return dara.Validate(s)
}
