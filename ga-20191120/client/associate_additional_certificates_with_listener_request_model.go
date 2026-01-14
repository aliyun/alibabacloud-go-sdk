// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateAdditionalCertificatesWithListenerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *AssociateAdditionalCertificatesWithListenerRequest
	GetAcceleratorId() *string
	SetCertificates(v []*AssociateAdditionalCertificatesWithListenerRequestCertificates) *AssociateAdditionalCertificatesWithListenerRequest
	GetCertificates() []*AssociateAdditionalCertificatesWithListenerRequestCertificates
	SetClientToken(v string) *AssociateAdditionalCertificatesWithListenerRequest
	GetClientToken() *string
	SetListenerId(v string) *AssociateAdditionalCertificatesWithListenerRequest
	GetListenerId() *string
	SetRegionId(v string) *AssociateAdditionalCertificatesWithListenerRequest
	GetRegionId() *string
}

type AssociateAdditionalCertificatesWithListenerRequest struct {
	// The GA instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The additional certificates.
	//
	// You can specify up to 10 certificates in each request.
	//
	// This parameter is required.
	Certificates []*AssociateAdditionalCertificatesWithListenerRequestCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The listener ID. Only HTTPS listeners are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The region ID of the GA instance. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
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

func (s *AssociateAdditionalCertificatesWithListenerRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *AssociateAdditionalCertificatesWithListenerRequest) GetCertificates() []*AssociateAdditionalCertificatesWithListenerRequestCertificates {
	return s.Certificates
}

func (s *AssociateAdditionalCertificatesWithListenerRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AssociateAdditionalCertificatesWithListenerRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *AssociateAdditionalCertificatesWithListenerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AssociateAdditionalCertificatesWithListenerRequest) SetAcceleratorId(v string) *AssociateAdditionalCertificatesWithListenerRequest {
	s.AcceleratorId = &v
	return s
}

func (s *AssociateAdditionalCertificatesWithListenerRequest) SetCertificates(v []*AssociateAdditionalCertificatesWithListenerRequestCertificates) *AssociateAdditionalCertificatesWithListenerRequest {
	s.Certificates = v
	return s
}

func (s *AssociateAdditionalCertificatesWithListenerRequest) SetClientToken(v string) *AssociateAdditionalCertificatesWithListenerRequest {
	s.ClientToken = &v
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
	if s.Certificates != nil {
		for _, item := range s.Certificates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AssociateAdditionalCertificatesWithListenerRequestCertificates struct {
	// The domain name that is specified by the certificate. You can associate each domain name with only one additional certificate.
	//
	// You can specify up to 10 domain names in each request.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the certificate. Only server certificates are supported.
	//
	// You can specify up to 10 certificate IDs in each request.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6092**-cn-hangzhou
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s AssociateAdditionalCertificatesWithListenerRequestCertificates) String() string {
	return dara.Prettify(s)
}

func (s AssociateAdditionalCertificatesWithListenerRequestCertificates) GoString() string {
	return s.String()
}

func (s *AssociateAdditionalCertificatesWithListenerRequestCertificates) GetDomain() *string {
	return s.Domain
}

func (s *AssociateAdditionalCertificatesWithListenerRequestCertificates) GetId() *string {
	return s.Id
}

func (s *AssociateAdditionalCertificatesWithListenerRequestCertificates) SetDomain(v string) *AssociateAdditionalCertificatesWithListenerRequestCertificates {
	s.Domain = &v
	return s
}

func (s *AssociateAdditionalCertificatesWithListenerRequestCertificates) SetId(v string) *AssociateAdditionalCertificatesWithListenerRequestCertificates {
	s.Id = &v
	return s
}

func (s *AssociateAdditionalCertificatesWithListenerRequestCertificates) Validate() error {
	return dara.Validate(s)
}
