// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEaiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetElasticAcceleratedInstanceId(v string) *DeleteEaiRequest
	GetElasticAcceleratedInstanceId() *string
	SetForce(v bool) *DeleteEaiRequest
	GetForce() *bool
	SetRegionId(v string) *DeleteEaiRequest
	GetRegionId() *string
}

type DeleteEaiRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eais-sz8t15a7gt7z7j7i****
	ElasticAcceleratedInstanceId *string `json:"ElasticAcceleratedInstanceId,omitempty" xml:"ElasticAcceleratedInstanceId,omitempty"`
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteEaiRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEaiRequest) GoString() string {
	return s.String()
}

func (s *DeleteEaiRequest) GetElasticAcceleratedInstanceId() *string {
	return s.ElasticAcceleratedInstanceId
}

func (s *DeleteEaiRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteEaiRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteEaiRequest) SetElasticAcceleratedInstanceId(v string) *DeleteEaiRequest {
	s.ElasticAcceleratedInstanceId = &v
	return s
}

func (s *DeleteEaiRequest) SetForce(v bool) *DeleteEaiRequest {
	s.Force = &v
	return s
}

func (s *DeleteEaiRequest) SetRegionId(v string) *DeleteEaiRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteEaiRequest) Validate() error {
	return dara.Validate(s)
}
