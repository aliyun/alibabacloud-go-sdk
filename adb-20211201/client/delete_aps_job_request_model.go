// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApsJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApsJobId(v string) *DeleteApsJobRequest
	GetApsJobId() *string
	SetRegionId(v string) *DeleteApsJobRequest
	GetRegionId() *string
}

type DeleteApsJobRequest struct {
	// The job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// aps-*****
	ApsJobId *string `json:"ApsJobId,omitempty" xml:"ApsJobId,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteApsJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteApsJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteApsJobRequest) GetApsJobId() *string {
	return s.ApsJobId
}

func (s *DeleteApsJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteApsJobRequest) SetApsJobId(v string) *DeleteApsJobRequest {
	s.ApsJobId = &v
	return s
}

func (s *DeleteApsJobRequest) SetRegionId(v string) *DeleteApsJobRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteApsJobRequest) Validate() error {
	return dara.Validate(s)
}
