// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMgsApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiStatus(v string) *ListMgsApiRequest
	GetApiStatus() *string
	SetApiType(v string) *ListMgsApiRequest
	GetApiType() *string
	SetAppId(v string) *ListMgsApiRequest
	GetAppId() *string
	SetFormat(v string) *ListMgsApiRequest
	GetFormat() *string
	SetHost(v string) *ListMgsApiRequest
	GetHost() *string
	SetNeedEncrypt(v string) *ListMgsApiRequest
	GetNeedEncrypt() *string
	SetNeedEtag(v string) *ListMgsApiRequest
	GetNeedEtag() *string
	SetNeedSign(v string) *ListMgsApiRequest
	GetNeedSign() *string
	SetOperationType(v string) *ListMgsApiRequest
	GetOperationType() *string
	SetOptFuzzy(v string) *ListMgsApiRequest
	GetOptFuzzy() *string
	SetPageIndex(v int64) *ListMgsApiRequest
	GetPageIndex() *int64
	SetPageSize(v int64) *ListMgsApiRequest
	GetPageSize() *int64
	SetSysId(v int64) *ListMgsApiRequest
	GetSysId() *int64
	SetSysName(v string) *ListMgsApiRequest
	GetSysName() *string
	SetTenantId(v string) *ListMgsApiRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *ListMgsApiRequest
	GetWorkspaceId() *string
}

type ListMgsApiRequest struct {
	ApiStatus     *string `json:"ApiStatus,omitempty" xml:"ApiStatus,omitempty"`
	ApiType       *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Format        *string `json:"Format,omitempty" xml:"Format,omitempty"`
	Host          *string `json:"Host,omitempty" xml:"Host,omitempty"`
	NeedEncrypt   *string `json:"NeedEncrypt,omitempty" xml:"NeedEncrypt,omitempty"`
	NeedEtag      *string `json:"NeedEtag,omitempty" xml:"NeedEtag,omitempty"`
	NeedSign      *string `json:"NeedSign,omitempty" xml:"NeedSign,omitempty"`
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	OptFuzzy      *string `json:"OptFuzzy,omitempty" xml:"OptFuzzy,omitempty"`
	PageIndex     *int64  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize      *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SysId         *int64  `json:"SysId,omitempty" xml:"SysId,omitempty"`
	SysName       *string `json:"SysName,omitempty" xml:"SysName,omitempty"`
	TenantId      *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	WorkspaceId   *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListMgsApiRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMgsApiRequest) GoString() string {
	return s.String()
}

func (s *ListMgsApiRequest) GetApiStatus() *string {
	return s.ApiStatus
}

func (s *ListMgsApiRequest) GetApiType() *string {
	return s.ApiType
}

func (s *ListMgsApiRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListMgsApiRequest) GetFormat() *string {
	return s.Format
}

func (s *ListMgsApiRequest) GetHost() *string {
	return s.Host
}

func (s *ListMgsApiRequest) GetNeedEncrypt() *string {
	return s.NeedEncrypt
}

func (s *ListMgsApiRequest) GetNeedEtag() *string {
	return s.NeedEtag
}

func (s *ListMgsApiRequest) GetNeedSign() *string {
	return s.NeedSign
}

func (s *ListMgsApiRequest) GetOperationType() *string {
	return s.OperationType
}

func (s *ListMgsApiRequest) GetOptFuzzy() *string {
	return s.OptFuzzy
}

func (s *ListMgsApiRequest) GetPageIndex() *int64 {
	return s.PageIndex
}

func (s *ListMgsApiRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListMgsApiRequest) GetSysId() *int64 {
	return s.SysId
}

func (s *ListMgsApiRequest) GetSysName() *string {
	return s.SysName
}

func (s *ListMgsApiRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListMgsApiRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListMgsApiRequest) SetApiStatus(v string) *ListMgsApiRequest {
	s.ApiStatus = &v
	return s
}

func (s *ListMgsApiRequest) SetApiType(v string) *ListMgsApiRequest {
	s.ApiType = &v
	return s
}

func (s *ListMgsApiRequest) SetAppId(v string) *ListMgsApiRequest {
	s.AppId = &v
	return s
}

func (s *ListMgsApiRequest) SetFormat(v string) *ListMgsApiRequest {
	s.Format = &v
	return s
}

func (s *ListMgsApiRequest) SetHost(v string) *ListMgsApiRequest {
	s.Host = &v
	return s
}

func (s *ListMgsApiRequest) SetNeedEncrypt(v string) *ListMgsApiRequest {
	s.NeedEncrypt = &v
	return s
}

func (s *ListMgsApiRequest) SetNeedEtag(v string) *ListMgsApiRequest {
	s.NeedEtag = &v
	return s
}

func (s *ListMgsApiRequest) SetNeedSign(v string) *ListMgsApiRequest {
	s.NeedSign = &v
	return s
}

func (s *ListMgsApiRequest) SetOperationType(v string) *ListMgsApiRequest {
	s.OperationType = &v
	return s
}

func (s *ListMgsApiRequest) SetOptFuzzy(v string) *ListMgsApiRequest {
	s.OptFuzzy = &v
	return s
}

func (s *ListMgsApiRequest) SetPageIndex(v int64) *ListMgsApiRequest {
	s.PageIndex = &v
	return s
}

func (s *ListMgsApiRequest) SetPageSize(v int64) *ListMgsApiRequest {
	s.PageSize = &v
	return s
}

func (s *ListMgsApiRequest) SetSysId(v int64) *ListMgsApiRequest {
	s.SysId = &v
	return s
}

func (s *ListMgsApiRequest) SetSysName(v string) *ListMgsApiRequest {
	s.SysName = &v
	return s
}

func (s *ListMgsApiRequest) SetTenantId(v string) *ListMgsApiRequest {
	s.TenantId = &v
	return s
}

func (s *ListMgsApiRequest) SetWorkspaceId(v string) *ListMgsApiRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListMgsApiRequest) Validate() error {
	return dara.Validate(s)
}
