// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRcsSignMenuResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *AddRcsSignMenuResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *AddRcsSignMenuResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *AddRcsSignMenuResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *AddRcsSignMenuResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddRcsSignMenuResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddRcsSignMenuResponseBody
	GetSuccess() *bool
}

type AddRcsSignMenuResponseBody struct {
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
	// 示例值示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddRcsSignMenuResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddRcsSignMenuResponseBody) GoString() string {
	return s.String()
}

func (s *AddRcsSignMenuResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *AddRcsSignMenuResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddRcsSignMenuResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *AddRcsSignMenuResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddRcsSignMenuResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddRcsSignMenuResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddRcsSignMenuResponseBody) SetAccessDeniedDetail(v string) *AddRcsSignMenuResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AddRcsSignMenuResponseBody) SetCode(v string) *AddRcsSignMenuResponseBody {
	s.Code = &v
	return s
}

func (s *AddRcsSignMenuResponseBody) SetData(v map[string]interface{}) *AddRcsSignMenuResponseBody {
	s.Data = v
	return s
}

func (s *AddRcsSignMenuResponseBody) SetMessage(v string) *AddRcsSignMenuResponseBody {
	s.Message = &v
	return s
}

func (s *AddRcsSignMenuResponseBody) SetRequestId(v string) *AddRcsSignMenuResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddRcsSignMenuResponseBody) SetSuccess(v bool) *AddRcsSignMenuResponseBody {
	s.Success = &v
	return s
}

func (s *AddRcsSignMenuResponseBody) Validate() error {
	return dara.Validate(s)
}
