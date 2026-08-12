// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudBasicMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBasicMonitors(v []*DescribeHybridCloudBasicMonitorResponseBodyBasicMonitors) *DescribeHybridCloudBasicMonitorResponseBody
	GetBasicMonitors() []*DescribeHybridCloudBasicMonitorResponseBodyBasicMonitors
	SetRequestId(v string) *DescribeHybridCloudBasicMonitorResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeHybridCloudBasicMonitorResponseBody
	GetTotalCount() *int32
}

type DescribeHybridCloudBasicMonitorResponseBody struct {
	// The list of basic monitoring items.
	BasicMonitors []*DescribeHybridCloudBasicMonitorResponseBodyBasicMonitors `json:"BasicMonitors,omitempty" xml:"BasicMonitors,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0B8AF42B-16A9-5762-AEF3-D148****FE5D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Deprecated
	//
	// > This parameter is deprecated and no longer returns meaningful data.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeHybridCloudBasicMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudBasicMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudBasicMonitorResponseBody) GetBasicMonitors() []*DescribeHybridCloudBasicMonitorResponseBodyBasicMonitors {
	return s.BasicMonitors
}

func (s *DescribeHybridCloudBasicMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHybridCloudBasicMonitorResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeHybridCloudBasicMonitorResponseBody) SetBasicMonitors(v []*DescribeHybridCloudBasicMonitorResponseBodyBasicMonitors) *DescribeHybridCloudBasicMonitorResponseBody {
	s.BasicMonitors = v
	return s
}

func (s *DescribeHybridCloudBasicMonitorResponseBody) SetRequestId(v string) *DescribeHybridCloudBasicMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHybridCloudBasicMonitorResponseBody) SetTotalCount(v int32) *DescribeHybridCloudBasicMonitorResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeHybridCloudBasicMonitorResponseBody) Validate() error {
	if s.BasicMonitors != nil {
		for _, item := range s.BasicMonitors {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeHybridCloudBasicMonitorResponseBodyBasicMonitors struct {
	// The status. Valid values:
	//
	// - **normal**: Normal.
	//
	// - `****`: abnormal (four asterisk characters are returned).
	//
	// example:
	//
	// normal
	Levle *string `json:"Levle,omitempty" xml:"Levle,omitempty"`
	// The monitoring item. Valid values:
	//
	// - **basic_monitor_cpu_usage**: CPU.
	//
	// - **basic_monitor_memory_usage**: memory.
	//
	// - **basic_monitor_disk_usage**: disk.
	//
	// example:
	//
	// basic_monitor_cpu_usage
	MonitorName *string `json:"MonitorName,omitempty" xml:"MonitorName,omitempty"`
	// The usage ratio.
	//
	// example:
	//
	// 20
	UseRatio *int64 `json:"UseRatio,omitempty" xml:"UseRatio,omitempty"`
}

func (s DescribeHybridCloudBasicMonitorResponseBodyBasicMonitors) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudBasicMonitorResponseBodyBasicMonitors) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudBasicMonitorResponseBodyBasicMonitors) GetLevle() *string {
	return s.Levle
}

func (s *DescribeHybridCloudBasicMonitorResponseBodyBasicMonitors) GetMonitorName() *string {
	return s.MonitorName
}

func (s *DescribeHybridCloudBasicMonitorResponseBodyBasicMonitors) GetUseRatio() *int64 {
	return s.UseRatio
}

func (s *DescribeHybridCloudBasicMonitorResponseBodyBasicMonitors) SetLevle(v string) *DescribeHybridCloudBasicMonitorResponseBodyBasicMonitors {
	s.Levle = &v
	return s
}

func (s *DescribeHybridCloudBasicMonitorResponseBodyBasicMonitors) SetMonitorName(v string) *DescribeHybridCloudBasicMonitorResponseBodyBasicMonitors {
	s.MonitorName = &v
	return s
}

func (s *DescribeHybridCloudBasicMonitorResponseBodyBasicMonitors) SetUseRatio(v int64) *DescribeHybridCloudBasicMonitorResponseBodyBasicMonitors {
	s.UseRatio = &v
	return s
}

func (s *DescribeHybridCloudBasicMonitorResponseBodyBasicMonitors) Validate() error {
	return dara.Validate(s)
}
