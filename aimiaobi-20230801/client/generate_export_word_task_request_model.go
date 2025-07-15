// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateExportWordTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *GenerateExportWordTaskRequest
	GetAgentKey() *string
	SetGeneratedContentId(v int64) *GenerateExportWordTaskRequest
	GetGeneratedContentId() *int64
}

type GenerateExportWordTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	GeneratedContentId *int64 `json:"GeneratedContentId,omitempty" xml:"GeneratedContentId,omitempty"`
}

func (s GenerateExportWordTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateExportWordTaskRequest) GoString() string {
	return s.String()
}

func (s *GenerateExportWordTaskRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *GenerateExportWordTaskRequest) GetGeneratedContentId() *int64 {
	return s.GeneratedContentId
}

func (s *GenerateExportWordTaskRequest) SetAgentKey(v string) *GenerateExportWordTaskRequest {
	s.AgentKey = &v
	return s
}

func (s *GenerateExportWordTaskRequest) SetGeneratedContentId(v int64) *GenerateExportWordTaskRequest {
	s.GeneratedContentId = &v
	return s
}

func (s *GenerateExportWordTaskRequest) Validate() error {
	return dara.Validate(s)
}
