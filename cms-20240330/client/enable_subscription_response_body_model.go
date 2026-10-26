// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSubscriptionResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableSubscriptionResponseBody
  GetRequestId() *string 
  SetSubscriptionId(v string) *EnableSubscriptionResponseBody
  GetSubscriptionId() *string 
}

type EnableSubscriptionResponseBody struct {
  // Id of the request
  // 
  // example:
  // 
  // 264C3E89-XXXX-XXXX-XXXX-CE9C2196C7DC
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // example:
  // 
  // 123123123123
  SubscriptionId *string `json:"subscriptionId,omitempty" xml:"subscriptionId,omitempty"`
}

func (s EnableSubscriptionResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableSubscriptionResponseBody) GoString() string {
  return s.String()
}

func (s *EnableSubscriptionResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableSubscriptionResponseBody) GetSubscriptionId() *string  {
  return s.SubscriptionId
}

func (s *EnableSubscriptionResponseBody) SetRequestId(v string) *EnableSubscriptionResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableSubscriptionResponseBody) SetSubscriptionId(v string) *EnableSubscriptionResponseBody {
  s.SubscriptionId = &v
  return s
}

func (s *EnableSubscriptionResponseBody) Validate() error {
  return dara.Validate(s)
}

