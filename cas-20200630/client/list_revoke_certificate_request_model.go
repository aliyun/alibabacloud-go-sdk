// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRevokeCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListRevokeCertificateRequest
	GetCurrentPage() *int32
	SetShowSize(v int32) *ListRevokeCertificateRequest
	GetShowSize() *int32
}

type ListRevokeCertificateRequest struct {
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of revoked certificates to return on each page. Default value: **20**.
	//
	// example:
	//
	// 20
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
}

func (s ListRevokeCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRevokeCertificateRequest) GoString() string {
	return s.String()
}

func (s *ListRevokeCertificateRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListRevokeCertificateRequest) GetShowSize() *int32 {
	return s.ShowSize
}

func (s *ListRevokeCertificateRequest) SetCurrentPage(v int32) *ListRevokeCertificateRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListRevokeCertificateRequest) SetShowSize(v int32) *ListRevokeCertificateRequest {
	s.ShowSize = &v
	return s
}

func (s *ListRevokeCertificateRequest) Validate() error {
	return dara.Validate(s)
}
