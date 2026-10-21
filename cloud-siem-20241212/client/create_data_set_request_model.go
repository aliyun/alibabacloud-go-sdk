// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSetDescription(v string) *CreateDataSetRequest
	GetDataSetDescription() *string
	SetDataSetFieldKeyName(v string) *CreateDataSetRequest
	GetDataSetFieldKeyName() *string
	SetDataSetFileName(v string) *CreateDataSetRequest
	GetDataSetFileName() *string
	SetDataSetName(v string) *CreateDataSetRequest
	GetDataSetName() *string
	SetDataSetStatus(v int32) *CreateDataSetRequest
	GetDataSetStatus() *int32
	SetDataSetType(v string) *CreateDataSetRequest
	GetDataSetType() *string
	SetIpWhitelistRecognizers(v []*CreateDataSetRequestIpWhitelistRecognizers) *CreateDataSetRequest
	GetIpWhitelistRecognizers() []*CreateDataSetRequestIpWhitelistRecognizers
	SetLang(v string) *CreateDataSetRequest
	GetLang() *string
	SetRegionId(v string) *CreateDataSetRequest
	GetRegionId() *string
	SetRoleFor(v int64) *CreateDataSetRequest
	GetRoleFor() *int64
}

type CreateDataSetRequest struct {
	// The description of the dataset.
	//
	// example:
	//
	// lmftest contains ip list
	DataSetDescription *string `json:"DataSetDescription,omitempty" xml:"DataSetDescription,omitempty"`
	// The unique key name of the dataset.
	//
	// This parameter is required.
	//
	// example:
	//
	// ip
	DataSetFieldKeyName *string `json:"DataSetFieldKeyName,omitempty" xml:"DataSetFieldKeyName,omitempty"`
	// The name of the uploaded dataset file.
	//
	// This parameter is required.
	//
	// example:
	//
	// cloudsiem-dataset/1358117679873357_17433*****.csv
	DataSetFileName *string `json:"DataSetFileName,omitempty" xml:"DataSetFileName,omitempty"`
	// The name of the dataset.
	//
	// This parameter is required.
	//
	// example:
	//
	// lmftest
	DataSetName *string `json:"DataSetName,omitempty" xml:"DataSetName,omitempty"`
	// The status of the dataset. Valid values:
	//
	// - 0: Deleted.
	//
	// - 1: Enabled.
	//
	// example:
	//
	// 1
	DataSetStatus *int32 `json:"DataSetStatus,omitempty" xml:"DataSetStatus,omitempty"`
	// The type of the dataset. Valid values:
	//
	// - custom: custom.
	//
	// - preset: predefined.
	//
	// example:
	//
	// custom
	DataSetType *string `json:"DataSetType,omitempty" xml:"DataSetType,omitempty"`
	// The list of recognizers.
	IpWhitelistRecognizers []*CreateDataSetRequestIpWhitelistRecognizers `json:"IpWhitelistRecognizers,omitempty" xml:"IpWhitelistRecognizers,omitempty" type:"Repeated"`
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// Note: This parameter is required. If you do not specify this parameter, the API returns a 400 error: "The specified Lang can not be empty."
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region where the threat analysis data management center resides. Specify the management center based on the region of your assets. Valid values:
	//
	// - cn-hangzhou: Your assets reside in the Chinese mainland.
	//
	// - ap-southeast-1: Your assets reside outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the member to which the administrator switches the view.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s CreateDataSetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataSetRequest) GoString() string {
	return s.String()
}

func (s *CreateDataSetRequest) GetDataSetDescription() *string {
	return s.DataSetDescription
}

func (s *CreateDataSetRequest) GetDataSetFieldKeyName() *string {
	return s.DataSetFieldKeyName
}

func (s *CreateDataSetRequest) GetDataSetFileName() *string {
	return s.DataSetFileName
}

func (s *CreateDataSetRequest) GetDataSetName() *string {
	return s.DataSetName
}

func (s *CreateDataSetRequest) GetDataSetStatus() *int32 {
	return s.DataSetStatus
}

func (s *CreateDataSetRequest) GetDataSetType() *string {
	return s.DataSetType
}

func (s *CreateDataSetRequest) GetIpWhitelistRecognizers() []*CreateDataSetRequestIpWhitelistRecognizers {
	return s.IpWhitelistRecognizers
}

func (s *CreateDataSetRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateDataSetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDataSetRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *CreateDataSetRequest) SetDataSetDescription(v string) *CreateDataSetRequest {
	s.DataSetDescription = &v
	return s
}

func (s *CreateDataSetRequest) SetDataSetFieldKeyName(v string) *CreateDataSetRequest {
	s.DataSetFieldKeyName = &v
	return s
}

