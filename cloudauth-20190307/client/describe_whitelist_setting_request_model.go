// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWhitelistSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertNo(v string) *DescribeWhitelistSettingRequest
	GetCertNo() *string
	SetCertifyId(v string) *DescribeWhitelistSettingRequest
	GetCertifyId() *string
	SetCurrentPage(v int32) *DescribeWhitelistSettingRequest
	GetCurrentPage() *int32
	SetLang(v string) *DescribeWhitelistSettingRequest
	GetLang() *string
	SetPageSize(v int32) *DescribeWhitelistSettingRequest
	GetPageSize() *int32
	SetSceneId(v int64) *DescribeWhitelistSettingRequest
	GetSceneId() *int64
	SetServiceCode(v string) *DescribeWhitelistSettingRequest
	GetServiceCode() *string
	SetSourceIp(v string) *DescribeWhitelistSettingRequest
	GetSourceIp() *string
	SetStatus(v string) *DescribeWhitelistSettingRequest
	GetStatus() *string
	SetValidEndDate(v int64) *DescribeWhitelistSettingRequest
	GetValidEndDate() *int64
	SetValidStartDate(v int64) *DescribeWhitelistSettingRequest
	GetValidStartDate() *int64
}

type DescribeWhitelistSettingRequest struct {
	// ID Number
	//
	// example:
	//
	// 320321XXXXXXXX701X
	CertNo *string `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	// Certification ID
	//
	// example:
	//
	// shsf57a4e0d9981c3bd66dc754f3d3cd
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	// Pagination parameter: current page number, default value is 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Specify the language to query. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Number of items per page for pagination.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Scene ID.
	//
	// example:
	//
	// 1000004530
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// Service Code:
	//
	// - **Enhanced Financial Grade**: cloudauthst
	//
	// - **Financial Grade**: antcloudauth
	//
	// This parameter is required.
	//
	// example:
	//
	// antcloudauth
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// Visitor\\"s source IP address.
	//
	// example:
	//
	// 113.140.85.74
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// Whitelist status:
	//
	// - **VALID**: Valid
	//
	// - **INVALID**: Invalid
	//
	// - **DELETED**: Deleted
	//
	// example:
	//
	// VALID
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Expiration date.
	//
	// example:
	//
	// 1730304000000
	ValidEndDate *int64 `json:"ValidEndDate,omitempty" xml:"ValidEndDate,omitempty"`
	// Effective start time (in seconds timestamp).
	//
	// example:
	//
	// 1759939200000
	ValidStartDate *int64 `json:"ValidStartDate,omitempty" xml:"ValidStartDate,omitempty"`
}

func (s DescribeWhitelistSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhitelistSettingRequest) GoString() string {
	return s.String()
}

func (s *DescribeWhitelistSettingRequest) GetCertNo() *string {
	return s.CertNo
}

func (s *DescribeWhitelistSettingRequest) GetCertifyId() *string {
	return s.CertifyId
}

func (s *DescribeWhitelistSettingRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeWhitelistSettingRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeWhitelistSettingRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeWhitelistSettingRequest) GetSceneId() *int64 {
	return s.SceneId
}

func (s *DescribeWhitelistSettingRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *DescribeWhitelistSettingRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeWhitelistSettingRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeWhitelistSettingRequest) GetValidEndDate() *int64 {
	return s.ValidEndDate
}

func (s *DescribeWhitelistSettingRequest) GetValidStartDate() *int64 {
	return s.ValidStartDate
}

func (s *DescribeWhitelistSettingRequest) SetCertNo(v string) *DescribeWhitelistSettingRequest {
	s.CertNo = &v
	return s
}

func (s *DescribeWhitelistSettingRequest) SetCertifyId(v string) *DescribeWhitelistSettingRequest {
	s.CertifyId = &v
	return s
}

func (s *DescribeWhitelistSettingRequest) SetCurrentPage(v int32) *DescribeWhitelistSettingRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWhitelistSettingRequest) SetLang(v string) *DescribeWhitelistSettingRequest {
	s.Lang = &v
	return s
}

func (s *DescribeWhitelistSettingRequest) SetPageSize(v int32) *DescribeWhitelistSettingRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeWhitelistSettingRequest) SetSceneId(v int64) *DescribeWhitelistSettingRequest {
	s.SceneId = &v
	return s
}

func (s *DescribeWhitelistSettingRequest) SetServiceCode(v string) *DescribeWhitelistSettingRequest {
	s.ServiceCode = &v
	return s
}

func (s *DescribeWhitelistSettingRequest) SetSourceIp(v string) *DescribeWhitelistSettingRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeWhitelistSettingRequest) SetStatus(v string) *DescribeWhitelistSettingRequest {
	s.Status = &v
	return s
}

func (s *DescribeWhitelistSettingRequest) SetValidEndDate(v int64) *DescribeWhitelistSettingRequest {
	s.ValidEndDate = &v
	return s
}

func (s *DescribeWhitelistSettingRequest) SetValidStartDate(v int64) *DescribeWhitelistSettingRequest {
	s.ValidStartDate = &v
	return s
}

func (s *DescribeWhitelistSettingRequest) Validate() error {
	return dara.Validate(s)
}
