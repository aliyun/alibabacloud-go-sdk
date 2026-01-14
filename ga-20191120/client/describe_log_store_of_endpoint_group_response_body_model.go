// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogStoreOfEndpointGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *DescribeLogStoreOfEndpointGroupResponseBody
	GetAcceleratorId() *string
	SetAccessLogRecordCustomizedHeaderList(v []*string) *DescribeLogStoreOfEndpointGroupResponseBody
	GetAccessLogRecordCustomizedHeaderList() []*string
	SetAccessLogRecordCustomizedHeadersEnabled(v bool) *DescribeLogStoreOfEndpointGroupResponseBody
	GetAccessLogRecordCustomizedHeadersEnabled() *bool
	SetEndpointGroupId(v string) *DescribeLogStoreOfEndpointGroupResponseBody
	GetEndpointGroupId() *string
	SetListenerId(v string) *DescribeLogStoreOfEndpointGroupResponseBody
	GetListenerId() *string
	SetRequestId(v string) *DescribeLogStoreOfEndpointGroupResponseBody
	GetRequestId() *string
	SetSlsLogStoreName(v string) *DescribeLogStoreOfEndpointGroupResponseBody
	GetSlsLogStoreName() *string
	SetSlsProjectName(v string) *DescribeLogStoreOfEndpointGroupResponseBody
	GetSlsProjectName() *string
	SetSlsRegionId(v string) *DescribeLogStoreOfEndpointGroupResponseBody
	GetSlsRegionId() *string
	SetStatus(v string) *DescribeLogStoreOfEndpointGroupResponseBody
	GetStatus() *string
}

type DescribeLogStoreOfEndpointGroupResponseBody struct {
	// The ID of the GA instance.
	//
	// example:
	//
	// ga-xxxxxxxxxxxxx
	AcceleratorId                           *string   `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	AccessLogRecordCustomizedHeaderList     []*string `json:"AccessLogRecordCustomizedHeaderList,omitempty" xml:"AccessLogRecordCustomizedHeaderList,omitempty" type:"Repeated"`
	AccessLogRecordCustomizedHeadersEnabled *bool     `json:"AccessLogRecordCustomizedHeadersEnabled,omitempty" xml:"AccessLogRecordCustomizedHeadersEnabled,omitempty"`
	// The ID of the endpoint group.
	//
	// example:
	//
	// epg-xxxxxxxxxxxxxxx
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The ID of the listener.
	//
	// example:
	//
	// lsr-xxxxxxxxxxxxxxx
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The response parameters.
	//
	// example:
	//
	// 1777E713-8456-55F1-9A69-9AD9EAE2B3B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the Logstore.
	//
	// example:
	//
	// ga_log
	SlsLogStoreName *string `json:"SlsLogStoreName,omitempty" xml:"SlsLogStoreName,omitempty"`
	// The name of the Simple Log Service project.
	//
	// example:
	//
	// ga_project_name
	SlsProjectName *string `json:"SlsProjectName,omitempty" xml:"SlsProjectName,omitempty"`
	// The region ID of the Simple Log Service project.
	//
	// example:
	//
	// cn-hangzhou
	SlsRegionId *string `json:"SlsRegionId,omitempty" xml:"SlsRegionId,omitempty"`
	// Indicates whether the endpoint group is bound to the Simple Log Service project.
	//
	// 	- **on:*	- The endpoint group is bound to the Simple Log Service project.
	//
	// 	- **off:*	- The endpoint group is not bound to the Simple Log Service project.
	//
	// example:
	//
	// on - binding
	//
	// off - unbinding
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeLogStoreOfEndpointGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogStoreOfEndpointGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogStoreOfEndpointGroupResponseBody) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *DescribeLogStoreOfEndpointGroupResponseBody) GetAccessLogRecordCustomizedHeaderList() []*string {
	return s.AccessLogRecordCustomizedHeaderList
}

func (s *DescribeLogStoreOfEndpointGroupResponseBody) GetAccessLogRecordCustomizedHeadersEnabled() *bool {
	return s.AccessLogRecordCustomizedHeadersEnabled
}

func (s *DescribeLogStoreOfEndpointGroupResponseBody) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *DescribeLogStoreOfEndpointGroupResponseBody) GetListenerId() *string {
	return s.ListenerId
}

func (s *DescribeLogStoreOfEndpointGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLogStoreOfEndpointGroupResponseBody) GetSlsLogStoreName() *string {
	return s.SlsLogStoreName
}

func (s *DescribeLogStoreOfEndpointGroupResponseBody) GetSlsProjectName() *string {
	return s.SlsProjectName
}

func (s *DescribeLogStoreOfEndpointGroupResponseBody) GetSlsRegionId() *string {
	return s.SlsRegionId
}

func (s *DescribeLogStoreOfEndpointGroupResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeLogStoreOfEndpointGroupResponseBody) SetAcceleratorId(v string) *DescribeLogStoreOfEndpointGroupResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *DescribeLogStoreOfEndpointGroupResponseBody) SetAccessLogRecordCustomizedHeaderList(v []*string) *DescribeLogStoreOfEndpointGroupResponseBody {
	s.AccessLogRecordCustomizedHeaderList = v
	return s
}

func (s *DescribeLogStoreOfEndpointGroupResponseBody) SetAccessLogRecordCustomizedHeadersEnabled(v bool) *DescribeLogStoreOfEndpointGroupResponseBody {
	s.AccessLogRecordCustomizedHeadersEnabled = &v
	return s
}

func (s *DescribeLogStoreOfEndpointGroupResponseBody) SetEndpointGroupId(v string) *DescribeLogStoreOfEndpointGroupResponseBody {
	s.EndpointGroupId = &v
	return s
}

func (s *DescribeLogStoreOfEndpointGroupResponseBody) SetListenerId(v string) *DescribeLogStoreOfEndpointGroupResponseBody {
	s.ListenerId = &v
	return s
}

func (s *DescribeLogStoreOfEndpointGroupResponseBody) SetRequestId(v string) *DescribeLogStoreOfEndpointGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogStoreOfEndpointGroupResponseBody) SetSlsLogStoreName(v string) *DescribeLogStoreOfEndpointGroupResponseBody {
	s.SlsLogStoreName = &v
	return s
}

func (s *DescribeLogStoreOfEndpointGroupResponseBody) SetSlsProjectName(v string) *DescribeLogStoreOfEndpointGroupResponseBody {
	s.SlsProjectName = &v
	return s
}

func (s *DescribeLogStoreOfEndpointGroupResponseBody) SetSlsRegionId(v string) *DescribeLogStoreOfEndpointGroupResponseBody {
	s.SlsRegionId = &v
	return s
}

func (s *DescribeLogStoreOfEndpointGroupResponseBody) SetStatus(v string) *DescribeLogStoreOfEndpointGroupResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeLogStoreOfEndpointGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
