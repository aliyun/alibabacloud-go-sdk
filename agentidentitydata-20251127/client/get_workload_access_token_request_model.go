// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkloadAccessTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkloadIdentityName(v string) *GetWorkloadAccessTokenRequest
	GetWorkloadIdentityName() *string
}

type GetWorkloadAccessTokenRequest struct {
	// example:
	//
	// test-workload-identity
	WorkloadIdentityName *string `json:"WorkloadIdentityName,omitempty" xml:"WorkloadIdentityName,omitempty"`
}

func (s GetWorkloadAccessTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkloadAccessTokenRequest) GoString() string {
	return s.String()
}

func (s *GetWorkloadAccessTokenRequest) GetWorkloadIdentityName() *string {
	return s.WorkloadIdentityName
}

func (s *GetWorkloadAccessTokenRequest) SetWorkloadIdentityName(v string) *GetWorkloadAccessTokenRequest {
	s.WorkloadIdentityName = &v
	return s
}

func (s *GetWorkloadAccessTokenRequest) Validate() error {
	return dara.Validate(s)
}
