// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainByCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeLiveDomainByCertificateRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveDomainByCertificateRequest
	GetRegionId() *string
	SetSSLPub(v string) *DescribeLiveDomainByCertificateRequest
	GetSSLPub() *string
	SetSSLStatus(v bool) *DescribeLiveDomainByCertificateRequest
	GetSSLStatus() *bool
}

type DescribeLiveDomainByCertificateRequest struct {
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The public key of the SSL certificate. You must Base64-encode the public key before you invoke the encodeURIComponent function to encode a URI component. The public key must be in the PEM format.
	//
	// This parameter is required.
	//
	// example:
	//
	// ******
	SSLPub *string `json:"SSLPub,omitempty" xml:"SSLPub,omitempty"`
	// Specifies whether to return only domain names with HTTPS enabled or disabled.
	//
	// 	- **true**: returns only domain names with HTTPS enabled.
	//
	// 	- **false**: The rule is disabled.
	//
	// example:
	//
	// true
	SSLStatus *bool `json:"SSLStatus,omitempty" xml:"SSLStatus,omitempty"`
}

func (s DescribeLiveDomainByCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainByCertificateRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainByCertificateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveDomainByCertificateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveDomainByCertificateRequest) GetSSLPub() *string {
	return s.SSLPub
}

func (s *DescribeLiveDomainByCertificateRequest) GetSSLStatus() *bool {
	return s.SSLStatus
}

func (s *DescribeLiveDomainByCertificateRequest) SetOwnerId(v int64) *DescribeLiveDomainByCertificateRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveDomainByCertificateRequest) SetRegionId(v string) *DescribeLiveDomainByCertificateRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveDomainByCertificateRequest) SetSSLPub(v string) *DescribeLiveDomainByCertificateRequest {
	s.SSLPub = &v
	return s
}

func (s *DescribeLiveDomainByCertificateRequest) SetSSLStatus(v bool) *DescribeLiveDomainByCertificateRequest {
	s.SSLStatus = &v
	return s
}

func (s *DescribeLiveDomainByCertificateRequest) Validate() error {
	return dara.Validate(s)
}
