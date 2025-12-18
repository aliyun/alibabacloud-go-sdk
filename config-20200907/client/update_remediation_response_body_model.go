// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRemediationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRemediationId(v string) *UpdateRemediationResponseBody
	GetRemediationId() *string
	SetRequestId(v string) *UpdateRemediationResponseBody
	GetRequestId() *string
}

type UpdateRemediationResponseBody struct {
	// The ID of the remediation setting.
	//
	// example:
	//
	// crr-7c2ba2d0236700a3****
	RemediationId *string `json:"RemediationId,omitempty" xml:"RemediationId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C7817373-78CB-4F9A-8AFA-E7A88E9D64A2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRemediationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRemediationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRemediationResponseBody) GetRemediationId() *string {
	return s.RemediationId
}

func (s *UpdateRemediationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRemediationResponseBody) SetRemediationId(v string) *UpdateRemediationResponseBody {
	s.RemediationId = &v
	return s
}

func (s *UpdateRemediationResponseBody) SetRequestId(v string) *UpdateRemediationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRemediationResponseBody) Validate() error {
	return dara.Validate(s)
}
