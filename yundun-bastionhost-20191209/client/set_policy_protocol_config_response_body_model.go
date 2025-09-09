// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPolicyProtocolConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetPolicyProtocolConfigResponseBody
	GetRequestId() *string
}

type SetPolicyProtocolConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 81500666-d7f5-4143-8329-0223cc738105
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetPolicyProtocolConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetPolicyProtocolConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetPolicyProtocolConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetPolicyProtocolConfigResponseBody) SetRequestId(v string) *SetPolicyProtocolConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetPolicyProtocolConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
