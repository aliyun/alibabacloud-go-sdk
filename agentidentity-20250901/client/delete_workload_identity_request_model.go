// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkloadIdentityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkloadIdentityName(v string) *DeleteWorkloadIdentityRequest
	GetWorkloadIdentityName() *string
}

type DeleteWorkloadIdentityRequest struct {
	// example:
	//
	// agent-101
	WorkloadIdentityName *string `json:"WorkloadIdentityName,omitempty" xml:"WorkloadIdentityName,omitempty"`
}

func (s DeleteWorkloadIdentityRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkloadIdentityRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkloadIdentityRequest) GetWorkloadIdentityName() *string {
	return s.WorkloadIdentityName
}

func (s *DeleteWorkloadIdentityRequest) SetWorkloadIdentityName(v string) *DeleteWorkloadIdentityRequest {
	s.WorkloadIdentityName = &v
	return s
}

func (s *DeleteWorkloadIdentityRequest) Validate() error {
	return dara.Validate(s)
}
