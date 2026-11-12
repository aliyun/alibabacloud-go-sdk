// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iE2BSnapshot interface {
  dara.Model
  String() string
  GoString() string
  SetCpuCount(v float32) *E2BSnapshot
  GetCpuCount() *float32 
  SetCreatedTime(v string) *E2BSnapshot
  GetCreatedTime() *string 
  SetExpiredTime(v string) *E2BSnapshot
  GetExpiredTime() *string 
  SetMemoryMB(v int32) *E2BSnapshot
  GetMemoryMB() *int32 
  SetName(v string) *E2BSnapshot
  GetName() *string 
  SetResourceGroupID(v string) *E2BSnapshot
  GetResourceGroupID() *string 
  SetResourceTags(v []*E2BSnapshotResourceTags) *E2BSnapshot
  GetResourceTags() []*E2BSnapshotResourceTags 
  SetSnapshotID(v string) *E2BSnapshot
  GetSnapshotID() *string 
  SetSourceSandboxID(v string) *E2BSnapshot
  GetSourceSandboxID() *string 
  SetStatus(v string) *E2BSnapshot
  GetStatus() *string 
  SetTag(v string) *E2BSnapshot
  GetTag() *string 
  SetTeamID(v string) *E2BSnapshot
  GetTeamID() *string 
  SetTeamName(v string) *E2BSnapshot
  GetTeamName() *string 
  SetUserID(v string) *E2BSnapshot
  GetUserID() *string 
}

type E2BSnapshot struct {
  // The CPU specifications of the snapshot.
  // 
  // example:
  // 
  // 2
  CpuCount *float32 `json:"cpuCount,omitempty" xml:"cpuCount,omitempty"`
  // The creation time of the snapshot.
  // 
  // example:
  // 
  // 1793685726000
  CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
  // The expiration time of the snapshot.
  // 
  // example:
  // 
  // 1796277726000
  ExpiredTime *string `json:"expiredTime,omitempty" xml:"expiredTime,omitempty"`
  // The memory specifications of the snapshot.
  // 
  // example:
  // 
  // 4
  MemoryMB *int32 `json:"memoryMB,omitempty" xml:"memoryMB,omitempty"`
  // The short name of the snapshot.
  // 
  // example:
  // 
  // name
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // The resource group ID.
  // 
  // example:
  // 
  // rg-****
  ResourceGroupID *string `json:"resourceGroupID,omitempty" xml:"resourceGroupID,omitempty"`
  // The resource tags.
  ResourceTags []*E2BSnapshotResourceTags `json:"resourceTags" xml:"resourceTags" type:"Repeated"`
  // The snapshot ID.
  // 
  // example:
  // 
  // 0d66****b374
  SnapshotID *string `json:"snapshotID,omitempty" xml:"snapshotID,omitempty"`
  // The source sandbox ID.
  // 
  // example:
  // 
  // sbx-****
  SourceSandboxID *string `json:"sourceSandboxID,omitempty" xml:"sourceSandboxID,omitempty"`
  // The snapshot status.
  // 
  // example:
  // 
  // Available
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  // The snapshot tag.
  // 
  // example:
  // 
  // tag
  Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
  // The team ID.
  // 
  // example:
  // 
  // 9A66****b892
  TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
  // The team name.
  // 
  // example:
  // 
  // teamName
  TeamName *string `json:"teamName,omitempty" xml:"teamName,omitempty"`
  // The Alibaba Cloud account ID.
  // 
  // example:
  // 
  // 155****
  UserID *string `json:"userID,omitempty" xml:"userID,omitempty"`
}

func (s E2BSnapshot) String() string {
  return dara.Prettify(s)
}

func (s E2BSnapshot) GoString() string {
  return s.String()
}

func (s *E2BSnapshot) GetCpuCount() *float32  {
  return s.CpuCount
}

func (s *E2BSnapshot) GetCreatedTime() *string  {
  return s.CreatedTime
}

func (s *E2BSnapshot) GetExpiredTime() *string  {
  return s.ExpiredTime
}

