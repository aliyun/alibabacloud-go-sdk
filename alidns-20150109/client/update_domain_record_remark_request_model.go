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
	// The language of the request and the response. The default value is **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the DNS record.<props="china"> Call the [DescribeDomainRecords](https://help.aliyun.com/zh/dns/api-alidns-2015-01-09-describedomainrecords?spm=a2c4g.11186623.help-menu-search-29697.d_0) operation to obtain the record ID.<props="intl"> Call the [DescribeDomainRecords](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describedomainrecords?spm=a2c63.p38356.help-menu-search-29697.d_0) operation to obtain the record ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12*****
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The remarks for the DNS record. The default value is empty. If this parameter is left empty, the original remarks are deleted.
	//
	// example:
	//
	// 我的第一个解析记录
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The client IP address.
	//
	// example:
	//
	// 192.0.2.1
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
