// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMgsApipageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiStatus(v string) *QueryMgsApipageRequest
	GetApiStatus() *string
	SetApiType(v string) *QueryMgsApipageRequest
	GetApiType() *string
	SetAppId(v string) *QueryMgsApipageRequest
	GetAppId() *string
	SetFormat(v string) *QueryMgsApipageRequest
	GetFormat() *string
	SetHost(v string) *QueryMgsApipageRequest
	GetHost() *string
	SetNeedEncrypt(v string) *QueryMgsApipageRequest
	GetNeedEncrypt() *string
	SetNeedEtag(v string) *QueryMgsApipageRequest
	GetNeedEtag() *string
	SetNeedSign(v string) *QueryMgsApipageRequest
	GetNeedSign() *string
	SetOperationType(v string) *QueryMgsApipageRequest
	GetOperationType() *string
	SetOptFuzzy(v string) *QueryMgsApipageRequest
	GetOptFuzzy() *string
	SetPageIndex(v int64) *QueryMgsApipageRequest
	GetPageIndex() *int64
	SetPageSize(v int64) *QueryMgsApipageRequest
	GetPageSize() *int64
	SetSysId(v int64) *QueryMgsApipageRequest
	GetSysId() *int64
	SetSysName(v string) *QueryMgsApipageRequest
	GetSysName() *string
	SetTenantId(v string) *QueryMgsApipageRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *QueryMgsApipageRequest
	GetWorkspaceId() *string
}

type QueryMgsApipageRequest struct {
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

func (s QueryMgsApipageRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMgsApipageRequest) GoString() string {
	return s.String()
}

func (s *QueryMgsApipageRequest) GetApiStatus() *string {
	return s.ApiStatus
}

func (s *QueryMgsApipageRequest) GetApiType() *string {
	return s.ApiType
}

func (s *QueryMgsApipageRequest) GetAppId() *string {
	return s.AppId
}

func (s *QueryMgsApipageRequest) GetFormat() *string {
	return s.Format
}

func (s *QueryMgsApipageRequest) GetHost() *string {
	return s.Host
}

func (s *QueryMgsApipageRequest) GetNeedEncrypt() *string {
	return s.NeedEncrypt
}

func (s *QueryMgsApipageRequest) GetNeedEtag() *string {
	return s.NeedEtag
}

func (s *QueryMgsApipageRequest) GetNeedSign() *string {
	return s.NeedSign
}

func (s *QueryMgsApipageRequest) GetOperationType() *string {
	return s.OperationType
}

func (s *QueryMgsApipageRequest) GetOptFuzzy() *string {
	return s.OptFuzzy
}

func (s *QueryMgsApipageRequest) GetPageIndex() *int64 {
	return s.PageIndex
}

func (s *QueryMgsApipageRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryMgsApipageRequest) GetSysId() *int64 {
	return s.SysId
}

func (s *QueryMgsApipageRequest) GetSysName() *string {
	return s.SysName
}

func (s *QueryMgsApipageRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryMgsApipageRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryMgsApipageRequest) SetApiStatus(v string) *QueryMgsApipageRequest {
	s.ApiStatus = &v
	return s
}

func (s *QueryMgsApipageRequest) SetApiType(v string) *QueryMgsApipageRequest {
	s.ApiType = &v
	return s
}

func (s *QueryMgsApipageRequest) SetAppId(v string) *QueryMgsApipageRequest {
	s.AppId = &v
	return s
}

func (s *QueryMgsApipageRequest) SetFormat(v string) *QueryMgsApipageRequest {
	s.Format = &v
	return s
}

func (s *QueryMgsApipageRequest) SetHost(v string) *QueryMgsApipageRequest {
	s.Host = &v
	return s
}

func (s *QueryMgsApipageRequest) SetNeedEncrypt(v string) *QueryMgsApipageRequest {
	s.NeedEncrypt = &v
	return s
}

func (s *QueryMgsApipageRequest) SetNeedEtag(v string) *QueryMgsApipageRequest {
	s.NeedEtag = &v
	return s
}

func (s *QueryMgsApipageRequest) SetNeedSign(v string) *QueryMgsApipageRequest {
	s.NeedSign = &v
	return s
}

func (s *QueryMgsApipageRequest) SetOperationType(v string) *QueryMgsApipageRequest {
	s.OperationType = &v
	return s
}

func (s *QueryMgsApipageRequest) SetOptFuzzy(v string) *QueryMgsApipageRequest {
	s.OptFuzzy = &v
	return s
}

func (s *QueryMgsApipageRequest) SetPageIndex(v int64) *QueryMgsApipageRequest {
	s.PageIndex = &v
	return s
}

func (s *QueryMgsApipageRequest) SetPageSize(v int64) *QueryMgsApipageRequest {
	s.PageSize = &v
	return s
}

func (s *QueryMgsApipageRequest) SetSysId(v int64) *QueryMgsApipageRequest {
	s.SysId = &v
	return s
}

func (s *QueryMgsApipageRequest) SetSysName(v string) *QueryMgsApipageRequest {
	s.SysName = &v
	return s
}

func (s *QueryMgsApipageRequest) SetTenantId(v string) *QueryMgsApipageRequest {
	s.TenantId = &v
	return s
}

func (s *QueryMgsApipageRequest) SetWorkspaceId(v string) *QueryMgsApipageRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryMgsApipageRequest) Validate() error {
	return dara.Validate(s)
}
