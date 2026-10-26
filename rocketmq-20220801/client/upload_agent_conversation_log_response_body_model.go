// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadAgentConversationLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *UploadAgentConversationLogResponseBody
	GetFileName() *string
	SetId(v string) *UploadAgentConversationLogResponseBody
	GetId() *string
	SetSandboxPath(v string) *UploadAgentConversationLogResponseBody
	GetSandboxPath() *string
	SetSize(v int64) *UploadAgentConversationLogResponseBody
	GetSize() *int64
}

type UploadAgentConversationLogResponseBody struct {
	// example:
	//
	// client.log
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// 0198c8a91f4d7f01a3c21e99
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// client-logs/0198c8a91f4d7f01a3c21e99/client.log
	SandboxPath *string `json:"sandboxPath,omitempty" xml:"sandboxPath,omitempty"`
	// example:
	//
	// 114
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s UploadAgentConversationLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadAgentConversationLogResponseBody) GoString() string {
	return s.String()
}

func (s *UploadAgentConversationLogResponseBody) GetFileName() *string {
	return s.FileName
}

func (s *UploadAgentConversationLogResponseBody) GetId() *string {
	return s.Id
}

func (s *UploadAgentConversationLogResponseBody) GetSandboxPath() *string {
	return s.SandboxPath
}

func (s *UploadAgentConversationLogResponseBody) GetSize() *int64 {
	return s.Size
}

func (s *UploadAgentConversationLogResponseBody) SetFileName(v string) *UploadAgentConversationLogResponseBody {
	s.FileName = &v
	return s
}

func (s *UploadAgentConversationLogResponseBody) SetId(v string) *UploadAgentConversationLogResponseBody {
	s.Id = &v
	return s
}

func (s *UploadAgentConversationLogResponseBody) SetSandboxPath(v string) *UploadAgentConversationLogResponseBody {
	s.SandboxPath = &v
	return s
}

func (s *UploadAgentConversationLogResponseBody) SetSize(v int64) *UploadAgentConversationLogResponseBody {
	s.Size = &v
	return s
}

func (s *UploadAgentConversationLogResponseBody) Validate() error {
	return dara.Validate(s)
}
