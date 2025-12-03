// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChangeRequestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppCodeRepoSn(v string) *CreateChangeRequestResponseBody
	GetAppCodeRepoSn() *string
	SetAppName(v string) *CreateChangeRequestResponseBody
	GetAppName() *string
	SetAutoDeleteBranchWhenEnd(v bool) *CreateChangeRequestResponseBody
	GetAutoDeleteBranchWhenEnd() *bool
	SetBranch(v string) *CreateChangeRequestResponseBody
	GetBranch() *string
	SetCreatorAccountId(v string) *CreateChangeRequestResponseBody
	GetCreatorAccountId() *string
	SetCreatorId(v string) *CreateChangeRequestResponseBody
	GetCreatorId() *string
	SetGmtCreate(v string) *CreateChangeRequestResponseBody
	GetGmtCreate() *string
	SetGmtModified(v string) *CreateChangeRequestResponseBody
	GetGmtModified() *string
	SetName(v string) *CreateChangeRequestResponseBody
	GetName() *string
	SetOriginBranch(v string) *CreateChangeRequestResponseBody
	GetOriginBranch() *string
	SetOriginBranchRevisionSha(v string) *CreateChangeRequestResponseBody
	GetOriginBranchRevisionSha() *string
	SetOwnerAccountId(v string) *CreateChangeRequestResponseBody
	GetOwnerAccountId() *string
	SetOwnerId(v string) *CreateChangeRequestResponseBody
	GetOwnerId() *string
	SetSn(v string) *CreateChangeRequestResponseBody
	GetSn() *string
	SetState(v string) *CreateChangeRequestResponseBody
	GetState() *string
	SetType(v string) *CreateChangeRequestResponseBody
	GetType() *string
}

type CreateChangeRequestResponseBody struct {
	// example:
	//
	// sn123
	AppCodeRepoSn *string `json:"appCodeRepoSn,omitempty" xml:"appCodeRepoSn,omitempty"`
	// example:
	//
	// app-name
	AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
	// example:
	//
	// false
	AutoDeleteBranchWhenEnd *bool `json:"autoDeleteBranchWhenEnd,omitempty" xml:"autoDeleteBranchWhenEnd,omitempty"`
	// example:
	//
	// hotfix/20240524
	Branch *string `json:"branch,omitempty" xml:"branch,omitempty"`
	// example:
	//
	// create-account-123
	CreatorAccountId *string `json:"creatorAccountId,omitempty" xml:"creatorAccountId,omitempty"`
	// example:
	//
	// create-id-123
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// master
	OriginBranch *string `json:"originBranch,omitempty" xml:"originBranch,omitempty"`
	// example:
	//
	// revision-123
	OriginBranchRevisionSha *string `json:"originBranchRevisionSha,omitempty" xml:"originBranchRevisionSha,omitempty"`
	// example:
	//
	// account-id-123
	OwnerAccountId *string `json:"ownerAccountId,omitempty" xml:"ownerAccountId,omitempty"`
	// example:
	//
	// owner-id-123
	OwnerId *string `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	// example:
	//
	// ce51b31b996246ecaf874736838360b2
	Sn *string `json:"sn,omitempty" xml:"sn,omitempty"`
	// example:
	//
	// DEVELOPING
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// example:
	//
	// APP
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateChangeRequestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateChangeRequestResponseBody) GoString() string {
	return s.String()
}

func (s *CreateChangeRequestResponseBody) GetAppCodeRepoSn() *string {
	return s.AppCodeRepoSn
}

func (s *CreateChangeRequestResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *CreateChangeRequestResponseBody) GetAutoDeleteBranchWhenEnd() *bool {
	return s.AutoDeleteBranchWhenEnd
}

func (s *CreateChangeRequestResponseBody) GetBranch() *string {
	return s.Branch
}

func (s *CreateChangeRequestResponseBody) GetCreatorAccountId() *string {
	return s.CreatorAccountId
}

func (s *CreateChangeRequestResponseBody) GetCreatorId() *string {
	return s.CreatorId
}

func (s *CreateChangeRequestResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreateChangeRequestResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *CreateChangeRequestResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateChangeRequestResponseBody) GetOriginBranch() *string {
	return s.OriginBranch
}

func (s *CreateChangeRequestResponseBody) GetOriginBranchRevisionSha() *string {
	return s.OriginBranchRevisionSha
}

func (s *CreateChangeRequestResponseBody) GetOwnerAccountId() *string {
	return s.OwnerAccountId
}

func (s *CreateChangeRequestResponseBody) GetOwnerId() *string {
	return s.OwnerId
}

func (s *CreateChangeRequestResponseBody) GetSn() *string {
	return s.Sn
}

func (s *CreateChangeRequestResponseBody) GetState() *string {
	return s.State
}

func (s *CreateChangeRequestResponseBody) GetType() *string {
	return s.Type
}

func (s *CreateChangeRequestResponseBody) SetAppCodeRepoSn(v string) *CreateChangeRequestResponseBody {
	s.AppCodeRepoSn = &v
	return s
}

func (s *CreateChangeRequestResponseBody) SetAppName(v string) *CreateChangeRequestResponseBody {
	s.AppName = &v
	return s
}

func (s *CreateChangeRequestResponseBody) SetAutoDeleteBranchWhenEnd(v bool) *CreateChangeRequestResponseBody {
	s.AutoDeleteBranchWhenEnd = &v
	return s
}

func (s *CreateChangeRequestResponseBody) SetBranch(v string) *CreateChangeRequestResponseBody {
	s.Branch = &v
	return s
}

func (s *CreateChangeRequestResponseBody) SetCreatorAccountId(v string) *CreateChangeRequestResponseBody {
	s.CreatorAccountId = &v
	return s
}

func (s *CreateChangeRequestResponseBody) SetCreatorId(v string) *CreateChangeRequestResponseBody {
	s.CreatorId = &v
	return s
}

func (s *CreateChangeRequestResponseBody) SetGmtCreate(v string) *CreateChangeRequestResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreateChangeRequestResponseBody) SetGmtModified(v string) *CreateChangeRequestResponseBody {
	s.GmtModified = &v
	return s
}

func (s *CreateChangeRequestResponseBody) SetName(v string) *CreateChangeRequestResponseBody {
	s.Name = &v
	return s
}

func (s *CreateChangeRequestResponseBody) SetOriginBranch(v string) *CreateChangeRequestResponseBody {
	s.OriginBranch = &v
	return s
}

func (s *CreateChangeRequestResponseBody) SetOriginBranchRevisionSha(v string) *CreateChangeRequestResponseBody {
	s.OriginBranchRevisionSha = &v
	return s
}

func (s *CreateChangeRequestResponseBody) SetOwnerAccountId(v string) *CreateChangeRequestResponseBody {
	s.OwnerAccountId = &v
	return s
}

func (s *CreateChangeRequestResponseBody) SetOwnerId(v string) *CreateChangeRequestResponseBody {
	s.OwnerId = &v
	return s
}

func (s *CreateChangeRequestResponseBody) SetSn(v string) *CreateChangeRequestResponseBody {
	s.Sn = &v
	return s
}

func (s *CreateChangeRequestResponseBody) SetState(v string) *CreateChangeRequestResponseBody {
	s.State = &v
	return s
}

func (s *CreateChangeRequestResponseBody) SetType(v string) *CreateChangeRequestResponseBody {
	s.Type = &v
	return s
}

func (s *CreateChangeRequestResponseBody) Validate() error {
	return dara.Validate(s)
}
