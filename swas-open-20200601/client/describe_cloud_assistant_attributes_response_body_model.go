// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudAssistantAttributesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCloudAssistant(v []*DescribeCloudAssistantAttributesResponseBodyCloudAssistant) *DescribeCloudAssistantAttributesResponseBody
	GetCloudAssistant() []*DescribeCloudAssistantAttributesResponseBodyCloudAssistant
	SetPageNumber(v int32) *DescribeCloudAssistantAttributesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCloudAssistantAttributesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeCloudAssistantAttributesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeCloudAssistantAttributesResponseBody
	GetTotalCount() *int32
}

type DescribeCloudAssistantAttributesResponseBody struct {
	// The installation status of Cloud Assistant on the instances.
	CloudAssistant []*DescribeCloudAssistantAttributesResponseBodyCloudAssistant `json:"CloudAssistant,omitempty" xml:"CloudAssistant,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCloudAssistantAttributesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudAssistantAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudAssistantAttributesResponseBody) GetCloudAssistant() []*DescribeCloudAssistantAttributesResponseBodyCloudAssistant {
	return s.CloudAssistant
}

func (s *DescribeCloudAssistantAttributesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCloudAssistantAttributesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCloudAssistantAttributesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudAssistantAttributesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCloudAssistantAttributesResponseBody) SetCloudAssistant(v []*DescribeCloudAssistantAttributesResponseBodyCloudAssistant) *DescribeCloudAssistantAttributesResponseBody {
	s.CloudAssistant = v
	return s
}

func (s *DescribeCloudAssistantAttributesResponseBody) SetPageNumber(v int32) *DescribeCloudAssistantAttributesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCloudAssistantAttributesResponseBody) SetPageSize(v int32) *DescribeCloudAssistantAttributesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudAssistantAttributesResponseBody) SetRequestId(v string) *DescribeCloudAssistantAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudAssistantAttributesResponseBody) SetTotalCount(v int32) *DescribeCloudAssistantAttributesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCloudAssistantAttributesResponseBody) Validate() error {
	if s.CloudAssistant != nil {
		for _, item := range s.CloudAssistant {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCloudAssistantAttributesResponseBodyCloudAssistant struct {
	// The number of running tasks of Cloud Assistant on the instance.
	//
	// example:
	//
	// 0
	ActiveTaskCount *int64 `json:"ActiveTaskCount,omitempty" xml:"ActiveTaskCount,omitempty"`
	// Indicates whether Cloud Assistant is running on the instance. Valid values:
	//
	// true: A heartbeat was detected within the last 2 minutes.
	//
	// false: No heartbeat was detected within the last 2 minutes.
	//
	// example:
	//
	// true
	CloudAssistantStatus *string `json:"CloudAssistantStatus,omitempty" xml:"CloudAssistantStatus,omitempty"`
	// The version number of the Cloud Assistant Agent. This value is empty if the Cloud Assistant Agent is not installed or is not running.
	//
	// example:
	//
	// 2.2.0.106
	CloudAssistantVersion *string `json:"CloudAssistantVersion,omitempty" xml:"CloudAssistantVersion,omitempty"`
	// The instance ID of the simple application server instance.
	//
	// example:
	//
	// 85dbe3e7cc7b49e1a3df4af3bfa4ebbf
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of completed tasks of Cloud Assistant on the instance.
	//
	// example:
	//
	// 4
	InvocationCount *int64 `json:"InvocationCount,omitempty" xml:"InvocationCount,omitempty"`
	// The time of the last heartbeat of Cloud Assistant. This value is updated approximately every minute, with a random fluctuation of up to 5 seconds (minimum 55 seconds, maximum 65 seconds).
	//
	// example:
	//
	// 2021-03-15T09:00:00Z
	LastHeartbeatTime *string `json:"LastHeartbeatTime,omitempty" xml:"LastHeartbeatTime,omitempty"`
	// The time when the last command was executed. The time is displayed in UTC (Z) in the YYYY-MM-DDThh:mm:ssZ format.
	//
	// example:
	//
	// 2021-03-15T08:00:00Z
	LastInvokedTime *string `json:"LastInvokedTime,omitempty" xml:"LastInvokedTime,omitempty"`
	// The operating system type of the instance. Valid values:
	//
	// - Windows.
	//
	// - Linux.
	//
	// - FreeBSD.
	//
	// example:
	//
	// Linux
	OSType *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	// Indicates whether Cloud Assistant on the instance supports the session management feature. If session management is not supported, the Cloud Assistant version is outdated. Update the client to the latest version.
	//
	// The minimum Cloud Assistant versions that support session management are as follows:
	//
	// Linux: 2.2.3.189
	//
	// Windows: 2.1.3.189
	//
	// example:
	//
	// true
	SupportSessionManager *bool `json:"SupportSessionManager,omitempty" xml:"SupportSessionManager,omitempty"`
}

func (s DescribeCloudAssistantAttributesResponseBodyCloudAssistant) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudAssistantAttributesResponseBodyCloudAssistant) GoString() string {
	return s.String()
}

func (s *DescribeCloudAssistantAttributesResponseBodyCloudAssistant) GetActiveTaskCount() *int64 {
	return s.ActiveTaskCount
}

func (s *DescribeCloudAssistantAttributesResponseBodyCloudAssistant) GetCloudAssistantStatus() *string {
	return s.CloudAssistantStatus
}

func (s *DescribeCloudAssistantAttributesResponseBodyCloudAssistant) GetCloudAssistantVersion() *string {
	return s.CloudAssistantVersion
}

func (s *DescribeCloudAssistantAttributesResponseBodyCloudAssistant) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCloudAssistantAttributesResponseBodyCloudAssistant) GetInvocationCount() *int64 {
	return s.InvocationCount
}

func (s *DescribeCloudAssistantAttributesResponseBodyCloudAssistant) GetLastHeartbeatTime() *string {
	return s.LastHeartbeatTime
}

func (s *DescribeCloudAssistantAttributesResponseBodyCloudAssistant) GetLastInvokedTime() *string {
	return s.LastInvokedTime
}

func (s *DescribeCloudAssistantAttributesResponseBodyCloudAssistant) GetOSType() *string {
	return s.OSType
}

func (s *DescribeCloudAssistantAttributesResponseBodyCloudAssistant) GetSupportSessionManager() *bool {
	return s.SupportSessionManager
}

func (s *DescribeCloudAssistantAttributesResponseBodyCloudAssistant) SetActiveTaskCount(v int64) *DescribeCloudAssistantAttributesResponseBodyCloudAssistant {
	s.ActiveTaskCount = &v
	return s
}

func (s *DescribeCloudAssistantAttributesResponseBodyCloudAssistant) SetCloudAssistantStatus(v string) *DescribeCloudAssistantAttributesResponseBodyCloudAssistant {
	s.CloudAssistantStatus = &v
	return s
}

func (s *DescribeCloudAssistantAttributesResponseBodyCloudAssistant) SetCloudAssistantVersion(v string) *DescribeCloudAssistantAttributesResponseBodyCloudAssistant {
	s.CloudAssistantVersion = &v
	return s
}

func (s *DescribeCloudAssistantAttributesResponseBodyCloudAssistant) SetInstanceId(v string) *DescribeCloudAssistantAttributesResponseBodyCloudAssistant {
	s.InstanceId = &v
	return s
}

func (s *DescribeCloudAssistantAttributesResponseBodyCloudAssistant) SetInvocationCount(v int64) *DescribeCloudAssistantAttributesResponseBodyCloudAssistant {
	s.InvocationCount = &v
	return s
}

func (s *DescribeCloudAssistantAttributesResponseBodyCloudAssistant) SetLastHeartbeatTime(v string) *DescribeCloudAssistantAttributesResponseBodyCloudAssistant {
	s.LastHeartbeatTime = &v
	return s
}

func (s *DescribeCloudAssistantAttributesResponseBodyCloudAssistant) SetLastInvokedTime(v string) *DescribeCloudAssistantAttributesResponseBodyCloudAssistant {
	s.LastInvokedTime = &v
	return s
}

func (s *DescribeCloudAssistantAttributesResponseBodyCloudAssistant) SetOSType(v string) *DescribeCloudAssistantAttributesResponseBodyCloudAssistant {
	s.OSType = &v
	return s
}

func (s *DescribeCloudAssistantAttributesResponseBodyCloudAssistant) SetSupportSessionManager(v bool) *DescribeCloudAssistantAttributesResponseBodyCloudAssistant {
	s.SupportSessionManager = &v
	return s
}

func (s *DescribeCloudAssistantAttributesResponseBodyCloudAssistant) Validate() error {
	return dara.Validate(s)
}
