// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAggregateRemediationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRemediationId(v string) *CreateAggregateRemediationResponseBody
	GetRemediationId() *string
	SetRequestId(v string) *CreateAggregateRemediationResponseBody
	GetRequestId() *string
}

type CreateAggregateRemediationResponseBody struct {
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

func (s CreateAggregateRemediationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAggregateRemediationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAggregateRemediationResponseBody) GetRemediationId() *string {
	return s.RemediationId
}

func (s *CreateAggregateRemediationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAggregateRemediationResponseBody) SetRemediationId(v string) *CreateAggregateRemediationResponseBody {
	s.RemediationId = &v
	return s
}

func (s *CreateAggregateRemediationResponseBody) SetRequestId(v string) *CreateAggregateRemediationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAggregateRemediationResponseBody) Validate() error {
	return dara.Validate(s)
}
