// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportInterveneFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ImportInterveneFileRequest
	GetAgentKey() *string
	SetDocName(v string) *ImportInterveneFileRequest
	GetDocName() *string
	SetFileKey(v string) *ImportInterveneFileRequest
	GetFileKey() *string
	SetFileUrl(v string) *ImportInterveneFileRequest
	GetFileUrl() *string
}

type ImportInterveneFileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// import.xls
	DocName *string `json:"DocName,omitempty" xml:"DocName,omitempty"`
	// example:
	//
	// import.xsl
	FileKey *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	// example:
	//
	// http://xxx/xxx.xls
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s ImportInterveneFileRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportInterveneFileRequest) GoString() string {
	return s.String()
}

func (s *ImportInterveneFileRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ImportInterveneFileRequest) GetDocName() *string {
	return s.DocName
}

func (s *ImportInterveneFileRequest) GetFileKey() *string {
	return s.FileKey
}

func (s *ImportInterveneFileRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *ImportInterveneFileRequest) SetAgentKey(v string) *ImportInterveneFileRequest {
	s.AgentKey = &v
	return s
}

func (s *ImportInterveneFileRequest) SetDocName(v string) *ImportInterveneFileRequest {
	s.DocName = &v
	return s
}

func (s *ImportInterveneFileRequest) SetFileKey(v string) *ImportInterveneFileRequest {
	s.FileKey = &v
	return s
}

func (s *ImportInterveneFileRequest) SetFileUrl(v string) *ImportInterveneFileRequest {
	s.FileUrl = &v
	return s
}

func (s *ImportInterveneFileRequest) Validate() error {
	return dara.Validate(s)
}
