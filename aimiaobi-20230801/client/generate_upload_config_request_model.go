// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateUploadConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *GenerateUploadConfigRequest
	GetAgentKey() *string
	SetFileName(v string) *GenerateUploadConfigRequest
	GetFileName() *string
	SetParentDir(v string) *GenerateUploadConfigRequest
	GetParentDir() *string
}

type GenerateUploadConfigRequest struct {
	// The unique identifier of the workspace. For more information, see [AgentKey](https://help.aliyun.com/document_detail/2587494.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The file name.
	//
	// example:
	//
	// test.docx
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The parent folder.
	//
	// - materialDocument: The Material Library for AI Writing Assistant.
	//
	// - datasetUpload: The dataset for AI Search.
	//
	// - intervenes: Interventions.
	//
	// - temp: A temporary upload folder. Files in this folder are released periodically.
	//
	// This parameter is required.
	//
	// example:
	//
	// dataset
	ParentDir *string `json:"ParentDir,omitempty" xml:"ParentDir,omitempty"`
}

func (s GenerateUploadConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateUploadConfigRequest) GoString() string {
	return s.String()
}

func (s *GenerateUploadConfigRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *GenerateUploadConfigRequest) GetFileName() *string {
	return s.FileName
}

func (s *GenerateUploadConfigRequest) GetParentDir() *string {
	return s.ParentDir
}

func (s *GenerateUploadConfigRequest) SetAgentKey(v string) *GenerateUploadConfigRequest {
	s.AgentKey = &v
	return s
}

func (s *GenerateUploadConfigRequest) SetFileName(v string) *GenerateUploadConfigRequest {
	s.FileName = &v
	return s
}

func (s *GenerateUploadConfigRequest) SetParentDir(v string) *GenerateUploadConfigRequest {
	s.ParentDir = &v
	return s
}

func (s *GenerateUploadConfigRequest) Validate() error {
	return dara.Validate(s)
}
