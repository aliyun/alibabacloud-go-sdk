// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProtectionPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteProtectionPolicyResponseBody
	GetRequestId() *string
}

type DeleteProtectionPolicyResponseBody struct {
	// The unique identifier of the request.
	//
	// example:
	//
	// 5B2F09BF-CEBD-5A7E-AC01-E7F86169A5E5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteProtectionPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteProtectionPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProtectionPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteProtectionPolicyResponseBody) SetRequestId(v string) *DeleteProtectionPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProtectionPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
