// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertySoftwareDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribePropertySoftwareDetailRequest
	GetCurrentPage() *int32
	SetExtend(v string) *DescribePropertySoftwareDetailRequest
	GetExtend() *string
	SetInstallTimeEnd(v int64) *DescribePropertySoftwareDetailRequest
	GetInstallTimeEnd() *int64
	SetInstallTimeStart(v int64) *DescribePropertySoftwareDetailRequest
	GetInstallTimeStart() *int64
	SetName(v string) *DescribePropertySoftwareDetailRequest
	GetName() *string
	SetNextToken(v string) *DescribePropertySoftwareDetailRequest
	GetNextToken() *string
	SetPageSize(v int32) *DescribePropertySoftwareDetailRequest
	GetPageSize() *int32
	SetPath(v string) *DescribePropertySoftwareDetailRequest
	GetPath() *string
	SetRemark(v string) *DescribePropertySoftwareDetailRequest
	GetRemark() *string
	SetSoftwareVersion(v string) *DescribePropertySoftwareDetailRequest
	GetSoftwareVersion() *string
	SetUseNextToken(v bool) *DescribePropertySoftwareDetailRequest
	GetUseNextToken() *bool
	SetUuid(v string) *DescribePropertySoftwareDetailRequest
	GetUuid() *string
}

type DescribePropertySoftwareDetailRequest struct {
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Specifies whether fuzzy search by software name is supported. If you want to use fuzzy search, set the parameter to 1. If you set the parameter to a different value or leave the parameter empty, fuzzy search is not supported.
	//
	// example:
	//
	// 1
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// The timestamp generated when the software update ends. Unit: milliseconds.
	//
	// example:
	//
	// 1650012695000
	InstallTimeEnd *int64 `json:"InstallTimeEnd,omitempty" xml:"InstallTimeEnd,omitempty"`
	// The timestamp generated when the software update starts. Unit: milliseconds.
	//
	// example:
	//
	// 1649321495000
	InstallTimeStart *int64 `json:"InstallTimeStart,omitempty" xml:"InstallTimeStart,omitempty"`
	// The name of the software.
	//
	// example:
	//
	// kernel
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The installation path of the software.
	//
	// example:
	//
	// /etc/test
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The name or IP address of the server.
	//
	// example:
	//
	// 192.168.XX.XX
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The version of the software.
	//
	// example:
	//
	// 3.10.0
	SoftwareVersion *string `json:"SoftwareVersion,omitempty" xml:"SoftwareVersion,omitempty"`
	UseNextToken    *bool   `json:"UseNextToken,omitempty" xml:"UseNextToken,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// 50d213b4-3a35-427a-b8a5-****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribePropertySoftwareDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertySoftwareDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribePropertySoftwareDetailRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribePropertySoftwareDetailRequest) GetExtend() *string {
	return s.Extend
}

func (s *DescribePropertySoftwareDetailRequest) GetInstallTimeEnd() *int64 {
	return s.InstallTimeEnd
}

func (s *DescribePropertySoftwareDetailRequest) GetInstallTimeStart() *int64 {
	return s.InstallTimeStart
}

func (s *DescribePropertySoftwareDetailRequest) GetName() *string {
	return s.Name
}

func (s *DescribePropertySoftwareDetailRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribePropertySoftwareDetailRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePropertySoftwareDetailRequest) GetPath() *string {
	return s.Path
}

func (s *DescribePropertySoftwareDetailRequest) GetRemark() *string {
	return s.Remark
}

func (s *DescribePropertySoftwareDetailRequest) GetSoftwareVersion() *string {
	return s.SoftwareVersion
}

func (s *DescribePropertySoftwareDetailRequest) GetUseNextToken() *bool {
	return s.UseNextToken
}

func (s *DescribePropertySoftwareDetailRequest) GetUuid() *string {
	return s.Uuid
}

func (s *DescribePropertySoftwareDetailRequest) SetCurrentPage(v int32) *DescribePropertySoftwareDetailRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertySoftwareDetailRequest) SetExtend(v string) *DescribePropertySoftwareDetailRequest {
	s.Extend = &v
	return s
}

func (s *DescribePropertySoftwareDetailRequest) SetInstallTimeEnd(v int64) *DescribePropertySoftwareDetailRequest {
	s.InstallTimeEnd = &v
	return s
}

func (s *DescribePropertySoftwareDetailRequest) SetInstallTimeStart(v int64) *DescribePropertySoftwareDetailRequest {
	s.InstallTimeStart = &v
	return s
}

func (s *DescribePropertySoftwareDetailRequest) SetName(v string) *DescribePropertySoftwareDetailRequest {
	s.Name = &v
	return s
}

func (s *DescribePropertySoftwareDetailRequest) SetNextToken(v string) *DescribePropertySoftwareDetailRequest {
	s.NextToken = &v
	return s
}

func (s *DescribePropertySoftwareDetailRequest) SetPageSize(v int32) *DescribePropertySoftwareDetailRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePropertySoftwareDetailRequest) SetPath(v string) *DescribePropertySoftwareDetailRequest {
	s.Path = &v
	return s
}

func (s *DescribePropertySoftwareDetailRequest) SetRemark(v string) *DescribePropertySoftwareDetailRequest {
	s.Remark = &v
	return s
}

func (s *DescribePropertySoftwareDetailRequest) SetSoftwareVersion(v string) *DescribePropertySoftwareDetailRequest {
	s.SoftwareVersion = &v
	return s
}

func (s *DescribePropertySoftwareDetailRequest) SetUseNextToken(v bool) *DescribePropertySoftwareDetailRequest {
	s.UseNextToken = &v
	return s
}

func (s *DescribePropertySoftwareDetailRequest) SetUuid(v string) *DescribePropertySoftwareDetailRequest {
	s.Uuid = &v
	return s
}

func (s *DescribePropertySoftwareDetailRequest) Validate() error {
	return dara.Validate(s)
}
