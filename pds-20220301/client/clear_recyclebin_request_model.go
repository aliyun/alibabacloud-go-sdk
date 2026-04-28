// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearRecyclebinRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *ClearRecyclebinRequest
	GetDriveId() *string
}

type ClearRecyclebinRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
}

func (s ClearRecyclebinRequest) String() string {
	return dara.Prettify(s)
}

func (s ClearRecyclebinRequest) GoString() string {
	return s.String()
}

func (s *ClearRecyclebinRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *ClearRecyclebinRequest) SetDriveId(v string) *ClearRecyclebinRequest {
	s.DriveId = &v
	return s
}

func (s *ClearRecyclebinRequest) Validate() error {
	return dara.Validate(s)
}
