// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeResourceServerScopesFromGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RevokeResourceServerScopesFromGroupResponseBody
	GetRequestId() *string
}

type RevokeResourceServerScopesFromGroupResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeResourceServerScopesFromGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeResourceServerScopesFromGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeResourceServerScopesFromGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeResourceServerScopesFromGroupResponseBody) SetRequestId(v string) *RevokeResourceServerScopesFromGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeResourceServerScopesFromGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
