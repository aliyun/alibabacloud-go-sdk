// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkloadAccessTokenForUserIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v string) *GetWorkloadAccessTokenForUserIdRequest
	GetUserId() *string
	SetWorkloadIdentityName(v string) *GetWorkloadAccessTokenForUserIdRequest
	GetWorkloadIdentityName() *string
}

type GetWorkloadAccessTokenForUserIdRequest struct {
	// example:
	//
	// externalUser1
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// test-workload-identity
	WorkloadIdentityName *string `json:"WorkloadIdentityName,omitempty" xml:"WorkloadIdentityName,omitempty"`
}

func (s GetWorkloadAccessTokenForUserIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkloadAccessTokenForUserIdRequest) GoString() string {
	return s.String()
}

func (s *GetWorkloadAccessTokenForUserIdRequest) GetUserId() *string {
	return s.UserId
}

func (s *GetWorkloadAccessTokenForUserIdRequest) GetWorkloadIdentityName() *string {
	return s.WorkloadIdentityName
}

func (s *GetWorkloadAccessTokenForUserIdRequest) SetUserId(v string) *GetWorkloadAccessTokenForUserIdRequest {
	s.UserId = &v
	return s
}

func (s *GetWorkloadAccessTokenForUserIdRequest) SetWorkloadIdentityName(v string) *GetWorkloadAccessTokenForUserIdRequest {
	s.WorkloadIdentityName = &v
	return s
}

func (s *GetWorkloadAccessTokenForUserIdRequest) Validate() error {
	return dara.Validate(s)
}
