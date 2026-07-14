// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveContactToGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *MoveContactToGroupResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *MoveContactToGroupResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *MoveContactToGroupResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *MoveContactToGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *MoveContactToGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MoveContactToGroupResponseBody
	GetSuccess() *bool
}

type MoveContactToGroupResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MoveContactToGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MoveContactToGroupResponseBody) GoString() string {
	return s.String()
}

func (s *MoveContactToGroupResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *MoveContactToGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *MoveContactToGroupResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *MoveContactToGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MoveContactToGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MoveContactToGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MoveContactToGroupResponseBody) SetAccessDeniedDetail(v string) *MoveContactToGroupResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *MoveContactToGroupResponseBody) SetCode(v string) *MoveContactToGroupResponseBody {
	s.Code = &v
	return s
}

func (s *MoveContactToGroupResponseBody) SetData(v map[string]interface{}) *MoveContactToGroupResponseBody {
	s.Data = v
	return s
}

func (s *MoveContactToGroupResponseBody) SetMessage(v string) *MoveContactToGroupResponseBody {
	s.Message = &v
	return s
}

func (s *MoveContactToGroupResponseBody) SetRequestId(v string) *MoveContactToGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveContactToGroupResponseBody) SetSuccess(v bool) *MoveContactToGroupResponseBody {
	s.Success = &v
	return s
}

func (s *MoveContactToGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
