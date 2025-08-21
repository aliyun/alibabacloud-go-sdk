// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSceneDefensePolicyResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableSceneDefensePolicyResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnableSceneDefensePolicyResponseBody
  GetSuccess() *bool 
}

type EnableSceneDefensePolicyResponseBody struct {
  // The ID of the request.
  // 
  // example:
  // 
  // F65DF043-E0EB-4796-9467-23DDCDF92C1D
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the request was successful. Valid values:
  // 
  // 	- **true**: yes
  // 
  // 	- **false**: no
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableSceneDefensePolicyResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableSceneDefensePolicyResponseBody) GoString() string {
  return s.String()
}

func (s *EnableSceneDefensePolicyResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableSceneDefensePolicyResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableSceneDefensePolicyResponseBody) SetRequestId(v string) *EnableSceneDefensePolicyResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableSceneDefensePolicyResponseBody) SetSuccess(v bool) *EnableSceneDefensePolicyResponseBody {
  s.Success = &v
  return s
}

func (s *EnableSceneDefensePolicyResponseBody) Validate() error {
  return dara.Validate(s)
}

