// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachEaisEiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEiInstanceId(v string) *DetachEaisEiRequest
	GetEiInstanceId() *string
	SetRegionId(v string) *DetachEaisEiRequest
	GetRegionId() *string
}

type DetachEaisEiRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eais-hzu00xufs1c8j5nn****
	EiInstanceId *string `json:"EiInstanceId,omitempty" xml:"EiInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DetachEaisEiRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachEaisEiRequest) GoString() string {
	return s.String()
}

func (s *DetachEaisEiRequest) GetEiInstanceId() *string {
	return s.EiInstanceId
}

func (s *DetachEaisEiRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachEaisEiRequest) SetEiInstanceId(v string) *DetachEaisEiRequest {
	s.EiInstanceId = &v
	return s
}

func (s *DetachEaisEiRequest) SetRegionId(v string) *DetachEaisEiRequest {
	s.RegionId = &v
	return s
}

func (s *DetachEaisEiRequest) Validate() error {
	return dara.Validate(s)
}
