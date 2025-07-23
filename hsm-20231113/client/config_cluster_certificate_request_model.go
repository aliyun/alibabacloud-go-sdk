// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigClusterCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterCertificate(v string) *ConfigClusterCertificateRequest
	GetClusterCertificate() *string
	SetClusterId(v string) *ConfigClusterCertificateRequest
	GetClusterId() *string
	SetIssuerCertificate(v string) *ConfigClusterCertificateRequest
	GetIssuerCertificate() *string
}

type ConfigClusterCertificateRequest struct {
	// The cluster certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	//
	// MIIDaTCCAlECAQEwDQYJKoZIhvcNAQELBQAwVTELMAkGA1UEBhMCY24xCzAJBgNV
	//
	// BAgMAnpqMQswCQYDVQQHDAJoejEWMBQGA1UECgwNQWxpYmFiYSBDbG91ZDEUMBIG
	//
	// A1UECwwLU2VjQ2xvdWRIc20wHhcNMjQwNzAzM****-----END CERTIFICATE-----
	ClusterCertificate *string `json:"ClusterCertificate,omitempty" xml:"ClusterCertificate,omitempty"`
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cluster-BqxX63Bsg****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The self-signed certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	//
	// MIIDfTCCAmWgAwIBAgIJAMRqQMr5if66MA0GCSqGSIb3DQEBCwUAMFUxCzAJBgNV
	//
	// BAYTAmNuMQswCQYDVQQIDAJ6ajELMAkGA1UEBwwCaHoxFjAUBgNVBAoMDUFsaWJh
	//
	// YmEgQ2xvdWQxFDA****
	//
	// -----END CERTIFICATE-----
	IssuerCertificate *string `json:"IssuerCertificate,omitempty" xml:"IssuerCertificate,omitempty"`
}

func (s ConfigClusterCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigClusterCertificateRequest) GoString() string {
	return s.String()
}

func (s *ConfigClusterCertificateRequest) GetClusterCertificate() *string {
	return s.ClusterCertificate
}

func (s *ConfigClusterCertificateRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ConfigClusterCertificateRequest) GetIssuerCertificate() *string {
	return s.IssuerCertificate
}

func (s *ConfigClusterCertificateRequest) SetClusterCertificate(v string) *ConfigClusterCertificateRequest {
	s.ClusterCertificate = &v
	return s
}

func (s *ConfigClusterCertificateRequest) SetClusterId(v string) *ConfigClusterCertificateRequest {
	s.ClusterId = &v
	return s
}

func (s *ConfigClusterCertificateRequest) SetIssuerCertificate(v string) *ConfigClusterCertificateRequest {
	s.IssuerCertificate = &v
	return s
}

func (s *ConfigClusterCertificateRequest) Validate() error {
	return dara.Validate(s)
}
