// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryUserProvisioningEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RetryUserProvisioningEventResponseBody
	GetRequestId() *string
}

type RetryUserProvisioningEventResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F6F90F3D-4502-5877-B80B-97476F6AE2CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RetryUserProvisioningEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RetryUserProvisioningEventResponseBody) GoString() string {
	return s.String()
}

func (s *RetryUserProvisioningEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RetryUserProvisioningEventResponseBody) SetRequestId(v string) *RetryUserProvisioningEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetryUserProvisioningEventResponseBody) Validate() error {
	return dara.Validate(s)
}
