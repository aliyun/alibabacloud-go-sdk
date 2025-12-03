// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEaisEiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEiInstanceId(v string) *DeleteEaisEiRequest
	GetEiInstanceId() *string
	SetForce(v bool) *DeleteEaisEiRequest
	GetForce() *bool
	SetRegionId(v string) *DeleteEaisEiRequest
	GetRegionId() *string
}

type DeleteEaisEiRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eais-hzu00xufs1c8j5nn****
	EiInstanceId *string `json:"EiInstanceId,omitempty" xml:"EiInstanceId,omitempty"`
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteEaisEiRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEaisEiRequest) GoString() string {
	return s.String()
}

func (s *DeleteEaisEiRequest) GetEiInstanceId() *string {
	return s.EiInstanceId
}

func (s *DeleteEaisEiRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteEaisEiRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteEaisEiRequest) SetEiInstanceId(v string) *DeleteEaisEiRequest {
	s.EiInstanceId = &v
	return s
}

func (s *DeleteEaisEiRequest) SetForce(v bool) *DeleteEaisEiRequest {
	s.Force = &v
	return s
}

func (s *DeleteEaisEiRequest) SetRegionId(v string) *DeleteEaisEiRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteEaisEiRequest) Validate() error {
	return dara.Validate(s)
}
