// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAdditionalCertificateWithListenerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *UpdateAdditionalCertificateWithListenerRequest
	GetAcceleratorId() *string
	SetCertificateId(v string) *UpdateAdditionalCertificateWithListenerRequest
	GetCertificateId() *string
	SetClientToken(v string) *UpdateAdditionalCertificateWithListenerRequest
	GetClientToken() *string
	SetDomain(v string) *UpdateAdditionalCertificateWithListenerRequest
	GetDomain() *string
	SetDryRun(v bool) *UpdateAdditionalCertificateWithListenerRequest
	GetDryRun() *bool
	SetListenerId(v string) *UpdateAdditionalCertificateWithListenerRequest
	GetListenerId() *string
	SetRegionId(v string) *UpdateAdditionalCertificateWithListenerRequest
	GetRegionId() *string
}

type UpdateAdditionalCertificateWithListenerRequest struct {
	// The ID of the GA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The ID of the replacement certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6092**-cn-hangzhou
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The domain name associated with the additional certificate that you want to replace.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// 	- **true**: performs a dry run. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false:*	- performs a dry run and sends the request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed. This is the default value.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the listener. Only HTTPS listeners are supported.
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

func (s UpdateAdditionalCertificateWithListenerRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAdditionalCertificateWithListenerRequest) GoString() string {
	return s.String()
}

func (s *UpdateAdditionalCertificateWithListenerRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *UpdateAdditionalCertificateWithListenerRequest) GetCertificateId() *string {
	return s.CertificateId
}

func (s *UpdateAdditionalCertificateWithListenerRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateAdditionalCertificateWithListenerRequest) GetDomain() *string {
	return s.Domain
}

func (s *UpdateAdditionalCertificateWithListenerRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateAdditionalCertificateWithListenerRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *UpdateAdditionalCertificateWithListenerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateAdditionalCertificateWithListenerRequest) SetAcceleratorId(v string) *UpdateAdditionalCertificateWithListenerRequest {
	s.AcceleratorId = &v
	return s
}

func (s *UpdateAdditionalCertificateWithListenerRequest) SetCertificateId(v string) *UpdateAdditionalCertificateWithListenerRequest {
	s.CertificateId = &v
	return s
}

func (s *UpdateAdditionalCertificateWithListenerRequest) SetClientToken(v string) *UpdateAdditionalCertificateWithListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAdditionalCertificateWithListenerRequest) SetDomain(v string) *UpdateAdditionalCertificateWithListenerRequest {
	s.Domain = &v
	return s
}

func (s *UpdateAdditionalCertificateWithListenerRequest) SetDryRun(v bool) *UpdateAdditionalCertificateWithListenerRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateAdditionalCertificateWithListenerRequest) SetListenerId(v string) *UpdateAdditionalCertificateWithListenerRequest {
	s.ListenerId = &v
	return s
}

func (s *UpdateAdditionalCertificateWithListenerRequest) SetRegionId(v string) *UpdateAdditionalCertificateWithListenerRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateAdditionalCertificateWithListenerRequest) Validate() error {
	return dara.Validate(s)
}
