// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAIQueryResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *GetAIQueryResultRequest
	GetXDebugId() *string
	SetAnalysisId(v string) *GetAIQueryResultRequest
	GetAnalysisId() *string
	SetXSysomInvokeSource(v string) *GetAIQueryResultRequest
	GetXSysomInvokeSource() *string
}

type GetAIQueryResultRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// The AI analysis ID.
	//
	// example:
	//
	// 16896fa8-37f6-4c70-bb32-67fa9817d426
	AnalysisId         *string `json:"analysisId,omitempty" xml:"analysisId,omitempty"`
	XSysomInvokeSource *string `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s GetAIQueryResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAIQueryResultRequest) GoString() string {
	return s.String()
}

func (s *GetAIQueryResultRequest) GetXDebugId() *string {
	return s.XDebugId
}

func (s *GetAIQueryResultRequest) GetAnalysisId() *string {
	return s.AnalysisId
}

func (s *GetAIQueryResultRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *GetAIQueryResultRequest) SetXDebugId(v string) *GetAIQueryResultRequest {
	s.XDebugId = &v
	return s
}

func (s *GetAIQueryResultRequest) SetAnalysisId(v string) *GetAIQueryResultRequest {
	s.AnalysisId = &v
	return s
}

func (s *GetAIQueryResultRequest) SetXSysomInvokeSource(v string) *GetAIQueryResultRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *GetAIQueryResultRequest) Validate() error {
	return dara.Validate(s)
}
