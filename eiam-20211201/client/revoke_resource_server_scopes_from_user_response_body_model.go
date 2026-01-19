// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeResourceServerScopesFromUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RevokeResourceServerScopesFromUserResponseBody
	GetRequestId() *string
}

type RevokeResourceServerScopesFromUserResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeResourceServerScopesFromUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeResourceServerScopesFromUserResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeResourceServerScopesFromUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeResourceServerScopesFromUserResponseBody) SetRequestId(v string) *RevokeResourceServerScopesFromUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeResourceServerScopesFromUserResponseBody) Validate() error {
	return dara.Validate(s)
}
