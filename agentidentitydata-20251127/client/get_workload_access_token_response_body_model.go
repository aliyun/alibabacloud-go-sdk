// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkloadAccessTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetWorkloadAccessTokenResponseBody
	GetRequestId() *string
	SetWorkloadAccessToken(v string) *GetWorkloadAccessTokenResponseBody
	GetWorkloadAccessToken() *string
	SetWorkloadIdentityArn(v string) *GetWorkloadAccessTokenResponseBody
	GetWorkloadIdentityArn() *string
}

type GetWorkloadAccessTokenResponseBody struct {
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WorkloadAccessToken *string `json:"WorkloadAccessToken,omitempty" xml:"WorkloadAccessToken,omitempty"`
	WorkloadIdentityArn *string `json:"WorkloadIdentityArn,omitempty" xml:"WorkloadIdentityArn,omitempty"`
}

func (s GetWorkloadAccessTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkloadAccessTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkloadAccessTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkloadAccessTokenResponseBody) GetWorkloadAccessToken() *string {
	return s.WorkloadAccessToken
}

func (s *GetWorkloadAccessTokenResponseBody) GetWorkloadIdentityArn() *string {
	return s.WorkloadIdentityArn
}

func (s *GetWorkloadAccessTokenResponseBody) SetRequestId(v string) *GetWorkloadAccessTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkloadAccessTokenResponseBody) SetWorkloadAccessToken(v string) *GetWorkloadAccessTokenResponseBody {
	s.WorkloadAccessToken = &v
	return s
}

func (s *GetWorkloadAccessTokenResponseBody) SetWorkloadIdentityArn(v string) *GetWorkloadAccessTokenResponseBody {
	s.WorkloadIdentityArn = &v
	return s
}

func (s *GetWorkloadAccessTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
