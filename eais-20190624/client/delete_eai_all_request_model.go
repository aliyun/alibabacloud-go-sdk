// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEaiAllRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientInstanceId(v string) *DeleteEaiAllRequest
	GetClientInstanceId() *string
	SetElasticAcceleratedInstanceId(v string) *DeleteEaiAllRequest
	GetElasticAcceleratedInstanceId() *string
	SetRegionId(v string) *DeleteEaiAllRequest
	GetRegionId() *string
}

type DeleteEaiAllRequest struct {
	// example:
	//
	// i-bp1fvhi60e1zizsp****
	ClientInstanceId *string `json:"ClientInstanceId,omitempty" xml:"ClientInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eais-hza1ahi0uuw0re33****
	ElasticAcceleratedInstanceId *string `json:"ElasticAcceleratedInstanceId,omitempty" xml:"ElasticAcceleratedInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteEaiAllRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEaiAllRequest) GoString() string {
	return s.String()
}

func (s *DeleteEaiAllRequest) GetClientInstanceId() *string {
	return s.ClientInstanceId
}

func (s *DeleteEaiAllRequest) GetElasticAcceleratedInstanceId() *string {
	return s.ElasticAcceleratedInstanceId
}

func (s *DeleteEaiAllRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteEaiAllRequest) SetClientInstanceId(v string) *DeleteEaiAllRequest {
	s.ClientInstanceId = &v
	return s
}

func (s *DeleteEaiAllRequest) SetElasticAcceleratedInstanceId(v string) *DeleteEaiAllRequest {
	s.ElasticAcceleratedInstanceId = &v
	return s
}

func (s *DeleteEaiAllRequest) SetRegionId(v string) *DeleteEaiAllRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteEaiAllRequest) Validate() error {
	return dara.Validate(s)
}
