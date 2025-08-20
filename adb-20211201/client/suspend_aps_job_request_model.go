// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendApsJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApsJobId(v string) *SuspendApsJobRequest
	GetApsJobId() *string
	SetRegionId(v string) *SuspendApsJobRequest
	GetRegionId() *string
}

type SuspendApsJobRequest struct {
	// The job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// aps-bj1xxxxxx
	ApsJobId *string `json:"ApsJobId,omitempty" xml:"ApsJobId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SuspendApsJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SuspendApsJobRequest) GoString() string {
	return s.String()
}

func (s *SuspendApsJobRequest) GetApsJobId() *string {
	return s.ApsJobId
}

func (s *SuspendApsJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SuspendApsJobRequest) SetApsJobId(v string) *SuspendApsJobRequest {
	s.ApsJobId = &v
	return s
}

func (s *SuspendApsJobRequest) SetRegionId(v string) *SuspendApsJobRequest {
	s.RegionId = &v
	return s
}

func (s *SuspendApsJobRequest) Validate() error {
	return dara.Validate(s)
}
