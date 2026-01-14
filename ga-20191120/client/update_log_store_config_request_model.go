// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLogStoreConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *UpdateLogStoreConfigRequest
	GetAcceleratorId() *string
	SetAccessLogRecordCustomizedHeaderList(v []*string) *UpdateLogStoreConfigRequest
	GetAccessLogRecordCustomizedHeaderList() []*string
	SetAccessLogRecordCustomizedHeadersEnabled(v bool) *UpdateLogStoreConfigRequest
	GetAccessLogRecordCustomizedHeadersEnabled() *bool
	SetClientToken(v string) *UpdateLogStoreConfigRequest
	GetClientToken() *string
	SetEndpointGroupId(v string) *UpdateLogStoreConfigRequest
	GetEndpointGroupId() *string
	SetListenerId(v string) *UpdateLogStoreConfigRequest
	GetListenerId() *string
	SetRegionId(v string) *UpdateLogStoreConfigRequest
	GetRegionId() *string
	SetSlsLogStoreName(v string) *UpdateLogStoreConfigRequest
	GetSlsLogStoreName() *string
	SetSlsProjectName(v string) *UpdateLogStoreConfigRequest
	GetSlsProjectName() *string
}

type UpdateLogStoreConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId                       *string   `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	AccessLogRecordCustomizedHeaderList []*string `json:"AccessLogRecordCustomizedHeaderList,omitempty" xml:"AccessLogRecordCustomizedHeaderList,omitempty" type:"Repeated"`
	// example:
	//
	// False
	AccessLogRecordCustomizedHeadersEnabled *bool `json:"AccessLogRecordCustomizedHeadersEnabled,omitempty" xml:"AccessLogRecordCustomizedHeadersEnabled,omitempty"`
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// epg-bp1dmlohjjz4kqaun****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// lsr-xxxxxxxxxxxxxxx
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ga-access-log-epg-01
	SlsLogStoreName *string `json:"SlsLogStoreName,omitempty" xml:"SlsLogStoreName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ga-access-log
	SlsProjectName *string `json:"SlsProjectName,omitempty" xml:"SlsProjectName,omitempty"`
}

func (s UpdateLogStoreConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLogStoreConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateLogStoreConfigRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *UpdateLogStoreConfigRequest) GetAccessLogRecordCustomizedHeaderList() []*string {
	return s.AccessLogRecordCustomizedHeaderList
}

func (s *UpdateLogStoreConfigRequest) GetAccessLogRecordCustomizedHeadersEnabled() *bool {
	return s.AccessLogRecordCustomizedHeadersEnabled
}

func (s *UpdateLogStoreConfigRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateLogStoreConfigRequest) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *UpdateLogStoreConfigRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *UpdateLogStoreConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateLogStoreConfigRequest) GetSlsLogStoreName() *string {
	return s.SlsLogStoreName
}

func (s *UpdateLogStoreConfigRequest) GetSlsProjectName() *string {
	return s.SlsProjectName
}

func (s *UpdateLogStoreConfigRequest) SetAcceleratorId(v string) *UpdateLogStoreConfigRequest {
	s.AcceleratorId = &v
	return s
}

func (s *UpdateLogStoreConfigRequest) SetAccessLogRecordCustomizedHeaderList(v []*string) *UpdateLogStoreConfigRequest {
	s.AccessLogRecordCustomizedHeaderList = v
	return s
}

func (s *UpdateLogStoreConfigRequest) SetAccessLogRecordCustomizedHeadersEnabled(v bool) *UpdateLogStoreConfigRequest {
	s.AccessLogRecordCustomizedHeadersEnabled = &v
	return s
}

func (s *UpdateLogStoreConfigRequest) SetClientToken(v string) *UpdateLogStoreConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateLogStoreConfigRequest) SetEndpointGroupId(v string) *UpdateLogStoreConfigRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *UpdateLogStoreConfigRequest) SetListenerId(v string) *UpdateLogStoreConfigRequest {
	s.ListenerId = &v
	return s
}

func (s *UpdateLogStoreConfigRequest) SetRegionId(v string) *UpdateLogStoreConfigRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateLogStoreConfigRequest) SetSlsLogStoreName(v string) *UpdateLogStoreConfigRequest {
	s.SlsLogStoreName = &v
	return s
}

func (s *UpdateLogStoreConfigRequest) SetSlsProjectName(v string) *UpdateLogStoreConfigRequest {
	s.SlsProjectName = &v
	return s
}

func (s *UpdateLogStoreConfigRequest) Validate() error {
	return dara.Validate(s)
}
