// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApsJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApsJobId(v string) *ModifyApsJobRequest
	GetApsJobId() *string
	SetDbList(v string) *ModifyApsJobRequest
	GetDbList() *string
	SetPartitionList(v string) *ModifyApsJobRequest
	GetPartitionList() *string
	SetRegionId(v string) *ModifyApsJobRequest
	GetRegionId() *string
}

type ModifyApsJobRequest struct {
	// The job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// aps-bj1xxxxxx
	ApsJobId *string `json:"ApsJobId,omitempty" xml:"ApsJobId,omitempty"`
	// The objects to be synchronized.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"EntireInstance":true}
	DbList *string `json:"DbList,omitempty" xml:"DbList,omitempty"`
	// The partitions.
	//
	// example:
	//
	// {}
	PartitionList *string `json:"PartitionList,omitempty" xml:"PartitionList,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyApsJobRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyApsJobRequest) GoString() string {
	return s.String()
}

func (s *ModifyApsJobRequest) GetApsJobId() *string {
	return s.ApsJobId
}

func (s *ModifyApsJobRequest) GetDbList() *string {
	return s.DbList
}

func (s *ModifyApsJobRequest) GetPartitionList() *string {
	return s.PartitionList
}

func (s *ModifyApsJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyApsJobRequest) SetApsJobId(v string) *ModifyApsJobRequest {
	s.ApsJobId = &v
	return s
}

func (s *ModifyApsJobRequest) SetDbList(v string) *ModifyApsJobRequest {
	s.DbList = &v
	return s
}

func (s *ModifyApsJobRequest) SetPartitionList(v string) *ModifyApsJobRequest {
	s.PartitionList = &v
	return s
}

func (s *ModifyApsJobRequest) SetRegionId(v string) *ModifyApsJobRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyApsJobRequest) Validate() error {
	return dara.Validate(s)
}
