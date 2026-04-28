// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *CreateStoryResponseBody
	GetDriveId() *string
}

type CreateStoryResponseBody struct {
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
}

func (s CreateStoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateStoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStoryResponseBody) GetDriveId() *string {
	return s.DriveId
}

func (s *CreateStoryResponseBody) SetDriveId(v string) *CreateStoryResponseBody {
	s.DriveId = &v
	return s
}

func (s *CreateStoryResponseBody) Validate() error {
	return dara.Validate(s)
}
