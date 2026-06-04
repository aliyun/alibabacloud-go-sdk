// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitParseDocumentLayoutTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitParseDocumentLayoutTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitParseDocumentLayoutTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitParseDocumentLayoutTaskResponseBody) *SubmitParseDocumentLayoutTaskResponse
	GetBody() *SubmitParseDocumentLayoutTaskResponseBody
}

type SubmitParseDocumentLayoutTaskResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitParseDocumentLayoutTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitParseDocumentLayoutTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitParseDocumentLayoutTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitParseDocumentLayoutTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitParseDocumentLayoutTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitParseDocumentLayoutTaskResponse) GetBody() *SubmitParseDocumentLayoutTaskResponseBody {
	return s.Body
}

func (s *SubmitParseDocumentLayoutTaskResponse) SetHeaders(v map[string]*string) *SubmitParseDocumentLayoutTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitParseDocumentLayoutTaskResponse) SetStatusCode(v int32) *SubmitParseDocumentLayoutTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitParseDocumentLayoutTaskResponse) SetBody(v *SubmitParseDocumentLayoutTaskResponseBody) *SubmitParseDocumentLayoutTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitParseDocumentLayoutTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
