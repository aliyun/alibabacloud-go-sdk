// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseUninvitedAdminInviteJoinEnterpriseResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponse
  GetStatusCode() *int32 
  SetBody(v *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponse
  GetBody() *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody 
}

type EnterpriseUninvitedAdminInviteJoinEnterpriseResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseUninvitedAdminInviteJoinEnterpriseResponse) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseUninvitedAdminInviteJoinEnterpriseResponse) GoString() string {
  return s.String()
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponse) GetBody() *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody  {
  return s.Body
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponse) SetHeaders(v map[string]*string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponse {
  s.Headers = v
  return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponse) SetStatusCode(v int32) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponse {
  s.StatusCode = &v
  return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponse) SetBody(v *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponse {
  s.Body = v
  return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

