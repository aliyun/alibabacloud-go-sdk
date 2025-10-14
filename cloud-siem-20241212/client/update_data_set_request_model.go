// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSetDescription(v string) *UpdateDataSetRequest
	GetDataSetDescription() *string
	SetDataSetFileName(v string) *UpdateDataSetRequest
	GetDataSetFileName() *string
	SetDataSetId(v string) *UpdateDataSetRequest
	GetDataSetId() *string
	SetDataSetName(v string) *UpdateDataSetRequest
	GetDataSetName() *string
	SetDataSetStatus(v int32) *UpdateDataSetRequest
	GetDataSetStatus() *int32
	SetIpWhitelistRecognizers(v []*UpdateDataSetRequestIpWhitelistRecognizers) *UpdateDataSetRequest
	GetIpWhitelistRecognizers() []*UpdateDataSetRequestIpWhitelistRecognizers
	SetLang(v string) *UpdateDataSetRequest
	GetLang() *string
	SetRegionId(v string) *UpdateDataSetRequest
	GetRegionId() *string
	SetRoleFor(v int64) *UpdateDataSetRequest
	GetRoleFor() *int64
}

type UpdateDataSetRequest struct {
	// example:
	//
	// lmftest desc
	DataSetDescription *string `json:"DataSetDescription,omitempty" xml:"DataSetDescription,omitempty"`
	DataSetFileName    *string `json:"DataSetFileName,omitempty" xml:"DataSetFileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dataset-10iy8mbifnb4gniv****
	DataSetId              *string                                       `json:"DataSetId,omitempty" xml:"DataSetId,omitempty"`
	DataSetName            *string                                       `json:"DataSetName,omitempty" xml:"DataSetName,omitempty"`
	DataSetStatus          *int32                                        `json:"DataSetStatus,omitempty" xml:"DataSetStatus,omitempty"`
	IpWhitelistRecognizers []*UpdateDataSetRequestIpWhitelistRecognizers `json:"IpWhitelistRecognizers,omitempty" xml:"IpWhitelistRecognizers,omitempty" type:"Repeated"`
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

func (s UpdateDataSetRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataSetRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataSetRequest) GetDataSetDescription() *string {
	return s.DataSetDescription
}

func (s *UpdateDataSetRequest) GetDataSetFileName() *string {
	return s.DataSetFileName
}

func (s *UpdateDataSetRequest) GetDataSetId() *string {
	return s.DataSetId
}

func (s *UpdateDataSetRequest) GetDataSetName() *string {
	return s.DataSetName
}

func (s *UpdateDataSetRequest) GetDataSetStatus() *int32 {
	return s.DataSetStatus
}

func (s *UpdateDataSetRequest) GetIpWhitelistRecognizers() []*UpdateDataSetRequestIpWhitelistRecognizers {
	return s.IpWhitelistRecognizers
}

func (s *UpdateDataSetRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateDataSetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateDataSetRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *UpdateDataSetRequest) SetDataSetDescription(v string) *UpdateDataSetRequest {
	s.DataSetDescription = &v
	return s
}

func (s *UpdateDataSetRequest) SetDataSetFileName(v string) *UpdateDataSetRequest {
	s.DataSetFileName = &v
	return s
}

func (s *UpdateDataSetRequest) SetDataSetId(v string) *UpdateDataSetRequest {
	s.DataSetId = &v
	return s
}

func (s *UpdateDataSetRequest) SetDataSetName(v string) *UpdateDataSetRequest {
	s.DataSetName = &v
	return s
}

func (s *UpdateDataSetRequest) SetDataSetStatus(v int32) *UpdateDataSetRequest {
	s.DataSetStatus = &v
	return s
}

func (s *UpdateDataSetRequest) SetIpWhitelistRecognizers(v []*UpdateDataSetRequestIpWhitelistRecognizers) *UpdateDataSetRequest {
	s.IpWhitelistRecognizers = v
	return s
}

func (s *UpdateDataSetRequest) SetLang(v string) *UpdateDataSetRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDataSetRequest) SetRegionId(v string) *UpdateDataSetRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateDataSetRequest) SetRoleFor(v int64) *UpdateDataSetRequest {
	s.RoleFor = &v
	return s
}

func (s *UpdateDataSetRequest) Validate() error {
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

type UpdateDataSetRequestIpWhitelistRecognizers struct {
	// example:
	//
	// enabled
	AutoRecognizeStatus *string `json:"AutoRecognizeStatus,omitempty" xml:"AutoRecognizeStatus,omitempty"`
	// example:
	//
	// cdn_back_source_ip
	IpWhitelistRecognizerType *string `json:"IpWhitelistRecognizerType,omitempty" xml:"IpWhitelistRecognizerType,omitempty"`
	// example:
	//
	// current_account
	RecognizeScope *string `json:"RecognizeScope,omitempty" xml:"RecognizeScope,omitempty"`
}

func (s UpdateDataSetRequestIpWhitelistRecognizers) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataSetRequestIpWhitelistRecognizers) GoString() string {
	return s.String()
}

func (s *UpdateDataSetRequestIpWhitelistRecognizers) GetAutoRecognizeStatus() *string {
	return s.AutoRecognizeStatus
}

func (s *UpdateDataSetRequestIpWhitelistRecognizers) GetIpWhitelistRecognizerType() *string {
	return s.IpWhitelistRecognizerType
}

func (s *UpdateDataSetRequestIpWhitelistRecognizers) GetRecognizeScope() *string {
	return s.RecognizeScope
}

func (s *UpdateDataSetRequestIpWhitelistRecognizers) SetAutoRecognizeStatus(v string) *UpdateDataSetRequestIpWhitelistRecognizers {
	s.AutoRecognizeStatus = &v
	return s
}

func (s *UpdateDataSetRequestIpWhitelistRecognizers) SetIpWhitelistRecognizerType(v string) *UpdateDataSetRequestIpWhitelistRecognizers {
	s.IpWhitelistRecognizerType = &v
	return s
}

func (s *UpdateDataSetRequestIpWhitelistRecognizers) SetRecognizeScope(v string) *UpdateDataSetRequestIpWhitelistRecognizers {
	s.RecognizeScope = &v
	return s
}

func (s *UpdateDataSetRequestIpWhitelistRecognizers) Validate() error {
	return dara.Validate(s)
}
