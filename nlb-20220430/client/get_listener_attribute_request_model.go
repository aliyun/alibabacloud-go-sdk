// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetListenerAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetListenerAttributeRequest
	GetClientToken() *string
	SetDryRun(v bool) *GetListenerAttributeRequest
	GetDryRun() *bool
	SetListenerId(v string) *GetListenerAttributeRequest
	GetListenerId() *string
	SetRegionId(v string) *GetListenerAttributeRequest
	GetRegionId() *string
}

type GetListenerAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate the parameter value from your client to ensure that it is unique among different requests. The ClientToken parameter supports only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to only precheck the request. Valid values:
	//
	// - **true**: sends a check request and does not query the listener details. Check items include whether required parameters are specified, the request format, and service limits. If the check fails, the corresponding error is returned. If the check succeeds, the error code `DryRunOperation` is returned.
	//
	// - **false*	- (default): sends a normal request. After the check succeeds, the HTTP 2xx status code is returned and the operation is directly performed.
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
	// lsn-bp1bpn0kn908w4nbw****@233
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The region ID of the Network Load Balancer (NLB) instance.
	//
	// You can invoke the [DescribeRegions](https://help.aliyun.com/document_detail/443657.html) operation to query region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetListenerAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetListenerAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetListenerAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *GetListenerAttributeRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *GetListenerAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetListenerAttributeRequest) SetClientToken(v string) *GetListenerAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *GetListenerAttributeRequest) SetDryRun(v bool) *GetListenerAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *GetListenerAttributeRequest) SetListenerId(v string) *GetListenerAttributeRequest {
	s.ListenerId = &v
	return s
}

func (s *GetListenerAttributeRequest) SetRegionId(v string) *GetListenerAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *GetListenerAttributeRequest) Validate() error {
	return dara.Validate(s)
}
