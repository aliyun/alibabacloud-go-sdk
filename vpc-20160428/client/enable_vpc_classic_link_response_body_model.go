// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableVpcClassicLinkResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableVpcClassicLinkResponseBody
  GetRequestId() *string 
}

type EnableVpcClassicLinkResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // 54B48E3D-DF70-471B-AA93-08E683A1B45
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableVpcClassicLinkResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableVpcClassicLinkResponseBody) GoString() string {
  return s.String()
}

func (s *EnableVpcClassicLinkResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableVpcClassicLinkResponseBody) SetRequestId(v string) *EnableVpcClassicLinkResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableVpcClassicLinkResponseBody) Validate() error {
  return dara.Validate(s)
}

