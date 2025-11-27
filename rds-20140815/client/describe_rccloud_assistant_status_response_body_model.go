// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCCloudAssistantStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceCloudAssistantStatusSet(v []*DescribeRCCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet) *DescribeRCCloudAssistantStatusResponseBody
	GetInstanceCloudAssistantStatusSet() []*DescribeRCCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet
	SetNextToken(v string) *DescribeRCCloudAssistantStatusResponseBody
	GetNextToken() *string
	SetPageNumber(v string) *DescribeRCCloudAssistantStatusResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *DescribeRCCloudAssistantStatusResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeRCCloudAssistantStatusResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeRCCloudAssistantStatusResponseBody
	GetTotalCount() *int32
}

type DescribeRCCloudAssistantStatusResponseBody struct {
	// Details about the installation status of Cloud Assistant on the instances.
	InstanceCloudAssistantStatusSet []*DescribeRCCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet `json:"InstanceCloudAssistantStatusSet,omitempty" xml:"InstanceCloudAssistantStatusSet,omitempty" type:"Repeated"`
	// The token that marks the end of the current returned page. If this parameter is empty, all data is retrieved.
	//
	// This parameter is required.
	//
	// example:
	//
	// AAAAAdDWBF2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0688F1D2-CDA8-5617-A43C-ADAC61D80D43
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of instances.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRCCloudAssistantStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCCloudAssistantStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRCCloudAssistantStatusResponseBody) GetInstanceCloudAssistantStatusSet() []*DescribeRCCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet {
	return s.InstanceCloudAssistantStatusSet
}

func (s *DescribeRCCloudAssistantStatusResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeRCCloudAssistantStatusResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeRCCloudAssistantStatusResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeRCCloudAssistantStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRCCloudAssistantStatusResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeRCCloudAssistantStatusResponseBody) SetInstanceCloudAssistantStatusSet(v []*DescribeRCCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet) *DescribeRCCloudAssistantStatusResponseBody {
	s.InstanceCloudAssistantStatusSet = v
	return s
}

func (s *DescribeRCCloudAssistantStatusResponseBody) SetNextToken(v string) *DescribeRCCloudAssistantStatusResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeRCCloudAssistantStatusResponseBody) SetPageNumber(v string) *DescribeRCCloudAssistantStatusResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRCCloudAssistantStatusResponseBody) SetPageSize(v string) *DescribeRCCloudAssistantStatusResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRCCloudAssistantStatusResponseBody) SetRequestId(v string) *DescribeRCCloudAssistantStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRCCloudAssistantStatusResponseBody) SetTotalCount(v int32) *DescribeRCCloudAssistantStatusResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRCCloudAssistantStatusResponseBody) Validate() error {
	if s.InstanceCloudAssistantStatusSet != nil {
		for _, item := range s.InstanceCloudAssistantStatusSet {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRCCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet struct {
	// The number of tasks that Cloud Assistant was running on the instance.
	//
	// example:
	//
	// 0
	ActiveTaskCount *int32 `json:"ActiveTaskCount,omitempty" xml:"ActiveTaskCount,omitempty"`
	// Indicates whether Cloud Assistant is running on the instance. Valid values:
	//
	// 	- **true**: Heartbeats are detected in the last 2 minutes.
	//
	// 	- **false**: No heartbeat is detected in the last 2 minutes.
	//
	// example:
	//
	// true
	CloudAssistantStatus *string `json:"CloudAssistantStatus,omitempty" xml:"CloudAssistantStatus,omitempty"`
	// The version number of Cloud Assistant Agent. This parameter is empty if Cloud Assistant Agent is not installed or is not running on the instance.
	//
	// example:
	//
	// 2.2.0.106
	CloudAssistantVersion *string `json:"CloudAssistantVersion,omitempty" xml:"CloudAssistantVersion,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rc-e2g521l55k038cr8****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of tasks that Cloud Assistant completed on the instance.
	//
	// example:
	//
	// 2
	InvocationCount *int32 `json:"InvocationCount,omitempty" xml:"InvocationCount,omitempty"`
	// The last heartbeat time of Cloud Assistant. The value is updated every minute on average. The interval can be 55, 60, or 65 seconds.
	//
	// example:
	//
	// 2025-03-15T09:00:00Z
	LastHeartbeatTime *string `json:"LastHeartbeatTime,omitempty" xml:"LastHeartbeatTime,omitempty"`
	// The time when commands were last run.
	//
	// example:
	//
	// 2025-03-15T09:00:00Z
	LastInvokedTime *string `json:"LastInvokedTime,omitempty" xml:"LastInvokedTime,omitempty"`
	// The operating system type of the instance.
	//
	// example:
	//
	// Linux
	OSType *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	// Indicates whether Cloud Assistant supports Session Manager on the instance. If Session Manager is not supported, the version of Cloud Assistant Agent is outdated. Update Cloud Assistant Agent to the latest version.
	//
	// To support Session Manager, the version of Cloud Assistant Agent cannot be earlier than 2.2.3.189.
	//
	// example:
	//
	// true
	SupportSessionManager *bool `json:"SupportSessionManager,omitempty" xml:"SupportSessionManager,omitempty"`
}

func (s DescribeRCCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet) GoString() string {
	return s.String()
}

func (s *DescribeRCCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet) GetActiveTaskCount() *int32 {
	return s.ActiveTaskCount
}

func (s *DescribeRCCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet) GetCloudAssistantStatus() *string {
	return s.CloudAssistantStatus
}

func (s *DescribeRCCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet) GetCloudAssistantVersion() *string {
	return s.CloudAssistantVersion
}

func (s *DescribeRCCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRCCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet) GetInvocationCount() *int32 {
	return s.InvocationCount
}

func (s *DescribeRCCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet) GetLastHeartbeatTime() *string {
	return s.LastHeartbeatTime
}

func (s *DescribeRCCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet) GetLastInvokedTime() *string {
	return s.LastInvokedTime
}

func (s *DescribeRCCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet) GetOSType() *string {
	return s.OSType
}

func (s *DescribeRCCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet) GetSupportSessionManager() *bool {
	return s.SupportSessionManager
}

func (s *DescribeRCCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet) SetActiveTaskCount(v int32) *DescribeRCCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet {
	s.ActiveTaskCount = &v
	return s
}

func (s *DescribeRCCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet) SetCloudAssistantStatus(v string) *DescribeRCCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet {
	s.CloudAssistantStatus = &v
	return s
}

func (s *DescribeRCCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet) SetCloudAssistantVersion(v string) *DescribeRCCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet {
	s.CloudAssistantVersion = &v
	return s
}

func (s *DescribeRCCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet) SetInstanceId(v string) *DescribeRCCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet {
	s.InstanceId = &v
	return s
}

func (s *DescribeRCCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet) SetInvocationCount(v int32) *DescribeRCCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet {
	s.InvocationCount = &v
	return s
}

func (s *DescribeRCCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet) SetLastHeartbeatTime(v string) *DescribeRCCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet {
	s.LastHeartbeatTime = &v
	return s
}

func (s *DescribeRCCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet) SetLastInvokedTime(v string) *DescribeRCCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet {
	s.LastInvokedTime = &v
	return s
}

func (s *DescribeRCCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet) SetOSType(v string) *DescribeRCCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet {
	s.OSType = &v
	return s
}

func (s *DescribeRCCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet) SetSupportSessionManager(v bool) *DescribeRCCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet {
	s.SupportSessionManager = &v
	return s
}

func (s *DescribeRCCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet) Validate() error {
	return dara.Validate(s)
}
