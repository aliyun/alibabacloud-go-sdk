// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRotationPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateRotationPolicyResponseBody
	GetRequestId() *string
}

type UpdateRotationPolicyResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// efb1cbbd-a093-4278-bc03-639dd4fcc207
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRotationPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRotationPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRotationPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRotationPolicyResponseBody) SetRequestId(v string) *UpdateRotationPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRotationPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
