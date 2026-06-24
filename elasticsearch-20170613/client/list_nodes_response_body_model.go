// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v *ListNodesResponseBodyHeaders) *ListNodesResponseBody
	GetHeaders() *ListNodesResponseBodyHeaders
	SetRequestId(v string) *ListNodesResponseBody
	GetRequestId() *string
	SetResult(v []*ListNodesResponseBodyResult) *ListNodesResponseBody
	GetResult() []*ListNodesResponseBodyResult
}

type ListNodesResponseBody struct {
	// The response headers.
	Headers *ListNodesResponseBodyHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// E1FD7642-7C40-4FF2-9C0F-21F1A1746F70
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned results.
	Result []*ListNodesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBody) GetHeaders() *ListNodesResponseBodyHeaders {
	return s.Headers
}

func (s *ListNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNodesResponseBody) GetResult() []*ListNodesResponseBodyResult {
	return s.Result
}

func (s *ListNodesResponseBody) SetHeaders(v *ListNodesResponseBodyHeaders) *ListNodesResponseBody {
	s.Headers = v
	return s
}

func (s *ListNodesResponseBody) SetRequestId(v string) *ListNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodesResponseBody) SetResult(v []*ListNodesResponseBodyResult) *ListNodesResponseBody {
	s.Result = v
	return s
}

func (s *ListNodesResponseBody) Validate() error {
	if s.Headers != nil {
		if err := s.Headers.Validate(); err != nil {
			return err
		}
	}
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNodesResponseBodyHeaders struct {
	// The total number of returned records.
	//
	// example:
	//
	// 10
	XTotalCount *int32 `json:"X-Total-Count,omitempty" xml:"X-Total-Count,omitempty"`
}

func (s ListNodesResponseBodyHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBodyHeaders) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyHeaders) GetXTotalCount() *int32 {
	return s.XTotalCount
}

func (s *ListNodesResponseBodyHeaders) SetXTotalCount(v int32) *ListNodesResponseBodyHeaders {
	s.XTotalCount = &v
	return s
}

func (s *ListNodesResponseBodyHeaders) Validate() error {
	return dara.Validate(s)
}

type ListNodesResponseBodyResult struct {
	// The status of each collector on the ECS instance. Valid values:
	//
	// - heartOk: normal heartbeat.
	//
	// - heartLost: abnormal heartbeat.
	//
	// - uninstalled: not installed.
	//
	// - failed: installation failed.
	//
	// example:
	//
	// heartOk
	AgentStatus *string `json:"agentStatus,omitempty" xml:"agentStatus,omitempty"`
	// Indicates whether Cloud Assistant is enabled. Valid values:
	//
	// - true: enabled.
	//
	// - false: not enabled.
	//
	// example:
	//
	// true
	CloudAssistantStatus *string `json:"cloudAssistantStatus,omitempty" xml:"cloudAssistantStatus,omitempty"`
	// The ECS instance ID.
	//
	// example:
	//
	// i-bp13y63575oypr****
	EcsInstanceId *string `json:"ecsInstanceId,omitempty" xml:"ecsInstanceId,omitempty"`
	// The ECS instance name.
	//
	// example:
	//
	// ECS_beat
	EcsInstanceName *string `json:"ecsInstanceName,omitempty" xml:"ecsInstanceName,omitempty"`
	// The list of IP address information of the ECS instance.
	IpAddress []*ListNodesResponseBodyResultIpAddress `json:"ipAddress,omitempty" xml:"ipAddress,omitempty" type:"Repeated"`
	// The operating system type of the ECS instance. Valid values:
	//
	// - windows: Windows Server.
	//
	// - linux: Linux.
	//
	// example:
	//
	// linux
	OsType *string `json:"osType,omitempty" xml:"osType,omitempty"`
	// The status of the ECS instance. Valid values:
	//
	// - running: running.
	//
	// - starting: starting.
	//
	// - stopping: stopping.
	//
	// - stopped: stopped.
	//
	// example:
	//
	// running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The tag information of the ECS instance.
	Tags []*ListNodesResponseBodyResultTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s ListNodesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyResult) GetAgentStatus() *string {
	return s.AgentStatus
}

