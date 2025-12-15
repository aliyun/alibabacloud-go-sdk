// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableResourceGroupNotificationResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableResourceGroupNotificationResponseBody
  GetRequestId() *string 
}

type EnableResourceGroupNotificationResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // F7701451-340B-5CB3-AEA7-7D831F7F38C0
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableResourceGroupNotificationResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableResourceGroupNotificationResponseBody) GoString() string {
  return s.String()
}

func (s *EnableResourceGroupNotificationResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableResourceGroupNotificationResponseBody) SetRequestId(v string) *EnableResourceGroupNotificationResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableResourceGroupNotificationResponseBody) Validate() error {
  return dara.Validate(s)
}

