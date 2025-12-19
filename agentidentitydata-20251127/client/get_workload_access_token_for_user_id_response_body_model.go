// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkloadAccessTokenForUserIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetWorkloadAccessTokenForUserIdResponseBody
	GetRequestId() *string
	SetWorkloadAccessToken(v string) *GetWorkloadAccessTokenForUserIdResponseBody
	GetWorkloadAccessToken() *string
}

type GetWorkloadAccessTokenForUserIdResponseBody struct {
	// example:
	//
	// 1E85BA86-0955-5841-9679-9C33867E460D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// eyAgImFsZyI6ICJSUzI1N****
	WorkloadAccessToken *string `json:"WorkloadAccessToken,omitempty" xml:"WorkloadAccessToken,omitempty"`
}

func (s GetWorkloadAccessTokenForUserIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkloadAccessTokenForUserIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkloadAccessTokenForUserIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkloadAccessTokenForUserIdResponseBody) GetWorkloadAccessToken() *string {
	return s.WorkloadAccessToken
}

func (s *GetWorkloadAccessTokenForUserIdResponseBody) SetRequestId(v string) *GetWorkloadAccessTokenForUserIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkloadAccessTokenForUserIdResponseBody) SetWorkloadAccessToken(v string) *GetWorkloadAccessTokenForUserIdResponseBody {
	s.WorkloadAccessToken = &v
	return s
}

func (s *GetWorkloadAccessTokenForUserIdResponseBody) Validate() error {
	return dara.Validate(s)
}
