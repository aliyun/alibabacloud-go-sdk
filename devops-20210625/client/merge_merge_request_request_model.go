// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMergeMergeRequestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *MergeMergeRequestRequest
	GetAccessToken() *string
	SetMergeMessage(v string) *MergeMergeRequestRequest
	GetMergeMessage() *string
	SetMergeType(v string) *MergeMergeRequestRequest
	GetMergeType() *string
	SetRemoveSourceBranch(v bool) *MergeMergeRequestRequest
	GetRemoveSourceBranch() *bool
	SetOrganizationId(v string) *MergeMergeRequestRequest
	GetOrganizationId() *string
}

type MergeMergeRequestRequest struct {
	// example:
	//
	// 0cf2c8458ac44d9481aab2dd6ec10596v3
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// example:
	//
	// ""
	MergeMessage *string `json:"mergeMessage,omitempty" xml:"mergeMessage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// no-fast-forward
	MergeType *string `json:"mergeType,omitempty" xml:"mergeType,omitempty"`
	// example:
	//
	// true
	RemoveSourceBranch *bool `json:"removeSourceBranch,omitempty" xml:"removeSourceBranch,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60de7a6852743a5162b5f957
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s MergeMergeRequestRequest) String() string {
	return dara.Prettify(s)
}

func (s MergeMergeRequestRequest) GoString() string {
	return s.String()
}

func (s *MergeMergeRequestRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *MergeMergeRequestRequest) GetMergeMessage() *string {
	return s.MergeMessage
}

func (s *MergeMergeRequestRequest) GetMergeType() *string {
	return s.MergeType
}

func (s *MergeMergeRequestRequest) GetRemoveSourceBranch() *bool {
	return s.RemoveSourceBranch
}

func (s *MergeMergeRequestRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *MergeMergeRequestRequest) SetAccessToken(v string) *MergeMergeRequestRequest {
	s.AccessToken = &v
	return s
}

func (s *MergeMergeRequestRequest) SetMergeMessage(v string) *MergeMergeRequestRequest {
	s.MergeMessage = &v
	return s
}

func (s *MergeMergeRequestRequest) SetMergeType(v string) *MergeMergeRequestRequest {
	s.MergeType = &v
	return s
}

func (s *MergeMergeRequestRequest) SetRemoveSourceBranch(v bool) *MergeMergeRequestRequest {
	s.RemoveSourceBranch = &v
	return s
}

func (s *MergeMergeRequestRequest) SetOrganizationId(v string) *MergeMergeRequestRequest {
	s.OrganizationId = &v
	return s
}

func (s *MergeMergeRequestRequest) Validate() error {
	return dara.Validate(s)
}
