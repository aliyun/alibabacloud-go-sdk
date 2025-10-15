// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunDialogAnalysisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSessionId(v string) *RunDialogAnalysisRequest
	GetSessionId() *string
}

type RunDialogAnalysisRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1759457905S001vejpvd6vej
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s RunDialogAnalysisRequest) String() string {
	return dara.Prettify(s)
}

func (s RunDialogAnalysisRequest) GoString() string {
	return s.String()
}

func (s *RunDialogAnalysisRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *RunDialogAnalysisRequest) SetSessionId(v string) *RunDialogAnalysisRequest {
	s.SessionId = &v
	return s
}

func (s *RunDialogAnalysisRequest) Validate() error {
	return dara.Validate(s)
}
