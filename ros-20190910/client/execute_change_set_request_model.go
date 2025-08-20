// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteChangeSetRequest interface {
  dara.Model
  String() string
  GoString() string
  SetChangeSetId(v string) *ExecuteChangeSetRequest
  GetChangeSetId() *string 
  SetClientToken(v string) *ExecuteChangeSetRequest
  GetClientToken() *string 
  SetRegionId(v string) *ExecuteChangeSetRequest
  GetRegionId() *string 
}

type ExecuteChangeSetRequest struct {
  // The ID of the change set.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 1f6521a4-05af-4975-afe9-bc4b45ad****
  ChangeSetId *string `json:"ChangeSetId,omitempty" xml:"ChangeSetId,omitempty"`
  // The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests.
  // 
  // The token can be up to 64 characters in length and can contain letters, digits, hyphens (-), and underscores (_).
  // 
  // For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/134212.html).
  // 
  // example:
  // 
  // 123e4567-e89b-12d3-a456-42665544****
  ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
  // The region ID of the change set. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ExecuteChangeSetRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteChangeSetRequest) GoString() string {
  return s.String()
}

func (s *ExecuteChangeSetRequest) GetChangeSetId() *string  {
  return s.ChangeSetId
}

func (s *ExecuteChangeSetRequest) GetClientToken() *string  {
  return s.ClientToken
}

func (s *ExecuteChangeSetRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *ExecuteChangeSetRequest) SetChangeSetId(v string) *ExecuteChangeSetRequest {
  s.ChangeSetId = &v
  return s
}

func (s *ExecuteChangeSetRequest) SetClientToken(v string) *ExecuteChangeSetRequest {
  s.ClientToken = &v
  return s
}

func (s *ExecuteChangeSetRequest) SetRegionId(v string) *ExecuteChangeSetRequest {
  s.RegionId = &v
  return s
}

func (s *ExecuteChangeSetRequest) Validate() error {
  return dara.Validate(s)
}

