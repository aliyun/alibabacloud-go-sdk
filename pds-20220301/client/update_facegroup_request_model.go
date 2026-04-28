// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFacegroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *UpdateFacegroupRequest
	GetDriveId() *string
	SetGroupCoverFaceId(v string) *UpdateFacegroupRequest
	GetGroupCoverFaceId() *string
	SetGroupId(v string) *UpdateFacegroupRequest
	GetGroupId() *string
	SetGroupName(v string) *UpdateFacegroupRequest
	GetGroupName() *string
	SetRemarks(v string) *UpdateFacegroupRequest
	GetRemarks() *string
}

type UpdateFacegroupRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The face ID of the thumbnail of the face-based group. You can obtain the face ID from the **image_media_metadata*	- parameter in the returned results of the GetFile, ListFile, or SearchFile operation.
	//
	// example:
	//
	// face1
	GroupCoverFaceId *string `json:"group_cover_face_id,omitempty" xml:"group_cover_face_id,omitempty"`
	// The ID of the face-based group. You can call the ListFacegroups operation to query the group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// group-abc
	GroupId *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
	// The name of the face-based group. The name can be up to 128 characters in length.
	GroupName *string `json:"group_name,omitempty" xml:"group_name,omitempty"`
	// The remarks. The remarks can be up to 128 characters in length.
	Remarks *string `json:"remarks,omitempty" xml:"remarks,omitempty"`
}

func (s UpdateFacegroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFacegroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateFacegroupRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *UpdateFacegroupRequest) GetGroupCoverFaceId() *string {
	return s.GroupCoverFaceId
}

func (s *UpdateFacegroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateFacegroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateFacegroupRequest) GetRemarks() *string {
	return s.Remarks
}

func (s *UpdateFacegroupRequest) SetDriveId(v string) *UpdateFacegroupRequest {
	s.DriveId = &v
	return s
}

func (s *UpdateFacegroupRequest) SetGroupCoverFaceId(v string) *UpdateFacegroupRequest {
	s.GroupCoverFaceId = &v
	return s
}

func (s *UpdateFacegroupRequest) SetGroupId(v string) *UpdateFacegroupRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateFacegroupRequest) SetGroupName(v string) *UpdateFacegroupRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateFacegroupRequest) SetRemarks(v string) *UpdateFacegroupRequest {
	s.Remarks = &v
	return s
}

func (s *UpdateFacegroupRequest) Validate() error {
	return dara.Validate(s)
}
