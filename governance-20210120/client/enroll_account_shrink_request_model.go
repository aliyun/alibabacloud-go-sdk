// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnrollAccountShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAccountNamePrefix(v string) *EnrollAccountShrinkRequest
  GetAccountNamePrefix() *string 
  SetAccountUid(v int64) *EnrollAccountShrinkRequest
  GetAccountUid() *int64 
  SetBaselineId(v string) *EnrollAccountShrinkRequest
  GetBaselineId() *string 
  SetBaselineItems(v []*EnrollAccountShrinkRequestBaselineItems) *EnrollAccountShrinkRequest
  GetBaselineItems() []*EnrollAccountShrinkRequestBaselineItems 
  SetDisplayName(v string) *EnrollAccountShrinkRequest
  GetDisplayName() *string 
  SetFolderId(v string) *EnrollAccountShrinkRequest
  GetFolderId() *string 
  SetPayerAccountUid(v int64) *EnrollAccountShrinkRequest
  GetPayerAccountUid() *int64 
  SetRegionId(v string) *EnrollAccountShrinkRequest
  GetRegionId() *string 
  SetResellAccountType(v string) *EnrollAccountShrinkRequest
  GetResellAccountType() *string 
  SetTagShrink(v string) *EnrollAccountShrinkRequest
  GetTagShrink() *string 
}

type EnrollAccountShrinkRequest struct {
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
  BaselineItems []*EnrollAccountShrinkRequestBaselineItems `json:"BaselineItems,omitempty" xml:"BaselineItems,omitempty" type:"Repeated"`
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
  TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s EnrollAccountShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s EnrollAccountShrinkRequest) GoString() string {
  return s.String()
}

func (s *EnrollAccountShrinkRequest) GetAccountNamePrefix() *string  {
  return s.AccountNamePrefix
}

func (s *EnrollAccountShrinkRequest) GetAccountUid() *int64  {
  return s.AccountUid
}

func (s *EnrollAccountShrinkRequest) GetBaselineId() *string  {
  return s.BaselineId
}

func (s *EnrollAccountShrinkRequest) GetBaselineItems() []*EnrollAccountShrinkRequestBaselineItems  {
  return s.BaselineItems
}

func (s *EnrollAccountShrinkRequest) GetDisplayName() *string  {
  return s.DisplayName
}

func (s *EnrollAccountShrinkRequest) GetFolderId() *string  {
  return s.FolderId
}

func (s *EnrollAccountShrinkRequest) GetPayerAccountUid() *int64  {
  return s.PayerAccountUid
}

func (s *EnrollAccountShrinkRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnrollAccountShrinkRequest) GetResellAccountType() *string  {
  return s.ResellAccountType
}

func (s *EnrollAccountShrinkRequest) GetTagShrink() *string  {
  return s.TagShrink
}

func (s *EnrollAccountShrinkRequest) SetAccountNamePrefix(v string) *EnrollAccountShrinkRequest {
  s.AccountNamePrefix = &v
  return s
}

func (s *EnrollAccountShrinkRequest) SetAccountUid(v int64) *EnrollAccountShrinkRequest {
  s.AccountUid = &v
  return s
}

func (s *EnrollAccountShrinkRequest) SetBaselineId(v string) *EnrollAccountShrinkRequest {
  s.BaselineId = &v
  return s
}

func (s *EnrollAccountShrinkRequest) SetBaselineItems(v []*EnrollAccountShrinkRequestBaselineItems) *EnrollAccountShrinkRequest {
  s.BaselineItems = v
  return s
}

func (s *EnrollAccountShrinkRequest) SetDisplayName(v string) *EnrollAccountShrinkRequest {
  s.DisplayName = &v
  return s
}

func (s *EnrollAccountShrinkRequest) SetFolderId(v string) *EnrollAccountShrinkRequest {
  s.FolderId = &v
  return s
}

func (s *EnrollAccountShrinkRequest) SetPayerAccountUid(v int64) *EnrollAccountShrinkRequest {
  s.PayerAccountUid = &v
  return s
}

func (s *EnrollAccountShrinkRequest) SetRegionId(v string) *EnrollAccountShrinkRequest {
  s.RegionId = &v
  return s
}

func (s *EnrollAccountShrinkRequest) SetResellAccountType(v string) *EnrollAccountShrinkRequest {
  s.ResellAccountType = &v
  return s
}

func (s *EnrollAccountShrinkRequest) SetTagShrink(v string) *EnrollAccountShrinkRequest {
  s.TagShrink = &v
  return s
}

func (s *EnrollAccountShrinkRequest) Validate() error {
  if s.BaselineItems != nil {
    for _, item := range s.BaselineItems {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type EnrollAccountShrinkRequestBaselineItems struct {
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

func (s EnrollAccountShrinkRequestBaselineItems) String() string {
  return dara.Prettify(s)
}

func (s EnrollAccountShrinkRequestBaselineItems) GoString() string {
  return s.String()
}

func (s *EnrollAccountShrinkRequestBaselineItems) GetConfig() *string  {
  return s.Config
}

func (s *EnrollAccountShrinkRequestBaselineItems) GetName() *string  {
  return s.Name
}

func (s *EnrollAccountShrinkRequestBaselineItems) GetSkip() *bool  {
  return s.Skip
}

func (s *EnrollAccountShrinkRequestBaselineItems) GetVersion() *string  {
  return s.Version
}

func (s *EnrollAccountShrinkRequestBaselineItems) SetConfig(v string) *EnrollAccountShrinkRequestBaselineItems {
  s.Config = &v
  return s
}

func (s *EnrollAccountShrinkRequestBaselineItems) SetName(v string) *EnrollAccountShrinkRequestBaselineItems {
  s.Name = &v
  return s
}

func (s *EnrollAccountShrinkRequestBaselineItems) SetSkip(v bool) *EnrollAccountShrinkRequestBaselineItems {
  s.Skip = &v
  return s
}

func (s *EnrollAccountShrinkRequestBaselineItems) SetVersion(v string) *EnrollAccountShrinkRequestBaselineItems {
  s.Version = &v
  return s
}

func (s *EnrollAccountShrinkRequestBaselineItems) Validate() error {
  return dara.Validate(s)
}

