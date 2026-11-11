// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateAdditionalCertificatesWithListenerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalCertificateIds(v []*string) *AssociateAdditionalCertificatesWithListenerRequest
	GetAdditionalCertificateIds() []*string
	SetClientToken(v string) *AssociateAdditionalCertificatesWithListenerRequest
	GetClientToken() *string
	SetDryRun(v bool) *AssociateAdditionalCertificatesWithListenerRequest
	GetDryRun() *bool
	SetListenerId(v string) *AssociateAdditionalCertificatesWithListenerRequest
	GetListenerId() *string
	SetRegionId(v string) *AssociateAdditionalCertificatesWithListenerRequest
	GetRegionId() *string
}

type AssociateAdditionalCertificatesWithListenerRequest struct {
	// The list of additional certificates. You can associate up to 15 additional certificates with a listener in a single call.
	//
	// This parameter is required.
	AdditionalCertificateIds []*string `json:"AdditionalCertificateIds,omitempty" xml:"AdditionalCertificateIds,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate a value for this parameter on your client to ensure that the value is unique across different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is unique.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to only precheck this request. Valid values:
	//
	// - **true**: sends a check request and does not add the additional certificates to the TCPSSL listener. The system checks whether the required parameters are specified, the request format is valid, and the service limits are met. If the request fails the check, the system returns an error. If the request passes the check, the system returns the error code `DryRunOperation`.
	//
	// - **false*	- (default): sends a standard request. If the request passes the check, the system returns an HTTP 2xx status code and performs the operation.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the listener. This operation only supports TCPSSL listener IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsn-bpn0kn908w4nbw****@80
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The ID of the region where the Network Load Balancer (NLB) instance is deployed.
	//
	// You can invoke the [DescribeRegions](https://help.aliyun.com/document_detail/443657.html) operation to query region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AssociateAdditionalCertificatesWithListenerRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateAdditionalCertificatesWithListenerRequest) GoString() string {
	return s.String()
}

func (s *AssociateAdditionalCertificatesWithListenerRequest) GetAdditionalCertificateIds() []*string {
	return s.AdditionalCertificateIds
}

func (s *AssociateAdditionalCertificatesWithListenerRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AssociateAdditionalCertificatesWithListenerRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *AssociateAdditionalCertificatesWithListenerRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *AssociateAdditionalCertificatesWithListenerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AssociateAdditionalCertificatesWithListenerRequest) SetAdditionalCertificateIds(v []*string) *AssociateAdditionalCertificatesWithListenerRequest {
	s.AdditionalCertificateIds = v
	return s
}

func (s *AssociateAdditionalCertificatesWithListenerRequest) SetClientToken(v string) *AssociateAdditionalCertificatesWithListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *AssociateAdditionalCertificatesWithListenerRequest) SetDryRun(v bool) *AssociateAdditionalCertificatesWithListenerRequest {
	s.DryRun = &v
	return s
}

func (s *AssociateAdditionalCertificatesWithListenerRequest) SetListenerId(v string) *AssociateAdditionalCertificatesWithListenerRequest {
	s.ListenerId = &v
	return s
}

func (s *AssociateAdditionalCertificatesWithListenerRequest) SetRegionId(v string) *AssociateAdditionalCertificatesWithListenerRequest {
	s.RegionId = &v
	return s
}

func (s *AssociateAdditionalCertificatesWithListenerRequest) Validate() error {
	return dara.Validate(s)
}
