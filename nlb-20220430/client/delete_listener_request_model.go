// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteListenerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteListenerRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteListenerRequest
	GetDryRun() *bool
	SetListenerId(v string) *DeleteListenerRequest
	GetListenerId() *string
	SetRegionId(v string) *DeleteListenerRequest
	GetRegionId() *string
}

type DeleteListenerRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate a value for this parameter on the client to ensure that the value is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is unique.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to only perform a dry run for the request. Valid values:
	//
	// - **true**: sends a dry run request and does not delete the listener. The system checks whether required parameters are specified, the request format is valid, and service limits are met. If the request fails the check, the corresponding error is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends a regular request. If the request passes the check, an HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The listener ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsn-bp1bpn0kn908w4nbw****@80
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The region ID of the Network Load Balancer (NLB) instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/443657.html) operation to obtain region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteListenerRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteListenerRequest) GoString() string {
	return s.String()
}

func (s *DeleteListenerRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteListenerRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteListenerRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *DeleteListenerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteListenerRequest) SetClientToken(v string) *DeleteListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteListenerRequest) SetDryRun(v bool) *DeleteListenerRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteListenerRequest) SetListenerId(v string) *DeleteListenerRequest {
	s.ListenerId = &v
	return s
}

func (s *DeleteListenerRequest) SetRegionId(v string) *DeleteListenerRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteListenerRequest) Validate() error {
	return dara.Validate(s)
}
