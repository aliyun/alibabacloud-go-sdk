// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachEaiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetElasticAcceleratedInstanceId(v string) *DetachEaiRequest
	GetElasticAcceleratedInstanceId() *string
	SetRegionId(v string) *DetachEaiRequest
	GetRegionId() *string
}

type DetachEaiRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eais-sz8t15a7gt7z7j7i****
	ElasticAcceleratedInstanceId *string `json:"ElasticAcceleratedInstanceId,omitempty" xml:"ElasticAcceleratedInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DetachEaiRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachEaiRequest) GoString() string {
	return s.String()
}

func (s *DetachEaiRequest) GetElasticAcceleratedInstanceId() *string {
	return s.ElasticAcceleratedInstanceId
}

func (s *DetachEaiRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachEaiRequest) SetElasticAcceleratedInstanceId(v string) *DetachEaiRequest {
	s.ElasticAcceleratedInstanceId = &v
	return s
}

func (s *DetachEaiRequest) SetRegionId(v string) *DetachEaiRequest {
	s.RegionId = &v
	return s
}

func (s *DetachEaiRequest) Validate() error {
	return dara.Validate(s)
}
