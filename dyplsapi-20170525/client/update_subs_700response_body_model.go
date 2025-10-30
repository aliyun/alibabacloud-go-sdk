// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSubs700ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateSubs700ResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UpdateSubs700ResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *UpdateSubs700ResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *UpdateSubs700ResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateSubs700ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateSubs700ResponseBody
	GetSuccess() *bool
}

type UpdateSubs700ResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {}
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

func (s UpdateSubs700ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSubs700ResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSubs700ResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateSubs700ResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateSubs700ResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *UpdateSubs700ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateSubs700ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSubs700ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateSubs700ResponseBody) SetAccessDeniedDetail(v string) *UpdateSubs700ResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateSubs700ResponseBody) SetCode(v string) *UpdateSubs700ResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSubs700ResponseBody) SetData(v map[string]interface{}) *UpdateSubs700ResponseBody {
	s.Data = v
	return s
}

func (s *UpdateSubs700ResponseBody) SetMessage(v string) *UpdateSubs700ResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSubs700ResponseBody) SetRequestId(v string) *UpdateSubs700ResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSubs700ResponseBody) SetSuccess(v bool) *UpdateSubs700ResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateSubs700ResponseBody) Validate() error {
	return dara.Validate(s)
}
