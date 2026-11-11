// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateClientCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *ActivateClientCertificateRequest
	GetId() *string
	SetSiteId(v int64) *ActivateClientCertificateRequest
	GetSiteId() *int64
}

type ActivateClientCertificateRequest struct {
	// The certificate ID. You can call the [ListClientCertificates](https://help.aliyun.com/document_detail/2852848.html) operation to obtain the certificate ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// babaded901474b9693acf530e0fb****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The site ID. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ActivateClientCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s ActivateClientCertificateRequest) GoString() string {
	return s.String()
}

func (s *ActivateClientCertificateRequest) GetId() *string {
	return s.Id
}

func (s *ActivateClientCertificateRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ActivateClientCertificateRequest) SetId(v string) *ActivateClientCertificateRequest {
	s.Id = &v
	return s
}

func (s *ActivateClientCertificateRequest) SetSiteId(v int64) *ActivateClientCertificateRequest {
	s.SiteId = &v
	return s
}

func (s *ActivateClientCertificateRequest) Validate() error {
	return dara.Validate(s)
}
