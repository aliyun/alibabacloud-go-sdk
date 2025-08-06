// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEditingProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCoverURL(v string) *UpdateEditingProjectRequest
	GetCoverURL() *string
	SetDescription(v string) *UpdateEditingProjectRequest
	GetDescription() *string
	SetOwnerAccount(v string) *UpdateEditingProjectRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *UpdateEditingProjectRequest
	GetOwnerId() *string
	SetProjectId(v string) *UpdateEditingProjectRequest
	GetProjectId() *string
	SetResourceOwnerAccount(v string) *UpdateEditingProjectRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *UpdateEditingProjectRequest
	GetResourceOwnerId() *string
	SetTimeline(v string) *UpdateEditingProjectRequest
	GetTimeline() *string
	SetTitle(v string) *UpdateEditingProjectRequest
	GetTitle() *string
}

type UpdateEditingProjectRequest struct {
	// The thumbnail URL of the online editing project.
	//
	// example:
	//
	// https://****.com/6AB4D0E1E1C7446888****.png
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The description of the online editing project.
	//
	// example:
	//
	// testtimeline001desciption
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the online editing project.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4ee4b97e27*****b525142a6b2
	ProjectId            *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The timeline of the online editing project. For more information about the structure, see [Timeline](https://help.aliyun.com/document_detail/52839.html).
	//
	// example:
	//
	// {"VideoTracks":[{"VideoTrackClips":[{"MediaId":"cc3308ac500c*****a54328bc3443"},{"MediaId":"da87a9cff64*****d88bc6d8326e4"}]}]}
	Timeline *string `json:"Timeline,omitempty" xml:"Timeline,omitempty"`
	// The title of the online editing project.
	//
	// example:
	//
	// testtimeline
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateEditingProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEditingProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateEditingProjectRequest) GetCoverURL() *string {
	return s.CoverURL
}

func (s *UpdateEditingProjectRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateEditingProjectRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateEditingProjectRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *UpdateEditingProjectRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *UpdateEditingProjectRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateEditingProjectRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *UpdateEditingProjectRequest) GetTimeline() *string {
	return s.Timeline
}

func (s *UpdateEditingProjectRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateEditingProjectRequest) SetCoverURL(v string) *UpdateEditingProjectRequest {
	s.CoverURL = &v
	return s
}

func (s *UpdateEditingProjectRequest) SetDescription(v string) *UpdateEditingProjectRequest {
	s.Description = &v
	return s
}

func (s *UpdateEditingProjectRequest) SetOwnerAccount(v string) *UpdateEditingProjectRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateEditingProjectRequest) SetOwnerId(v string) *UpdateEditingProjectRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateEditingProjectRequest) SetProjectId(v string) *UpdateEditingProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateEditingProjectRequest) SetResourceOwnerAccount(v string) *UpdateEditingProjectRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateEditingProjectRequest) SetResourceOwnerId(v string) *UpdateEditingProjectRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateEditingProjectRequest) SetTimeline(v string) *UpdateEditingProjectRequest {
	s.Timeline = &v
	return s
}

func (s *UpdateEditingProjectRequest) SetTitle(v string) *UpdateEditingProjectRequest {
	s.Title = &v
	return s
}

func (s *UpdateEditingProjectRequest) Validate() error {
	return dara.Validate(s)
}
