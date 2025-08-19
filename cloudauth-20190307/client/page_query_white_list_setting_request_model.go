// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPageQueryWhiteListSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertNo(v string) *PageQueryWhiteListSettingRequest
	GetCertNo() *string
	SetCertifyId(v string) *PageQueryWhiteListSettingRequest
	GetCertifyId() *string
	SetCurrentPage(v int32) *PageQueryWhiteListSettingRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *PageQueryWhiteListSettingRequest
	GetPageSize() *int32
	SetSceneId(v int64) *PageQueryWhiteListSettingRequest
	GetSceneId() *int64
	SetServiceCode(v string) *PageQueryWhiteListSettingRequest
	GetServiceCode() *string
	SetStatus(v string) *PageQueryWhiteListSettingRequest
	GetStatus() *string
	SetValidEndDate(v string) *PageQueryWhiteListSettingRequest
	GetValidEndDate() *string
	SetValidStartDate(v string) *PageQueryWhiteListSettingRequest
	GetValidStartDate() *string
}

type PageQueryWhiteListSettingRequest struct {
	// ID number.
	//
	// example:
	//
	// 330103xxxxxxxxxxxx
	CertNo *string `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	// Unique identifier for real person authentication.
	//
	// example:
	//
	// sha75b4e19a1ddda059b920757b0e12b
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	// Current page number, default is 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Number of items per page, default is 10
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Authentication scene ID. This ID is automatically generated after creating an authentication scene in the console. For how to create an authentication scene, see Adding an Authentication Scene.
	//
	// example:
	//
	// 1000000xxx
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// ServiceCode of the real person cloud product, value: **antcloudauth**.
	//
	// example:
	//
	// antcloudauth
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// Status:
	//
	// - DELETE: Deleted
	//
	// - VALID: Not deleted and within the validity period, valid
	//
	// - INVALID: Not deleted but outside the validity period, invalid
	//
	// example:
	//
	// VALID
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// End date of validity (timestamp in milliseconds)
	//
	// example:
	//
	// 1725379200000
	ValidEndDate *string `json:"ValidEndDate,omitempty" xml:"ValidEndDate,omitempty"`
	// Start date of validity (timestamp in milliseconds)
	//
	// example:
	//
	// 1725120000000
	ValidStartDate *string `json:"ValidStartDate,omitempty" xml:"ValidStartDate,omitempty"`
}

func (s PageQueryWhiteListSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s PageQueryWhiteListSettingRequest) GoString() string {
	return s.String()
}

func (s *PageQueryWhiteListSettingRequest) GetCertNo() *string {
	return s.CertNo
}

func (s *PageQueryWhiteListSettingRequest) GetCertifyId() *string {
	return s.CertifyId
}

func (s *PageQueryWhiteListSettingRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *PageQueryWhiteListSettingRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *PageQueryWhiteListSettingRequest) GetSceneId() *int64 {
	return s.SceneId
}

func (s *PageQueryWhiteListSettingRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *PageQueryWhiteListSettingRequest) GetStatus() *string {
	return s.Status
}

func (s *PageQueryWhiteListSettingRequest) GetValidEndDate() *string {
	return s.ValidEndDate
}

func (s *PageQueryWhiteListSettingRequest) GetValidStartDate() *string {
	return s.ValidStartDate
}

func (s *PageQueryWhiteListSettingRequest) SetCertNo(v string) *PageQueryWhiteListSettingRequest {
	s.CertNo = &v
	return s
}

func (s *PageQueryWhiteListSettingRequest) SetCertifyId(v string) *PageQueryWhiteListSettingRequest {
	s.CertifyId = &v
	return s
}

func (s *PageQueryWhiteListSettingRequest) SetCurrentPage(v int32) *PageQueryWhiteListSettingRequest {
	s.CurrentPage = &v
	return s
}

func (s *PageQueryWhiteListSettingRequest) SetPageSize(v int32) *PageQueryWhiteListSettingRequest {
	s.PageSize = &v
	return s
}

func (s *PageQueryWhiteListSettingRequest) SetSceneId(v int64) *PageQueryWhiteListSettingRequest {
	s.SceneId = &v
	return s
}

func (s *PageQueryWhiteListSettingRequest) SetServiceCode(v string) *PageQueryWhiteListSettingRequest {
	s.ServiceCode = &v
	return s
}

func (s *PageQueryWhiteListSettingRequest) SetStatus(v string) *PageQueryWhiteListSettingRequest {
	s.Status = &v
	return s
}

func (s *PageQueryWhiteListSettingRequest) SetValidEndDate(v string) *PageQueryWhiteListSettingRequest {
	s.ValidEndDate = &v
	return s
}

func (s *PageQueryWhiteListSettingRequest) SetValidStartDate(v string) *PageQueryWhiteListSettingRequest {
	s.ValidStartDate = &v
	return s
}

func (s *PageQueryWhiteListSettingRequest) Validate() error {
	return dara.Validate(s)
}
