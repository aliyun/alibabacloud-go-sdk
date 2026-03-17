// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAdvancedMonitorStateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetAdvancedMonitorStateRequest
	GetRegionId() *string
	SetSagId(v string) *GetAdvancedMonitorStateRequest
	GetSagId() *string
}

type GetAdvancedMonitorStateRequest struct {
	// The region ID of the SAG instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/69813.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-asdfz6ac74oj5v****
	SagId *string `json:"SagId,omitempty" xml:"SagId,omitempty"`
}

func (s GetAdvancedMonitorStateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAdvancedMonitorStateRequest) GoString() string {
	return s.String()
}

func (s *GetAdvancedMonitorStateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAdvancedMonitorStateRequest) GetSagId() *string {
	return s.SagId
}

func (s *GetAdvancedMonitorStateRequest) SetRegionId(v string) *GetAdvancedMonitorStateRequest {
	s.RegionId = &v
	return s
}

func (s *GetAdvancedMonitorStateRequest) SetSagId(v string) *GetAdvancedMonitorStateRequest {
	s.SagId = &v
	return s
}

func (s *GetAdvancedMonitorStateRequest) Validate() error {
	return dara.Validate(s)
}
