// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunContractExtractShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *RunContractExtractShrinkRequest
	GetAppId() *string
	SetFieldsToExtractShrink(v string) *RunContractExtractShrinkRequest
	GetFieldsToExtractShrink() *string
	SetFileOssUrl(v string) *RunContractExtractShrinkRequest
	GetFileOssUrl() *string
	SetRegionId(v string) *RunContractExtractShrinkRequest
	GetRegionId() *string
}

type RunContractExtractShrinkRequest struct {
	// example:
	//
	// farui
	AppId                 *string `json:"appId,omitempty" xml:"appId,omitempty"`
	FieldsToExtractShrink *string `json:"fieldsToExtract,omitempty" xml:"fieldsToExtract,omitempty"`
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

func (s RunContractExtractShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunContractExtractShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunContractExtractShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *RunContractExtractShrinkRequest) GetFieldsToExtractShrink() *string {
	return s.FieldsToExtractShrink
}

func (s *RunContractExtractShrinkRequest) GetFileOssUrl() *string {
	return s.FileOssUrl
}

func (s *RunContractExtractShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RunContractExtractShrinkRequest) SetAppId(v string) *RunContractExtractShrinkRequest {
	s.AppId = &v
	return s
}

func (s *RunContractExtractShrinkRequest) SetFieldsToExtractShrink(v string) *RunContractExtractShrinkRequest {
	s.FieldsToExtractShrink = &v
	return s
}

func (s *RunContractExtractShrinkRequest) SetFileOssUrl(v string) *RunContractExtractShrinkRequest {
	s.FileOssUrl = &v
	return s
}

func (s *RunContractExtractShrinkRequest) SetRegionId(v string) *RunContractExtractShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *RunContractExtractShrinkRequest) Validate() error {
	return dara.Validate(s)
}
