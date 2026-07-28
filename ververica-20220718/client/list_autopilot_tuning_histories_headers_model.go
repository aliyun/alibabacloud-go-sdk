// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAutopilotTuningHistoriesHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListAutopilotTuningHistoriesHeaders
	GetCommonHeaders() map[string]*string
	SetAcceptLanguage(v string) *ListAutopilotTuningHistoriesHeaders
	GetAcceptLanguage() *string
	SetWorkspace(v string) *ListAutopilotTuningHistoriesHeaders
	GetWorkspace() *string
}

type ListAutopilotTuningHistoriesHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The language type. Default value: en-US. Set this to zh-CN for Chinese.
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"Accept-Language,omitempty" xml:"Accept-Language,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListAutopilotTuningHistoriesHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListAutopilotTuningHistoriesHeaders) GoString() string {
	return s.String()
}

func (s *ListAutopilotTuningHistoriesHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListAutopilotTuningHistoriesHeaders) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListAutopilotTuningHistoriesHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListAutopilotTuningHistoriesHeaders) SetCommonHeaders(v map[string]*string) *ListAutopilotTuningHistoriesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListAutopilotTuningHistoriesHeaders) SetAcceptLanguage(v string) *ListAutopilotTuningHistoriesHeaders {
	s.AcceptLanguage = &v
	return s
}

func (s *ListAutopilotTuningHistoriesHeaders) SetWorkspace(v string) *ListAutopilotTuningHistoriesHeaders {
	s.Workspace = &v
	return s
}

func (s *ListAutopilotTuningHistoriesHeaders) Validate() error {
	return dara.Validate(s)
}
