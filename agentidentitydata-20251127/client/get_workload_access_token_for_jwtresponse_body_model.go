// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkloadAccessTokenForJWTResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetWorkloadAccessTokenForJWTResponseBody
	GetRequestId() *string
	SetWorkloadAccessToken(v string) *GetWorkloadAccessTokenForJWTResponseBody
	GetWorkloadAccessToken() *string
}

type GetWorkloadAccessTokenForJWTResponseBody struct {
	// example:
	//
	// D679769C-957B-586A-AD00-5C2064DAFA7D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// eyAgImFsZyI6ICJSUzI1N****
	WorkloadAccessToken *string `json:"WorkloadAccessToken,omitempty" xml:"WorkloadAccessToken,omitempty"`
}

func (s GetWorkloadAccessTokenForJWTResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkloadAccessTokenForJWTResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkloadAccessTokenForJWTResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkloadAccessTokenForJWTResponseBody) GetWorkloadAccessToken() *string {
	return s.WorkloadAccessToken
}

func (s *GetWorkloadAccessTokenForJWTResponseBody) SetRequestId(v string) *GetWorkloadAccessTokenForJWTResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkloadAccessTokenForJWTResponseBody) SetWorkloadAccessToken(v string) *GetWorkloadAccessTokenForJWTResponseBody {
	s.WorkloadAccessToken = &v
	return s
}

func (s *GetWorkloadAccessTokenForJWTResponseBody) Validate() error {
	return dara.Validate(s)
}
