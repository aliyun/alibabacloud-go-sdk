// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachLogStoreToEndpointGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *AttachLogStoreToEndpointGroupRequest
	GetAcceleratorId() *string
	SetAccessLogRecordCustomizedHeaderList(v []*string) *AttachLogStoreToEndpointGroupRequest
	GetAccessLogRecordCustomizedHeaderList() []*string
	SetAccessLogRecordCustomizedHeadersEnabled(v bool) *AttachLogStoreToEndpointGroupRequest
	GetAccessLogRecordCustomizedHeadersEnabled() *bool
	SetClientToken(v string) *AttachLogStoreToEndpointGroupRequest
	GetClientToken() *string
	SetEndpointGroupIds(v []*string) *AttachLogStoreToEndpointGroupRequest
	GetEndpointGroupIds() []*string
	SetListenerId(v string) *AttachLogStoreToEndpointGroupRequest
	GetListenerId() *string
	SetRegionId(v string) *AttachLogStoreToEndpointGroupRequest
	GetRegionId() *string
	SetSlsLogStoreName(v string) *AttachLogStoreToEndpointGroupRequest
	GetSlsLogStoreName() *string
	SetSlsProjectName(v string) *AttachLogStoreToEndpointGroupRequest
	GetSlsProjectName() *string
	SetSlsRegionId(v string) *AttachLogStoreToEndpointGroupRequest
	GetSlsRegionId() *string
}

type AttachLogStoreToEndpointGroupRequest struct {
	// The ID of the GA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId                           *string   `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	AccessLogRecordCustomizedHeaderList     []*string `json:"AccessLogRecordCustomizedHeaderList,omitempty" xml:"AccessLogRecordCustomizedHeaderList,omitempty" type:"Repeated"`
	AccessLogRecordCustomizedHeadersEnabled *bool     `json:"AccessLogRecordCustomizedHeadersEnabled,omitempty" xml:"AccessLogRecordCustomizedHeadersEnabled,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. ClientToken can contain only ASCII characters.
	//
	// >  If you do not set this parameter, **ClientToken*	- is set to the value of **RequestId**. The value of **RequestId*	- for each API request may be different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The IDs of the endpoint groups.
	//
	// This parameter is required.
	EndpointGroupIds []*string `json:"EndpointGroupIds,omitempty" xml:"EndpointGroupIds,omitempty" type:"Repeated"`
	// The ID of the listener.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The ID of the region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the Logstore.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsn-01
	SlsLogStoreName *string `json:"SlsLogStoreName,omitempty" xml:"SlsLogStoreName,omitempty"`
	// The name of the Log Service project.
	//
	// This parameter is required.
	//
	// example:
	//
	// pn-01
	SlsProjectName *string `json:"SlsProjectName,omitempty" xml:"SlsProjectName,omitempty"`
	// The region ID of the Log Service project.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	SlsRegionId *string `json:"SlsRegionId,omitempty" xml:"SlsRegionId,omitempty"`
}

func (s AttachLogStoreToEndpointGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachLogStoreToEndpointGroupRequest) GoString() string {
	return s.String()
}

func (s *AttachLogStoreToEndpointGroupRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *AttachLogStoreToEndpointGroupRequest) GetAccessLogRecordCustomizedHeaderList() []*string {
	return s.AccessLogRecordCustomizedHeaderList
}

func (s *AttachLogStoreToEndpointGroupRequest) GetAccessLogRecordCustomizedHeadersEnabled() *bool {
	return s.AccessLogRecordCustomizedHeadersEnabled
}

func (s *AttachLogStoreToEndpointGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AttachLogStoreToEndpointGroupRequest) GetEndpointGroupIds() []*string {
	return s.EndpointGroupIds
}

func (s *AttachLogStoreToEndpointGroupRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *AttachLogStoreToEndpointGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachLogStoreToEndpointGroupRequest) GetSlsLogStoreName() *string {
	return s.SlsLogStoreName
}

func (s *AttachLogStoreToEndpointGroupRequest) GetSlsProjectName() *string {
	return s.SlsProjectName
}

func (s *AttachLogStoreToEndpointGroupRequest) GetSlsRegionId() *string {
	return s.SlsRegionId
}

func (s *AttachLogStoreToEndpointGroupRequest) SetAcceleratorId(v string) *AttachLogStoreToEndpointGroupRequest {
	s.AcceleratorId = &v
	return s
}

func (s *AttachLogStoreToEndpointGroupRequest) SetAccessLogRecordCustomizedHeaderList(v []*string) *AttachLogStoreToEndpointGroupRequest {
	s.AccessLogRecordCustomizedHeaderList = v
	return s
}

func (s *AttachLogStoreToEndpointGroupRequest) SetAccessLogRecordCustomizedHeadersEnabled(v bool) *AttachLogStoreToEndpointGroupRequest {
	s.AccessLogRecordCustomizedHeadersEnabled = &v
	return s
}

func (s *AttachLogStoreToEndpointGroupRequest) SetClientToken(v string) *AttachLogStoreToEndpointGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *AttachLogStoreToEndpointGroupRequest) SetEndpointGroupIds(v []*string) *AttachLogStoreToEndpointGroupRequest {
	s.EndpointGroupIds = v
	return s
}

func (s *AttachLogStoreToEndpointGroupRequest) SetListenerId(v string) *AttachLogStoreToEndpointGroupRequest {
	s.ListenerId = &v
	return s
}

func (s *AttachLogStoreToEndpointGroupRequest) SetRegionId(v string) *AttachLogStoreToEndpointGroupRequest {
	s.RegionId = &v
	return s
}

func (s *AttachLogStoreToEndpointGroupRequest) SetSlsLogStoreName(v string) *AttachLogStoreToEndpointGroupRequest {
	s.SlsLogStoreName = &v
	return s
}

func (s *AttachLogStoreToEndpointGroupRequest) SetSlsProjectName(v string) *AttachLogStoreToEndpointGroupRequest {
	s.SlsProjectName = &v
	return s
}

func (s *AttachLogStoreToEndpointGroupRequest) SetSlsRegionId(v string) *AttachLogStoreToEndpointGroupRequest {
	s.SlsRegionId = &v
	return s
}

func (s *AttachLogStoreToEndpointGroupRequest) Validate() error {
	return dara.Validate(s)
}
