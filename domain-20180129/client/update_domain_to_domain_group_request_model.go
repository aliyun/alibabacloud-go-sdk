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
	// The data source for the domain names. Valid values:
	//
	// - **1**: custom input.
	//
	// - **2**: file upload.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DataSource *int32 `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	// The ID of the domain name group. Call the [QueryDomainGroupList](https://help.aliyun.com/document_detail/69362.html) API to get this ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	DomainGroupId *int64 `json:"DomainGroupId,omitempty" xml:"DomainGroupId,omitempty"`
	// An array of domain names. This parameter is required when DataSource is set to 1 (custom input).
	//
	// example:
	//
	// example.com
	DomainName []*string `json:"DomainName,omitempty" xml:"DomainName,omitempty" type:"Repeated"`
	// The Base64-encoded content of a file. This parameter is required if you set DataSource to 2. The file must be in **.xls*	- or **.xlsx*	- format, contain one domain name per line, and not exceed 2 MB.
	//
	// example:
	//
	// dGVzdA==
	FileToUpload *string `json:"FileToUpload,omitempty" xml:"FileToUpload,omitempty"`
	// The language of API error messages. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// Default value: **en**.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Specifies whether to replace the existing domain names in the group. Valid values:
	//
	// - **false**: Adds the new domain names to the group.
	//
	// - **true**: Replaces all existing domain names in the group with the new ones.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	Replace *bool `json:"Replace,omitempty" xml:"Replace,omitempty"`
	// The user IP address. You can set this parameter to **127.0.0.1**.
	//
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
