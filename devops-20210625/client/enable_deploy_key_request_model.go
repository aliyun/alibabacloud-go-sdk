// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDeployKeyRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAccessToken(v string) *EnableDeployKeyRequest
  GetAccessToken() *string 
  SetOrganizationId(v string) *EnableDeployKeyRequest
  GetOrganizationId() *string 
}

type EnableDeployKeyRequest struct {
  // example:
  // 
  // f0b1e61db5961df5975a93f9129d2513
  AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 60de7a6852743a5162b5f957
  OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s EnableDeployKeyRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableDeployKeyRequest) GoString() string {
  return s.String()
}

func (s *EnableDeployKeyRequest) GetAccessToken() *string  {
  return s.AccessToken
}

func (s *EnableDeployKeyRequest) GetOrganizationId() *string  {
  return s.OrganizationId
}

func (s *EnableDeployKeyRequest) SetAccessToken(v string) *EnableDeployKeyRequest {
  s.AccessToken = &v
  return s
}

func (s *EnableDeployKeyRequest) SetOrganizationId(v string) *EnableDeployKeyRequest {
  s.OrganizationId = &v
  return s
}

func (s *EnableDeployKeyRequest) Validate() error {
  return dara.Validate(s)
}

