// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkloadIdentityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateWorkloadIdentityResponseBody
	GetRequestId() *string
}

type UpdateWorkloadIdentityResponseBody struct {
	// example:
	//
	// 69469F02-DFA8-5733-960F-AEE7CBC7C436
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateWorkloadIdentityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkloadIdentityResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkloadIdentityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWorkloadIdentityResponseBody) SetRequestId(v string) *UpdateWorkloadIdentityResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkloadIdentityResponseBody) Validate() error {
	return dara.Validate(s)
}