func (s *CreateDataSetRequest) SetDataSetFileName(v string) *CreateDataSetRequest {
	s.DataSetFileName = &v
	return s
}

func (s *CreateDataSetRequest) SetDataSetName(v string) *CreateDataSetRequest {
	s.DataSetName = &v
	return s
}

func (s *CreateDataSetRequest) SetDataSetStatus(v int32) *CreateDataSetRequest {
	s.DataSetStatus = &v
	return s
}

func (s *CreateDataSetRequest) SetDataSetType(v string) *CreateDataSetRequest {
	s.DataSetType = &v
	return s
}

func (s *CreateDataSetRequest) SetIpWhitelistRecognizers(v []*CreateDataSetRequestIpWhitelistRecognizers) *CreateDataSetRequest {
	s.IpWhitelistRecognizers = v
	return s
}

func (s *CreateDataSetRequest) SetLang(v string) *CreateDataSetRequest {
	s.Lang = &v
	return s
}

func (s *CreateDataSetRequest) SetRegionId(v string) *CreateDataSetRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDataSetRequest) SetRoleFor(v int64) *CreateDataSetRequest {
	s.RoleFor = &v
	return s
}

func (s *CreateDataSetRequest) Validate() error {
	if s.IpWhitelistRecognizers != nil {
		for _, item := range s.IpWhitelistRecognizers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateDataSetRequestIpWhitelistRecognizers struct {
	// The automatic recognition status. Valid values:
	//
	// - enabled: Enabled.
	//
	// - disabled: Disabled.
	//
	// example:
	//
	// enabled
	AutoRecognizeStatus *string `json:"AutoRecognizeStatus,omitempty" xml:"AutoRecognizeStatus,omitempty"`
	// The IP address type identified by the recognizer. Valid values:
	//
	// - sas_vulnerability_scanner_ip: IP addresses of the Security Center vulnerability web scanner.
	//
	// - waf_back_source_ip: back-to-origin IP addresses of Web Application Firewall (WAF).
	//
	// - ddos_back_source_ip: back-to-origin IP addresses of Anti-DDoS.
	//
	// - esa_back_source_ip: back-to-origin node IP addresses of Edge secure acceleration (ESA).
	//
	// - ecs_public_ip: public IP addresses of Elastic Compute Service (ECS) instances.
	//
	// - slb_public_ip: public IP addresses of Server Load Balancer (SLB) instances.
	//
	// - vpc_eip: elastic IP addresses (EIPs).
	//
	// - cdn_back_source_ip: back-to-origin IP addresses of content delivery network (CDN).
	//
	// - ga_back_source_ip: back-to-origin IP addresses of Global Accelerator (GA).
	//
	// example:
	//
	// waf_back_source_ip
	IpWhitelistRecognizerType *string `json:"IpWhitelistRecognizerType,omitempty" xml:"IpWhitelistRecognizerType,omitempty"`
	// The recognition scope. Valid values:
	//
	// - current_account: the current account only.
	//
	// - rd_accounts: multi-account mode enabled.
	//
	// example:
	//
	// current_account
	RecognizeScope *string `json:"RecognizeScope,omitempty" xml:"RecognizeScope,omitempty"`
}

func (s CreateDataSetRequestIpWhitelistRecognizers) String() string {
	return dara.Prettify(s)
}

func (s CreateDataSetRequestIpWhitelistRecognizers) GoString() string {
	return s.String()
}

func (s *CreateDataSetRequestIpWhitelistRecognizers) GetAutoRecognizeStatus() *string {
	return s.AutoRecognizeStatus
}

func (s *CreateDataSetRequestIpWhitelistRecognizers) GetIpWhitelistRecognizerType() *string {
	return s.IpWhitelistRecognizerType
}

func (s *CreateDataSetRequestIpWhitelistRecognizers) GetRecognizeScope() *string {
	return s.RecognizeScope
}

func (s *CreateDataSetRequestIpWhitelistRecognizers) SetAutoRecognizeStatus(v string) *CreateDataSetRequestIpWhitelistRecognizers {
	s.AutoRecognizeStatus = &v
	return s
}

func (s *CreateDataSetRequestIpWhitelistRecognizers) SetIpWhitelistRecognizerType(v string) *CreateDataSetRequestIpWhitelistRecognizers {
	s.IpWhitelistRecognizerType = &v
	return s
}

func (s *CreateDataSetRequestIpWhitelistRecognizers) SetRecognizeScope(v string) *CreateDataSetRequestIpWhitelistRecognizers {
	s.RecognizeScope = &v
	return s
}

func (s *CreateDataSetRequestIpWhitelistRecognizers) Validate() error {
	return dara.Validate(s)
}
