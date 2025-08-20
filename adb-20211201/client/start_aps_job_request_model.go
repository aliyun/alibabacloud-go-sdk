// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartApsJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApsJobId(v string) *StartApsJobRequest
	GetApsJobId() *string
	SetRegionId(v string) *StartApsJobRequest
	GetRegionId() *string
}

type StartApsJobRequest struct {
	// The job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// aps-******
	ApsJobId *string `json:"ApsJobId,omitempty" xml:"ApsJobId,omitempty"`
	// The region ID
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StartApsJobRequest) String() string {
	return dara.Prettify(s)
}

func (s StartApsJobRequest) GoString() string {
	return s.String()
}

func (s *StartApsJobRequest) GetApsJobId() *string {
	return s.ApsJobId
}

func (s *StartApsJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartApsJobRequest) SetApsJobId(v string) *StartApsJobRequest {
	s.ApsJobId = &v
	return s
}

func (s *StartApsJobRequest) SetRegionId(v string) *StartApsJobRequest {
	s.RegionId = &v
	return s
}

func (s *StartApsJobRequest) Validate() error {
	return dara.Validate(s)
}
