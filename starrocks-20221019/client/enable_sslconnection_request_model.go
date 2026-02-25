// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSSLConnectionRequest interface {
  dara.Model
  String() string
  GoString() string
  SetCustomSSLCertificate(v string) *EnableSSLConnectionRequest
  GetCustomSSLCertificate() *string 
  SetEnableCustom(v bool) *EnableSSLConnectionRequest
  GetEnableCustom() *bool 
  SetInstanceId(v string) *EnableSSLConnectionRequest
  GetInstanceId() *string 
  SetRenewal(v bool) *EnableSSLConnectionRequest
  GetRenewal() *bool 
  SetSslKeyPassword(v string) *EnableSSLConnectionRequest
  GetSslKeyPassword() *string 
  SetSslKeystorePassword(v string) *EnableSSLConnectionRequest
  GetSslKeystorePassword() *string 
}

type EnableSSLConnectionRequest struct {
  // example:
  // 
  // MIIP0wIBAzCCD4wGCSqGSIb3DQEHAaCCD30Egg95MIIPdTCCBbwGCSqGSIb3DQEHAaCCBa0EggWpMIIFpTCCBaEGCyqGSIb3DQEMCgECoIIFQDCCBTwwZgYJKoZIhvcNAQUNMFkwOAYJKoZIhvcNAQUMMCsEFHkQJTDaeFabOsRTB4Q7hgW6if7hAgInEAIBIDAMBggqhkiG9w0CCQU
  CustomSSLCertificate *string `json:"CustomSSLCertificate,omitempty" xml:"CustomSSLCertificate,omitempty"`
  // example:
  // 
  // true
  EnableCustom *bool `json:"EnableCustom,omitempty" xml:"EnableCustom,omitempty"`
  // example:
  // 
  // c-0104730e9de40215
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  // example:
  // 
  // true
  Renewal *bool `json:"Renewal,omitempty" xml:"Renewal,omitempty"`
  // example:
  // 
  // KoZ13vcNAQ
  SslKeyPassword *string `json:"SslKeyPassword,omitempty" xml:"SslKeyPassword,omitempty"`
  // example:
  // 
  // 21esSd9Ao
  SslKeystorePassword *string `json:"SslKeystorePassword,omitempty" xml:"SslKeystorePassword,omitempty"`
}

func (s EnableSSLConnectionRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableSSLConnectionRequest) GoString() string {
  return s.String()
}

func (s *EnableSSLConnectionRequest) GetCustomSSLCertificate() *string  {
  return s.CustomSSLCertificate
}

func (s *EnableSSLConnectionRequest) GetEnableCustom() *bool  {
  return s.EnableCustom
}

func (s *EnableSSLConnectionRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableSSLConnectionRequest) GetRenewal() *bool  {
  return s.Renewal
}

func (s *EnableSSLConnectionRequest) GetSslKeyPassword() *string  {
  return s.SslKeyPassword
}

func (s *EnableSSLConnectionRequest) GetSslKeystorePassword() *string  {
  return s.SslKeystorePassword
}

func (s *EnableSSLConnectionRequest) SetCustomSSLCertificate(v string) *EnableSSLConnectionRequest {
  s.CustomSSLCertificate = &v
  return s
}

func (s *EnableSSLConnectionRequest) SetEnableCustom(v bool) *EnableSSLConnectionRequest {
  s.EnableCustom = &v
  return s
}

func (s *EnableSSLConnectionRequest) SetInstanceId(v string) *EnableSSLConnectionRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableSSLConnectionRequest) SetRenewal(v bool) *EnableSSLConnectionRequest {
  s.Renewal = &v
  return s
}

func (s *EnableSSLConnectionRequest) SetSslKeyPassword(v string) *EnableSSLConnectionRequest {
  s.SslKeyPassword = &v
  return s
}

func (s *EnableSSLConnectionRequest) SetSslKeystorePassword(v string) *EnableSSLConnectionRequest {
  s.SslKeystorePassword = &v
  return s
}

func (s *EnableSSLConnectionRequest) Validate() error {
  return dara.Validate(s)
}

