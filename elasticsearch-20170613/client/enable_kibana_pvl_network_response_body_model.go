// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableKibanaPvlNetworkResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableKibanaPvlNetworkResponseBody
  GetRequestId() *string 
  SetResult(v bool) *EnableKibanaPvlNetworkResponseBody
  GetResult() *bool 
}

type EnableKibanaPvlNetworkResponseBody struct {
  // example:
  // 
  // 0DC92CFE-62AF-51AF-9D5B-F1078D7C451E
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // example:
  // 
  // true
  Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s EnableKibanaPvlNetworkResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableKibanaPvlNetworkResponseBody) GoString() string {
  return s.String()
}

func (s *EnableKibanaPvlNetworkResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableKibanaPvlNetworkResponseBody) GetResult() *bool  {
  return s.Result
}

func (s *EnableKibanaPvlNetworkResponseBody) SetRequestId(v string) *EnableKibanaPvlNetworkResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableKibanaPvlNetworkResponseBody) SetResult(v bool) *EnableKibanaPvlNetworkResponseBody {
  s.Result = &v
  return s
}

func (s *EnableKibanaPvlNetworkResponseBody) Validate() error {
  return dara.Validate(s)
}