func (s *E2BSnapshot) GetMemoryMB() *int32  {
  return s.MemoryMB
}

func (s *E2BSnapshot) GetName() *string  {
  return s.Name
}

func (s *E2BSnapshot) GetResourceGroupID() *string  {
  return s.ResourceGroupID
}

func (s *E2BSnapshot) GetResourceTags() []*E2BSnapshotResourceTags  {
  return s.ResourceTags
}

func (s *E2BSnapshot) GetSnapshotID() *string  {
  return s.SnapshotID
}

func (s *E2BSnapshot) GetSourceSandboxID() *string  {
  return s.SourceSandboxID
}

func (s *E2BSnapshot) GetStatus() *string  {
  return s.Status
}

func (s *E2BSnapshot) GetTag() *string  {
  return s.Tag
}

func (s *E2BSnapshot) GetTeamID() *string  {
  return s.TeamID
}

func (s *E2BSnapshot) GetTeamName() *string  {
  return s.TeamName
}

func (s *E2BSnapshot) GetUserID() *string  {
  return s.UserID
}

func (s *E2BSnapshot) SetCpuCount(v float32) *E2BSnapshot {
  s.CpuCount = &v
  return s
}

func (s *E2BSnapshot) SetCreatedTime(v string) *E2BSnapshot {
  s.CreatedTime = &v
  return s
}

func (s *E2BSnapshot) SetExpiredTime(v string) *E2BSnapshot {
  s.ExpiredTime = &v
  return s
}

func (s *E2BSnapshot) SetMemoryMB(v int32) *E2BSnapshot {
  s.MemoryMB = &v
  return s
}

func (s *E2BSnapshot) SetName(v string) *E2BSnapshot {
  s.Name = &v
  return s
}

func (s *E2BSnapshot) SetResourceGroupID(v string) *E2BSnapshot {
  s.ResourceGroupID = &v
  return s
}

func (s *E2BSnapshot) SetResourceTags(v []*E2BSnapshotResourceTags) *E2BSnapshot {
  s.ResourceTags = v
  return s
}

func (s *E2BSnapshot) SetSnapshotID(v string) *E2BSnapshot {
  s.SnapshotID = &v
  return s
}

func (s *E2BSnapshot) SetSourceSandboxID(v string) *E2BSnapshot {
  s.SourceSandboxID = &v
  return s
}

func (s *E2BSnapshot) SetStatus(v string) *E2BSnapshot {
  s.Status = &v
  return s
}

func (s *E2BSnapshot) SetTag(v string) *E2BSnapshot {
  s.Tag = &v
  return s
}

func (s *E2BSnapshot) SetTeamID(v string) *E2BSnapshot {
  s.TeamID = &v
  return s
}

func (s *E2BSnapshot) SetTeamName(v string) *E2BSnapshot {
  s.TeamName = &v
  return s
}

func (s *E2BSnapshot) SetUserID(v string) *E2BSnapshot {
  s.UserID = &v
  return s
}

func (s *E2BSnapshot) Validate() error {
  if s.ResourceTags != nil {
    for _, item := range s.ResourceTags {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type E2BSnapshotResourceTags struct {
  // The tag key.
  // 
  // example:
  // 
  // project
  TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
  // The tag value.
  // 
  // example:
  // 
  // demo
  TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s E2BSnapshotResourceTags) String() string {
  return dara.Prettify(s)
}

func (s E2BSnapshotResourceTags) GoString() string {
  return s.String()
}

func (s *E2BSnapshotResourceTags) GetTagKey() *string  {
  return s.TagKey
}

func (s *E2BSnapshotResourceTags) GetTagValue() *string  {
  return s.TagValue
}

func (s *E2BSnapshotResourceTags) SetTagKey(v string) *E2BSnapshotResourceTags {
  s.TagKey = &v
  return s
}

func (s *E2BSnapshotResourceTags) SetTagValue(v string) *E2BSnapshotResourceTags {
  s.TagValue = &v
  return s
}

func (s *E2BSnapshotResourceTags) Validate() error {
  return dara.Validate(s)
}

