// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDomainRecordRemarkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *UpdateDomainRecordRemarkRequest
	GetLang() *string
	SetRecordId(v string) *UpdateDomainRecordRemarkRequest
	GetRecordId() *string
	SetRemark(v string) *UpdateDomainRecordRemarkRequest
	GetRemark() *string
	SetUserClientIp(v string) *UpdateDomainRecordRemarkRequest
	GetUserClientIp() *string
}

type UpdateDomainRecordRemarkRequest struct {
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the DNS record. You can call the [DescribeDomainRecords](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describedomainrecords?spm=a2c63.p38356.help-menu-search-29697.d_0) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The description of the DNS record. This parameter is empty by default. If this parameter is empty, the original remarks are deleted.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 192.0.2.0
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s UpdateDomainRecordRemarkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDomainRecordRemarkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDomainRecordRemarkRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateDomainRecordRemarkRequest) GetRecordId() *string {
	return s.RecordId
}

func (s *UpdateDomainRecordRemarkRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateDomainRecordRemarkRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *UpdateDomainRecordRemarkRequest) SetLang(v string) *UpdateDomainRecordRemarkRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDomainRecordRemarkRequest) SetRecordId(v string) *UpdateDomainRecordRemarkRequest {
	s.RecordId = &v
	return s
}

func (s *UpdateDomainRecordRemarkRequest) SetRemark(v string) *UpdateDomainRecordRemarkRequest {
	s.Remark = &v
	return s
}

func (s *UpdateDomainRecordRemarkRequest) SetUserClientIp(v string) *UpdateDomainRecordRemarkRequest {
	s.UserClientIp = &v
	return s
}

func (s *UpdateDomainRecordRemarkRequest) Validate() error {
	return dara.Validate(s)
}
