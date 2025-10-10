// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateAdditionalCertificatesWithListenerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertificates(v []*AssociateAdditionalCertificatesWithListenerRequestCertificates) *AssociateAdditionalCertificatesWithListenerRequest
	GetCertificates() []*AssociateAdditionalCertificatesWithListenerRequestCertificates
	SetClientToken(v string) *AssociateAdditionalCertificatesWithListenerRequest
	GetClientToken() *string
	SetDryRun(v bool) *AssociateAdditionalCertificatesWithListenerRequest
	GetDryRun() *bool
	SetListenerId(v string) *AssociateAdditionalCertificatesWithListenerRequest
	GetListenerId() *string
}

type AssociateAdditionalCertificatesWithListenerRequest struct {
	// The extended validation certificates that you want to add to the listener.
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
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a `2xx HTTP` status code is returned and the operation is performed.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The listener ID. This parameter is supported only by HTTPS and QUIC listeners.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
}

func (s AssociateAdditionalCertificatesWithListenerRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateAdditionalCertificatesWithListenerRequest) GoString() string {
	return s.String()
}

func (s *AssociateAdditionalCertificatesWithListenerRequest) GetCertificates() []*AssociateAdditionalCertificatesWithListenerRequestCertificates {
	return s.Certificates
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

func (s *AssociateAdditionalCertificatesWithListenerRequest) SetCertificates(v []*AssociateAdditionalCertificatesWithListenerRequestCertificates) *AssociateAdditionalCertificatesWithListenerRequest {
	s.Certificates = v
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

func (s *AssociateAdditionalCertificatesWithListenerRequest) Validate() error {
	return dara.Validate(s)
}

type AssociateAdditionalCertificatesWithListenerRequestCertificates struct {
	// The ID of the certificate. Only server certificates are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// cert-123
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
}

func (s AssociateAdditionalCertificatesWithListenerRequestCertificates) String() string {
	return dara.Prettify(s)
}

func (s AssociateAdditionalCertificatesWithListenerRequestCertificates) GoString() string {
	return s.String()
}

func (s *AssociateAdditionalCertificatesWithListenerRequestCertificates) GetCertificateId() *string {
	return s.CertificateId
}

func (s *AssociateAdditionalCertificatesWithListenerRequestCertificates) SetCertificateId(v string) *AssociateAdditionalCertificatesWithListenerRequestCertificates {
	s.CertificateId = &v
	return s
}

func (s *AssociateAdditionalCertificatesWithListenerRequestCertificates) Validate() error {
	return dara.Validate(s)
}
