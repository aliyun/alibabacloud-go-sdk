// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentAnalyzeTaskStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDocumentAnalyzeTaskStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDocumentAnalyzeTaskStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetDocumentAnalyzeTaskStatusResponseBody) *GetDocumentAnalyzeTaskStatusResponse
	GetBody() *GetDocumentAnalyzeTaskStatusResponseBody
}

type GetDocumentAnalyzeTaskStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocumentAnalyzeTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocumentAnalyzeTaskStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentAnalyzeTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *GetDocumentAnalyzeTaskStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDocumentAnalyzeTaskStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDocumentAnalyzeTaskStatusResponse) GetBody() *GetDocumentAnalyzeTaskStatusResponseBody {
	return s.Body
}

func (s *GetDocumentAnalyzeTaskStatusResponse) SetHeaders(v map[string]*string) *GetDocumentAnalyzeTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *GetDocumentAnalyzeTaskStatusResponse) SetStatusCode(v int32) *GetDocumentAnalyzeTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocumentAnalyzeTaskStatusResponse) SetBody(v *GetDocumentAnalyzeTaskStatusResponseBody) *GetDocumentAnalyzeTaskStatusResponse {
	s.Body = v
	return s
}

func (s *GetDocumentAnalyzeTaskStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
