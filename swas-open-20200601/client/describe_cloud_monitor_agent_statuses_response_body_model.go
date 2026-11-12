// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudMonitorAgentStatusesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceStatusList(v []*DescribeCloudMonitorAgentStatusesResponseBodyInstanceStatusList) *DescribeCloudMonitorAgentStatusesResponseBody
	GetInstanceStatusList() []*DescribeCloudMonitorAgentStatusesResponseBodyInstanceStatusList
	SetRequestId(v string) *DescribeCloudMonitorAgentStatusesResponseBody
	GetRequestId() *string
}

type DescribeCloudMonitorAgentStatusesResponseBody struct {
	// Indicates whether the Cloud Monitor agent was automatically installed on the simple application server.
	InstanceStatusList []*DescribeCloudMonitorAgentStatusesResponseBodyInstanceStatusList `json:"InstanceStatusList,omitempty" xml:"InstanceStatusList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E1FEE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCloudMonitorAgentStatusesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudMonitorAgentStatusesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudMonitorAgentStatusesResponseBody) GetInstanceStatusList() []*DescribeCloudMonitorAgentStatusesResponseBodyInstanceStatusList {
	return s.InstanceStatusList
}

func (s *DescribeCloudMonitorAgentStatusesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudMonitorAgentStatusesResponseBody) SetInstanceStatusList(v []*DescribeCloudMonitorAgentStatusesResponseBodyInstanceStatusList) *DescribeCloudMonitorAgentStatusesResponseBody {
	s.InstanceStatusList = v
	return s
}

func (s *DescribeCloudMonitorAgentStatusesResponseBody) SetRequestId(v string) *DescribeCloudMonitorAgentStatusesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudMonitorAgentStatusesResponseBody) Validate() error {
	if s.InstanceStatusList != nil {
		for _, item := range s.InstanceStatusList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCloudMonitorAgentStatusesResponseBodyInstanceStatusList struct {
	// Indicates whether the Cloud Monitor agent was automatically installed on the simple application server. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// true
	AutoInstall *bool `json:"AutoInstall,omitempty" xml:"AutoInstall,omitempty"`
	// The ID of the simple application server.
	//
	// example:
	//
	// c854dc6f07e74953830bb5808d0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The running status of the Cloud Monitoring plug-in. Possible values are:
	//
	// - running: Cloud Monitoring plug-in running.
	//
	// - stopped: Cloud Monitoring plug-in stopped.
	//
	// - installing: Cloud Monitoring plug-in installing.
	//
	// - install_faild: Cloud Monitoring plug-in installation failed.
	//
	// - abnormal: Cloud Monitoring plug-in installation abnormal.
	//
	// - not_installed: Cloud Monitoring plug-in not installed.
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCloudMonitorAgentStatusesResponseBodyInstanceStatusList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudMonitorAgentStatusesResponseBodyInstanceStatusList) GoString() string {
	return s.String()
}

func (s *DescribeCloudMonitorAgentStatusesResponseBodyInstanceStatusList) GetAutoInstall() *bool {
	return s.AutoInstall
}

func (s *DescribeCloudMonitorAgentStatusesResponseBodyInstanceStatusList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCloudMonitorAgentStatusesResponseBodyInstanceStatusList) GetStatus() *string {
	return s.Status
}

func (s *DescribeCloudMonitorAgentStatusesResponseBodyInstanceStatusList) SetAutoInstall(v bool) *DescribeCloudMonitorAgentStatusesResponseBodyInstanceStatusList {
	s.AutoInstall = &v
	return s
}

func (s *DescribeCloudMonitorAgentStatusesResponseBodyInstanceStatusList) SetInstanceId(v string) *DescribeCloudMonitorAgentStatusesResponseBodyInstanceStatusList {
	s.InstanceId = &v
	return s
}

func (s *DescribeCloudMonitorAgentStatusesResponseBodyInstanceStatusList) SetStatus(v string) *DescribeCloudMonitorAgentStatusesResponseBodyInstanceStatusList {
	s.Status = &v
	return s
}

func (s *DescribeCloudMonitorAgentStatusesResponseBodyInstanceStatusList) Validate() error {
	return dara.Validate(s)
}
