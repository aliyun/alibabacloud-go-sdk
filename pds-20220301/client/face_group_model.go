// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFaceGroup interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v string) *FaceGroup
	GetCreatedAt() *string
	SetGroupCoverFaceBoundary(v *FaceGroupGroupCoverFaceBoundary) *FaceGroup
	GetGroupCoverFaceBoundary() *FaceGroupGroupCoverFaceBoundary
	SetGroupCoverFileId(v string) *FaceGroup
	GetGroupCoverFileId() *string
	SetGroupCoverHeight(v int64) *FaceGroup
	GetGroupCoverHeight() *int64
	SetGroupCoverUrl(v string) *FaceGroup
	GetGroupCoverUrl() *string
	SetGroupCoverWidth(v int64) *FaceGroup
	GetGroupCoverWidth() *int64
	SetGroupId(v string) *FaceGroup
	GetGroupId() *string
	SetGroupName(v string) *FaceGroup
	GetGroupName() *string
	SetImageCount(v int64) *FaceGroup
	GetImageCount() *int64
	SetRemarks(v string) *FaceGroup
	GetRemarks() *string
	SetUpdatedAt(v string) *FaceGroup
	GetUpdatedAt() *string
}

type FaceGroup struct {
	// The time when the face-based group was generated.
	//
	// example:
	//
	// 2022-01-14T10:10:52.83948013+08:00
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// The border of the image used as the cover the face-based group.
	GroupCoverFaceBoundary *FaceGroupGroupCoverFaceBoundary `json:"group_cover_face_boundary,omitempty" xml:"group_cover_face_boundary,omitempty" type:"Struct"`
	// The ID of the file used as the cover of the face-based group.
	//
	// example:
	//
	// 6549c959640fbd517c9b4d93b3b36aecc45xxxxx
	GroupCoverFileId *string `json:"group_cover_file_id,omitempty" xml:"group_cover_file_id,omitempty"`
	// The height of the image used as the cover of the face-based group.
	//
	// example:
	//
	// 1080
	GroupCoverHeight *int64 `json:"group_cover_height,omitempty" xml:"group_cover_height,omitempty"`
	// The URL of the image used as the cover of the face-based group.
	//
	// example:
	//
	// https://xxx
	GroupCoverUrl *string `json:"group_cover_url,omitempty" xml:"group_cover_url,omitempty"`
	// The width of the image used as the cover of the face-based group.
	//
	// example:
	//
	// 1920
	GroupCoverWidth *int64 `json:"group_cover_width,omitempty" xml:"group_cover_width,omitempty"`
	// The ID of the face-based group.
	//
	// example:
	//
	// Cluster-ae6e3472-999e-410b-b54e-cd5dba****
	GroupId *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
	// The name of the face-based group.
	//
	// example:
	//
	// name
	GroupName *string `json:"group_name,omitempty" xml:"group_name,omitempty"`
	// The number of photos in the face-based group.
	//
	// example:
	//
	// 10
	ImageCount *int64 `json:"image_count,omitempty" xml:"image_count,omitempty"`
	// The remarks.
	Remarks *string `json:"remarks,omitempty" xml:"remarks,omitempty"`
	// The time when the face-based group was modified.
	//
	// example:
	//
	// 2022-01-14T10:10:52.83948013+08:00
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

func (s FaceGroup) String() string {
	return dara.Prettify(s)
}

func (s FaceGroup) GoString() string {
	return s.String()
}

func (s *FaceGroup) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *FaceGroup) GetGroupCoverFaceBoundary() *FaceGroupGroupCoverFaceBoundary {
	return s.GroupCoverFaceBoundary
}

func (s *FaceGroup) GetGroupCoverFileId() *string {
	return s.GroupCoverFileId
}

func (s *FaceGroup) GetGroupCoverHeight() *int64 {
	return s.GroupCoverHeight
}

func (s *FaceGroup) GetGroupCoverUrl() *string {
	return s.GroupCoverUrl
}

