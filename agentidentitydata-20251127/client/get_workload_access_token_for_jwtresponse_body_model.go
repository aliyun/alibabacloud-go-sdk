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
	SetUserContext(v *GetWorkloadAccessTokenForJWTResponseBodyUserContext) *GetWorkloadAccessTokenForJWTResponseBody
	GetUserContext() *GetWorkloadAccessTokenForJWTResponseBodyUserContext
	SetWorkloadAccessToken(v string) *GetWorkloadAccessTokenForJWTResponseBody
	GetWorkloadAccessToken() *string
	SetWorkloadIdentityArn(v string) *GetWorkloadAccessTokenForJWTResponseBody
	GetWorkloadIdentityArn() *string
}

type GetWorkloadAccessTokenForJWTResponseBody struct {
	RequestId           *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserContext         *GetWorkloadAccessTokenForJWTResponseBodyUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
	WorkloadAccessToken *string                                              `json:"WorkloadAccessToken,omitempty" xml:"WorkloadAccessToken,omitempty"`
	WorkloadIdentityArn *string                                              `json:"WorkloadIdentityArn,omitempty" xml:"WorkloadIdentityArn,omitempty"`
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

func (s *GetWorkloadAccessTokenForJWTResponseBody) GetUserContext() *GetWorkloadAccessTokenForJWTResponseBodyUserContext {
	return s.UserContext
}

func (s *GetWorkloadAccessTokenForJWTResponseBody) GetWorkloadAccessToken() *string {
	return s.WorkloadAccessToken
}

func (s *GetWorkloadAccessTokenForJWTResponseBody) GetWorkloadIdentityArn() *string {
	return s.WorkloadIdentityArn
}

func (s *GetWorkloadAccessTokenForJWTResponseBody) SetRequestId(v string) *GetWorkloadAccessTokenForJWTResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkloadAccessTokenForJWTResponseBody) SetUserContext(v *GetWorkloadAccessTokenForJWTResponseBodyUserContext) *GetWorkloadAccessTokenForJWTResponseBody {
	s.UserContext = v
	return s
}

func (s *GetWorkloadAccessTokenForJWTResponseBody) SetWorkloadAccessToken(v string) *GetWorkloadAccessTokenForJWTResponseBody {
	s.WorkloadAccessToken = &v
	return s
}

func (s *GetWorkloadAccessTokenForJWTResponseBody) SetWorkloadIdentityArn(v string) *GetWorkloadAccessTokenForJWTResponseBody {
	s.WorkloadIdentityArn = &v
	return s
}

func (s *GetWorkloadAccessTokenForJWTResponseBody) Validate() error {
	if s.UserContext != nil {
		if err := s.UserContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkloadAccessTokenForJWTResponseBodyUserContext struct {
	JwtClaims *GetWorkloadAccessTokenForJWTResponseBodyUserContextJwtClaims `json:"JwtClaims,omitempty" xml:"JwtClaims,omitempty" type:"Struct"`
}

func (s GetWorkloadAccessTokenForJWTResponseBodyUserContext) String() string {
	return dara.Prettify(s)
}

func (s GetWorkloadAccessTokenForJWTResponseBodyUserContext) GoString() string {
	return s.String()
}

func (s *GetWorkloadAccessTokenForJWTResponseBodyUserContext) GetJwtClaims() *GetWorkloadAccessTokenForJWTResponseBodyUserContextJwtClaims {
	return s.JwtClaims
}

func (s *GetWorkloadAccessTokenForJWTResponseBodyUserContext) SetJwtClaims(v *GetWorkloadAccessTokenForJWTResponseBodyUserContextJwtClaims) *GetWorkloadAccessTokenForJWTResponseBodyUserContext {
	s.JwtClaims = v
	return s
}

func (s *GetWorkloadAccessTokenForJWTResponseBodyUserContext) Validate() error {
	if s.JwtClaims != nil {
		if err := s.JwtClaims.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkloadAccessTokenForJWTResponseBodyUserContextJwtClaims struct {
	Audiences *string `json:"Audiences,omitempty" xml:"Audiences,omitempty"`
	Issuer    *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	JTI       *string `json:"JTI,omitempty" xml:"JTI,omitempty"`
	Subject   *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
}

func (s GetWorkloadAccessTokenForJWTResponseBodyUserContextJwtClaims) String() string {
	return dara.Prettify(s)
}

func (s GetWorkloadAccessTokenForJWTResponseBodyUserContextJwtClaims) GoString() string {
	return s.String()
}

func (s *GetWorkloadAccessTokenForJWTResponseBodyUserContextJwtClaims) GetAudiences() *string {
	return s.Audiences
}

func (s *GetWorkloadAccessTokenForJWTResponseBodyUserContextJwtClaims) GetIssuer() *string {
	return s.Issuer
}

func (s *GetWorkloadAccessTokenForJWTResponseBodyUserContextJwtClaims) GetJTI() *string {
	return s.JTI
}

func (s *GetWorkloadAccessTokenForJWTResponseBodyUserContextJwtClaims) GetSubject() *string {
	return s.Subject
}

func (s *GetWorkloadAccessTokenForJWTResponseBodyUserContextJwtClaims) SetAudiences(v string) *GetWorkloadAccessTokenForJWTResponseBodyUserContextJwtClaims {
	s.Audiences = &v
	return s
}

func (s *GetWorkloadAccessTokenForJWTResponseBodyUserContextJwtClaims) SetIssuer(v string) *GetWorkloadAccessTokenForJWTResponseBodyUserContextJwtClaims {
	s.Issuer = &v
	return s
}

func (s *GetWorkloadAccessTokenForJWTResponseBodyUserContextJwtClaims) SetJTI(v string) *GetWorkloadAccessTokenForJWTResponseBodyUserContextJwtClaims {
	s.JTI = &v
	return s
}

func (s *GetWorkloadAccessTokenForJWTResponseBodyUserContextJwtClaims) SetSubject(v string) *GetWorkloadAccessTokenForJWTResponseBodyUserContextJwtClaims {
	s.Subject = &v
	return s
}

func (s *GetWorkloadAccessTokenForJWTResponseBodyUserContextJwtClaims) Validate() error {
	return dara.Validate(s)
}
