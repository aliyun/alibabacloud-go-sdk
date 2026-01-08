// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *AddGroupResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *AddGroupResponseBody
	GetCode() *string
	SetData(v string) *AddGroupResponseBody
	GetData() *string
	SetMessage(v string) *AddGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddGroupResponseBody
	GetSuccess() *bool
}

type AddGroupResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// true
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// dgdf5-bvcv**
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddGroupResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *AddGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddGroupResponseBody) GetData() *string {
	return s.Data
}

func (s *AddGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddGroupResponseBody) SetAccessDeniedDetail(v string) *AddGroupResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AddGroupResponseBody) SetCode(v string) *AddGroupResponseBody {
	s.Code = &v
	return s
}

func (s *AddGroupResponseBody) SetData(v string) *AddGroupResponseBody {
	s.Data = &v
	return s
}

func (s *AddGroupResponseBody) SetMessage(v string) *AddGroupResponseBody {
	s.Message = &v
	return s
}

func (s *AddGroupResponseBody) SetRequestId(v string) *AddGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddGroupResponseBody) SetSuccess(v bool) *AddGroupResponseBody {
	s.Success = &v
	return s
}

func (s *AddGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
