// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachEaisEiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientInstanceId(v string) *AttachEaisEiRequest
	GetClientInstanceId() *string
	SetEiInstanceId(v string) *AttachEaisEiRequest
	GetEiInstanceId() *string
	SetEiInstanceType(v string) *AttachEaisEiRequest
	GetEiInstanceType() *string
	SetRegionId(v string) *AttachEaisEiRequest
	GetRegionId() *string
}

type AttachEaisEiRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// i-bp14ws9hbt1oe0u9****
	ClientInstanceId *string `json:"ClientInstanceId,omitempty" xml:"ClientInstanceId,omitempty"`
	// example:
	//
	// eais-hzu00xufs1c8j5nn****
	EiInstanceId *string `json:"EiInstanceId,omitempty" xml:"EiInstanceId,omitempty"`
	// example:
	//
	// eais.ei-a6.2xlarge
	EiInstanceType *string `json:"EiInstanceType,omitempty" xml:"EiInstanceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AttachEaisEiRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachEaisEiRequest) GoString() string {
	return s.String()
}

func (s *AttachEaisEiRequest) GetClientInstanceId() *string {
	return s.ClientInstanceId
}

func (s *AttachEaisEiRequest) GetEiInstanceId() *string {
	return s.EiInstanceId
}

func (s *AttachEaisEiRequest) GetEiInstanceType() *string {
	return s.EiInstanceType
}

func (s *AttachEaisEiRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachEaisEiRequest) SetClientInstanceId(v string) *AttachEaisEiRequest {
	s.ClientInstanceId = &v
	return s
}

func (s *AttachEaisEiRequest) SetEiInstanceId(v string) *AttachEaisEiRequest {
	s.EiInstanceId = &v
	return s
}

func (s *AttachEaisEiRequest) SetEiInstanceType(v string) *AttachEaisEiRequest {
	s.EiInstanceType = &v
	return s
}

func (s *AttachEaisEiRequest) SetRegionId(v string) *AttachEaisEiRequest {
	s.RegionId = &v
	return s
}

func (s *AttachEaisEiRequest) Validate() error {
	return dara.Validate(s)
}
