// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayAuthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteGatewayAuthResponseBody
	GetCode() *int32
	SetData(v *DeleteGatewayAuthResponseBodyData) *DeleteGatewayAuthResponseBody
	GetData() *DeleteGatewayAuthResponseBodyData
	SetHttpStatusCode(v int32) *DeleteGatewayAuthResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteGatewayAuthResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteGatewayAuthResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteGatewayAuthResponseBody
	GetSuccess() *bool
}

type DeleteGatewayAuthResponseBody struct {
	// example:
	//
	// 200
	Code *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DeleteGatewayAuthResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 316F5F64-F73D-42DC-8632-01E308B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteGatewayAuthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayAuthResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewayAuthResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteGatewayAuthResponseBody) GetData() *DeleteGatewayAuthResponseBodyData {
	return s.Data
}

func (s *DeleteGatewayAuthResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteGatewayAuthResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteGatewayAuthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGatewayAuthResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteGatewayAuthResponseBody) SetCode(v int32) *DeleteGatewayAuthResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGatewayAuthResponseBody) SetData(v *DeleteGatewayAuthResponseBodyData) *DeleteGatewayAuthResponseBody {
	s.Data = v
	return s
}

func (s *DeleteGatewayAuthResponseBody) SetHttpStatusCode(v int32) *DeleteGatewayAuthResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteGatewayAuthResponseBody) SetMessage(v string) *DeleteGatewayAuthResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGatewayAuthResponseBody) SetRequestId(v string) *DeleteGatewayAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGatewayAuthResponseBody) SetSuccess(v bool) *DeleteGatewayAuthResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteGatewayAuthResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteGatewayAuthResponseBodyData struct {
	// example:
	//
	// example-app
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// xxxxx
	ClientSecret *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
	// example:
	//
	// test.com
	CookieDomain *string `json:"CookieDomain,omitempty" xml:"CookieDomain,omitempty"`
	// example:
	//
	// 399
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// example:
	//
	// gw-e2d226bba4b2445c9e29fa7f8216****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// example:
	//
	// 2022-01-07 18:07:57
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2022-01-07 18:07:57
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 120
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// true
	IsWhite *bool `json:"IsWhite,omitempty" xml:"IsWhite,omitempty"`
	// example:
	//
	// https://example.com/auth
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// example:
	//
	// {
	//
	//       "keys": [
	//
	//             {
	//
	//                   "e": "AQAB",
	//
	//                   "kid": "DHFbpoIUqrY8t2zpA2qXfCmr5VO5ZEr4RzHU_-envvQ",
	//
	//                   "kty": "RSA",
	//
	//                   "n": "xAE7eB6qugXyCAG3yhh7pkDkT65pHymX-P7KfIupjf59vsdo91bSP9C8H07pSAGQO1MV_xFj9VswgsCg4R6otmg5PV2He95lZdHtOcU5DXIg_pbhLdKXbi66GlVeK6ABZOUW3WYtnNHD-91gVuoeJT_DwtGGcp4ignkgXfkiEm4sw-4sfb4qdt5oLbyVpmW6x9cfa7vs2WTfURiCrBoUqgBo_-4WTiULmmHSGZHOjzwa8WtrtOQGsAFjIbno85jp6MnGGGZPYZbDAa_b3y5u-YpW7ypZrvD8BgtKVjgtQgZhLAGezMt0ua3DRrWnKqTZ0BJ_EyxOGuHJrLsn00fnMQ"
	//
	//             }
	//
	//       ]
	//
	// }
	Jwks *string `json:"Jwks,omitempty" xml:"Jwks,omitempty"`
	// example:
	//
	// https://daxxxxcn.aliyunidaas.com/
	LoginUrl *string `json:"LoginUrl,omitempty" xml:"LoginUrl,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// https://yourdomain/path
	RedirectUrl *string   `json:"RedirectUrl,omitempty" xml:"RedirectUrl,omitempty"`
	ScopesList  []*string `json:"ScopesList,omitempty" xml:"ScopesList,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// Authorization
	TokenName *string `json:"TokenName,omitempty" xml:"TokenName,omitempty"`
	// example:
	//
	// Bearer
	TokenNamePrefix *string `json:"TokenNamePrefix,omitempty" xml:"TokenNamePrefix,omitempty"`
	// example:
	//
	// true
	TokenPass *bool `json:"TokenPass,omitempty" xml:"TokenPass,omitempty"`
	// example:
	//
	// HEADER
	TokenPosition *string `json:"TokenPosition,omitempty" xml:"TokenPosition,omitempty"`
	// example:
	//
	// JWT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DeleteGatewayAuthResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayAuthResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteGatewayAuthResponseBodyData) GetClientId() *string {
	return s.ClientId
}

func (s *DeleteGatewayAuthResponseBodyData) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *DeleteGatewayAuthResponseBodyData) GetCookieDomain() *string {
	return s.CookieDomain
}

func (s *DeleteGatewayAuthResponseBodyData) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *DeleteGatewayAuthResponseBodyData) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *DeleteGatewayAuthResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DeleteGatewayAuthResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DeleteGatewayAuthResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *DeleteGatewayAuthResponseBodyData) GetIsWhite() *bool {
	return s.IsWhite
}

