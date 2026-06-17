// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransportEncryptionInfo interface {
	dara.Model
	String() string
	GoString() string
	SetCertificates(v []*TransportCertificateInfo) *TransportEncryptionInfo
	GetCertificates() []*TransportCertificateInfo
	SetDeployError(v string) *TransportEncryptionInfo
	GetDeployError() *string
	SetDeployStatus(v string) *TransportEncryptionInfo
	GetDeployStatus() *string
	SetHttp2Enabled(v bool) *TransportEncryptionInfo
	GetHttp2Enabled() *bool
	SetTlsPolicy(v string) *TransportEncryptionInfo
	GetTlsPolicy() *string
}

type TransportEncryptionInfo struct {
	Certificates []*TransportCertificateInfo `json:"certificates,omitempty" xml:"certificates,omitempty" type:"Repeated"`
	DeployError  *string                     `json:"deployError,omitempty" xml:"deployError,omitempty"`
	DeployStatus *string                     `json:"deployStatus,omitempty" xml:"deployStatus,omitempty"`
	Http2Enabled *bool                       `json:"http2Enabled,omitempty" xml:"http2Enabled,omitempty"`
	TlsPolicy    *string                     `json:"tlsPolicy,omitempty" xml:"tlsPolicy,omitempty"`
}

func (s TransportEncryptionInfo) String() string {
	return dara.Prettify(s)
}

func (s TransportEncryptionInfo) GoString() string {
	return s.String()
}

func (s *TransportEncryptionInfo) GetCertificates() []*TransportCertificateInfo {
	return s.Certificates
}

func (s *TransportEncryptionInfo) GetDeployError() *string {
	return s.DeployError
}

func (s *TransportEncryptionInfo) GetDeployStatus() *string {
	return s.DeployStatus
}

func (s *TransportEncryptionInfo) GetHttp2Enabled() *bool {
	return s.Http2Enabled
}

func (s *TransportEncryptionInfo) GetTlsPolicy() *string {
	return s.TlsPolicy
}

func (s *TransportEncryptionInfo) SetCertificates(v []*TransportCertificateInfo) *TransportEncryptionInfo {
	s.Certificates = v
	return s
}

func (s *TransportEncryptionInfo) SetDeployError(v string) *TransportEncryptionInfo {
	s.DeployError = &v
	return s
}

func (s *TransportEncryptionInfo) SetDeployStatus(v string) *TransportEncryptionInfo {
	s.DeployStatus = &v
	return s
}

func (s *TransportEncryptionInfo) SetHttp2Enabled(v bool) *TransportEncryptionInfo {
	s.Http2Enabled = &v
	return s
}

func (s *TransportEncryptionInfo) SetTlsPolicy(v string) *TransportEncryptionInfo {
	s.TlsPolicy = &v
	return s
}

func (s *TransportEncryptionInfo) Validate() error {
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
