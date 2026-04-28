// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDriveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *DeleteDriveRequest
	GetDriveId() *string
}

type DeleteDriveRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
}

func (s DeleteDriveRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDriveRequest) GoString() string {
	return s.String()
}

func (s *DeleteDriveRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *DeleteDriveRequest) SetDriveId(v string) *DeleteDriveRequest {
	s.DriveId = &v
	return s
}

func (s *DeleteDriveRequest) Validate() error {
	return dara.Validate(s)
}
