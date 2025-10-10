// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecurityPolicyAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *UpdateSecurityPolicyAttributeResponseBody
	GetJobId() *string
	SetRequestId(v string) *UpdateSecurityPolicyAttributeResponseBody
	GetRequestId() *string
}

type UpdateSecurityPolicyAttributeResponseBody struct {
	// The asynchronous task ID.
	//
	// example:
	//
	// 72dcd26b-f12d-4c27-b3af-18f6aed5****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 593B0448-D13E-4C56-AC0D-FDF0FDE0E9A3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSecurityPolicyAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecurityPolicyAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSecurityPolicyAttributeResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *UpdateSecurityPolicyAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSecurityPolicyAttributeResponseBody) SetJobId(v string) *UpdateSecurityPolicyAttributeResponseBody {
	s.JobId = &v
	return s
}

func (s *UpdateSecurityPolicyAttributeResponseBody) SetRequestId(v string) *UpdateSecurityPolicyAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSecurityPolicyAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
