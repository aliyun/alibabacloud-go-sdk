// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKillProcessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComputingGroupId(v string) *KillProcessRequest
	GetComputingGroupId() *string
	SetDBInstanceId(v string) *KillProcessRequest
	GetDBInstanceId() *string
	SetInitialQueryId(v string) *KillProcessRequest
	GetInitialQueryId() *string
	SetRegionId(v string) *KillProcessRequest
	GetRegionId() *string
}

type KillProcessRequest struct {
	ComputingGroupId *string `json:"ComputingGroupId,omitempty" xml:"ComputingGroupId,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The query ID.
	//
	// example:
	//
	// 1
	InitialQueryId *string `json:"InitialQueryId,omitempty" xml:"InitialQueryId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s KillProcessRequest) String() string {
	return dara.Prettify(s)
}

func (s KillProcessRequest) GoString() string {
	return s.String()
}

func (s *KillProcessRequest) GetComputingGroupId() *string {
	return s.ComputingGroupId
}

func (s *KillProcessRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *KillProcessRequest) GetInitialQueryId() *string {
	return s.InitialQueryId
}

func (s *KillProcessRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *KillProcessRequest) SetComputingGroupId(v string) *KillProcessRequest {
	s.ComputingGroupId = &v
	return s
}

func (s *KillProcessRequest) SetDBInstanceId(v string) *KillProcessRequest {
	s.DBInstanceId = &v
	return s
}

func (s *KillProcessRequest) SetInitialQueryId(v string) *KillProcessRequest {
	s.InitialQueryId = &v
	return s
}

func (s *KillProcessRequest) SetRegionId(v string) *KillProcessRequest {
	s.RegionId = &v
	return s
}

func (s *KillProcessRequest) Validate() error {
	return dara.Validate(s)
}
