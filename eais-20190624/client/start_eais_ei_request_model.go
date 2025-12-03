// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartEaisEiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEiInstanceId(v string) *StartEaisEiRequest
	GetEiInstanceId() *string
	SetRegionId(v string) *StartEaisEiRequest
	GetRegionId() *string
}

type StartEaisEiRequest struct {
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

func (s StartEaisEiRequest) String() string {
	return dara.Prettify(s)
}

func (s StartEaisEiRequest) GoString() string {
	return s.String()
}

func (s *StartEaisEiRequest) GetEiInstanceId() *string {
	return s.EiInstanceId
}

func (s *StartEaisEiRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartEaisEiRequest) SetEiInstanceId(v string) *StartEaisEiRequest {
	s.EiInstanceId = &v
	return s
}

func (s *StartEaisEiRequest) SetRegionId(v string) *StartEaisEiRequest {
	s.RegionId = &v
	return s
}

func (s *StartEaisEiRequest) Validate() error {
	return dara.Validate(s)
}
