// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportInterveneFileAsyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ImportInterveneFileAsyncRequest
	GetAgentKey() *string
	SetDocName(v string) *ImportInterveneFileAsyncRequest
	GetDocName() *string
	SetFileKey(v string) *ImportInterveneFileAsyncRequest
	GetFileKey() *string
	SetFileUrl(v string) *ImportInterveneFileAsyncRequest
	GetFileUrl() *string
}

type ImportInterveneFileAsyncRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c160c841c8e54295bf2f441432785944_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// import.xls
	DocName *string `json:"DocName,omitempty" xml:"DocName,omitempty"`
	// example:
	//
	// import.xls
	FileKey *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	// example:
	//
	// https://xxx/import.xls
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s ImportInterveneFileAsyncRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportInterveneFileAsyncRequest) GoString() string {
	return s.String()
}

func (s *ImportInterveneFileAsyncRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ImportInterveneFileAsyncRequest) GetDocName() *string {
	return s.DocName
}

func (s *ImportInterveneFileAsyncRequest) GetFileKey() *string {
	return s.FileKey
}

func (s *ImportInterveneFileAsyncRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *ImportInterveneFileAsyncRequest) SetAgentKey(v string) *ImportInterveneFileAsyncRequest {
	s.AgentKey = &v
	return s
}

func (s *ImportInterveneFileAsyncRequest) SetDocName(v string) *ImportInterveneFileAsyncRequest {
	s.DocName = &v
	return s
}

func (s *ImportInterveneFileAsyncRequest) SetFileKey(v string) *ImportInterveneFileAsyncRequest {
	s.FileKey = &v
	return s
}

func (s *ImportInterveneFileAsyncRequest) SetFileUrl(v string) *ImportInterveneFileAsyncRequest {
	s.FileUrl = &v
	return s
}

func (s *ImportInterveneFileAsyncRequest) Validate() error {
	return dara.Validate(s)
}
