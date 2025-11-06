// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDomainToDomainGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSource(v int32) *UpdateDomainToDomainGroupRequest
	GetDataSource() *int32
	SetDomainGroupId(v int64) *UpdateDomainToDomainGroupRequest
	GetDomainGroupId() *int64
	SetDomainName(v []*string) *UpdateDomainToDomainGroupRequest
	GetDomainName() []*string
	SetFileToUpload(v string) *UpdateDomainToDomainGroupRequest
	GetFileToUpload() *string
	SetLang(v string) *UpdateDomainToDomainGroupRequest
	GetLang() *string
	SetReplace(v bool) *UpdateDomainToDomainGroupRequest
	GetReplace() *bool
	SetUserClientIp(v string) *UpdateDomainToDomainGroupRequest
	GetUserClientIp() *string
}

type UpdateDomainToDomainGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	DataSource *int32 `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234
	DomainGroupId *int64 `json:"DomainGroupId,omitempty" xml:"DomainGroupId,omitempty"`
	// example:
	//
	// example.com
	DomainName []*string `json:"DomainName,omitempty" xml:"DomainName,omitempty" type:"Repeated"`
	// example:
	//
	// dGVzdA==
	FileToUpload *string `json:"FileToUpload,omitempty" xml:"FileToUpload,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	Replace *bool `json:"Replace,omitempty" xml:"Replace,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s UpdateDomainToDomainGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDomainToDomainGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateDomainToDomainGroupRequest) GetDataSource() *int32 {
	return s.DataSource
}

func (s *UpdateDomainToDomainGroupRequest) GetDomainGroupId() *int64 {
	return s.DomainGroupId
}

func (s *UpdateDomainToDomainGroupRequest) GetDomainName() []*string {
	return s.DomainName
}

func (s *UpdateDomainToDomainGroupRequest) GetFileToUpload() *string {
	return s.FileToUpload
}

func (s *UpdateDomainToDomainGroupRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateDomainToDomainGroupRequest) GetReplace() *bool {
	return s.Replace
}

func (s *UpdateDomainToDomainGroupRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *UpdateDomainToDomainGroupRequest) SetDataSource(v int32) *UpdateDomainToDomainGroupRequest {
	s.DataSource = &v
	return s
}

func (s *UpdateDomainToDomainGroupRequest) SetDomainGroupId(v int64) *UpdateDomainToDomainGroupRequest {
	s.DomainGroupId = &v
	return s
}

func (s *UpdateDomainToDomainGroupRequest) SetDomainName(v []*string) *UpdateDomainToDomainGroupRequest {
	s.DomainName = v
	return s
}

func (s *UpdateDomainToDomainGroupRequest) SetFileToUpload(v string) *UpdateDomainToDomainGroupRequest {
	s.FileToUpload = &v
	return s
}

func (s *UpdateDomainToDomainGroupRequest) SetLang(v string) *UpdateDomainToDomainGroupRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDomainToDomainGroupRequest) SetReplace(v bool) *UpdateDomainToDomainGroupRequest {
	s.Replace = &v
	return s
}

func (s *UpdateDomainToDomainGroupRequest) SetUserClientIp(v string) *UpdateDomainToDomainGroupRequest {
	s.UserClientIp = &v
	return s
}

func (s *UpdateDomainToDomainGroupRequest) Validate() error {
	return dara.Validate(s)
}
