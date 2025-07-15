// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitEnterpriseVocAnalysisTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitEnterpriseVocAnalysisTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitEnterpriseVocAnalysisTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitEnterpriseVocAnalysisTaskResponseBody) *SubmitEnterpriseVocAnalysisTaskResponse
	GetBody() *SubmitEnterpriseVocAnalysisTaskResponseBody
}

type SubmitEnterpriseVocAnalysisTaskResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitEnterpriseVocAnalysisTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitEnterpriseVocAnalysisTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitEnterpriseVocAnalysisTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitEnterpriseVocAnalysisTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitEnterpriseVocAnalysisTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitEnterpriseVocAnalysisTaskResponse) GetBody() *SubmitEnterpriseVocAnalysisTaskResponseBody {
	return s.Body
}

func (s *SubmitEnterpriseVocAnalysisTaskResponse) SetHeaders(v map[string]*string) *SubmitEnterpriseVocAnalysisTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskResponse) SetStatusCode(v int32) *SubmitEnterpriseVocAnalysisTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskResponse) SetBody(v *SubmitEnterpriseVocAnalysisTaskResponseBody) *SubmitEnterpriseVocAnalysisTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskResponse) Validate() error {
	return dara.Validate(s)
}
