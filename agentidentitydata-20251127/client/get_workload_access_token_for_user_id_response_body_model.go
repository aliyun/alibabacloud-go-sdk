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
	SetUserContext(v *GetWorkloadAccessTokenForUserIdResponseBodyUserContext) *GetWorkloadAccessTokenForUserIdResponseBody
	GetUserContext() *GetWorkloadAccessTokenForUserIdResponseBodyUserContext
	SetWorkloadAccessToken(v string) *GetWorkloadAccessTokenForUserIdResponseBody
	GetWorkloadAccessToken() *string
	SetWorkloadIdentityArn(v string) *GetWorkloadAccessTokenForUserIdResponseBody
	GetWorkloadIdentityArn() *string
}

type GetWorkloadAccessTokenForUserIdResponseBody struct {
	RequestId           *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserContext         *GetWorkloadAccessTokenForUserIdResponseBodyUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
	WorkloadAccessToken *string                                                 `json:"WorkloadAccessToken,omitempty" xml:"WorkloadAccessToken,omitempty"`
	WorkloadIdentityArn *string                                                 `json:"WorkloadIdentityArn,omitempty" xml:"WorkloadIdentityArn,omitempty"`
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

func (s *GetWorkloadAccessTokenForUserIdResponseBody) GetUserContext() *GetWorkloadAccessTokenForUserIdResponseBodyUserContext {
	return s.UserContext
}

func (s *GetWorkloadAccessTokenForUserIdResponseBody) GetWorkloadAccessToken() *string {
	return s.WorkloadAccessToken
}

func (s *GetWorkloadAccessTokenForUserIdResponseBody) GetWorkloadIdentityArn() *string {
	return s.WorkloadIdentityArn
}

func (s *GetWorkloadAccessTokenForUserIdResponseBody) SetRequestId(v string) *GetWorkloadAccessTokenForUserIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkloadAccessTokenForUserIdResponseBody) SetUserContext(v *GetWorkloadAccessTokenForUserIdResponseBodyUserContext) *GetWorkloadAccessTokenForUserIdResponseBody {
	s.UserContext = v
	return s
}

func (s *GetWorkloadAccessTokenForUserIdResponseBody) SetWorkloadAccessToken(v string) *GetWorkloadAccessTokenForUserIdResponseBody {
	s.WorkloadAccessToken = &v
	return s
}

func (s *GetWorkloadAccessTokenForUserIdResponseBody) SetWorkloadIdentityArn(v string) *GetWorkloadAccessTokenForUserIdResponseBody {
	s.WorkloadIdentityArn = &v
	return s
}

func (s *GetWorkloadAccessTokenForUserIdResponseBody) Validate() error {
	if s.UserContext != nil {
		if err := s.UserContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkloadAccessTokenForUserIdResponseBodyUserContext struct {
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetWorkloadAccessTokenForUserIdResponseBodyUserContext) String() string {
	return dara.Prettify(s)
}

func (s GetWorkloadAccessTokenForUserIdResponseBodyUserContext) GoString() string {
	return s.String()
}

func (s *GetWorkloadAccessTokenForUserIdResponseBodyUserContext) GetUserId() *string {
	return s.UserId
}

func (s *GetWorkloadAccessTokenForUserIdResponseBodyUserContext) SetUserId(v string) *GetWorkloadAccessTokenForUserIdResponseBodyUserContext {
	s.UserId = &v
	return s
}

func (s *GetWorkloadAccessTokenForUserIdResponseBodyUserContext) Validate() error {
	return dara.Validate(s)
}