func (s *ListNodesResponseBodyResult) GetCloudAssistantStatus() *string {
	return s.CloudAssistantStatus
}

func (s *ListNodesResponseBodyResult) GetEcsInstanceId() *string {
	return s.EcsInstanceId
}

func (s *ListNodesResponseBodyResult) GetEcsInstanceName() *string {
	return s.EcsInstanceName
}

func (s *ListNodesResponseBodyResult) GetIpAddress() []*ListNodesResponseBodyResultIpAddress {
	return s.IpAddress
}

func (s *ListNodesResponseBodyResult) GetOsType() *string {
	return s.OsType
}

func (s *ListNodesResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ListNodesResponseBodyResult) GetTags() []*ListNodesResponseBodyResultTags {
	return s.Tags
}

func (s *ListNodesResponseBodyResult) SetAgentStatus(v string) *ListNodesResponseBodyResult {
	s.AgentStatus = &v
	return s
}

func (s *ListNodesResponseBodyResult) SetCloudAssistantStatus(v string) *ListNodesResponseBodyResult {
	s.CloudAssistantStatus = &v
	return s
}

func (s *ListNodesResponseBodyResult) SetEcsInstanceId(v string) *ListNodesResponseBodyResult {
	s.EcsInstanceId = &v
	return s
}

func (s *ListNodesResponseBodyResult) SetEcsInstanceName(v string) *ListNodesResponseBodyResult {
	s.EcsInstanceName = &v
	return s
}

func (s *ListNodesResponseBodyResult) SetIpAddress(v []*ListNodesResponseBodyResultIpAddress) *ListNodesResponseBodyResult {
	s.IpAddress = v
	return s
}

func (s *ListNodesResponseBodyResult) SetOsType(v string) *ListNodesResponseBodyResult {
	s.OsType = &v
	return s
}

func (s *ListNodesResponseBodyResult) SetStatus(v string) *ListNodesResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListNodesResponseBodyResult) SetTags(v []*ListNodesResponseBodyResultTags) *ListNodesResponseBodyResult {
	s.Tags = v
	return s
}

func (s *ListNodesResponseBodyResult) Validate() error {
	if s.IpAddress != nil {
		for _, item := range s.IpAddress {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNodesResponseBodyResultIpAddress struct {
	// The IP address.
	//
	// example:
	//
	// 192.168.xx.xx
	Host *string `json:"host,omitempty" xml:"host,omitempty"`
	// The type of the IP address. Valid values:
	//
	// - public: public IP address.
	//
	// - private: private IP address.
	//
	// example:
	//
	// public
	IpType *string `json:"ipType,omitempty" xml:"ipType,omitempty"`
}

func (s ListNodesResponseBodyResultIpAddress) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBodyResultIpAddress) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyResultIpAddress) GetHost() *string {
	return s.Host
}

func (s *ListNodesResponseBodyResultIpAddress) GetIpType() *string {
	return s.IpType
}

func (s *ListNodesResponseBodyResultIpAddress) SetHost(v string) *ListNodesResponseBodyResultIpAddress {
	s.Host = &v
	return s
}

func (s *ListNodesResponseBodyResultIpAddress) SetIpType(v string) *ListNodesResponseBodyResultIpAddress {
	s.IpType = &v
	return s
}

func (s *ListNodesResponseBodyResultIpAddress) Validate() error {
	return dara.Validate(s)
}

type ListNodesResponseBodyResultTags struct {
	// The tag key.
	//
	// example:
	//
	// abc
	TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// xyz
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s ListNodesResponseBodyResultTags) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBodyResultTags) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyResultTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ListNodesResponseBodyResultTags) GetTagValue() *string {
	return s.TagValue
}

func (s *ListNodesResponseBodyResultTags) SetTagKey(v string) *ListNodesResponseBodyResultTags {
	s.TagKey = &v
	return s
}

func (s *ListNodesResponseBodyResultTags) SetTagValue(v string) *ListNodesResponseBodyResultTags {
	s.TagValue = &v
	return s
}

func (s *ListNodesResponseBodyResultTags) Validate() error {
	return dara.Validate(s)
}
