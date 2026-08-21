// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCopilotHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *GetCopilotHistoryRequest
	GetXDebugId() *string
	SetCount(v int64) *GetCopilotHistoryRequest
	GetCount() *int64
	SetXSysomInvokeSource(v string) *GetCopilotHistoryRequest
	GetXSysomInvokeSource() *string
}

type GetCopilotHistoryRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// The number of historical chat records to request. The value is generally less than 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	Count              *int64  `json:"count,omitempty" xml:"count,omitempty"`
	XSysomInvokeSource *string `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s GetCopilotHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCopilotHistoryRequest) GoString() string {
	return s.String()
}

func (s *GetCopilotHistoryRequest) GetXDebugId() *string {
	return s.XDebugId
}

func (s *GetCopilotHistoryRequest) GetCount() *int64 {
	return s.Count
}

func (s *GetCopilotHistoryRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *GetCopilotHistoryRequest) SetXDebugId(v string) *GetCopilotHistoryRequest {
	s.XDebugId = &v
	return s
}

func (s *GetCopilotHistoryRequest) SetCount(v int64) *GetCopilotHistoryRequest {
	s.Count = &v
	return s
}

func (s *GetCopilotHistoryRequest) SetXSysomInvokeSource(v string) *GetCopilotHistoryRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *GetCopilotHistoryRequest) Validate() error {
	return dara.Validate(s)
}
