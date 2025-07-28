// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemberAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateMemberAccountsResponseBody
	GetRequestId() *string
}

type CreateMemberAccountsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 66A98669-ER12-WE34-23PO-301469*****E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMemberAccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMemberAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMemberAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMemberAccountsResponseBody) SetRequestId(v string) *CreateMemberAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMemberAccountsResponseBody) Validate() error {
	return dara.Validate(s)
}
