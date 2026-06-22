// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebLockBindListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeWebLockBindListRequest
	GetCurrentPage() *int32
	SetLang(v string) *DescribeWebLockBindListRequest
	GetLang() *string
	SetPageSize(v int32) *DescribeWebLockBindListRequest
	GetPageSize() *int32
	SetRemark(v string) *DescribeWebLockBindListRequest
	GetRemark() *string
	SetSourceIp(v string) *DescribeWebLockBindListRequest
	GetSourceIp() *string
	SetStatus(v string) *DescribeWebLockBindListRequest
	GetStatus() *string
	SetUuid(v string) *DescribeWebLockBindListRequest
	GetUuid() *string
}

type DescribeWebLockBindListRequest struct {
	// The page number of the current page in a paging query. Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The language of the request and response. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of entries per page in a paging query. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The fuzzy match field for the server. The value can be a server name or IP address.
	//
	// example:
	//
	// 192.168.XX.XX
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The IP address of the access source.
	//
	// example:
	//
	// 116.30.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The protection status of the servers that you want to query. Valid values:
	//
	// - **on**: Protection is enabled.
	//
	// - **off**: Protection is disabled.
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The UUID of the asset that you want to query.
	//
	// > Call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to obtain this parameter.
	//
	// example:
	//
	// 7151f27e-1d51-4e98-a540-8936a****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeWebLockBindListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebLockBindListRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebLockBindListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeWebLockBindListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeWebLockBindListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeWebLockBindListRequest) GetRemark() *string {
	return s.Remark
}

func (s *DescribeWebLockBindListRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeWebLockBindListRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeWebLockBindListRequest) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeWebLockBindListRequest) SetCurrentPage(v int32) *DescribeWebLockBindListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWebLockBindListRequest) SetLang(v string) *DescribeWebLockBindListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeWebLockBindListRequest) SetPageSize(v int32) *DescribeWebLockBindListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeWebLockBindListRequest) SetRemark(v string) *DescribeWebLockBindListRequest {
	s.Remark = &v
	return s
}

func (s *DescribeWebLockBindListRequest) SetSourceIp(v string) *DescribeWebLockBindListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeWebLockBindListRequest) SetStatus(v string) *DescribeWebLockBindListRequest {
	s.Status = &v
	return s
}

func (s *DescribeWebLockBindListRequest) SetUuid(v string) *DescribeWebLockBindListRequest {
	s.Uuid = &v
	return s
}

func (s *DescribeWebLockBindListRequest) Validate() error {
	return dara.Validate(s)
}
