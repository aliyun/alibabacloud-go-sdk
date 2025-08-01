// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAIQueryResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnalysisId(v string) *GetAIQueryResultRequest
	GetAnalysisId() *string
}

type GetAIQueryResultRequest struct {
	// example:
	//
	// 16896fa8-37f6-4c70-bb32-67fa9817d426
	AnalysisId *string `json:"analysisId,omitempty" xml:"analysisId,omitempty"`
}

func (s GetAIQueryResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAIQueryResultRequest) GoString() string {
	return s.String()
}

func (s *GetAIQueryResultRequest) GetAnalysisId() *string {
	return s.AnalysisId
}

func (s *GetAIQueryResultRequest) SetAnalysisId(v string) *GetAIQueryResultRequest {
	s.AnalysisId = &v
	return s
}

func (s *GetAIQueryResultRequest) Validate() error {
	return dara.Validate(s)
}
