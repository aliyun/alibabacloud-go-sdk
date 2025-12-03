// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableUserHeaders interface {
  dara.Model
  String() string
  GoString() string
  SetCommonHeaders(v map[string]*string) *EnableUserHeaders
  GetCommonHeaders() map[string]*string 
  SetAuthorization(v string) *EnableUserHeaders
  GetAuthorization() *string 
}

type EnableUserHeaders struct {
  CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
  // The authentication information. Format: Bearer ${access_token}. Example: Bearer ATxxxx.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // Bearer xxxx
  Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s EnableUserHeaders) String() string {
  return dara.Prettify(s)
}

func (s EnableUserHeaders) GoString() string {
  return s.String()
}

func (s *EnableUserHeaders) GetCommonHeaders() map[string]*string  {
  return s.CommonHeaders
}

func (s *EnableUserHeaders) GetAuthorization() *string  {
  return s.Authorization
}

func (s *EnableUserHeaders) SetCommonHeaders(v map[string]*string) *EnableUserHeaders {
  s.CommonHeaders = v
  return s
}

func (s *EnableUserHeaders) SetAuthorization(v string) *EnableUserHeaders {
  s.Authorization = &v
  return s
}

func (s *EnableUserHeaders) Validate() error {
  return dara.Validate(s)
}

