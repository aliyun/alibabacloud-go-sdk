// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunContractExtractRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *RunContractExtractRequest
	GetAppId() *string
	SetFieldsToExtract(v []*RunContractExtractRequestFieldsToExtract) *RunContractExtractRequest
	GetFieldsToExtract() []*RunContractExtractRequestFieldsToExtract
	SetFileOssUrl(v string) *RunContractExtractRequest
	GetFileOssUrl() *string
	SetRegionId(v string) *RunContractExtractRequest
	GetRegionId() *string
}

type RunContractExtractRequest struct {
	// example:
	//
	// farui
	AppId           *string                                     `json:"appId,omitempty" xml:"appId,omitempty"`
	FieldsToExtract []*RunContractExtractRequestFieldsToExtract `json:"fieldsToExtract,omitempty" xml:"fieldsToExtract,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// ttps://xxxxx.oss-cn-hangzhou.aliyuncs.com/legalmind/userdownload/4a83e0fe-baee-41d5-89f6-e33c8d462839/contract/report/9ce843d2-a05e-4351-9d69-15ae96bd910a_1713348901026.pdf
	FileOssUrl *string `json:"fileOssUrl,omitempty" xml:"fileOssUrl,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s RunContractExtractRequest) String() string {
	return dara.Prettify(s)
}

func (s RunContractExtractRequest) GoString() string {
	return s.String()
}

func (s *RunContractExtractRequest) GetAppId() *string {
	return s.AppId
}

func (s *RunContractExtractRequest) GetFieldsToExtract() []*RunContractExtractRequestFieldsToExtract {
	return s.FieldsToExtract
}

func (s *RunContractExtractRequest) GetFileOssUrl() *string {
	return s.FileOssUrl
}

func (s *RunContractExtractRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RunContractExtractRequest) SetAppId(v string) *RunContractExtractRequest {
	s.AppId = &v
	return s
}

func (s *RunContractExtractRequest) SetFieldsToExtract(v []*RunContractExtractRequestFieldsToExtract) *RunContractExtractRequest {
	s.FieldsToExtract = v
	return s
}

func (s *RunContractExtractRequest) SetFileOssUrl(v string) *RunContractExtractRequest {
	s.FileOssUrl = &v
	return s
}

func (s *RunContractExtractRequest) SetRegionId(v string) *RunContractExtractRequest {
	s.RegionId = &v
	return s
}

func (s *RunContractExtractRequest) Validate() error {
	if s.FieldsToExtract != nil {
		for _, item := range s.FieldsToExtract {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunContractExtractRequestFieldsToExtract struct {
	Desc        *string   `json:"desc,omitempty" xml:"desc,omitempty"`
	ExtractItem *string   `json:"extractItem,omitempty" xml:"extractItem,omitempty"`
	Option      []*string `json:"option,omitempty" xml:"option,omitempty" type:"Repeated"`
}

func (s RunContractExtractRequestFieldsToExtract) String() string {
	return dara.Prettify(s)
}

func (s RunContractExtractRequestFieldsToExtract) GoString() string {
	return s.String()
}

func (s *RunContractExtractRequestFieldsToExtract) GetDesc() *string {
	return s.Desc
}

func (s *RunContractExtractRequestFieldsToExtract) GetExtractItem() *string {
	return s.ExtractItem
}

func (s *RunContractExtractRequestFieldsToExtract) GetOption() []*string {
	return s.Option
}

func (s *RunContractExtractRequestFieldsToExtract) SetDesc(v string) *RunContractExtractRequestFieldsToExtract {
	s.Desc = &v
	return s
}

func (s *RunContractExtractRequestFieldsToExtract) SetExtractItem(v string) *RunContractExtractRequestFieldsToExtract {
	s.ExtractItem = &v
	return s
}

func (s *RunContractExtractRequestFieldsToExtract) SetOption(v []*string) *RunContractExtractRequestFieldsToExtract {
	s.Option = v
	return s
}

func (s *RunContractExtractRequestFieldsToExtract) Validate() error {
	return dara.Validate(s)
}
