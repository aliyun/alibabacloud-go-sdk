// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactsByIdsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteContactsByIdsResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DeleteContactsByIdsResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteContactsByIdsResponseBody
	GetMessage() *string
	SetModel(v map[string]interface{}) *DeleteContactsByIdsResponseBody
	GetModel() map[string]interface{}
	SetRequestId(v string) *DeleteContactsByIdsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteContactsByIdsResponseBody
	GetSuccess() *bool
}

type DeleteContactsByIdsResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// True
	Model map[string]interface{} `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// 示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteContactsByIdsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactsByIdsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContactsByIdsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteContactsByIdsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteContactsByIdsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteContactsByIdsResponseBody) GetModel() map[string]interface{} {
	return s.Model
}

func (s *DeleteContactsByIdsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteContactsByIdsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteContactsByIdsResponseBody) SetAccessDeniedDetail(v string) *DeleteContactsByIdsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteContactsByIdsResponseBody) SetCode(v string) *DeleteContactsByIdsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteContactsByIdsResponseBody) SetMessage(v string) *DeleteContactsByIdsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteContactsByIdsResponseBody) SetModel(v map[string]interface{}) *DeleteContactsByIdsResponseBody {
	s.Model = v
	return s
}

func (s *DeleteContactsByIdsResponseBody) SetRequestId(v string) *DeleteContactsByIdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteContactsByIdsResponseBody) SetSuccess(v bool) *DeleteContactsByIdsResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteContactsByIdsResponseBody) Validate() error {
	return dara.Validate(s)
}
