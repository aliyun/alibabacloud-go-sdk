// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNamePrefix(v string) *CreateResourceAccountRequest
	GetAccountNamePrefix() *string
	SetDisplayName(v string) *CreateResourceAccountRequest
	GetDisplayName() *string
	SetDryRun(v bool) *CreateResourceAccountRequest
	GetDryRun() *bool
	SetParentFolderId(v string) *CreateResourceAccountRequest
	GetParentFolderId() *string
	SetPayerAccountId(v string) *CreateResourceAccountRequest
	GetPayerAccountId() *string
	SetResellAccountType(v string) *CreateResourceAccountRequest
	GetResellAccountType() *string
	SetTag(v []*CreateResourceAccountRequestTag) *CreateResourceAccountRequest
	GetTag() []*CreateResourceAccountRequestTag
}

type CreateResourceAccountRequest struct {
	// The prefix for the Alibaba Cloud account name of the member. If you leave this parameter empty, the system randomly generates a prefix.
	//
	// The prefix must be 2 to 37 characters in length.
	//
	// The prefix can contain letters, digits, and special characters but cannot contain consecutive special characters. The prefix must start with a letter or digit and end with a letter or digit. Valid special characters include underscores (`_`), periods (.), and hyphens (`-`).
	//
	// The complete Alibaba Cloud account name of a member in a resource directory is in the @.aliyunid.com format, such as `alice@rd-3G****.aliyunid.com`.
	//
	// Each name must be unique in the resource directory.
	//
	// example:
	//
	// alice
	AccountNamePrefix *string `json:"AccountNamePrefix,omitempty" xml:"AccountNamePrefix,omitempty"`
	// The display name of the member.
	//
	// The name must be 2 to 50 characters in length.
	//
	// The name can contain letters, digits, underscores (_), periods (.), hyphens (-), and spaces.
	//
	// The name must be unique in the resource directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// Dev
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- true: performs only a dry run. The system checks whether an identity type can be specified for the member. If the request does not pass the dry run, an error code is returned.
	//
	// 	- false (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the parent folder.
	//
	// example:
	//
	// fd-r23M55****
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
	// The ID of the billing account. If you leave this parameter empty, the newly created member is used as its billing account.
	//
	// example:
	//
	// 12323344****
	PayerAccountId *string `json:"PayerAccountId,omitempty" xml:"PayerAccountId,omitempty"`
	// The identity type of the member. Valid values:
	//
	// 	- resell: The member is an account for a reseller. This is the default value. A relationship is automatically established between the member and the reseller. The management account of the resource directory must be used as the billing account of the member.
	//
	// 	- non_resell: The member is not an account for a reseller. The member is an account that is not associated with a reseller. You can directly use the account to purchase Alibaba Cloud resources. The member is used as its own billing account.
	//
	// > This parameter is available only for resellers at the international site (alibabacloud.com).
	//
	// example:
	//
	// resell
	ResellAccountType *string `json:"ResellAccountType,omitempty" xml:"ResellAccountType,omitempty"`
	// The tag of the member.
	Tag []*CreateResourceAccountRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateResourceAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceAccountRequest) GetAccountNamePrefix() *string {
	return s.AccountNamePrefix
}

func (s *CreateResourceAccountRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateResourceAccountRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateResourceAccountRequest) GetParentFolderId() *string {
	return s.ParentFolderId
}

func (s *CreateResourceAccountRequest) GetPayerAccountId() *string {
	return s.PayerAccountId
}

func (s *CreateResourceAccountRequest) GetResellAccountType() *string {
	return s.ResellAccountType
}

func (s *CreateResourceAccountRequest) GetTag() []*CreateResourceAccountRequestTag {
	return s.Tag
}

func (s *CreateResourceAccountRequest) SetAccountNamePrefix(v string) *CreateResourceAccountRequest {
	s.AccountNamePrefix = &v
	return s
}

func (s *CreateResourceAccountRequest) SetDisplayName(v string) *CreateResourceAccountRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateResourceAccountRequest) SetDryRun(v bool) *CreateResourceAccountRequest {
	s.DryRun = &v
	return s
}

func (s *CreateResourceAccountRequest) SetParentFolderId(v string) *CreateResourceAccountRequest {
	s.ParentFolderId = &v
	return s
}

func (s *CreateResourceAccountRequest) SetPayerAccountId(v string) *CreateResourceAccountRequest {
	s.PayerAccountId = &v
	return s
}

func (s *CreateResourceAccountRequest) SetResellAccountType(v string) *CreateResourceAccountRequest {
	s.ResellAccountType = &v
	return s
}

func (s *CreateResourceAccountRequest) SetTag(v []*CreateResourceAccountRequestTag) *CreateResourceAccountRequest {
	s.Tag = v
	return s
}

func (s *CreateResourceAccountRequest) Validate() error {
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

type CreateResourceAccountRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateResourceAccountRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceAccountRequestTag) GoString() string {
	return s.String()
}

func (s *CreateResourceAccountRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateResourceAccountRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateResourceAccountRequestTag) SetKey(v string) *CreateResourceAccountRequestTag {
	s.Key = &v
	return s
}

func (s *CreateResourceAccountRequestTag) SetValue(v string) *CreateResourceAccountRequestTag {
	s.Value = &v
	return s
}

func (s *CreateResourceAccountRequestTag) Validate() error {
	return dara.Validate(s)
}
