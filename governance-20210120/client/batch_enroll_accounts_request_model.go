// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchEnrollAccountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccounts(v []*BatchEnrollAccountsRequestAccounts) *BatchEnrollAccountsRequest
	GetAccounts() []*BatchEnrollAccountsRequestAccounts
	SetBaselineId(v string) *BatchEnrollAccountsRequest
	GetBaselineId() *string
	SetBaselineItems(v []*BatchEnrollAccountsRequestBaselineItems) *BatchEnrollAccountsRequest
	GetBaselineItems() []*BatchEnrollAccountsRequestBaselineItems
	SetRegionId(v string) *BatchEnrollAccountsRequest
	GetRegionId() *string
}

type BatchEnrollAccountsRequest struct {
	// The resource accounts.
	Accounts []*BatchEnrollAccountsRequestAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
	// The baseline ID.
	//
	// If this parameter is left empty, the default baseline is used.
	//
	// example:
	//
	// afb-bp1durvn3lgqe28v****
	BaselineId *string `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The baseline items.
	//
	// If this parameter is specified, the configurations of the baseline items are merged with the baseline applied to the specified account. The configurations of the same baseline items are subject to the configurations of this parameter. We recommend that you leave this parameter empty and configure the `BaselineId` parameter to specify an account baseline and apply the configurations of the account baseline to the account.
	BaselineItems []*BatchEnrollAccountsRequestBaselineItems `json:"BaselineItems,omitempty" xml:"BaselineItems,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s BatchEnrollAccountsRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchEnrollAccountsRequest) GoString() string {
	return s.String()
}

func (s *BatchEnrollAccountsRequest) GetAccounts() []*BatchEnrollAccountsRequestAccounts {
	return s.Accounts
}

func (s *BatchEnrollAccountsRequest) GetBaselineId() *string {
	return s.BaselineId
}

func (s *BatchEnrollAccountsRequest) GetBaselineItems() []*BatchEnrollAccountsRequestBaselineItems {
	return s.BaselineItems
}

func (s *BatchEnrollAccountsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *BatchEnrollAccountsRequest) SetAccounts(v []*BatchEnrollAccountsRequestAccounts) *BatchEnrollAccountsRequest {
	s.Accounts = v
	return s
}

func (s *BatchEnrollAccountsRequest) SetBaselineId(v string) *BatchEnrollAccountsRequest {
	s.BaselineId = &v
	return s
}

func (s *BatchEnrollAccountsRequest) SetBaselineItems(v []*BatchEnrollAccountsRequestBaselineItems) *BatchEnrollAccountsRequest {
	s.BaselineItems = v
	return s
}

func (s *BatchEnrollAccountsRequest) SetRegionId(v string) *BatchEnrollAccountsRequest {
	s.RegionId = &v
	return s
}

func (s *BatchEnrollAccountsRequest) Validate() error {
	if s.Accounts != nil {
		for _, item := range s.Accounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type BatchEnrollAccountsRequestAccounts struct {
	// The account ID. This parameter is required.
	//
	// example:
	//
	// 12868156179****
	AccountUid *int64 `json:"AccountUid,omitempty" xml:"AccountUid,omitempty"`
}

func (s BatchEnrollAccountsRequestAccounts) String() string {
	return dara.Prettify(s)
}

func (s BatchEnrollAccountsRequestAccounts) GoString() string {
	return s.String()
}

func (s *BatchEnrollAccountsRequestAccounts) GetAccountUid() *int64 {
	return s.AccountUid
}

func (s *BatchEnrollAccountsRequestAccounts) SetAccountUid(v int64) *BatchEnrollAccountsRequestAccounts {
	s.AccountUid = &v
	return s
}

func (s *BatchEnrollAccountsRequestAccounts) Validate() error {
	return dara.Validate(s)
}

type BatchEnrollAccountsRequestBaselineItems struct {
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
	// 	- false
	//
	// 	- true
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

func (s BatchEnrollAccountsRequestBaselineItems) String() string {
	return dara.Prettify(s)
}

func (s BatchEnrollAccountsRequestBaselineItems) GoString() string {
	return s.String()
}

func (s *BatchEnrollAccountsRequestBaselineItems) GetConfig() *string {
	return s.Config
}

func (s *BatchEnrollAccountsRequestBaselineItems) GetName() *string {
	return s.Name
}

func (s *BatchEnrollAccountsRequestBaselineItems) GetSkip() *bool {
	return s.Skip
}

func (s *BatchEnrollAccountsRequestBaselineItems) GetVersion() *string {
	return s.Version
}

func (s *BatchEnrollAccountsRequestBaselineItems) SetConfig(v string) *BatchEnrollAccountsRequestBaselineItems {
	s.Config = &v
	return s
}

func (s *BatchEnrollAccountsRequestBaselineItems) SetName(v string) *BatchEnrollAccountsRequestBaselineItems {
	s.Name = &v
	return s
}

func (s *BatchEnrollAccountsRequestBaselineItems) SetSkip(v bool) *BatchEnrollAccountsRequestBaselineItems {
	s.Skip = &v
	return s
}

func (s *BatchEnrollAccountsRequestBaselineItems) SetVersion(v string) *BatchEnrollAccountsRequestBaselineItems {
	s.Version = &v
	return s
}

func (s *BatchEnrollAccountsRequestBaselineItems) Validate() error {
	return dara.Validate(s)
}
