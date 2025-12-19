// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkloadAccessTokenForJWTRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserToken(v string) *GetWorkloadAccessTokenForJWTRequest
	GetUserToken() *string
	SetWorkloadIdentityName(v string) *GetWorkloadAccessTokenForJWTRequest
	GetWorkloadIdentityName() *string
}

type GetWorkloadAccessTokenForJWTRequest struct {
	// example:
	//
	// eyAgImFsZyI6ICJSUzI1NiIsI*******
	UserToken *string `json:"UserToken,omitempty" xml:"UserToken,omitempty"`
	// example:
	//
	// test-workload-identity
	WorkloadIdentityName *string `json:"WorkloadIdentityName,omitempty" xml:"WorkloadIdentityName,omitempty"`
}

func (s GetWorkloadAccessTokenForJWTRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkloadAccessTokenForJWTRequest) GoString() string {
	return s.String()
}

func (s *GetWorkloadAccessTokenForJWTRequest) GetUserToken() *string {
	return s.UserToken
}

func (s *GetWorkloadAccessTokenForJWTRequest) GetWorkloadIdentityName() *string {
	return s.WorkloadIdentityName
}

func (s *GetWorkloadAccessTokenForJWTRequest) SetUserToken(v string) *GetWorkloadAccessTokenForJWTRequest {
	s.UserToken = &v
	return s
}

func (s *GetWorkloadAccessTokenForJWTRequest) SetWorkloadIdentityName(v string) *GetWorkloadAccessTokenForJWTRequest {
	s.WorkloadIdentityName = &v
	return s
}

func (s *GetWorkloadAccessTokenForJWTRequest) Validate() error {
	return dara.Validate(s)
}
