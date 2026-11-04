// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveFaceUserFromUserGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RemoveFaceUserFromUserGroupResponseBody
	GetCode() *string
	SetErrorMessage(v string) *RemoveFaceUserFromUserGroupResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *RemoveFaceUserFromUserGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RemoveFaceUserFromUserGroupResponseBody
	GetSuccess() *bool
}

type RemoveFaceUserFromUserGroupResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// success
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0A13279A-5640-45E7-87AA-83750541AD0E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveFaceUserFromUserGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveFaceUserFromUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveFaceUserFromUserGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *RemoveFaceUserFromUserGroupResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RemoveFaceUserFromUserGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveFaceUserFromUserGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveFaceUserFromUserGroupResponseBody) SetCode(v string) *RemoveFaceUserFromUserGroupResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveFaceUserFromUserGroupResponseBody) SetErrorMessage(v string) *RemoveFaceUserFromUserGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RemoveFaceUserFromUserGroupResponseBody) SetRequestId(v string) *RemoveFaceUserFromUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveFaceUserFromUserGroupResponseBody) SetSuccess(v bool) *RemoveFaceUserFromUserGroupResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveFaceUserFromUserGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
