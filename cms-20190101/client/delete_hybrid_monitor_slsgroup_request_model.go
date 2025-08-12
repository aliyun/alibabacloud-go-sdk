// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHybridMonitorSLSGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeleteHybridMonitorSLSGroupRequest
	GetRegionId() *string
	SetSLSGroupName(v string) *DeleteHybridMonitorSLSGroupRequest
	GetSLSGroupName() *string
}

type DeleteHybridMonitorSLSGroupRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the Logstore group.
	//
	// For information about how to obtain the name of a Logstore group, see [DescribeHybridMonitorSLSGroup](https://help.aliyun.com/document_detail/429526.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// Logstore_test
	SLSGroupName *string `json:"SLSGroupName,omitempty" xml:"SLSGroupName,omitempty"`
}

func (s DeleteHybridMonitorSLSGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHybridMonitorSLSGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteHybridMonitorSLSGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteHybridMonitorSLSGroupRequest) GetSLSGroupName() *string {
	return s.SLSGroupName
}

func (s *DeleteHybridMonitorSLSGroupRequest) SetRegionId(v string) *DeleteHybridMonitorSLSGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteHybridMonitorSLSGroupRequest) SetSLSGroupName(v string) *DeleteHybridMonitorSLSGroupRequest {
	s.SLSGroupName = &v
	return s
}

func (s *DeleteHybridMonitorSLSGroupRequest) Validate() error {
	return dara.Validate(s)
}
