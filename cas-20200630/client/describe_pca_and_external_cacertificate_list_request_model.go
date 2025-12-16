// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePcaAndExternalCACertificateListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribePcaAndExternalCACertificateListRequest
	GetCurrentPage() *int32
	SetShowSize(v int32) *DescribePcaAndExternalCACertificateListRequest
	GetShowSize() *int32
}

type DescribePcaAndExternalCACertificateListRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 50
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
}

func (s DescribePcaAndExternalCACertificateListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePcaAndExternalCACertificateListRequest) GoString() string {
	return s.String()
}

func (s *DescribePcaAndExternalCACertificateListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribePcaAndExternalCACertificateListRequest) GetShowSize() *int32 {
	return s.ShowSize
}

func (s *DescribePcaAndExternalCACertificateListRequest) SetCurrentPage(v int32) *DescribePcaAndExternalCACertificateListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePcaAndExternalCACertificateListRequest) SetShowSize(v int32) *DescribePcaAndExternalCACertificateListRequest {
	s.ShowSize = &v
	return s
}

func (s *DescribePcaAndExternalCACertificateListRequest) Validate() error {
	return dara.Validate(s)
}
