// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentExtractResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDocumentExtractResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDocumentExtractResultResponse
	GetStatusCode() *int32
	SetBody(v *GetDocumentExtractResultResponseBody) *GetDocumentExtractResultResponse
	GetBody() *GetDocumentExtractResultResponseBody
}

type GetDocumentExtractResultResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocumentExtractResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocumentExtractResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentExtractResultResponse) GoString() string {
	return s.String()
}

func (s *GetDocumentExtractResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDocumentExtractResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDocumentExtractResultResponse) GetBody() *GetDocumentExtractResultResponseBody {
	return s.Body
}

func (s *GetDocumentExtractResultResponse) SetHeaders(v map[string]*string) *GetDocumentExtractResultResponse {
	s.Headers = v
	return s
}

func (s *GetDocumentExtractResultResponse) SetStatusCode(v int32) *GetDocumentExtractResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocumentExtractResultResponse) SetBody(v *GetDocumentExtractResultResponseBody) *GetDocumentExtractResultResponse {
	s.Body = v
	return s
}

func (s *GetDocumentExtractResultResponse) Validate() error {
	return dara.Validate(s)
}
