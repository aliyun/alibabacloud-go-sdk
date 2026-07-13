// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProtectionPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateProtectionPolicyResponseBody
	GetRequestId() *string
}

type UpdateProtectionPolicyResponseBody struct {
	// The unique ID of the request.
	//
	// example:
	//
	// 86DEBAC9-AB6A-59AB-9E5C-A540E579ECC9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateProtectionPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateProtectionPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProtectionPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateProtectionPolicyResponseBody) SetRequestId(v string) *UpdateProtectionPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProtectionPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
