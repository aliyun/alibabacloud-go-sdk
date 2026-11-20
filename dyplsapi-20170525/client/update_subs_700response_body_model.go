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
	// The details of the access denial.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The status code of the API request. Valid values:
	//
	// - OK: The request is successful.
	//
	// - Other error codes: For more information, see [API error codes](https://help.aliyun.com/document_detail/59892.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The default response format. This field does not contain an actual value.
	//
	// example:
	//
	// {}
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9297B722-A016-43FB-B51A-E54050******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call is successful. Valid values:
	//
	// - true: The call is successful.
	//
	// - false: The call fails.
	//
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
