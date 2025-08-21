// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableWebAccessLogConfigResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableWebAccessLogConfigResponseBody
  GetRequestId() *string 
}

type EnableWebAccessLogConfigResponseBody struct {
  // The ID of the request.
  // 
  // example:
  // 
  // CF33B4C3-196E-4015-AADD-5CAD00057B80
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableWebAccessLogConfigResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableWebAccessLogConfigResponseBody) GoString() string {
  return s.String()
}

func (s *EnableWebAccessLogConfigResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableWebAccessLogConfigResponseBody) SetRequestId(v string) *EnableWebAccessLogConfigResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableWebAccessLogConfigResponseBody) Validate() error {
  return dara.Validate(s)
}

