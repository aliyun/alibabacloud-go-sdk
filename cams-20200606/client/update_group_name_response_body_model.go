// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGroupNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateGroupNameResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UpdateGroupNameResponseBody
	GetCode() *string
	SetData(v string) *UpdateGroupNameResponseBody
	GetData() *string
	SetMessage(v string) *UpdateGroupNameResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateGroupNameResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateGroupNameResponseBody
	GetSuccess() *bool
}

type UpdateGroupNameResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s UpdateGroupNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGroupNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupNameResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateGroupNameResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateGroupNameResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateGroupNameResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGroupNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGroupNameResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateGroupNameResponseBody) SetAccessDeniedDetail(v string) *UpdateGroupNameResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateGroupNameResponseBody) SetCode(v string) *UpdateGroupNameResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGroupNameResponseBody) SetData(v string) *UpdateGroupNameResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateGroupNameResponseBody) SetMessage(v string) *UpdateGroupNameResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGroupNameResponseBody) SetRequestId(v string) *UpdateGroupNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGroupNameResponseBody) SetSuccess(v bool) *UpdateGroupNameResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateGroupNameResponseBody) Validate() error {
	return dara.Validate(s)
}
