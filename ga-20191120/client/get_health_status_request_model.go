// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHealthStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *GetHealthStatusRequest
	GetAcceleratorId() *string
	SetClientToken(v string) *GetHealthStatusRequest
	GetClientToken() *string
	SetDryRun(v bool) *GetHealthStatusRequest
	GetDryRun() *bool
	SetListenerId(v string) *GetHealthStatusRequest
	GetListenerId() *string
	SetRegionId(v string) *GetHealthStatusRequest
	GetRegionId() *string
}

type GetHealthStatusRequest struct {
	// The ID of the GA instance that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether only to precheck the request. Valid values:
	//
	// 	- **true**: prechecks the request only. The health status of the listener is not queried. The system prechecks the required parameters, request syntax, and limits. If the request fails the precheck, an error message is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): sends a normal request. If the request passes the precheck, a 2xx HTTP status code is returned and the health status of the listener is obtained.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the listener.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The ID of the region where the Global Accelerator (GA) instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetHealthStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHealthStatusRequest) GoString() string {
	return s.String()
}

func (s *GetHealthStatusRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *GetHealthStatusRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetHealthStatusRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *GetHealthStatusRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *GetHealthStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetHealthStatusRequest) SetAcceleratorId(v string) *GetHealthStatusRequest {
	s.AcceleratorId = &v
	return s
}

func (s *GetHealthStatusRequest) SetClientToken(v string) *GetHealthStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *GetHealthStatusRequest) SetDryRun(v bool) *GetHealthStatusRequest {
	s.DryRun = &v
	return s
}

func (s *GetHealthStatusRequest) SetListenerId(v string) *GetHealthStatusRequest {
	s.ListenerId = &v
	return s
}

func (s *GetHealthStatusRequest) SetRegionId(v string) *GetHealthStatusRequest {
	s.RegionId = &v
	return s
}

func (s *GetHealthStatusRequest) Validate() error {
	return dara.Validate(s)
}
