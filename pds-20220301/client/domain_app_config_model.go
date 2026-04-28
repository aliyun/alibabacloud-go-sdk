// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDomainAppConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAllowUploadCustomFileExtList(v []*string) *DomainAppConfig
	GetAllowUploadCustomFileExtList() []*string
	SetAllowUploadFileCategoryList(v []*string) *DomainAppConfig
	GetAllowUploadFileCategoryList() []*string
	SetSameNameFileUploadMode(v string) *DomainAppConfig
	GetSameNameFileUploadMode() *string
	SetSingleFileUploadSizeLimit(v int64) *DomainAppConfig
	GetSingleFileUploadSizeLimit() *int64
	SetWebClientDownloadMode(v string) *DomainAppConfig
	GetWebClientDownloadMode() *string
}

type DomainAppConfig struct {
	AllowUploadCustomFileExtList []*string `json:"allow_upload_custom_file_ext_list,omitempty" xml:"allow_upload_custom_file_ext_list,omitempty" type:"Repeated"`
	AllowUploadFileCategoryList  []*string `json:"allow_upload_file_category_list,omitempty" xml:"allow_upload_file_category_list,omitempty" type:"Repeated"`
	SameNameFileUploadMode       *string   `json:"same_name_file_upload_mode,omitempty" xml:"same_name_file_upload_mode,omitempty"`
	SingleFileUploadSizeLimit    *int64    `json:"single_file_upload_size_limit,omitempty" xml:"single_file_upload_size_limit,omitempty"`
	WebClientDownloadMode        *string   `json:"web_client_download_mode,omitempty" xml:"web_client_download_mode,omitempty"`
}

func (s DomainAppConfig) String() string {
	return dara.Prettify(s)
}

func (s DomainAppConfig) GoString() string {
	return s.String()
}

func (s *DomainAppConfig) GetAllowUploadCustomFileExtList() []*string {
	return s.AllowUploadCustomFileExtList
}

func (s *DomainAppConfig) GetAllowUploadFileCategoryList() []*string {
	return s.AllowUploadFileCategoryList
}

func (s *DomainAppConfig) GetSameNameFileUploadMode() *string {
	return s.SameNameFileUploadMode
}

func (s *DomainAppConfig) GetSingleFileUploadSizeLimit() *int64 {
	return s.SingleFileUploadSizeLimit
}

func (s *DomainAppConfig) GetWebClientDownloadMode() *string {
	return s.WebClientDownloadMode
}

func (s *DomainAppConfig) SetAllowUploadCustomFileExtList(v []*string) *DomainAppConfig {
	s.AllowUploadCustomFileExtList = v
	return s
}

func (s *DomainAppConfig) SetAllowUploadFileCategoryList(v []*string) *DomainAppConfig {
	s.AllowUploadFileCategoryList = v
	return s
}

func (s *DomainAppConfig) SetSameNameFileUploadMode(v string) *DomainAppConfig {
	s.SameNameFileUploadMode = &v
	return s
}

func (s *DomainAppConfig) SetSingleFileUploadSizeLimit(v int64) *DomainAppConfig {
	s.SingleFileUploadSizeLimit = &v
	return s
}

func (s *DomainAppConfig) SetWebClientDownloadMode(v string) *DomainAppConfig {
	s.WebClientDownloadMode = &v
	return s
}

func (s *DomainAppConfig) Validate() error {
	return dara.Validate(s)
}
