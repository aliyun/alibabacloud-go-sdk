// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDeletionProtectionRequest interface {
  dara.Model
  String() string
  GoString() string
  SetClientToken(v string) *EnableDeletionProtectionRequest
  GetClientToken() *string 
  SetDryRun(v bool) *EnableDeletionProtectionRequest
  GetDryRun() *bool 
  SetResourceId(v string) *EnableDeletionProtectionRequest
  GetResourceId() *string 
}

type EnableDeletionProtectionRequest struct {
  // The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
  // 
  // example:
  // 
  // 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
  ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
  // Specifies whether to perform only a dry run, without performing the actual request. Valid values:
  // 
  // 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
  // 
  // 	- **false**: (default): performs a dry run and performs the actual request. If the request passes the dry run, a `2xx HTTP` status code is returned and the operation is performed.
  // 
  // example:
  // 
  // true
  DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
  // The Application Load Balancer (ALB) instance ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // re-atstuj3rtop****
  ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s EnableDeletionProtectionRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableDeletionProtectionRequest) GoString() string {
  return s.String()
}

func (s *EnableDeletionProtectionRequest) GetClientToken() *string  {
  return s.ClientToken
}

func (s *EnableDeletionProtectionRequest) GetDryRun() *bool  {
  return s.DryRun
}

func (s *EnableDeletionProtectionRequest) GetResourceId() *string  {
  return s.ResourceId
}

func (s *EnableDeletionProtectionRequest) SetClientToken(v string) *EnableDeletionProtectionRequest {
  s.ClientToken = &v
  return s
}

func (s *EnableDeletionProtectionRequest) SetDryRun(v bool) *EnableDeletionProtectionRequest {
  s.DryRun = &v
  return s
}

func (s *EnableDeletionProtectionRequest) SetResourceId(v string) *EnableDeletionProtectionRequest {
  s.ResourceId = &v
  return s
}

func (s *EnableDeletionProtectionRequest) Validate() error {
  return dara.Validate(s)
}

