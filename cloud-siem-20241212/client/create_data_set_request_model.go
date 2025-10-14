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
	// example:
	//
	// lmftest contains ip list
	DataSetDescription *string `json:"DataSetDescription,omitempty" xml:"DataSetDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ip
	DataSetFieldKeyName *string `json:"DataSetFieldKeyName,omitempty" xml:"DataSetFieldKeyName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cloudsiem-dataset/1358117679873357_17433*****.csv
	DataSetFileName *string `json:"DataSetFileName,omitempty" xml:"DataSetFileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// lmftest
	DataSetName            *string                                       `json:"DataSetName,omitempty" xml:"DataSetName,omitempty"`
	DataSetStatus          *int32                                        `json:"DataSetStatus,omitempty" xml:"DataSetStatus,omitempty"`
	DataSetType            *string                                       `json:"DataSetType,omitempty" xml:"DataSetType,omitempty"`
	IpWhitelistRecognizers []*CreateDataSetRequestIpWhitelistRecognizers `json:"IpWhitelistRecognizers,omitempty" xml:"IpWhitelistRecognizers,omitempty" type:"Repeated"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// example:
	//
	// enabled
	AutoRecognizeStatus *string `json:"AutoRecognizeStatus,omitempty" xml:"AutoRecognizeStatus,omitempty"`
	// example:
	//
	// waf_back_source_ip
	IpWhitelistRecognizerType *string `json:"IpWhitelistRecognizerType,omitempty" xml:"IpWhitelistRecognizerType,omitempty"`
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
