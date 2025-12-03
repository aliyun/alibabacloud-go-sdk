// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLinkMergeRequestLabelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *LinkMergeRequestLabelRequest
	GetAccessToken() *string
	SetLabelIds(v []*string) *LinkMergeRequestLabelRequest
	GetLabelIds() []*string
	SetLocalId(v int64) *LinkMergeRequestLabelRequest
	GetLocalId() *int64
	SetOrganizationId(v string) *LinkMergeRequestLabelRequest
	GetOrganizationId() *string
	SetRepositoryIdentity(v string) *LinkMergeRequestLabelRequest
	GetRepositoryIdentity() *string
}

type LinkMergeRequestLabelRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	LabelIds []*string `json:"labelIds,omitempty" xml:"labelIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	LocalId *int64 `json:"localId,omitempty" xml:"localId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5ebbc0228123212b59xxxxx
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// This parameter is required.
	RepositoryIdentity *string `json:"repositoryIdentity,omitempty" xml:"repositoryIdentity,omitempty"`
}

func (s LinkMergeRequestLabelRequest) String() string {
	return dara.Prettify(s)
}

func (s LinkMergeRequestLabelRequest) GoString() string {
	return s.String()
}

func (s *LinkMergeRequestLabelRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *LinkMergeRequestLabelRequest) GetLabelIds() []*string {
	return s.LabelIds
}

func (s *LinkMergeRequestLabelRequest) GetLocalId() *int64 {
	return s.LocalId
}

func (s *LinkMergeRequestLabelRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *LinkMergeRequestLabelRequest) GetRepositoryIdentity() *string {
	return s.RepositoryIdentity
}

func (s *LinkMergeRequestLabelRequest) SetAccessToken(v string) *LinkMergeRequestLabelRequest {
	s.AccessToken = &v
	return s
}

func (s *LinkMergeRequestLabelRequest) SetLabelIds(v []*string) *LinkMergeRequestLabelRequest {
	s.LabelIds = v
	return s
}

func (s *LinkMergeRequestLabelRequest) SetLocalId(v int64) *LinkMergeRequestLabelRequest {
	s.LocalId = &v
	return s
}

func (s *LinkMergeRequestLabelRequest) SetOrganizationId(v string) *LinkMergeRequestLabelRequest {
	s.OrganizationId = &v
	return s
}

func (s *LinkMergeRequestLabelRequest) SetRepositoryIdentity(v string) *LinkMergeRequestLabelRequest {
	s.RepositoryIdentity = &v
	return s
}

func (s *LinkMergeRequestLabelRequest) Validate() error {
	return dara.Validate(s)
}
