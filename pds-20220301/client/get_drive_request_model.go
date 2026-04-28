// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDriveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *GetDriveRequest
	GetDriveId() *string
}

type GetDriveRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
}

func (s GetDriveRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDriveRequest) GoString() string {
	return s.String()
}

func (s *GetDriveRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *GetDriveRequest) SetDriveId(v string) *GetDriveRequest {
	s.DriveId = &v
	return s
}

func (s *GetDriveRequest) Validate() error {
	return dara.Validate(s)
}
