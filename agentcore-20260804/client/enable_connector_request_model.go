// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableConnectorRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBody(v *EnableConnectorRequestBody) *EnableConnectorRequest
  GetBody() *EnableConnectorRequestBody 
}

type EnableConnectorRequest struct {
  // The enable request body.
  // 
  // This parameter is required.
  Body *EnableConnectorRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s EnableConnectorRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableConnectorRequest) GoString() string {
  return s.String()
}

func (s *EnableConnectorRequest) GetBody() *EnableConnectorRequestBody  {
  return s.Body
}

func (s *EnableConnectorRequest) SetBody(v *EnableConnectorRequestBody) *EnableConnectorRequest {
  s.Body = v
  return s
}

func (s *EnableConnectorRequest) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EnableConnectorRequestBody struct {
  // The Connector configuration JSON string. Set site to global or cn. apiKey is required. serviceAccountKeys must contain at least one named service account key. organizationId is optional.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // {"site":"global","organizationId":"org-xxxx","apiKey":"ak-xxxx","serviceAccountKeys":[{"name":"default","serviceAccountKey":"sk-xxxx"}]}
  Metadata *string `json:"metadata,omitempty" xml:"metadata,omitempty"`
}

func (s EnableConnectorRequestBody) String() string {
  return dara.Prettify(s)
}

func (s EnableConnectorRequestBody) GoString() string {
  return s.String()
}

func (s *EnableConnectorRequestBody) GetMetadata() *string  {
  return s.Metadata
}

func (s *EnableConnectorRequestBody) SetMetadata(v string) *EnableConnectorRequestBody {
  s.Metadata = &v
  return s
}

func (s *EnableConnectorRequestBody) Validate() error {
  return dara.Validate(s)
}

