// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSharedAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddSharedAccountsResponseBody
	GetRequestId() *string
}

type AddSharedAccountsResponseBody struct {
	// example:
	//
	// 0D797DC3-FF04-5C21-81EB-XXXXXXXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AddSharedAccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddSharedAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *AddSharedAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddSharedAccountsResponseBody) SetRequestId(v string) *AddSharedAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddSharedAccountsResponseBody) Validate() error {
	return dara.Validate(s)
}
