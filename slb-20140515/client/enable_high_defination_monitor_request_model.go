// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableHighDefinationMonitorRequest interface {
  dara.Model
  String() string
  GoString() string
  SetLogProject(v string) *EnableHighDefinationMonitorRequest
  GetLogProject() *string 
  SetLogStore(v string) *EnableHighDefinationMonitorRequest
  GetLogStore() *string 
  SetOwnerAccount(v string) *EnableHighDefinationMonitorRequest
  GetOwnerAccount() *string 
  SetOwnerId(v int64) *EnableHighDefinationMonitorRequest
  GetOwnerId() *int64 
  SetRegionId(v string) *EnableHighDefinationMonitorRequest
  GetRegionId() *string 
  SetResourceOwnerAccount(v string) *EnableHighDefinationMonitorRequest
  GetResourceOwnerAccount() *string 
  SetResourceOwnerId(v int64) *EnableHighDefinationMonitorRequest
  GetResourceOwnerId() *int64 
  SetTags(v string) *EnableHighDefinationMonitorRequest
  GetTags() *string 
}

type EnableHighDefinationMonitorRequest struct {
  // The name of the project of Log Service. The name must be 4 to 63 characters in length, and can contain digits and lowercase letters. It must start and end with a digit or a letter.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // my-project
  LogProject *string `json:"LogProject,omitempty" xml:"LogProject,omitempty"`
  // The name of the Logstore of Log Service. The name must be 2 to 64 characters in length and can contain digits, lowercase letters, hyphens (-) and underscores (_). It must start and end with a digit or a letter.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // my-log-store
  LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
  OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  // The ID of the region where the Classic Load Balancer (CLB) instance is deployed.
  // 
  // You can call the [DescribeRegions](https://help.aliyun.com/document_detail/27584.html) operation to query the most recent region list.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
  ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
  // The tags of the logs. The tags must be key-value pairs that are contained in a JSON dictionary.
  // 
  // example:
  // 
  // [{"tagKey":"Key1","tagValue":"Value1"}]
  Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s EnableHighDefinationMonitorRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableHighDefinationMonitorRequest) GoString() string {
  return s.String()
}

func (s *EnableHighDefinationMonitorRequest) GetLogProject() *string  {
  return s.LogProject
}

func (s *EnableHighDefinationMonitorRequest) GetLogStore() *string  {
  return s.LogStore
}

func (s *EnableHighDefinationMonitorRequest) GetOwnerAccount() *string  {
  return s.OwnerAccount
}

func (s *EnableHighDefinationMonitorRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *EnableHighDefinationMonitorRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableHighDefinationMonitorRequest) GetResourceOwnerAccount() *string  {
  return s.ResourceOwnerAccount
}

func (s *EnableHighDefinationMonitorRequest) GetResourceOwnerId() *int64  {
  return s.ResourceOwnerId
}

func (s *EnableHighDefinationMonitorRequest) GetTags() *string  {
  return s.Tags
}

func (s *EnableHighDefinationMonitorRequest) SetLogProject(v string) *EnableHighDefinationMonitorRequest {
  s.LogProject = &v
  return s
}

func (s *EnableHighDefinationMonitorRequest) SetLogStore(v string) *EnableHighDefinationMonitorRequest {
  s.LogStore = &v
  return s
}

func (s *EnableHighDefinationMonitorRequest) SetOwnerAccount(v string) *EnableHighDefinationMonitorRequest {
  s.OwnerAccount = &v
  return s
}

func (s *EnableHighDefinationMonitorRequest) SetOwnerId(v int64) *EnableHighDefinationMonitorRequest {
  s.OwnerId = &v
  return s
}

func (s *EnableHighDefinationMonitorRequest) SetRegionId(v string) *EnableHighDefinationMonitorRequest {
  s.RegionId = &v
  return s
}

func (s *EnableHighDefinationMonitorRequest) SetResourceOwnerAccount(v string) *EnableHighDefinationMonitorRequest {
  s.ResourceOwnerAccount = &v
  return s
}

func (s *EnableHighDefinationMonitorRequest) SetResourceOwnerId(v int64) *EnableHighDefinationMonitorRequest {
  s.ResourceOwnerId = &v
  return s
}

func (s *EnableHighDefinationMonitorRequest) SetTags(v string) *EnableHighDefinationMonitorRequest {
  s.Tags = &v
  return s
}

func (s *EnableHighDefinationMonitorRequest) Validate() error {
  return dara.Validate(s)
}

