// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudProcessMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProcessMonitors(v []*DescribeHybridCloudProcessMonitorResponseBodyProcessMonitors) *DescribeHybridCloudProcessMonitorResponseBody
	GetProcessMonitors() []*DescribeHybridCloudProcessMonitorResponseBodyProcessMonitors
	SetRequestId(v string) *DescribeHybridCloudProcessMonitorResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeHybridCloudProcessMonitorResponseBody
	GetTotalCount() *int32
}

type DescribeHybridCloudProcessMonitorResponseBody struct {
	// The application status list.
	ProcessMonitors []*DescribeHybridCloudProcessMonitorResponseBodyProcessMonitors `json:"ProcessMonitors,omitempty" xml:"ProcessMonitors,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// DBF79169-B6A0-5C8E-86B2-CFE3****496E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- This parameter is deprecated and no longer returns meaningful data.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeHybridCloudProcessMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudProcessMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudProcessMonitorResponseBody) GetProcessMonitors() []*DescribeHybridCloudProcessMonitorResponseBodyProcessMonitors {
	return s.ProcessMonitors
}

func (s *DescribeHybridCloudProcessMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHybridCloudProcessMonitorResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeHybridCloudProcessMonitorResponseBody) SetProcessMonitors(v []*DescribeHybridCloudProcessMonitorResponseBodyProcessMonitors) *DescribeHybridCloudProcessMonitorResponseBody {
	s.ProcessMonitors = v
	return s
}

func (s *DescribeHybridCloudProcessMonitorResponseBody) SetRequestId(v string) *DescribeHybridCloudProcessMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHybridCloudProcessMonitorResponseBody) SetTotalCount(v int32) *DescribeHybridCloudProcessMonitorResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeHybridCloudProcessMonitorResponseBody) Validate() error {
	if s.ProcessMonitors != nil {
		for _, item := range s.ProcessMonitors {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeHybridCloudProcessMonitorResponseBodyProcessMonitors struct {
	// The status. Valid values:
	//
	// - **normal**: Normal.
	//
	// - **critical**: abnormal.
	//
	// example:
	//
	// normal
	Levle *string `json:"Levle,omitempty" xml:"Levle,omitempty"`
	// The service. Valid values:
	//
	// - **tianqingproxy**: centralized management service.
	//
	// - **redis**: storage service.
	//
	// - **scc**: traffic computing service.
	//
	// - **keeper**: threat intelligence service.
	//
	// - **node_exporter**: application log upload service.
	//
	// - **xagent**: traffic detection service.
	//
	// - **noproxy**: traffic forwarding service.
	//
	// - **xloge**: attack log upload service.
	//
	// - **ilogtail**: log collection service.
	//
	// - **xlogd**: log analysis feature.
	//
	// example:
	//
	// keeper
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	// The running status. Valid values:
	//
	// - **0**: abnormal.
	//
	// - **1**: Normal.
	//
	// example:
	//
	// 1
	ProcessStatus *int64 `json:"ProcessStatus,omitempty" xml:"ProcessStatus,omitempty"`
}

func (s DescribeHybridCloudProcessMonitorResponseBodyProcessMonitors) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudProcessMonitorResponseBodyProcessMonitors) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudProcessMonitorResponseBodyProcessMonitors) GetLevle() *string {
	return s.Levle
}

func (s *DescribeHybridCloudProcessMonitorResponseBodyProcessMonitors) GetProcessName() *string {
	return s.ProcessName
}

func (s *DescribeHybridCloudProcessMonitorResponseBodyProcessMonitors) GetProcessStatus() *int64 {
	return s.ProcessStatus
}

func (s *DescribeHybridCloudProcessMonitorResponseBodyProcessMonitors) SetLevle(v string) *DescribeHybridCloudProcessMonitorResponseBodyProcessMonitors {
	s.Levle = &v
	return s
}

func (s *DescribeHybridCloudProcessMonitorResponseBodyProcessMonitors) SetProcessName(v string) *DescribeHybridCloudProcessMonitorResponseBodyProcessMonitors {
	s.ProcessName = &v
	return s
}

func (s *DescribeHybridCloudProcessMonitorResponseBodyProcessMonitors) SetProcessStatus(v int64) *DescribeHybridCloudProcessMonitorResponseBodyProcessMonitors {
	s.ProcessStatus = &v
	return s
}

func (s *DescribeHybridCloudProcessMonitorResponseBodyProcessMonitors) Validate() error {
	return dara.Validate(s)
}
