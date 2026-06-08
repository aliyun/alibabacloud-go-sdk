// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableProcessDefinitionRequest interface {
  dara.Model
  String() string
  GoString() string
  SetClientToken(v string) *EnableProcessDefinitionRequest
  GetClientToken() *string 
  SetId(v string) *EnableProcessDefinitionRequest
  GetId() *string 
}

type EnableProcessDefinitionRequest struct {
  // example:
  // 
  // 1AFAE64E-D1BE-432B-A9****
  ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
  // example:
  // 
  // f0d6d578-a305-40ac-ba1e-0a09f64cbc69
  Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s EnableProcessDefinitionRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableProcessDefinitionRequest) GoString() string {
  return s.String()
}

func (s *EnableProcessDefinitionRequest) GetClientToken() *string  {
  return s.ClientToken
}

func (s *EnableProcessDefinitionRequest) GetId() *string  {
  return s.Id
}

func (s *EnableProcessDefinitionRequest) SetClientToken(v string) *EnableProcessDefinitionRequest {
  s.ClientToken = &v
  return s
}

func (s *EnableProcessDefinitionRequest) SetId(v string) *EnableProcessDefinitionRequest {
  s.Id = &v
  return s
}

func (s *EnableProcessDefinitionRequest) Validate() error {
  return dara.Validate(s)
}

