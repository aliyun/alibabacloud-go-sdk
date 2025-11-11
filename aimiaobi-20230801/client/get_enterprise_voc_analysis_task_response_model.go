// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEnterpriseVocAnalysisTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEnterpriseVocAnalysisTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEnterpriseVocAnalysisTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetEnterpriseVocAnalysisTaskResponseBody) *GetEnterpriseVocAnalysisTaskResponse
	GetBody() *GetEnterpriseVocAnalysisTaskResponseBody
}

type GetEnterpriseVocAnalysisTaskResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEnterpriseVocAnalysisTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEnterpriseVocAnalysisTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEnterpriseVocAnalysisTaskResponse) GoString() string {
	return s.String()
}

func (s *GetEnterpriseVocAnalysisTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEnterpriseVocAnalysisTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEnterpriseVocAnalysisTaskResponse) GetBody() *GetEnterpriseVocAnalysisTaskResponseBody {
	return s.Body
}

func (s *GetEnterpriseVocAnalysisTaskResponse) SetHeaders(v map[string]*string) *GetEnterpriseVocAnalysisTaskResponse {
	s.Headers = v
	return s
}

func (s *GetEnterpriseVocAnalysisTaskResponse) SetStatusCode(v int32) *GetEnterpriseVocAnalysisTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEnterpriseVocAnalysisTaskResponse) SetBody(v *GetEnterpriseVocAnalysisTaskResponseBody) *GetEnterpriseVocAnalysisTaskResponse {
	s.Body = v
	return s
}

func (s *GetEnterpriseVocAnalysisTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
