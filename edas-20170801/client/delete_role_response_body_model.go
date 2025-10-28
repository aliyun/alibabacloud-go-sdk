// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteRoleResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteRoleResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteRoleResponseBody
	GetRequestId() *string
}

type DeleteRoleResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 042F329B-F518-4CC1-****-**********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRoleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRoleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteRoleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRoleResponseBody) SetCode(v int32) *DeleteRoleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRoleResponseBody) SetMessage(v string) *DeleteRoleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteRoleResponseBody) SetRequestId(v string) *DeleteRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
