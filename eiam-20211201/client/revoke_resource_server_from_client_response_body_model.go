// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeResourceServerFromClientResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RevokeResourceServerFromClientResponseBody
	GetRequestId() *string
}

type RevokeResourceServerFromClientResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeResourceServerFromClientResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeResourceServerFromClientResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeResourceServerFromClientResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeResourceServerFromClientResponseBody) SetRequestId(v string) *RevokeResourceServerFromClientResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeResourceServerFromClientResponseBody) Validate() error {
	return dara.Validate(s)
}
