// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExpireLoginTokenRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEndUserId(v string) *ExpireLoginTokenRequest
  GetEndUserId() *string 
  SetLoginToken(v string) *ExpireLoginTokenRequest
  GetLoginToken() *string 
  SetOfficeSiteId(v string) *ExpireLoginTokenRequest
  GetOfficeSiteId() *string 
  SetSessionId(v string) *ExpireLoginTokenRequest
  GetSessionId() *string 
}

type ExpireLoginTokenRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // testuser
  EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // v185fdd7f6d39fa7861981639366085772e150a390a5bb7b43c4e62440d94fc392b945770e1596cebe90085ce0af4d****
  LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
  // example:
  // 
  // cn-beijing+dir-172301****
  OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // a863f4c3-2f1d-4971-8cf7-e2b92ae9****
  SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s ExpireLoginTokenRequest) String() string {
  return dara.Prettify(s)
}

func (s ExpireLoginTokenRequest) GoString() string {
  return s.String()
}

func (s *ExpireLoginTokenRequest) GetEndUserId() *string  {
  return s.EndUserId
}

func (s *ExpireLoginTokenRequest) GetLoginToken() *string  {
  return s.LoginToken
}

func (s *ExpireLoginTokenRequest) GetOfficeSiteId() *string  {
  return s.OfficeSiteId
}

func (s *ExpireLoginTokenRequest) GetSessionId() *string  {
  return s.SessionId
}

func (s *ExpireLoginTokenRequest) SetEndUserId(v string) *ExpireLoginTokenRequest {
  s.EndUserId = &v
  return s
}

func (s *ExpireLoginTokenRequest) SetLoginToken(v string) *ExpireLoginTokenRequest {
  s.LoginToken = &v
  return s
}

func (s *ExpireLoginTokenRequest) SetOfficeSiteId(v string) *ExpireLoginTokenRequest {
  s.OfficeSiteId = &v
  return s
}

func (s *ExpireLoginTokenRequest) SetSessionId(v string) *ExpireLoginTokenRequest {
  s.SessionId = &v
  return s
}

func (s *ExpireLoginTokenRequest) Validate() error {
  return dara.Validate(s)
}

