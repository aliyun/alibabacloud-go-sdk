// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMemberAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyMemberAccountResponseBody
	GetRequestId() *string
}

type ModifyMemberAccountResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19****5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyMemberAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyMemberAccountResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMemberAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyMemberAccountResponseBody) SetRequestId(v string) *ModifyMemberAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyMemberAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
