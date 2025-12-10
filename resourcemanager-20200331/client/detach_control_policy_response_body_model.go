// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachControlPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachControlPolicyResponseBody
	GetRequestId() *string
}

type DetachControlPolicyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9EA4F962-1A2E-4AFE-BE0C-B14736FC46CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachControlPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DetachControlPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachControlPolicyResponseBody) SetRequestId(v string) *DetachControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachControlPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
