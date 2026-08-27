// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RemoveUserResponseBody
	GetCode() *string
	SetMessage(v string) *RemoveUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *RemoveUserResponseBody
	GetRequestId() *string
	SetWnUserId(v string) *RemoveUserResponseBody
	GetWnUserId() *string
}

type RemoveUserResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The status code description.
	//
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The WINNEXO platform user ID.
	//
	// example:
	//
	// 1
	WnUserId *string `json:"wnUserId,omitempty" xml:"wnUserId,omitempty"`
}

func (s RemoveUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *RemoveUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemoveUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveUserResponseBody) GetWnUserId() *string {
	return s.WnUserId
}

func (s *RemoveUserResponseBody) SetCode(v string) *RemoveUserResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveUserResponseBody) SetMessage(v string) *RemoveUserResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveUserResponseBody) SetRequestId(v string) *RemoveUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveUserResponseBody) SetWnUserId(v string) *RemoveUserResponseBody {
	s.WnUserId = &v
	return s
}

func (s *RemoveUserResponseBody) Validate() error {
	return dara.Validate(s)
}
