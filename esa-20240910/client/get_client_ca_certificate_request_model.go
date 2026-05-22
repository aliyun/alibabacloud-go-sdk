// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClientCaCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetClientCaCertificateRequest
	GetId() *string
	SetSiteId(v int64) *GetClientCaCertificateRequest
	GetSiteId() *int64
}

type GetClientCaCertificateRequest struct {
	// The certificate ID, which can be obtained by calling the [ListClientCaCertificates](https://help.aliyun.com/document_detail/2860651.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// babab9db65ee5efcca9f3d41d4b5****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetClientCaCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetClientCaCertificateRequest) GoString() string {
	return s.String()
}

func (s *GetClientCaCertificateRequest) GetId() *string {
	return s.Id
}

func (s *GetClientCaCertificateRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetClientCaCertificateRequest) SetId(v string) *GetClientCaCertificateRequest {
	s.Id = &v
	return s
}

func (s *GetClientCaCertificateRequest) SetSiteId(v int64) *GetClientCaCertificateRequest {
	s.SiteId = &v
	return s
}

func (s *GetClientCaCertificateRequest) Validate() error {
	return dara.Validate(s)
}
