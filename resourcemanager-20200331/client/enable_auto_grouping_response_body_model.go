// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableAutoGroupingResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableAutoGroupingResponseBody
  GetRequestId() *string 
}

type EnableAutoGroupingResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // DF5D5C52-7BD0-40E7-94C6-23A1505038A2
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableAutoGroupingResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableAutoGroupingResponseBody) GoString() string {
  return s.String()
}

func (s *EnableAutoGroupingResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableAutoGroupingResponseBody) SetRequestId(v string) *EnableAutoGroupingResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableAutoGroupingResponseBody) Validate() error {
  return dara.Validate(s)
}

