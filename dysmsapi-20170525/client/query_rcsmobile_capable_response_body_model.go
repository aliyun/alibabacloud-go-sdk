// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRCSMobileCapableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryRCSMobileCapableResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QueryRCSMobileCapableResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *QueryRCSMobileCapableResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *QueryRCSMobileCapableResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryRCSMobileCapableResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryRCSMobileCapableResponseBody
	GetSuccess() *bool
}

type QueryRCSMobileCapableResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值示例值
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
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryRCSMobileCapableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryRCSMobileCapableResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRCSMobileCapableResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryRCSMobileCapableResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryRCSMobileCapableResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *QueryRCSMobileCapableResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryRCSMobileCapableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryRCSMobileCapableResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryRCSMobileCapableResponseBody) SetAccessDeniedDetail(v string) *QueryRCSMobileCapableResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryRCSMobileCapableResponseBody) SetCode(v string) *QueryRCSMobileCapableResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRCSMobileCapableResponseBody) SetData(v map[string]interface{}) *QueryRCSMobileCapableResponseBody {
	s.Data = v
	return s
}

func (s *QueryRCSMobileCapableResponseBody) SetMessage(v string) *QueryRCSMobileCapableResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRCSMobileCapableResponseBody) SetRequestId(v string) *QueryRCSMobileCapableResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRCSMobileCapableResponseBody) SetSuccess(v bool) *QueryRCSMobileCapableResponseBody {
	s.Success = &v
	return s
}

func (s *QueryRCSMobileCapableResponseBody) Validate() error {
	return dara.Validate(s)
}
