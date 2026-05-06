// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClientCaCertificateHostnamesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetClientCaCertificateHostnamesRequest
	GetId() *string
	SetSiteId(v int64) *GetClientCaCertificateHostnamesRequest
	GetSiteId() *int64
}

type GetClientCaCertificateHostnamesRequest struct {
	// example:
	//
	// babaded901474b9693acf530e0fb****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetClientCaCertificateHostnamesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetClientCaCertificateHostnamesRequest) GoString() string {
	return s.String()
}

func (s *GetClientCaCertificateHostnamesRequest) GetId() *string {
	return s.Id
}

func (s *GetClientCaCertificateHostnamesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetClientCaCertificateHostnamesRequest) SetId(v string) *GetClientCaCertificateHostnamesRequest {
	s.Id = &v
	return s
}

func (s *GetClientCaCertificateHostnamesRequest) SetSiteId(v int64) *GetClientCaCertificateHostnamesRequest {
	s.SiteId = &v
	return s
}

func (s *GetClientCaCertificateHostnamesRequest) Validate() error {
	return dara.Validate(s)
}
