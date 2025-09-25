// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecallDocumentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecallDocumentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecallDocumentResponse
	GetStatusCode() *int32
	SetBody(v *RecallDocumentResponseBody) *RecallDocumentResponse
	GetBody() *RecallDocumentResponseBody
}

type RecallDocumentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecallDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecallDocumentResponse) String() string {
	return dara.Prettify(s)
}

func (s RecallDocumentResponse) GoString() string {
	return s.String()
}

func (s *RecallDocumentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecallDocumentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecallDocumentResponse) GetBody() *RecallDocumentResponseBody {
	return s.Body
}

func (s *RecallDocumentResponse) SetHeaders(v map[string]*string) *RecallDocumentResponse {
	s.Headers = v
	return s
}

func (s *RecallDocumentResponse) SetStatusCode(v int32) *RecallDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *RecallDocumentResponse) SetBody(v *RecallDocumentResponseBody) *RecallDocumentResponse {
	s.Body = v
	return s
}

func (s *RecallDocumentResponse) Validate() error {
	return dara.Validate(s)
}
