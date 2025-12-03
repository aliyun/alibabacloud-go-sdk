// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachEaiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientInstanceId(v string) *AttachEaiRequest
	GetClientInstanceId() *string
	SetElasticAcceleratedInstanceId(v string) *AttachEaiRequest
	GetElasticAcceleratedInstanceId() *string
	SetRegionId(v string) *AttachEaiRequest
	GetRegionId() *string
}

type AttachEaiRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// i-wz93g6pyat2g7t7o****
	ClientInstanceId *string `json:"ClientInstanceId,omitempty" xml:"ClientInstanceId,omitempty"`
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

func (s AttachEaiRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachEaiRequest) GoString() string {
	return s.String()
}

func (s *AttachEaiRequest) GetClientInstanceId() *string {
	return s.ClientInstanceId
}

func (s *AttachEaiRequest) GetElasticAcceleratedInstanceId() *string {
	return s.ElasticAcceleratedInstanceId
}

func (s *AttachEaiRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachEaiRequest) SetClientInstanceId(v string) *AttachEaiRequest {
	s.ClientInstanceId = &v
	return s
}

func (s *AttachEaiRequest) SetElasticAcceleratedInstanceId(v string) *AttachEaiRequest {
	s.ElasticAcceleratedInstanceId = &v
	return s
}

func (s *AttachEaiRequest) SetRegionId(v string) *AttachEaiRequest {
	s.RegionId = &v
	return s
}

func (s *AttachEaiRequest) Validate() error {
	return dara.Validate(s)
}
