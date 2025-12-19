// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkloadIdentityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkloadIdentityName(v string) *GetWorkloadIdentityRequest
	GetWorkloadIdentityName() *string
}

type GetWorkloadIdentityRequest struct {
	// example:
	//
	// agent-101
	WorkloadIdentityName *string `json:"WorkloadIdentityName,omitempty" xml:"WorkloadIdentityName,omitempty"`
}

func (s GetWorkloadIdentityRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkloadIdentityRequest) GoString() string {
	return s.String()
}

func (s *GetWorkloadIdentityRequest) GetWorkloadIdentityName() *string {
	return s.WorkloadIdentityName
}

func (s *GetWorkloadIdentityRequest) SetWorkloadIdentityName(v string) *GetWorkloadIdentityRequest {
	s.WorkloadIdentityName = &v
	return s
}

func (s *GetWorkloadIdentityRequest) Validate() error {
	return dara.Validate(s)
}
