// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentAnalyzeResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDocumentAnalyzeResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDocumentAnalyzeResultResponse
	GetStatusCode() *int32
	SetBody(v *GetDocumentAnalyzeResultResponseBody) *GetDocumentAnalyzeResultResponse
	GetBody() *GetDocumentAnalyzeResultResponseBody
}

type GetDocumentAnalyzeResultResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocumentAnalyzeResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocumentAnalyzeResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentAnalyzeResultResponse) GoString() string {
	return s.String()
}

func (s *GetDocumentAnalyzeResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDocumentAnalyzeResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDocumentAnalyzeResultResponse) GetBody() *GetDocumentAnalyzeResultResponseBody {
	return s.Body
}

func (s *GetDocumentAnalyzeResultResponse) SetHeaders(v map[string]*string) *GetDocumentAnalyzeResultResponse {
	s.Headers = v
	return s
}

func (s *GetDocumentAnalyzeResultResponse) SetStatusCode(v int32) *GetDocumentAnalyzeResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocumentAnalyzeResultResponse) SetBody(v *GetDocumentAnalyzeResultResponseBody) *GetDocumentAnalyzeResultResponse {
	s.Body = v
	return s
}

func (s *GetDocumentAnalyzeResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
