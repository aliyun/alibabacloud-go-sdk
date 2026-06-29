// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveOrganizationMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIds(v []*string) *RemoveOrganizationMemberRequest
	GetAccountIds() []*string
	SetLocale(v string) *RemoveOrganizationMemberRequest
	GetLocale() *string
}

type RemoveOrganizationMemberRequest struct {
	// The list of member account IDs to be removed.
	AccountIds []*string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// The language. Valid values: zh-CN and en-US.
	//
	// example:
	//
	// zh-CN
	Locale *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
}

func (s RemoveOrganizationMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveOrganizationMemberRequest) GoString() string {
	return s.String()
}

func (s *RemoveOrganizationMemberRequest) GetAccountIds() []*string {
	return s.AccountIds
}

func (s *RemoveOrganizationMemberRequest) GetLocale() *string {
	return s.Locale
}

func (s *RemoveOrganizationMemberRequest) SetAccountIds(v []*string) *RemoveOrganizationMemberRequest {
	s.AccountIds = v
	return s
}

func (s *RemoveOrganizationMemberRequest) SetLocale(v string) *RemoveOrganizationMemberRequest {
	s.Locale = &v
	return s
}

func (s *RemoveOrganizationMemberRequest) Validate() error {
	return dara.Validate(s)
}
