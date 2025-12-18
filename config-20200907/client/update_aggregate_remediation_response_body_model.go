// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAggregateRemediationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRemediationId(v string) *UpdateAggregateRemediationResponseBody
	GetRemediationId() *string
	SetRequestId(v string) *UpdateAggregateRemediationResponseBody
	GetRequestId() *string
}

type UpdateAggregateRemediationResponseBody struct {
	// The ID of the remediation setting.
	//
	// example:
	//
	// crr-909ba2d4716700eb****
	RemediationId *string `json:"RemediationId,omitempty" xml:"RemediationId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C7817373-78CB-4F9A-8AFA-E7A88E9D64A2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAggregateRemediationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAggregateRemediationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAggregateRemediationResponseBody) GetRemediationId() *string {
	return s.RemediationId
}

func (s *UpdateAggregateRemediationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAggregateRemediationResponseBody) SetRemediationId(v string) *UpdateAggregateRemediationResponseBody {
	s.RemediationId = &v
	return s
}

func (s *UpdateAggregateRemediationResponseBody) SetRequestId(v string) *UpdateAggregateRemediationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAggregateRemediationResponseBody) Validate() error {
	return dara.Validate(s)
}
