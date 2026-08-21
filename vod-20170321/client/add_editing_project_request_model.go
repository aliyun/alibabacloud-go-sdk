// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddEditingProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCoverURL(v string) *AddEditingProjectRequest
	GetCoverURL() *string
	SetDescription(v string) *AddEditingProjectRequest
	GetDescription() *string
	SetDivision(v string) *AddEditingProjectRequest
	GetDivision() *string
	SetOwnerAccount(v string) *AddEditingProjectRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *AddEditingProjectRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *AddEditingProjectRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *AddEditingProjectRequest
	GetResourceOwnerId() *string
	SetTimeline(v string) *AddEditingProjectRequest
	GetTimeline() *string
	SetTitle(v string) *AddEditingProjectRequest
	GetTitle() *string
}

type AddEditingProjectRequest struct {
	// The thumbnail of the online editing project. If this parameter is left empty and the video track on the timeline already contains materials, the thumbnail of the first material on the timeline is used by default.
	//
	// example:
	//
	// https://demo.aliyundoc.com/6AB4D0E1E1C74468883516C2349D1FC2-6-2.png
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The description of the online editing project.
	//
	// example:
	//
	// testtimeline001desciption
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The service region.
	//
	// example:
	//
	// cn-shanghai
	Division             *string `json:"Division,omitempty" xml:"Division,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The timeline of the online editing project. JSON format. For the specific structure definition, see [Timeline](https://help.aliyun.com/document_detail/52839.html).
	//
	// If this parameter is left empty, an empty timeline is created and the total duration of the online editing project is 0.
	//
	// example:
	//
	// {"VideoTracks":[{"VideoTrackClips":[{"MediaId":"cc3308ac5006aed55a54328bc3443****"},{"MediaId":"95948ddba24446b6aed5db985e78****"}]}]}
	Timeline *string `json:"Timeline,omitempty" xml:"Timeline,omitempty"`
	// The title of the online editing project.
	//
	// This parameter is required.
	//
	// example:
	//
	// testtimeline
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s AddEditingProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s AddEditingProjectRequest) GoString() string {
	return s.String()
}

func (s *AddEditingProjectRequest) GetCoverURL() *string {
	return s.CoverURL
}

func (s *AddEditingProjectRequest) GetDescription() *string {
	return s.Description
}

func (s *AddEditingProjectRequest) GetDivision() *string {
	return s.Division
}

func (s *AddEditingProjectRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddEditingProjectRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *AddEditingProjectRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddEditingProjectRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *AddEditingProjectRequest) GetTimeline() *string {
	return s.Timeline
}

func (s *AddEditingProjectRequest) GetTitle() *string {
	return s.Title
}

func (s *AddEditingProjectRequest) SetCoverURL(v string) *AddEditingProjectRequest {
	s.CoverURL = &v
	return s
}

func (s *AddEditingProjectRequest) SetDescription(v string) *AddEditingProjectRequest {
	s.Description = &v
	return s
}

func (s *AddEditingProjectRequest) SetDivision(v string) *AddEditingProjectRequest {
	s.Division = &v
	return s
}

func (s *AddEditingProjectRequest) SetOwnerAccount(v string) *AddEditingProjectRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddEditingProjectRequest) SetOwnerId(v string) *AddEditingProjectRequest {
	s.OwnerId = &v
	return s
}

func (s *AddEditingProjectRequest) SetResourceOwnerAccount(v string) *AddEditingProjectRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddEditingProjectRequest) SetResourceOwnerId(v string) *AddEditingProjectRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddEditingProjectRequest) SetTimeline(v string) *AddEditingProjectRequest {
	s.Timeline = &v
	return s
}

func (s *AddEditingProjectRequest) SetTitle(v string) *AddEditingProjectRequest {
	s.Title = &v
	return s
}

func (s *AddEditingProjectRequest) Validate() error {
	return dara.Validate(s)
}
