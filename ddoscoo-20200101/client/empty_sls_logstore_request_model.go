// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEmptySlsLogstoreRequest interface {
  dara.Model
  String() string
  GoString() string
  SetResourceGroupId(v string) *EmptySlsLogstoreRequest
  GetResourceGroupId() *string 
}

type EmptySlsLogstoreRequest struct {
  // The resource group ID to which the Anti-DDoS Proxy instance belongs in Resource Management. The default value is empty, which indicates that the instance belongs to the default resource group.
  // 
  // example:
  // 
  // default
  ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s EmptySlsLogstoreRequest) String() string {
  return dara.Prettify(s)
}

func (s EmptySlsLogstoreRequest) GoString() string {
  return s.String()
}

func (s *EmptySlsLogstoreRequest) GetResourceGroupId() *string  {
  return s.ResourceGroupId
}

func (s *EmptySlsLogstoreRequest) SetResourceGroupId(v string) *EmptySlsLogstoreRequest {
  s.ResourceGroupId = &v
  return s
}

func (s *EmptySlsLogstoreRequest) Validate() error {
  return dara.Validate(s)
}

