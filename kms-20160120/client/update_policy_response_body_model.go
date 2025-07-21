// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdatePolicyResponseBody
	GetRequestId() *string
}

type UpdatePolicyResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// f455324b-e229-4066-9f58-9c1cf3fe83a8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePolicyResponseBody) SetRequestId(v string) *UpdatePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
