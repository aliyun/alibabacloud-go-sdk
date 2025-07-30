// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDocParserJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *CreateDocParserJobRequest
	GetFileName() *string
	SetFileUrl(v string) *CreateDocParserJobRequest
	GetFileUrl() *string
	SetRegionId(v string) *CreateDocParserJobRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateDocParserJobRequest
	GetResourceGroupId() *string
	SetResultType(v string) *CreateDocParserJobRequest
	GetResultType() *string
}

type CreateDocParserJobRequest struct {
	// example:
	//
	// 2.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// https://oss-cn-hangzhou.aliyuncs.com/storage/pdf/40184458-fbb0-44cf-a391-350628ceccdd17375122****
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// zip
	ResultType *string `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
}

func (s CreateDocParserJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDocParserJobRequest) GoString() string {
	return s.String()
}

func (s *CreateDocParserJobRequest) GetFileName() *string {
	return s.FileName
}

func (s *CreateDocParserJobRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *CreateDocParserJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDocParserJobRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDocParserJobRequest) GetResultType() *string {
	return s.ResultType
}

func (s *CreateDocParserJobRequest) SetFileName(v string) *CreateDocParserJobRequest {
	s.FileName = &v
	return s
}

func (s *CreateDocParserJobRequest) SetFileUrl(v string) *CreateDocParserJobRequest {
	s.FileUrl = &v
	return s
}

func (s *CreateDocParserJobRequest) SetRegionId(v string) *CreateDocParserJobRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDocParserJobRequest) SetResourceGroupId(v string) *CreateDocParserJobRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDocParserJobRequest) SetResultType(v string) *CreateDocParserJobRequest {
	s.ResultType = &v
	return s
}

func (s *CreateDocParserJobRequest) Validate() error {
	return dara.Validate(s)
}
