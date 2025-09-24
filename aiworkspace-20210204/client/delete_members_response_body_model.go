// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteMembersResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteMembersResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteMembersResponseBody
	GetRequestId() *string
}

type DeleteMembersResponseBody struct {
	// The error code returned if the call failed.
	//
	// example:
	//
	// 100600017
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned if the call failed.
	//
	// example:
	//
	// Owner not allowed to delete
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D5BFFEE3-6025-443F-8A03-02D619B5C4B9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMembersResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMembersResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteMembersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMembersResponseBody) SetCode(v string) *DeleteMembersResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMembersResponseBody) SetMessage(v string) *DeleteMembersResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMembersResponseBody) SetRequestId(v string) *DeleteMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMembersResponseBody) Validate() error {
	return dara.Validate(s)
}
