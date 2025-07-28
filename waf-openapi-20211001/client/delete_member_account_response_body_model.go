// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMemberAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteMemberAccountResponseBody
	GetRequestId() *string
}

type DeleteMemberAccountResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5ABE714C-8890-5D7E-A08B-45CB****5473
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMemberAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMemberAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMemberAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMemberAccountResponseBody) SetRequestId(v string) *DeleteMemberAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMemberAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
