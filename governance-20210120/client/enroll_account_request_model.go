// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnrollAccountRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAccountNamePrefix(v string) *EnrollAccountRequest
  GetAccountNamePrefix() *string 
  SetAccountUid(v int64) *EnrollAccountRequest
  GetAccountUid() *int64 
  SetBaselineId(v string) *EnrollAccountRequest
  GetBaselineId() *string 
  SetBaselineItems(v []*EnrollAccountRequestBaselineItems) *EnrollAccountRequest
  GetBaselineItems() []*EnrollAccountRequestBaselineItems 
  SetDisplayName(v string) *EnrollAccountRequest
  GetDisplayName() *string 
  SetFolderId(v string) *EnrollAccountRequest
  GetFolderId() *string 
  SetPayerAccountUid(v int64) *EnrollAccountRequest
  GetPayerAccountUid() *int64 
  SetRegionId(v string) *EnrollAccountRequest
  GetRegionId() *string 
  SetResellAccountType(v string) *EnrollAccountRequest
  GetResellAccountType() *string 
  SetTag(v []*EnrollAccountRequestTag) *EnrollAccountRequest
  GetTag() []*EnrollAccountRequestTag 
}

type EnrollAccountRequest struct {
  // The prefix for the account name.
  // 
  // - If you are creating a new resource account, this parameter is required.
  // 
  // - If you are enrolling an existing account, this parameter is not required.
  // 
  // example:
  // 
  // test-account
  AccountNamePrefix *string `json:"AccountNamePrefix,omitempty" xml:"AccountNamePrefix,omitempty"`
  // The ID of the account to enroll.
  // 
  // - If you are creating a new resource account, this parameter is not required.
  // 
  // - If you are enrolling an existing account, this parameter is required.
  // 
  // example:
  // 
  // 12868156179****
  AccountUid *int64 `json:"AccountUid,omitempty" xml:"AccountUid,omitempty"`
  // The ID of the baseline. If you leave this parameter empty, the default baseline is used.
  // 
  // example:
  // 
  // afb-bp1durvn3lgqe28v****
  BaselineId *string `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
  // The baseline items.
  // 
  // If you specify this parameter, the baseline item configurations are merged with the configurations of the baseline specified by `BaselineId`. For duplicate baseline items, the configurations in this parameter take precedence. We recommend that you leave this parameter empty and use `BaselineId` to apply baseline configurations.
  BaselineItems []*EnrollAccountRequestBaselineItems `json:"BaselineItems,omitempty" xml:"BaselineItems,omitempty" type:"Repeated"`
  // The display name of the account.
  // 
  // - If you are creating a new resource account, this parameter is required.
  // 
  // - If you are enrolling an existing account, this parameter is not required.
  // 
  // example:
  // 
  // test-account
  DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
  // The ID of the parent folder.
  // 
  // - If you are creating a new resource account and do not specify this parameter, the account is created in the Root folder.
  // 
  // - If you are enrolling an existing account, this parameter is not required.
  // 
  // example:
  // 
  // fd-5ESoku****
  FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
  // The ID of the billing account.
  // 
  // - If you are creating a new resource account and do not specify this parameter, the self-pay settlement method is used.
  // 
  // - If you are enrolling an existing account, this parameter is not required.
  // 
  // example:
  // 
  // 19534534552****
  PayerAccountUid *int64 `json:"PayerAccountUid,omitempty" xml:"PayerAccountUid,omitempty"`
  // The region ID.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The identity type of the member. Valid values:
  // 
  // - resell (default): The member is a reseller account. A reseller relationship is automatically established between the member and the reseller. The management account of the resource directory is used as the billing account of the member.
  // 
  // - non_resell: The member is a non-reseller account. The member is not associated with a reseller and can directly purchase Alibaba Cloud resources. The member is used as its own billing account.
  // 
  // > This parameter is available only for resellers at the international site (alibabacloud.com).
  // 
  // example:
  // 
  // resell
  ResellAccountType *string `json:"ResellAccountType,omitempty" xml:"ResellAccountType,omitempty"`
  // The tags. You can specify up to 20 tags.
  Tag []*EnrollAccountRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s EnrollAccountRequest) String() string {
  return dara.Prettify(s)
}

func (s EnrollAccountRequest) GoString() string {
  return s.String()
}

func (s *EnrollAccountRequest) GetAccountNamePrefix() *string  {
  return s.AccountNamePrefix
}

func (s *EnrollAccountRequest) GetAccountUid() *int64  {
  return s.AccountUid
}

func (s *EnrollAccountRequest) GetBaselineId() *string  {
  return s.BaselineId
}

func (s *EnrollAccountRequest) GetBaselineItems() []*EnrollAccountRequestBaselineItems  {
  return s.BaselineItems
}

func (s *EnrollAccountRequest) GetDisplayName() *string  {
  return s.DisplayName
}

func (s *EnrollAccountRequest) GetFolderId() *string  {
  return s.FolderId
}

func (s *EnrollAccountRequest) GetPayerAccountUid() *int64  {
  return s.PayerAccountUid
}

func (s *EnrollAccountRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnrollAccountRequest) GetResellAccountType() *string  {
  return s.ResellAccountType
}

func (s *EnrollAccountRequest) GetTag() []*EnrollAccountRequestTag  {
  return s.Tag
}

func (s *EnrollAccountRequest) SetAccountNamePrefix(v string) *EnrollAccountRequest {
  s.AccountNamePrefix = &v
  return s
}

func (s *EnrollAccountRequest) SetAccountUid(v int64) *EnrollAccountRequest {
  s.AccountUid = &v
  return s
}

func (s *EnrollAccountRequest) SetBaselineId(v string) *EnrollAccountRequest {
  s.BaselineId = &v
  return s
}

func (s *EnrollAccountRequest) SetBaselineItems(v []*EnrollAccountRequestBaselineItems) *EnrollAccountRequest {
  s.BaselineItems = v
  return s
}

func (s *EnrollAccountRequest) SetDisplayName(v string) *EnrollAccountRequest {
  s.DisplayName = &v
  return s
}

func (s *EnrollAccountRequest) SetFolderId(v string) *EnrollAccountRequest {
  s.FolderId = &v
  return s
}

func (s *EnrollAccountRequest) SetPayerAccountUid(v int64) *EnrollAccountRequest {
  s.PayerAccountUid = &v
  return s
}

func (s *EnrollAccountRequest) SetRegionId(v string) *EnrollAccountRequest {
  s.RegionId = &v
  return s
}

func (s *EnrollAccountRequest) SetResellAccountType(v string) *EnrollAccountRequest {
  s.ResellAccountType = &v
  return s
}

func (s *EnrollAccountRequest) SetTag(v []*EnrollAccountRequestTag) *EnrollAccountRequest {
  s.Tag = v
  return s
}

func (s *EnrollAccountRequest) Validate() error {
  if s.BaselineItems != nil {
    for _, item := range s.BaselineItems {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  if s.Tag != nil {
    for _, item := range s.Tag {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type EnrollAccountRequestBaselineItems struct {
  // The configurations of the baseline item.
  // 
  // example:
  // 
  // {\\"Notifications\\":[{\\"GroupKey\\":\\"account_msg\\",\\"Contacts\\":[{\\"Name\\":\\"aa\\"}],\\"PmsgStatus\\":1,\\"EmailStatus\\":1,\\"SmsStatus\\":1}]}
  Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
  // The name of the baseline item.
  // 
  // example:
  // 
  // ACS-BP_ACCOUNT_FACTORY_VPC
  Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
  // Specifies whether to skip the baseline item. Valid values:
  // 
  // - false (default): does not skip the baseline item.
  // 
  // - true: skips the baseline item.
  // 
  // example:
  // 
  // false
  Skip *bool `json:"Skip,omitempty" xml:"Skip,omitempty"`
  // The version of the baseline item.
  // 
  // example:
  // 
  // 1.0
  Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s EnrollAccountRequestBaselineItems) String() string {
  return dara.Prettify(s)
}

func (s EnrollAccountRequestBaselineItems) GoString() string {
  return s.String()
}

func (s *EnrollAccountRequestBaselineItems) GetConfig() *string  {
  return s.Config
}

func (s *EnrollAccountRequestBaselineItems) GetName() *string  {
  return s.Name
}

func (s *EnrollAccountRequestBaselineItems) GetSkip() *bool  {
  return s.Skip
}

func (s *EnrollAccountRequestBaselineItems) GetVersion() *string  {
  return s.Version
}

func (s *EnrollAccountRequestBaselineItems) SetConfig(v string) *EnrollAccountRequestBaselineItems {
  s.Config = &v
  return s
}

func (s *EnrollAccountRequestBaselineItems) SetName(v string) *EnrollAccountRequestBaselineItems {
  s.Name = &v
  return s
}

func (s *EnrollAccountRequestBaselineItems) SetSkip(v bool) *EnrollAccountRequestBaselineItems {
  s.Skip = &v
  return s
}

func (s *EnrollAccountRequestBaselineItems) SetVersion(v string) *EnrollAccountRequestBaselineItems {
  s.Version = &v
  return s
}

func (s *EnrollAccountRequestBaselineItems) Validate() error {
  return dara.Validate(s)
}

type EnrollAccountRequestTag struct {
  // The tag key.
  // 
  // example:
  // 
  // tagKey
  Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
  // The tag value.
  // 
  // example:
  // 
  // tagValue
  Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s EnrollAccountRequestTag) String() string {
  return dara.Prettify(s)
}

func (s EnrollAccountRequestTag) GoString() string {
  return s.String()
}

func (s *EnrollAccountRequestTag) GetKey() *string  {
  return s.Key
}

func (s *EnrollAccountRequestTag) GetValue() *string  {
  return s.Value
}

func (s *EnrollAccountRequestTag) SetKey(v string) *EnrollAccountRequestTag {
  s.Key = &v
  return s
}

func (s *EnrollAccountRequestTag) SetValue(v string) *EnrollAccountRequestTag {
  s.Value = &v
  return s
}

func (s *EnrollAccountRequestTag) Validate() error {
  return dara.Validate(s)
}

