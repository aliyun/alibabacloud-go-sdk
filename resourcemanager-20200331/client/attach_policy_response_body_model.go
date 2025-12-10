// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachPolicyResponseBody
	GetRequestId() *string
}

type AttachPolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 697852FB-50D7-44D9-9774-530C31EAC572
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *AttachPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachPolicyResponseBody) SetRequestId(v string) *AttachPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
