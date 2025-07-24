// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddMembersResponseBody
	GetRequestId() *string
}

type AddMembersResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AddMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddMembersResponseBody) GoString() string {
	return s.String()
}

func (s *AddMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddMembersResponseBody) SetRequestId(v string) *AddMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddMembersResponseBody) Validate() error {
	return dara.Validate(s)
}
