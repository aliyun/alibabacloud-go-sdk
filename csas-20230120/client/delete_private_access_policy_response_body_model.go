// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrivateAccessPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePrivateAccessPolicyResponseBody
	GetRequestId() *string
}

type DeletePrivateAccessPolicyResponseBody struct {
	// example:
	//
	// 4D169859-A4F2-5EC8-853B-8447787C0D8A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePrivateAccessPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePrivateAccessPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePrivateAccessPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePrivateAccessPolicyResponseBody) SetRequestId(v string) *DeletePrivateAccessPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePrivateAccessPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
