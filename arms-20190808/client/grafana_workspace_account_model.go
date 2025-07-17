// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrafanaWorkspaceAccount interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v int64) *GrafanaWorkspaceAccount
	GetAccountId() *int64
	SetAccountNotes(v string) *GrafanaWorkspaceAccount
	GetAccountNotes() *string
	SetAliyunUid(v string) *GrafanaWorkspaceAccount
	GetAliyunUid() *string
	SetAliyunUserName(v string) *GrafanaWorkspaceAccount
	GetAliyunUserName() *string
	SetGmtCreate(v float32) *GrafanaWorkspaceAccount
	GetGmtCreate() *float32
	SetOrgs(v []*GrafanaWorkspaceUserOrg) *GrafanaWorkspaceAccount
	GetOrgs() []*GrafanaWorkspaceUserOrg
	SetType(v string) *GrafanaWorkspaceAccount
	GetType() *string
}

type GrafanaWorkspaceAccount struct {
	// example:
	//
	// 1
	AccountId    *int64  `json:"accountId,omitempty" xml:"accountId,omitempty"`
	AccountNotes *string `json:"accountNotes,omitempty" xml:"accountNotes,omitempty"`
	// example:
	//
	// 131243781293
	AliyunUid *string `json:"aliyunUid,omitempty" xml:"aliyunUid,omitempty"`
	// example:
	//
	// test
	AliyunUserName *string `json:"aliyunUserName,omitempty" xml:"aliyunUserName,omitempty"`
	// example:
	//
	// 创建时间
	GmtCreate *float32                   `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	Orgs      []*GrafanaWorkspaceUserOrg `json:"orgs,omitempty" xml:"orgs,omitempty" type:"Repeated"`
	// example:
	//
	// aliyun
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GrafanaWorkspaceAccount) String() string {
	return dara.Prettify(s)
}

func (s GrafanaWorkspaceAccount) GoString() string {
	return s.String()
}

func (s *GrafanaWorkspaceAccount) GetAccountId() *int64 {
	return s.AccountId
}

func (s *GrafanaWorkspaceAccount) GetAccountNotes() *string {
	return s.AccountNotes
}

func (s *GrafanaWorkspaceAccount) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *GrafanaWorkspaceAccount) GetAliyunUserName() *string {
	return s.AliyunUserName
}

func (s *GrafanaWorkspaceAccount) GetGmtCreate() *float32 {
	return s.GmtCreate
}

func (s *GrafanaWorkspaceAccount) GetOrgs() []*GrafanaWorkspaceUserOrg {
	return s.Orgs
}

func (s *GrafanaWorkspaceAccount) GetType() *string {
	return s.Type
}

func (s *GrafanaWorkspaceAccount) SetAccountId(v int64) *GrafanaWorkspaceAccount {
	s.AccountId = &v
	return s
}

func (s *GrafanaWorkspaceAccount) SetAccountNotes(v string) *GrafanaWorkspaceAccount {
	s.AccountNotes = &v
	return s
}

func (s *GrafanaWorkspaceAccount) SetAliyunUid(v string) *GrafanaWorkspaceAccount {
	s.AliyunUid = &v
	return s
}

func (s *GrafanaWorkspaceAccount) SetAliyunUserName(v string) *GrafanaWorkspaceAccount {
	s.AliyunUserName = &v
	return s
}

func (s *GrafanaWorkspaceAccount) SetGmtCreate(v float32) *GrafanaWorkspaceAccount {
	s.GmtCreate = &v
	return s
}

func (s *GrafanaWorkspaceAccount) SetOrgs(v []*GrafanaWorkspaceUserOrg) *GrafanaWorkspaceAccount {
	s.Orgs = v
	return s
}

func (s *GrafanaWorkspaceAccount) SetType(v string) *GrafanaWorkspaceAccount {
	s.Type = &v
	return s
}

func (s *GrafanaWorkspaceAccount) Validate() error {
	return dara.Validate(s)
}
