// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableConnectorShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBodyShrink(v string) *EnableConnectorShrinkRequest
  GetBodyShrink() *string 
}

type EnableConnectorShrinkRequest struct {
  // The enable request body.
  // 
  // This parameter is required.
  BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableConnectorShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableConnectorShrinkRequest) GoString() string {
  return s.String()
}

func (s *EnableConnectorShrinkRequest) GetBodyShrink() *string  {
  return s.BodyShrink
}

func (s *EnableConnectorShrinkRequest) SetBodyShrink(v string) *EnableConnectorShrinkRequest {
  s.BodyShrink = &v
  return s
}

func (s *EnableConnectorShrinkRequest) Validate() error {
  return dara.Validate(s)
}

