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
  // The prefix for the account name of the member.
  // 
  // 	- If the account baseline is applied to an account that is newly created, you must configure this parameter.
  // 
  // 	- If the account baseline is applied to an existing account, you do not need to configure this parameter.
  // 
  // example:
  // 
  // test-account
  AccountNamePrefix *string `json:"AccountNamePrefix,omitempty" xml:"AccountNamePrefix,omitempty"`
  // The account ID.
  // 
  // 	- If the account baseline is applied to an account that is newly created, you do not need to configure this parameter.
  // 
  // 	- If the account baseline is applied to an existing account, you must configure this parameter.
  // 
  // example:
  // 
  // 12868156179****
  AccountUid *int64 `json:"AccountUid,omitempty" xml:"AccountUid,omitempty"`
  // The baseline ID.
  // 
  // If this parameter is left empty, the default baseline is used.
  // 
  // example:
  // 
  // afb-bp1durvn3lgqe28v****
  BaselineId *string `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
  // The array that contains baseline items.
  // 
  // If this parameter is specified, the configurations of the baseline items are merged with the baseline applied to the specified account. The configurations of the same baseline items are subject to the configurations of this parameter. We recommend that you leave this parameter empty and configure the `BaselineId` parameter to specify an account baseline and apply the configurations of the account baseline to the account.
  BaselineItems []*EnrollAccountRequestBaselineItems `json:"BaselineItems,omitempty" xml:"BaselineItems,omitempty" type:"Repeated"`
  // The display name of the account.
  // 
  // 	- If the account baseline is applied to an account that is newly created, you must configure this parameter.
  // 
  // 	- If the account baseline is applied to an existing account, you do not need to configure this parameter.
  // 
  // example:
  // 
  // test-account
  DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
  // The ID of the parent folder.
  // 
  // 	- If the account baseline is applied to an account that is newly created, you need to specify a parent folder. If you do not configure this parameter, the account is created in the Root folder.
  // 
  // 	- If the account baseline is applied to an existing account, you do not need to configure this parameter.
  // 
  // example:
  // 
  // fd-5ESoku****
  FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
  // The ID of the billing account.
  // 
  // 	- If the account baseline is applied to an account that is newly created, you need to specify a billing account. If you do not configure this parameter, the self-pay settlement method is used for the account.
  // 
  // 	- If the account baseline is applied to an existing account, you do not need to configure this parameter.
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
  // 	- resell (default): The member is an account for a reseller. A relationship is automatically established between the member and the reseller. The management account of the resource directory must be used as the billing account of the member.
  // 
  // 	- non_resell: The member is not an account for a reseller. The member is an account that is not associated with a reseller. You can directly use the account to purchase Alibaba Cloud resources. The member is used as its own billing account.
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
  // Whether to skip the baseline item. Valid values:
  // 
  // 	- false: The baseline item is not skipped.
  // 
  // 	- true: The baseline item is skipped.
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

