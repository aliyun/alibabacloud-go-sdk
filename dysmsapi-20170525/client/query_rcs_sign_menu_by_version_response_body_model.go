// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRcsSignMenuByVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryRcsSignMenuByVersionResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QueryRcsSignMenuByVersionResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *QueryRcsSignMenuByVersionResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *QueryRcsSignMenuByVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryRcsSignMenuByVersionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryRcsSignMenuByVersionResponseBody
	GetSuccess() *bool
}

type QueryRcsSignMenuByVersionResponseBody struct {
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
	// 示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryRcsSignMenuByVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryRcsSignMenuByVersionResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRcsSignMenuByVersionResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryRcsSignMenuByVersionResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryRcsSignMenuByVersionResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *QueryRcsSignMenuByVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryRcsSignMenuByVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryRcsSignMenuByVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryRcsSignMenuByVersionResponseBody) SetAccessDeniedDetail(v string) *QueryRcsSignMenuByVersionResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryRcsSignMenuByVersionResponseBody) SetCode(v string) *QueryRcsSignMenuByVersionResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRcsSignMenuByVersionResponseBody) SetData(v map[string]interface{}) *QueryRcsSignMenuByVersionResponseBody {
	s.Data = v
	return s
}

func (s *QueryRcsSignMenuByVersionResponseBody) SetMessage(v string) *QueryRcsSignMenuByVersionResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRcsSignMenuByVersionResponseBody) SetRequestId(v string) *QueryRcsSignMenuByVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRcsSignMenuByVersionResponseBody) SetSuccess(v bool) *QueryRcsSignMenuByVersionResponseBody {
	s.Success = &v
	return s
}

func (s *QueryRcsSignMenuByVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
