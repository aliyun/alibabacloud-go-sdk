// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateAdditionalCertificatesFromListenerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *DissociateAdditionalCertificatesFromListenerRequest
	GetAcceleratorId() *string
	SetClientToken(v string) *DissociateAdditionalCertificatesFromListenerRequest
	GetClientToken() *string
	SetDomains(v []*string) *DissociateAdditionalCertificatesFromListenerRequest
	GetDomains() []*string
	SetListenerId(v string) *DissociateAdditionalCertificatesFromListenerRequest
	GetListenerId() *string
	SetRegionId(v string) *DissociateAdditionalCertificatesFromListenerRequest
	GetRegionId() *string
}

type DissociateAdditionalCertificatesFromListenerRequest struct {
	// The ID of the GA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among all requests. ClientToken can contain only ASCII characters.
	//
	// >  If you do not set this parameter, **ClientToken*	- is set to the value of **RequestId**. The value of **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The domain name associated with the additional certificate.
	//
	// You can specify up to 10 domain names in each request.
	//
	// This parameter is required.
	Domains []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
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
}

func (s DissociateAdditionalCertificatesFromListenerRequest) String() string {
	return dara.Prettify(s)
}

func (s DissociateAdditionalCertificatesFromListenerRequest) GoString() string {
	return s.String()
}

func (s *DissociateAdditionalCertificatesFromListenerRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *DissociateAdditionalCertificatesFromListenerRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DissociateAdditionalCertificatesFromListenerRequest) GetDomains() []*string {
	return s.Domains
}

func (s *DissociateAdditionalCertificatesFromListenerRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *DissociateAdditionalCertificatesFromListenerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DissociateAdditionalCertificatesFromListenerRequest) SetAcceleratorId(v string) *DissociateAdditionalCertificatesFromListenerRequest {
	s.AcceleratorId = &v
	return s
}

func (s *DissociateAdditionalCertificatesFromListenerRequest) SetClientToken(v string) *DissociateAdditionalCertificatesFromListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *DissociateAdditionalCertificatesFromListenerRequest) SetDomains(v []*string) *DissociateAdditionalCertificatesFromListenerRequest {
	s.Domains = v
	return s
}

func (s *DissociateAdditionalCertificatesFromListenerRequest) SetListenerId(v string) *DissociateAdditionalCertificatesFromListenerRequest {
	s.ListenerId = &v
	return s
}

func (s *DissociateAdditionalCertificatesFromListenerRequest) SetRegionId(v string) *DissociateAdditionalCertificatesFromListenerRequest {
	s.RegionId = &v
	return s
}

func (s *DissociateAdditionalCertificatesFromListenerRequest) Validate() error {
	return dara.Validate(s)
}
