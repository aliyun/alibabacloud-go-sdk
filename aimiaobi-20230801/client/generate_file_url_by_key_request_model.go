// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateFileUrlByKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *GenerateFileUrlByKeyRequest
	GetAgentKey() *string
	SetFileKey(v string) *GenerateFileUrlByKeyRequest
	GetFileKey() *string
	SetFileName(v string) *GenerateFileUrlByKeyRequest
	GetFileName() *string
}

type GenerateFileUrlByKeyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// oss://default/oss-bucket-name/aimiaobi/2021/07/01/1625126400000/1.docx
	FileKey  *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s GenerateFileUrlByKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateFileUrlByKeyRequest) GoString() string {
	return s.String()
}

func (s *GenerateFileUrlByKeyRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *GenerateFileUrlByKeyRequest) GetFileKey() *string {
	return s.FileKey
}

func (s *GenerateFileUrlByKeyRequest) GetFileName() *string {
	return s.FileName
}

func (s *GenerateFileUrlByKeyRequest) SetAgentKey(v string) *GenerateFileUrlByKeyRequest {
	s.AgentKey = &v
	return s
}

func (s *GenerateFileUrlByKeyRequest) SetFileKey(v string) *GenerateFileUrlByKeyRequest {
	s.FileKey = &v
	return s
}

func (s *GenerateFileUrlByKeyRequest) SetFileName(v string) *GenerateFileUrlByKeyRequest {
	s.FileName = &v
	return s
}

func (s *GenerateFileUrlByKeyRequest) Validate() error {
	return dara.Validate(s)
}