func (s *DeleteGatewayAuthResponseBodyData) GetIssuer() *string {
	return s.Issuer
}

func (s *DeleteGatewayAuthResponseBodyData) GetJwks() *string {
	return s.Jwks
}

func (s *DeleteGatewayAuthResponseBodyData) GetLoginUrl() *string {
	return s.LoginUrl
}

func (s *DeleteGatewayAuthResponseBodyData) GetName() *string {
	return s.Name
}

func (s *DeleteGatewayAuthResponseBodyData) GetRedirectUrl() *string {
	return s.RedirectUrl
}

func (s *DeleteGatewayAuthResponseBodyData) GetScopesList() []*string {
	return s.ScopesList
}

func (s *DeleteGatewayAuthResponseBodyData) GetStatus() *bool {
	return s.Status
}

func (s *DeleteGatewayAuthResponseBodyData) GetTokenName() *string {
	return s.TokenName
}

func (s *DeleteGatewayAuthResponseBodyData) GetTokenNamePrefix() *string {
	return s.TokenNamePrefix
}

func (s *DeleteGatewayAuthResponseBodyData) GetTokenPass() *bool {
	return s.TokenPass
}

func (s *DeleteGatewayAuthResponseBodyData) GetTokenPosition() *string {
	return s.TokenPosition
}

func (s *DeleteGatewayAuthResponseBodyData) GetType() *string {
	return s.Type
}

func (s *DeleteGatewayAuthResponseBodyData) SetClientId(v string) *DeleteGatewayAuthResponseBodyData {
	s.ClientId = &v
	return s
}

func (s *DeleteGatewayAuthResponseBodyData) SetClientSecret(v string) *DeleteGatewayAuthResponseBodyData {
	s.ClientSecret = &v
	return s
}

func (s *DeleteGatewayAuthResponseBodyData) SetCookieDomain(v string) *DeleteGatewayAuthResponseBodyData {
	s.CookieDomain = &v
	return s
}

func (s *DeleteGatewayAuthResponseBodyData) SetGatewayId(v int64) *DeleteGatewayAuthResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *DeleteGatewayAuthResponseBodyData) SetGatewayUniqueId(v string) *DeleteGatewayAuthResponseBodyData {
	s.GatewayUniqueId = &v
	return s
}

func (s *DeleteGatewayAuthResponseBodyData) SetGmtCreate(v string) *DeleteGatewayAuthResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *DeleteGatewayAuthResponseBodyData) SetGmtModified(v string) *DeleteGatewayAuthResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *DeleteGatewayAuthResponseBodyData) SetId(v int64) *DeleteGatewayAuthResponseBodyData {
	s.Id = &v
	return s
}

func (s *DeleteGatewayAuthResponseBodyData) SetIsWhite(v bool) *DeleteGatewayAuthResponseBodyData {
	s.IsWhite = &v
	return s
}

func (s *DeleteGatewayAuthResponseBodyData) SetIssuer(v string) *DeleteGatewayAuthResponseBodyData {
	s.Issuer = &v
	return s
}

func (s *DeleteGatewayAuthResponseBodyData) SetJwks(v string) *DeleteGatewayAuthResponseBodyData {
	s.Jwks = &v
	return s
}

func (s *DeleteGatewayAuthResponseBodyData) SetLoginUrl(v string) *DeleteGatewayAuthResponseBodyData {
	s.LoginUrl = &v
	return s
}

func (s *DeleteGatewayAuthResponseBodyData) SetName(v string) *DeleteGatewayAuthResponseBodyData {
	s.Name = &v
	return s
}

func (s *DeleteGatewayAuthResponseBodyData) SetRedirectUrl(v string) *DeleteGatewayAuthResponseBodyData {
	s.RedirectUrl = &v
	return s
}

func (s *DeleteGatewayAuthResponseBodyData) SetScopesList(v []*string) *DeleteGatewayAuthResponseBodyData {
	s.ScopesList = v
	return s
}

func (s *DeleteGatewayAuthResponseBodyData) SetStatus(v bool) *DeleteGatewayAuthResponseBodyData {
	s.Status = &v
	return s
}

func (s *DeleteGatewayAuthResponseBodyData) SetTokenName(v string) *DeleteGatewayAuthResponseBodyData {
	s.TokenName = &v
	return s
}

func (s *DeleteGatewayAuthResponseBodyData) SetTokenNamePrefix(v string) *DeleteGatewayAuthResponseBodyData {
	s.TokenNamePrefix = &v
	return s
}

func (s *DeleteGatewayAuthResponseBodyData) SetTokenPass(v bool) *DeleteGatewayAuthResponseBodyData {
	s.TokenPass = &v
	return s
}

func (s *DeleteGatewayAuthResponseBodyData) SetTokenPosition(v string) *DeleteGatewayAuthResponseBodyData {
	s.TokenPosition = &v
	return s
}

func (s *DeleteGatewayAuthResponseBodyData) SetType(v string) *DeleteGatewayAuthResponseBodyData {
	s.Type = &v
	return s
}

func (s *DeleteGatewayAuthResponseBodyData) Validate() error {
	return dara.Validate(s)
}
