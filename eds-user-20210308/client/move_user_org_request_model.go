// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveUserOrgRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessChannel(v string) *MoveUserOrgRequest
	GetBusinessChannel() *string
	SetEndUserIds(v []*string) *MoveUserOrgRequest
	GetEndUserIds() []*string
	SetOrgId(v string) *MoveUserOrgRequest
	GetOrgId() *string
}

type MoveUserOrgRequest struct {
	// example:
	//
	// ENTERPRISE
	BusinessChannel *string `json:"BusinessChannel,omitempty" xml:"BusinessChannel,omitempty"`
	// The user IDs.
	//
	// This parameter is required.
	EndUserIds []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	// The organization ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// org-4mdgc1cocc59z****
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
}

func (s MoveUserOrgRequest) String() string {
	return dara.Prettify(s)
}

func (s MoveUserOrgRequest) GoString() string {
	return s.String()
}

func (s *MoveUserOrgRequest) GetBusinessChannel() *string {
	return s.BusinessChannel
}

func (s *MoveUserOrgRequest) GetEndUserIds() []*string {
	return s.EndUserIds
}

func (s *MoveUserOrgRequest) GetOrgId() *string {
	return s.OrgId
}

func (s *MoveUserOrgRequest) SetBusinessChannel(v string) *MoveUserOrgRequest {
	s.BusinessChannel = &v
	return s
}

func (s *MoveUserOrgRequest) SetEndUserIds(v []*string) *MoveUserOrgRequest {
	s.EndUserIds = v
	return s
}

func (s *MoveUserOrgRequest) SetOrgId(v string) *MoveUserOrgRequest {
	s.OrgId = &v
	return s
}

func (s *MoveUserOrgRequest) Validate() error {
	return dara.Validate(s)
}
