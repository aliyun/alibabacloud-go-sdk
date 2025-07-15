// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveAIStudioRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DeleteLiveAIStudioRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteLiveAIStudioRequest
	GetRegionId() *string
	SetStudioId(v string) *DeleteLiveAIStudioRequest
	GetStudioId() *string
}

type DeleteLiveAIStudioRequest struct {
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the template. You can obtain the ID from the response to the CreateLiveAIStudio operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 369ced1f-c33a-49e5-91da-bdaae3d6c1c2
	StudioId *string `json:"StudioId,omitempty" xml:"StudioId,omitempty"`
}

func (s DeleteLiveAIStudioRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveAIStudioRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveAIStudioRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLiveAIStudioRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLiveAIStudioRequest) GetStudioId() *string {
	return s.StudioId
}

func (s *DeleteLiveAIStudioRequest) SetOwnerId(v int64) *DeleteLiveAIStudioRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLiveAIStudioRequest) SetRegionId(v string) *DeleteLiveAIStudioRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLiveAIStudioRequest) SetStudioId(v string) *DeleteLiveAIStudioRequest {
	s.StudioId = &v
	return s
}

func (s *DeleteLiveAIStudioRequest) Validate() error {
	return dara.Validate(s)
}
