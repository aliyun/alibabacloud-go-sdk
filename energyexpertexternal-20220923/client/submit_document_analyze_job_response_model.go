// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDocumentAnalyzeJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitDocumentAnalyzeJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitDocumentAnalyzeJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitDocumentAnalyzeJobResponseBody) *SubmitDocumentAnalyzeJobResponse
	GetBody() *SubmitDocumentAnalyzeJobResponseBody
}

type SubmitDocumentAnalyzeJobResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitDocumentAnalyzeJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitDocumentAnalyzeJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocumentAnalyzeJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitDocumentAnalyzeJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitDocumentAnalyzeJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitDocumentAnalyzeJobResponse) GetBody() *SubmitDocumentAnalyzeJobResponseBody {
	return s.Body
}

func (s *SubmitDocumentAnalyzeJobResponse) SetHeaders(v map[string]*string) *SubmitDocumentAnalyzeJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitDocumentAnalyzeJobResponse) SetStatusCode(v int32) *SubmitDocumentAnalyzeJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitDocumentAnalyzeJobResponse) SetBody(v *SubmitDocumentAnalyzeJobResponseBody) *SubmitDocumentAnalyzeJobResponse {
	s.Body = v
	return s
}

func (s *SubmitDocumentAnalyzeJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
