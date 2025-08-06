// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPageByCondLicenseInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContractNo(v string) *GetPageByCondLicenseInstanceRequest
	GetContractNo() *string
	SetInstanceId(v string) *GetPageByCondLicenseInstanceRequest
	GetInstanceId() *string
	SetNeedTotalCount(v bool) *GetPageByCondLicenseInstanceRequest
	GetNeedTotalCount() *bool
	SetPageNo(v int64) *GetPageByCondLicenseInstanceRequest
	GetPageNo() *int64
	SetPageSize(v int64) *GetPageByCondLicenseInstanceRequest
	GetPageSize() *int64
	SetSortBy(v string) *GetPageByCondLicenseInstanceRequest
	GetSortBy() *string
}

type GetPageByCondLicenseInstanceRequest struct {
	ContractNo *string `json:"ContractNo,omitempty" xml:"ContractNo,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// true
	NeedTotalCount *bool `json:"NeedTotalCount,omitempty" xml:"NeedTotalCount,omitempty"`
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortBy   *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s GetPageByCondLicenseInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPageByCondLicenseInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetPageByCondLicenseInstanceRequest) GetContractNo() *string {
	return s.ContractNo
}

func (s *GetPageByCondLicenseInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetPageByCondLicenseInstanceRequest) GetNeedTotalCount() *bool {
	return s.NeedTotalCount
}

func (s *GetPageByCondLicenseInstanceRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *GetPageByCondLicenseInstanceRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetPageByCondLicenseInstanceRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *GetPageByCondLicenseInstanceRequest) SetContractNo(v string) *GetPageByCondLicenseInstanceRequest {
	s.ContractNo = &v
	return s
}

func (s *GetPageByCondLicenseInstanceRequest) SetInstanceId(v string) *GetPageByCondLicenseInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *GetPageByCondLicenseInstanceRequest) SetNeedTotalCount(v bool) *GetPageByCondLicenseInstanceRequest {
	s.NeedTotalCount = &v
	return s
}

func (s *GetPageByCondLicenseInstanceRequest) SetPageNo(v int64) *GetPageByCondLicenseInstanceRequest {
	s.PageNo = &v
	return s
}

func (s *GetPageByCondLicenseInstanceRequest) SetPageSize(v int64) *GetPageByCondLicenseInstanceRequest {
	s.PageSize = &v
	return s
}

func (s *GetPageByCondLicenseInstanceRequest) SetSortBy(v string) *GetPageByCondLicenseInstanceRequest {
	s.SortBy = &v
	return s
}

func (s *GetPageByCondLicenseInstanceRequest) Validate() error {
	return dara.Validate(s)
}
