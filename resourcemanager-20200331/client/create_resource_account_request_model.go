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
	// The prefix of the account name. If you leave this parameter empty, the system randomly generates a prefix.
	//
	// <props="china">Length: 2 to 50 characters.
	//
	// <props="intl">Length: 2 to 37 characters.
	//
	// Format: Can contain letters, digits, and special characters `_.-`. Must start and end with a letter or digit. Consecutive special characters `_.-` are not allowed.
	//
	// The full account name is in the format of <AccountNamePrefix>@<ResourceDirectoryId>.aliyunid.com. Example: `alice@rd-3G****.aliyunid.com`.
	//
	// The account name must be unique within the resource directory.
	//
	// example:
	//
	// alice
	AccountNamePrefix *string `json:"AccountNamePrefix,omitempty" xml:"AccountNamePrefix,omitempty"`
	// The display name of the member.
	//
	// Length: 2 to 50 characters.
	//
	// Format: Can contain letters, digits, underscores (_), periods (.), hyphens (-), and spaces.
	//
	// The display name must be unique within the resource directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// Dev
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the parent folder.
	//
	// example:
	//
	// fd-r23M55****
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
	// The ID of the billing account. If you leave this parameter empty, the newly created member pays for its own resource consumption.
	//
	// example:
	//
	// 12323344****
	PayerAccountId *string `json:"PayerAccountId,omitempty" xml:"PayerAccountId,omitempty"`
	// The identity type of the member. Valid values:
	//
	// - resell (default): reseller account. The created member automatically establishes a reseller relationship with your distributor. You must use the management account to pay for the new member.
	//
	// - non_resell: non-reseller account. The created member is a regular Alibaba Cloud account that has no relationship with your current distributor. This account purchases resources directly from Alibaba Cloud and pays for its own resource consumption.
	//
	// > This parameter is supported only for China site (Chinese mainland) distributor customers.
	//
	// example:
	//
	// resell
	ResellAccountType *string `json:"ResellAccountType,omitempty" xml:"ResellAccountType,omitempty"`
	// The tags of the member.
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
	// The tag key.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
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
