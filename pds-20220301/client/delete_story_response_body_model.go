// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *DeleteStoryResponseBody
	GetDriveId() *string
}

type DeleteStoryResponseBody struct {
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
}

func (s DeleteStoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteStoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStoryResponseBody) GetDriveId() *string {
	return s.DriveId
}

func (s *DeleteStoryResponseBody) SetDriveId(v string) *DeleteStoryResponseBody {
	s.DriveId = &v
	return s
}

func (s *DeleteStoryResponseBody) Validate() error {
	return dara.Validate(s)
}
