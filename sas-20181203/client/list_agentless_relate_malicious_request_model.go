// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentlessRelateMaliciousRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListAgentlessRelateMaliciousRequest
	GetCurrentPage() *int32
	SetEventId(v int64) *ListAgentlessRelateMaliciousRequest
	GetEventId() *int64
	SetLang(v string) *ListAgentlessRelateMaliciousRequest
	GetLang() *string
	SetPageSize(v string) *ListAgentlessRelateMaliciousRequest
	GetPageSize() *string
	SetScenario(v string) *ListAgentlessRelateMaliciousRequest
	GetScenario() *string
}

type ListAgentlessRelateMaliciousRequest struct {
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the event.
	//
	// example:
	//
	// 80****
	EventId *int64 `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese.
	//
	// 	- **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The scenarios of batch handling. Valid values:
	//
	// 	- same_file_md5: the same MD5 hash value.
	//
	// 	- default: the same alert type. This is the default value.
	//
	// example:
	//
	// default
	Scenario *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
}

func (s ListAgentlessRelateMaliciousRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentlessRelateMaliciousRequest) GoString() string {
	return s.String()
}

func (s *ListAgentlessRelateMaliciousRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListAgentlessRelateMaliciousRequest) GetEventId() *int64 {
	return s.EventId
}

func (s *ListAgentlessRelateMaliciousRequest) GetLang() *string {
	return s.Lang
}

func (s *ListAgentlessRelateMaliciousRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListAgentlessRelateMaliciousRequest) GetScenario() *string {
	return s.Scenario
}

func (s *ListAgentlessRelateMaliciousRequest) SetCurrentPage(v int32) *ListAgentlessRelateMaliciousRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListAgentlessRelateMaliciousRequest) SetEventId(v int64) *ListAgentlessRelateMaliciousRequest {
	s.EventId = &v
	return s
}

func (s *ListAgentlessRelateMaliciousRequest) SetLang(v string) *ListAgentlessRelateMaliciousRequest {
	s.Lang = &v
	return s
}

func (s *ListAgentlessRelateMaliciousRequest) SetPageSize(v string) *ListAgentlessRelateMaliciousRequest {
	s.PageSize = &v
	return s
}

func (s *ListAgentlessRelateMaliciousRequest) SetScenario(v string) *ListAgentlessRelateMaliciousRequest {
	s.Scenario = &v
	return s
}

func (s *ListAgentlessRelateMaliciousRequest) Validate() error {
	return dara.Validate(s)
}
