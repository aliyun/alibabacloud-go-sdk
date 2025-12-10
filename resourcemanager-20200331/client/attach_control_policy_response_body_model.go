// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachControlPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachControlPolicyResponseBody
	GetRequestId() *string
}

type AttachControlPolicyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 95060F1D-6990-4645-8920-A81D1BBFE992
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachControlPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *AttachControlPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachControlPolicyResponseBody) SetRequestId(v string) *AttachControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachControlPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
