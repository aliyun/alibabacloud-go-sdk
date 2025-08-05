// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddInstanceMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddInstanceMembersResponseBody
	GetRequestId() *string
}

type AddInstanceMembersResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B40A54DF-C142-44F7-8441-B31C1EADB36E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddInstanceMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddInstanceMembersResponseBody) GoString() string {
	return s.String()
}

func (s *AddInstanceMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddInstanceMembersResponseBody) SetRequestId(v string) *AddInstanceMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddInstanceMembersResponseBody) Validate() error {
	return dara.Validate(s)
}
