// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerKnowledgeFileParseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *TriggerKnowledgeFileParseRequest
	GetDBClusterId() *string
	SetFileId(v string) *TriggerKnowledgeFileParseRequest
	GetFileId() *string
	SetFileLocation(v string) *TriggerKnowledgeFileParseRequest
	GetFileLocation() *string
}

type TriggerKnowledgeFileParseRequest struct {
	// The ADB instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp19aaaaaa****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the knowledge base file. You must specify either this parameter or FileLocation, but not both.
	//
	// example:
	//
	// 123
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The URL of a single knowledge base file. You must specify either this parameter or FileId, but not both. You cannot specify a directory. The URL must uniquely point to an existing file.
	//
	// example:
	//
	// oss://bucket/doc.pdf
	FileLocation *string `json:"FileLocation,omitempty" xml:"FileLocation,omitempty"`
}

func (s TriggerKnowledgeFileParseRequest) String() string {
	return dara.Prettify(s)
}

func (s TriggerKnowledgeFileParseRequest) GoString() string {
	return s.String()
}

func (s *TriggerKnowledgeFileParseRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *TriggerKnowledgeFileParseRequest) GetFileId() *string {
	return s.FileId
}

func (s *TriggerKnowledgeFileParseRequest) GetFileLocation() *string {
	return s.FileLocation
}

func (s *TriggerKnowledgeFileParseRequest) SetDBClusterId(v string) *TriggerKnowledgeFileParseRequest {
	s.DBClusterId = &v
	return s
}

func (s *TriggerKnowledgeFileParseRequest) SetFileId(v string) *TriggerKnowledgeFileParseRequest {
	s.FileId = &v
	return s
}

func (s *TriggerKnowledgeFileParseRequest) SetFileLocation(v string) *TriggerKnowledgeFileParseRequest {
	s.FileLocation = &v
	return s
}

func (s *TriggerKnowledgeFileParseRequest) Validate() error {
	return dara.Validate(s)
}
