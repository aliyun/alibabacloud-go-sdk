// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnablePolicyTypeResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnablePolicyTypeResponseBody
  GetRequestId() *string 
}

type EnablePolicyTypeResponseBody struct {
  // example:
  // 
  // 6E27F22C-EDA3-132E-A53F-77DE3BC2343D
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnablePolicyTypeResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnablePolicyTypeResponseBody) GoString() string {
  return s.String()
}

func (s *EnablePolicyTypeResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnablePolicyTypeResponseBody) SetRequestId(v string) *EnablePolicyTypeResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnablePolicyTypeResponseBody) Validate() error {
  return dara.Validate(s)
}

