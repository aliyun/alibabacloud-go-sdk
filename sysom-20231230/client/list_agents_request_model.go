// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *ListAgentsRequest
	GetXDebugId() *string
	SetCurrent(v int64) *ListAgentsRequest
	GetCurrent() *int64
	SetName(v string) *ListAgentsRequest
	GetName() *string
	SetPageSize(v int64) *ListAgentsRequest
	GetPageSize() *int64
	SetType(v string) *ListAgentsRequest
	GetType() *string
	SetXSysomInvokeSource(v string) *ListAgentsRequest
	GetXSysomInvokeSource() *string
}

type ListAgentsRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// The current page number (starting from page 1).
	//
	// example:
	//
	// 1
	Current *int64 `json:"current,omitempty" xml:"current,omitempty"`
	// Filters plug-ins by plug-in name.
	//
	// example:
	//
	// SysOM
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Filters the list by Agent type. For example, pass control to retrieve all Agents of the control type.
	//
	// example:
	//
	// control
	Type               *string `json:"type,omitempty" xml:"type,omitempty"`
	XSysomInvokeSource *string `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s ListAgentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentsRequest) GoString() string {
	return s.String()
}

func (s *ListAgentsRequest) GetXDebugId() *string {
	return s.XDebugId
}

func (s *ListAgentsRequest) GetCurrent() *int64 {
	return s.Current
}

func (s *ListAgentsRequest) GetName() *string {
	return s.Name
}

func (s *ListAgentsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAgentsRequest) GetType() *string {
	return s.Type
}

func (s *ListAgentsRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *ListAgentsRequest) SetXDebugId(v string) *ListAgentsRequest {
	s.XDebugId = &v
	return s
}

func (s *ListAgentsRequest) SetCurrent(v int64) *ListAgentsRequest {
	s.Current = &v
	return s
}

func (s *ListAgentsRequest) SetName(v string) *ListAgentsRequest {
	s.Name = &v
	return s
}

func (s *ListAgentsRequest) SetPageSize(v int64) *ListAgentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAgentsRequest) SetType(v string) *ListAgentsRequest {
	s.Type = &v
	return s
}

func (s *ListAgentsRequest) SetXSysomInvokeSource(v string) *ListAgentsRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *ListAgentsRequest) Validate() error {
	return dara.Validate(s)
}
