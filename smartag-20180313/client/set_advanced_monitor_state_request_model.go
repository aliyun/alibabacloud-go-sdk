// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAdvancedMonitorStateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v bool) *SetAdvancedMonitorStateRequest
	GetEnable() *bool
	SetRegionId(v string) *SetAdvancedMonitorStateRequest
	GetRegionId() *string
	SetSagId(v string) *SetAdvancedMonitorStateRequest
	GetSagId() *string
}

type SetAdvancedMonitorStateRequest struct {
	// Specifies whether to enable the DPI feature. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
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
	// sag-****
	SagId *string `json:"SagId,omitempty" xml:"SagId,omitempty"`
}

func (s SetAdvancedMonitorStateRequest) String() string {
	return dara.Prettify(s)
}

func (s SetAdvancedMonitorStateRequest) GoString() string {
	return s.String()
}

func (s *SetAdvancedMonitorStateRequest) GetEnable() *bool {
	return s.Enable
}

func (s *SetAdvancedMonitorStateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetAdvancedMonitorStateRequest) GetSagId() *string {
	return s.SagId
}

func (s *SetAdvancedMonitorStateRequest) SetEnable(v bool) *SetAdvancedMonitorStateRequest {
	s.Enable = &v
	return s
}

func (s *SetAdvancedMonitorStateRequest) SetRegionId(v string) *SetAdvancedMonitorStateRequest {
	s.RegionId = &v
	return s
}

func (s *SetAdvancedMonitorStateRequest) SetSagId(v string) *SetAdvancedMonitorStateRequest {
	s.SagId = &v
	return s
}

func (s *SetAdvancedMonitorStateRequest) Validate() error {
	return dara.Validate(s)
}
