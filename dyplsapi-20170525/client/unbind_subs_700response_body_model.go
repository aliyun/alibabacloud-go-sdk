// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindSubs700ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UnbindSubs700ResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UnbindSubs700ResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *UnbindSubs700ResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *UnbindSubs700ResponseBody
	GetMessage() *string
	SetRequestId(v string) *UnbindSubs700ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UnbindSubs700ResponseBody
	GetSuccess() *bool
}

type UnbindSubs700ResponseBody struct {
	// The details of the permission check failure.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The status code of the API request. Valid values:
	//
	// - OK: The request is successful.
	//
	// - Other error codes. For more information, see [API error codes](https://help.aliyun.com/document_detail/59892.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The default response format. This parameter has no actual value.
	//
	// example:
	//
	// {}
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The response message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9297B722-A016-43FB-B51A-E54050******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Specifies whether the API call is successful. Valid values:
	//
	// - true: The API call is successful.
	//
	// - false: The API call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnbindSubs700ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindSubs700ResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindSubs700ResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UnbindSubs700ResponseBody) GetCode() *string {
	return s.Code
}

func (s *UnbindSubs700ResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *UnbindSubs700ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UnbindSubs700ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindSubs700ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UnbindSubs700ResponseBody) SetAccessDeniedDetail(v string) *UnbindSubs700ResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UnbindSubs700ResponseBody) SetCode(v string) *UnbindSubs700ResponseBody {
	s.Code = &v
	return s
}

func (s *UnbindSubs700ResponseBody) SetData(v map[string]interface{}) *UnbindSubs700ResponseBody {
	s.Data = v
	return s
}

func (s *UnbindSubs700ResponseBody) SetMessage(v string) *UnbindSubs700ResponseBody {
	s.Message = &v
	return s
}

func (s *UnbindSubs700ResponseBody) SetRequestId(v string) *UnbindSubs700ResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindSubs700ResponseBody) SetSuccess(v bool) *UnbindSubs700ResponseBody {
	s.Success = &v
	return s
}

func (s *UnbindSubs700ResponseBody) Validate() error {
	return dara.Validate(s)
}
