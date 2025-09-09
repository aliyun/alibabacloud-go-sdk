// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchGlobalBroadcastTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *SwitchGlobalBroadcastTypeRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *SwitchGlobalBroadcastTypeRequest
	GetDrdsInstanceId() *string
	SetRegionId(v string) *SwitchGlobalBroadcastTypeRequest
	GetRegionId() *string
}

type SwitchGlobalBroadcastTypeRequest struct {
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds********
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SwitchGlobalBroadcastTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchGlobalBroadcastTypeRequest) GoString() string {
	return s.String()
}

func (s *SwitchGlobalBroadcastTypeRequest) GetDbName() *string {
	return s.DbName
}

func (s *SwitchGlobalBroadcastTypeRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *SwitchGlobalBroadcastTypeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SwitchGlobalBroadcastTypeRequest) SetDbName(v string) *SwitchGlobalBroadcastTypeRequest {
	s.DbName = &v
	return s
}

func (s *SwitchGlobalBroadcastTypeRequest) SetDrdsInstanceId(v string) *SwitchGlobalBroadcastTypeRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SwitchGlobalBroadcastTypeRequest) SetRegionId(v string) *SwitchGlobalBroadcastTypeRequest {
	s.RegionId = &v
	return s
}

func (s *SwitchGlobalBroadcastTypeRequest) Validate() error {
	return dara.Validate(s)
}