func (s *FaceGroup) GetGroupCoverWidth() *int64 {
	return s.GroupCoverWidth
}

func (s *FaceGroup) GetGroupId() *string {
	return s.GroupId
}

func (s *FaceGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *FaceGroup) GetImageCount() *int64 {
	return s.ImageCount
}

func (s *FaceGroup) GetRemarks() *string {
	return s.Remarks
}

func (s *FaceGroup) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *FaceGroup) SetCreatedAt(v string) *FaceGroup {
	s.CreatedAt = &v
	return s
}

func (s *FaceGroup) SetGroupCoverFaceBoundary(v *FaceGroupGroupCoverFaceBoundary) *FaceGroup {
	s.GroupCoverFaceBoundary = v
	return s
}

func (s *FaceGroup) SetGroupCoverFileId(v string) *FaceGroup {
	s.GroupCoverFileId = &v
	return s
}

func (s *FaceGroup) SetGroupCoverHeight(v int64) *FaceGroup {
	s.GroupCoverHeight = &v
	return s
}

func (s *FaceGroup) SetGroupCoverUrl(v string) *FaceGroup {
	s.GroupCoverUrl = &v
	return s
}

func (s *FaceGroup) SetGroupCoverWidth(v int64) *FaceGroup {
	s.GroupCoverWidth = &v
	return s
}

func (s *FaceGroup) SetGroupId(v string) *FaceGroup {
	s.GroupId = &v
	return s
}

func (s *FaceGroup) SetGroupName(v string) *FaceGroup {
	s.GroupName = &v
	return s
}

func (s *FaceGroup) SetImageCount(v int64) *FaceGroup {
	s.ImageCount = &v
	return s
}

func (s *FaceGroup) SetRemarks(v string) *FaceGroup {
	s.Remarks = &v
	return s
}

func (s *FaceGroup) SetUpdatedAt(v string) *FaceGroup {
	s.UpdatedAt = &v
	return s
}

func (s *FaceGroup) Validate() error {
	if s.GroupCoverFaceBoundary != nil {
		if err := s.GroupCoverFaceBoundary.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FaceGroupGroupCoverFaceBoundary struct {
	// The height of the border. Unit: pixel.
	//
	// example:
	//
	// 300
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The distance from the left side of the photo to the border. Unit: pixel.
	//
	// example:
	//
	// 10
	Left *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	// The distance from the top of the photo to the border. Unit: pixel.
	//
	// example:
	//
	// 30
	Top *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	// The width of the border. Unit: pixel.
	//
	// example:
	//
	// 200
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s FaceGroupGroupCoverFaceBoundary) String() string {
	return dara.Prettify(s)
}

func (s FaceGroupGroupCoverFaceBoundary) GoString() string {
	return s.String()
}

func (s *FaceGroupGroupCoverFaceBoundary) GetHeight() *int32 {
	return s.Height
}

func (s *FaceGroupGroupCoverFaceBoundary) GetLeft() *int32 {
	return s.Left
}

func (s *FaceGroupGroupCoverFaceBoundary) GetTop() *int32 {
	return s.Top
}

func (s *FaceGroupGroupCoverFaceBoundary) GetWidth() *int32 {
	return s.Width
}

func (s *FaceGroupGroupCoverFaceBoundary) SetHeight(v int32) *FaceGroupGroupCoverFaceBoundary {
	s.Height = &v
	return s
}

func (s *FaceGroupGroupCoverFaceBoundary) SetLeft(v int32) *FaceGroupGroupCoverFaceBoundary {
	s.Left = &v
	return s
}

func (s *FaceGroupGroupCoverFaceBoundary) SetTop(v int32) *FaceGroupGroupCoverFaceBoundary {
	s.Top = &v
	return s
}

func (s *FaceGroupGroupCoverFaceBoundary) SetWidth(v int32) *FaceGroupGroupCoverFaceBoundary {
	s.Width = &v
	return s
}

func (s *FaceGroupGroupCoverFaceBoundary) Validate() error {
	return dara.Validate(s)
}
