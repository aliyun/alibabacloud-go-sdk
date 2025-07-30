// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEncryptPasswordRequest interface {
  dara.Model
  String() string
  GoString() string
  SetClientId(v string) *EncryptPasswordRequest
  GetClientId() *string 
  SetDirectoryId(v string) *EncryptPasswordRequest
  GetDirectoryId() *string 
  SetLoginToken(v string) *EncryptPasswordRequest
  GetLoginToken() *string 
  SetOfficeSiteId(v string) *EncryptPasswordRequest
  GetOfficeSiteId() *string 
  SetPassword(v string) *EncryptPasswordRequest
  GetPassword() *string 
  SetRegionId(v string) *EncryptPasswordRequest
  GetRegionId() *string 
  SetSessionId(v string) *EncryptPasswordRequest
  GetSessionId() *string 
}

type EncryptPasswordRequest struct {
  // The ID of the client. The system generates a unique ID for each client.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 1d40776f-e9cb-4e2b-a8da-308d10e8****
  ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
  // The directory ID.
  // 
  // example:
  // 
  // cn-beijing+dir-131196****
  DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
  // The logon token.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // v1b16dcff3ab21a6c5ec01652238375511cff5a1db59fd4dc49afb37e2ea7a626af6f38109fd0498b6abd9de1af7743****
  LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
  // The office network ID.
  // 
  // example:
  // 
  // cn-beijing+dir-131196****
  OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
  // The password that you want to encrypt.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // Ab123456
  Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
  // The region ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-beijing
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The session ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // c78e2e52-23d9-4401-a648-e67ac6ff****
  SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s EncryptPasswordRequest) String() string {
  return dara.Prettify(s)
}

func (s EncryptPasswordRequest) GoString() string {
  return s.String()
}

func (s *EncryptPasswordRequest) GetClientId() *string  {
  return s.ClientId
}

func (s *EncryptPasswordRequest) GetDirectoryId() *string  {
  return s.DirectoryId
}

func (s *EncryptPasswordRequest) GetLoginToken() *string  {
  return s.LoginToken
}

func (s *EncryptPasswordRequest) GetOfficeSiteId() *string  {
  return s.OfficeSiteId
}

func (s *EncryptPasswordRequest) GetPassword() *string  {
  return s.Password
}

func (s *EncryptPasswordRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EncryptPasswordRequest) GetSessionId() *string  {
  return s.SessionId
}

func (s *EncryptPasswordRequest) SetClientId(v string) *EncryptPasswordRequest {
  s.ClientId = &v
  return s
}

func (s *EncryptPasswordRequest) SetDirectoryId(v string) *EncryptPasswordRequest {
  s.DirectoryId = &v
  return s
}

func (s *EncryptPasswordRequest) SetLoginToken(v string) *EncryptPasswordRequest {
  s.LoginToken = &v
  return s
}

func (s *EncryptPasswordRequest) SetOfficeSiteId(v string) *EncryptPasswordRequest {
  s.OfficeSiteId = &v
  return s
}

func (s *EncryptPasswordRequest) SetPassword(v string) *EncryptPasswordRequest {
  s.Password = &v
  return s
}

func (s *EncryptPasswordRequest) SetRegionId(v string) *EncryptPasswordRequest {
  s.RegionId = &v
  return s
}

func (s *EncryptPasswordRequest) SetSessionId(v string) *EncryptPasswordRequest {
  s.SessionId = &v
  return s
}

func (s *EncryptPasswordRequest) Validate() error {
  return dara.Validate(s)
}

