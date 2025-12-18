// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRemediationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRemediationId(v string) *CreateRemediationResponseBody
	GetRemediationId() *string
	SetRequestId(v string) *CreateRemediationResponseBody
	GetRequestId() *string
}

type CreateRemediationResponseBody struct {
	// The ID of the remediation template.
	//
	// example:
	//
	// crr-909ba2d4716700eb****
	RemediationId *string `json:"RemediationId,omitempty" xml:"RemediationId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C7817373-78CB-4F9A-8AFA-E7A88E9D64A2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRemediationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRemediationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRemediationResponseBody) GetRemediationId() *string {
	return s.RemediationId
}

func (s *CreateRemediationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRemediationResponseBody) SetRemediationId(v string) *CreateRemediationResponseBody {
	s.RemediationId = &v
	return s
}

func (s *CreateRemediationResponseBody) SetRequestId(v string) *CreateRemediationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRemediationResponseBody) Validate() error {
	return dara.Validate(s)
}
