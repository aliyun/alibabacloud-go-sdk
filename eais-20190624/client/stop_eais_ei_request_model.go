// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopEaisEiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEiInstanceId(v string) *StopEaisEiRequest
	GetEiInstanceId() *string
	SetRegionId(v string) *StopEaisEiRequest
	GetRegionId() *string
}

type StopEaisEiRequest struct {
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

func (s StopEaisEiRequest) String() string {
	return dara.Prettify(s)
}

func (s StopEaisEiRequest) GoString() string {
	return s.String()
}

func (s *StopEaisEiRequest) GetEiInstanceId() *string {
	return s.EiInstanceId
}

func (s *StopEaisEiRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopEaisEiRequest) SetEiInstanceId(v string) *StopEaisEiRequest {
	s.EiInstanceId = &v
	return s
}

func (s *StopEaisEiRequest) SetRegionId(v string) *StopEaisEiRequest {
	s.RegionId = &v
	return s
}

func (s *StopEaisEiRequest) Validate() error {
	return dara.Validate(s)
}
