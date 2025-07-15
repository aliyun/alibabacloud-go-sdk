// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEduRoomRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEduRoomId(v string) *DeleteEduRoomRequest
	GetEduRoomId() *string
	SetRegionId(v string) *DeleteEduRoomRequest
	GetRegionId() *string
}

type DeleteEduRoomRequest struct {
	// This parameter is required.
	EduRoomId *string `json:"EduRoomId,omitempty" xml:"EduRoomId,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteEduRoomRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEduRoomRequest) GoString() string {
	return s.String()
}

func (s *DeleteEduRoomRequest) GetEduRoomId() *string {
	return s.EduRoomId
}

func (s *DeleteEduRoomRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteEduRoomRequest) SetEduRoomId(v string) *DeleteEduRoomRequest {
	s.EduRoomId = &v
	return s
}

func (s *DeleteEduRoomRequest) SetRegionId(v string) *DeleteEduRoomRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteEduRoomRequest) Validate() error {
	return dara.Validate(s)
}
