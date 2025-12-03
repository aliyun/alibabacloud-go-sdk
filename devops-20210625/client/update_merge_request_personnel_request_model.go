// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMergeRequestPersonnelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *UpdateMergeRequestPersonnelRequest
	GetAccessToken() *string
	SetNewUserIdList(v []*string) *UpdateMergeRequestPersonnelRequest
	GetNewUserIdList() []*string
	SetOrganizationId(v string) *UpdateMergeRequestPersonnelRequest
	GetOrganizationId() *string
}

type UpdateMergeRequestPersonnelRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken   *string   `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	NewUserIdList []*string `json:"newUserIdList,omitempty" xml:"newUserIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 5ebbc0228123212b59xxxxx
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s UpdateMergeRequestPersonnelRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMergeRequestPersonnelRequest) GoString() string {
	return s.String()
}

func (s *UpdateMergeRequestPersonnelRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *UpdateMergeRequestPersonnelRequest) GetNewUserIdList() []*string {
	return s.NewUserIdList
}

func (s *UpdateMergeRequestPersonnelRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *UpdateMergeRequestPersonnelRequest) SetAccessToken(v string) *UpdateMergeRequestPersonnelRequest {
	s.AccessToken = &v
	return s
}

func (s *UpdateMergeRequestPersonnelRequest) SetNewUserIdList(v []*string) *UpdateMergeRequestPersonnelRequest {
	s.NewUserIdList = v
	return s
}

func (s *UpdateMergeRequestPersonnelRequest) SetOrganizationId(v string) *UpdateMergeRequestPersonnelRequest {
	s.OrganizationId = &v
	return s
}

func (s *UpdateMergeRequestPersonnelRequest) Validate() error {
	return dara.Validate(s)
}
