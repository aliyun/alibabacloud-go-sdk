// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDocumentAnalyzeTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDocumentAnalyzeTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDocumentAnalyzeTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateDocumentAnalyzeTaskResponseBody) *CreateDocumentAnalyzeTaskResponse
	GetBody() *CreateDocumentAnalyzeTaskResponseBody
}

type CreateDocumentAnalyzeTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDocumentAnalyzeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDocumentAnalyzeTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDocumentAnalyzeTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateDocumentAnalyzeTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDocumentAnalyzeTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDocumentAnalyzeTaskResponse) GetBody() *CreateDocumentAnalyzeTaskResponseBody {
	return s.Body
}

func (s *CreateDocumentAnalyzeTaskResponse) SetHeaders(v map[string]*string) *CreateDocumentAnalyzeTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateDocumentAnalyzeTaskResponse) SetStatusCode(v int32) *CreateDocumentAnalyzeTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDocumentAnalyzeTaskResponse) SetBody(v *CreateDocumentAnalyzeTaskResponseBody) *CreateDocumentAnalyzeTaskResponse {
	s.Body = v
	return s
}

func (s *CreateDocumentAnalyzeTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
