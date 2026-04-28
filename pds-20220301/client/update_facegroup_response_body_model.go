// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFacegroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *UpdateFacegroupResponseBody
	GetDriveId() *string
	SetGroupId(v string) *UpdateFacegroupResponseBody
	GetGroupId() *string
}

type UpdateFacegroupResponseBody struct {
	// The drive ID.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The group ID.
	//
	// example:
	//
	// group-abc
	GroupId *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
}

func (s UpdateFacegroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFacegroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFacegroupResponseBody) GetDriveId() *string {
	return s.DriveId
}

func (s *UpdateFacegroupResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateFacegroupResponseBody) SetDriveId(v string) *UpdateFacegroupResponseBody {
	s.DriveId = &v
	return s
}

func (s *UpdateFacegroupResponseBody) SetGroupId(v string) *UpdateFacegroupResponseBody {
	s.GroupId = &v
	return s
}

func (s *UpdateFacegroupResponseBody) Validate() error {
	return dara.Validate(s)
}
