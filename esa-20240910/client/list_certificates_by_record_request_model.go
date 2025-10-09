// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCertificatesByRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDetail(v bool) *ListCertificatesByRecordRequest
	GetDetail() *bool
	SetRecordName(v string) *ListCertificatesByRecordRequest
	GetRecordName() *string
	SetSiteId(v int64) *ListCertificatesByRecordRequest
	GetSiteId() *int64
	SetValidOnly(v bool) *ListCertificatesByRecordRequest
	GetValidOnly() *bool
}

type ListCertificatesByRecordRequest struct {
	// Specifies whether to return the certificate details. 0 indicates that the certificate details are not returned. 1 indicates that the certificate details are returned.
	//
	// example:
	//
	// 1
	Detail *bool `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The record name.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](~~ListSites~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// Specifies whether to return only valid certificates. 0 indicates that all matched certificates are returned. 1 indicates that only valid certificates are returned.
	//
	// example:
	//
	// 1
	ValidOnly *bool `json:"ValidOnly,omitempty" xml:"ValidOnly,omitempty"`
}

func (s ListCertificatesByRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCertificatesByRecordRequest) GoString() string {
	return s.String()
}

func (s *ListCertificatesByRecordRequest) GetDetail() *bool {
	return s.Detail
}

func (s *ListCertificatesByRecordRequest) GetRecordName() *string {
	return s.RecordName
}

func (s *ListCertificatesByRecordRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListCertificatesByRecordRequest) GetValidOnly() *bool {
	return s.ValidOnly
}

func (s *ListCertificatesByRecordRequest) SetDetail(v bool) *ListCertificatesByRecordRequest {
	s.Detail = &v
	return s
}

func (s *ListCertificatesByRecordRequest) SetRecordName(v string) *ListCertificatesByRecordRequest {
	s.RecordName = &v
	return s
}

func (s *ListCertificatesByRecordRequest) SetSiteId(v int64) *ListCertificatesByRecordRequest {
	s.SiteId = &v
	return s
}

func (s *ListCertificatesByRecordRequest) SetValidOnly(v bool) *ListCertificatesByRecordRequest {
	s.ValidOnly = &v
	return s
}

func (s *ListCertificatesByRecordRequest) Validate() error {
	return dara.Validate(s)
}
