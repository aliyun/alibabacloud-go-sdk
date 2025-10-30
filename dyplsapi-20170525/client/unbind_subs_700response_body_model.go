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
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 200
	Code *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9297B722-A016-43FB-B51A-E54050******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
