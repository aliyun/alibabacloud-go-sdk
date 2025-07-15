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
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// test.docx
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// datasetUpload
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
