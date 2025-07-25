// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDNSSLBWeightRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *UpdateDNSSLBWeightRequest
	GetLang() *string
	SetRecordId(v string) *UpdateDNSSLBWeightRequest
	GetRecordId() *string
	SetUserClientIp(v string) *UpdateDNSSLBWeightRequest
	GetUserClientIp() *string
	SetWeight(v int32) *UpdateDNSSLBWeightRequest
	GetWeight() *int32
}

type UpdateDNSSLBWeightRequest struct {
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the DNS record. You can call the [DescribeDomainRecords](https://www.alibabacloud.com/help/en/dns/api-alidns-2015-01-09-describedomainrecords?spm=a2c63.p38356.help-menu-search-29697.d_0) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9999985
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 1.1.1.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// The weight of the DNS record that you want to specify. Valid values: `1 to 100`.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s UpdateDNSSLBWeightRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDNSSLBWeightRequest) GoString() string {
	return s.String()
}

func (s *UpdateDNSSLBWeightRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateDNSSLBWeightRequest) GetRecordId() *string {
	return s.RecordId
}

func (s *UpdateDNSSLBWeightRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *UpdateDNSSLBWeightRequest) GetWeight() *int32 {
	return s.Weight
}

func (s *UpdateDNSSLBWeightRequest) SetLang(v string) *UpdateDNSSLBWeightRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDNSSLBWeightRequest) SetRecordId(v string) *UpdateDNSSLBWeightRequest {
	s.RecordId = &v
	return s
}

func (s *UpdateDNSSLBWeightRequest) SetUserClientIp(v string) *UpdateDNSSLBWeightRequest {
	s.UserClientIp = &v
	return s
}

func (s *UpdateDNSSLBWeightRequest) SetWeight(v int32) *UpdateDNSSLBWeightRequest {
	s.Weight = &v
	return s
}

func (s *UpdateDNSSLBWeightRequest) Validate() error {
	return dara.Validate(s)
}
