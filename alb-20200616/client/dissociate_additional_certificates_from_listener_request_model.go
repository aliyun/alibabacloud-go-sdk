// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateAdditionalCertificatesFromListenerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertificates(v []*DissociateAdditionalCertificatesFromListenerRequestCertificates) *DissociateAdditionalCertificatesFromListenerRequest
	GetCertificates() []*DissociateAdditionalCertificatesFromListenerRequestCertificates
	SetClientToken(v string) *DissociateAdditionalCertificatesFromListenerRequest
	GetClientToken() *string
	SetDryRun(v bool) *DissociateAdditionalCertificatesFromListenerRequest
	GetDryRun() *bool
	SetListenerId(v string) *DissociateAdditionalCertificatesFromListenerRequest
	GetListenerId() *string
}

type DissociateAdditionalCertificatesFromListenerRequest struct {
	// The additional certificates. Only server certificates are supported. You can specify at most 20 certificates.
	//
	// This parameter is required.
	Certificates []*DissociateAdditionalCertificatesFromListenerRequestCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a **2xx HTTP*	- status code is returned and the operation is performed.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The listener ID. You must specify the ID of an HTTPS listener or a QUIC listener.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
}

func (s DissociateAdditionalCertificatesFromListenerRequest) String() string {
	return dara.Prettify(s)
}

func (s DissociateAdditionalCertificatesFromListenerRequest) GoString() string {
	return s.String()
}

func (s *DissociateAdditionalCertificatesFromListenerRequest) GetCertificates() []*DissociateAdditionalCertificatesFromListenerRequestCertificates {
	return s.Certificates
}

func (s *DissociateAdditionalCertificatesFromListenerRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DissociateAdditionalCertificatesFromListenerRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DissociateAdditionalCertificatesFromListenerRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *DissociateAdditionalCertificatesFromListenerRequest) SetCertificates(v []*DissociateAdditionalCertificatesFromListenerRequestCertificates) *DissociateAdditionalCertificatesFromListenerRequest {
	s.Certificates = v
	return s
}

func (s *DissociateAdditionalCertificatesFromListenerRequest) SetClientToken(v string) *DissociateAdditionalCertificatesFromListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *DissociateAdditionalCertificatesFromListenerRequest) SetDryRun(v bool) *DissociateAdditionalCertificatesFromListenerRequest {
	s.DryRun = &v
	return s
}

func (s *DissociateAdditionalCertificatesFromListenerRequest) SetListenerId(v string) *DissociateAdditionalCertificatesFromListenerRequest {
	s.ListenerId = &v
	return s
}

func (s *DissociateAdditionalCertificatesFromListenerRequest) Validate() error {
	return dara.Validate(s)
}

type DissociateAdditionalCertificatesFromListenerRequestCertificates struct {
	// The certificate ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12315790343_166f8204689_1714763408_70998****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
}

func (s DissociateAdditionalCertificatesFromListenerRequestCertificates) String() string {
	return dara.Prettify(s)
}

func (s DissociateAdditionalCertificatesFromListenerRequestCertificates) GoString() string {
	return s.String()
}

func (s *DissociateAdditionalCertificatesFromListenerRequestCertificates) GetCertificateId() *string {
	return s.CertificateId
}

func (s *DissociateAdditionalCertificatesFromListenerRequestCertificates) SetCertificateId(v string) *DissociateAdditionalCertificatesFromListenerRequestCertificates {
	s.CertificateId = &v
	return s
}

func (s *DissociateAdditionalCertificatesFromListenerRequestCertificates) Validate() error {
	return dara.Validate(s)
}
