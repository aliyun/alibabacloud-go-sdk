// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iCreateDocParserJobAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *CreateDocParserJobAdvanceRequest
	GetFileName() *string
	SetFileUrlObject(v io.Reader) *CreateDocParserJobAdvanceRequest
	GetFileUrlObject() io.Reader
	SetRagInstanceId(v string) *CreateDocParserJobAdvanceRequest
	GetRagInstanceId() *string
	SetRegionId(v string) *CreateDocParserJobAdvanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateDocParserJobAdvanceRequest
	GetResourceGroupId() *string
	SetResultType(v string) *CreateDocParserJobAdvanceRequest
	GetResultType() *string
}

type CreateDocParserJobAdvanceRequest struct {
	// The name of the document to be parsed.
	//
	// > The name must include the file name extension. Currently, only .pdf is supported.
	//
	// example:
	//
	// 2.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The OSS URL of the document to be parsed.
	//
	// > This parameter is automatically populated when you call this operation by using an SDK.
	//
	// example:
	//
	// https://oss-cn-hangzhou.aliyuncs.com/storage/pdf/40184458-fbb0-44cf-a391-350628ceccdd17375122****
	FileUrlObject io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	RagInstanceId *string   `json:"RagInstanceId,omitempty" xml:"RagInstanceId,omitempty"`
	// The region ID of the document parsing task. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The output format of the parsing result after the task is complete. Valid values:
	//
	// - **zip**: a ZIP compressed file.
	//
	// - **content**: plain text.
	//
	// example:
	//
	// zip
	ResultType *string `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
}

func (s CreateDocParserJobAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDocParserJobAdvanceRequest) GoString() string {
	return s.String()
}

func (s *CreateDocParserJobAdvanceRequest) GetFileName() *string {
	return s.FileName
}

func (s *CreateDocParserJobAdvanceRequest) GetFileUrlObject() io.Reader {
	return s.FileUrlObject
}

func (s *CreateDocParserJobAdvanceRequest) GetRagInstanceId() *string {
	return s.RagInstanceId
}

func (s *CreateDocParserJobAdvanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDocParserJobAdvanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDocParserJobAdvanceRequest) GetResultType() *string {
	return s.ResultType
}

func (s *CreateDocParserJobAdvanceRequest) SetFileName(v string) *CreateDocParserJobAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) SetFileUrlObject(v io.Reader) *CreateDocParserJobAdvanceRequest {
	s.FileUrlObject = v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) SetRagInstanceId(v string) *CreateDocParserJobAdvanceRequest {
	s.RagInstanceId = &v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) SetRegionId(v string) *CreateDocParserJobAdvanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) SetResourceGroupId(v string) *CreateDocParserJobAdvanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) SetResultType(v string) *CreateDocParserJobAdvanceRequest {
	s.ResultType = &v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
