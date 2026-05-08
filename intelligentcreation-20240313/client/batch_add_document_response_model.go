// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchAddDocumentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchAddDocumentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchAddDocumentResponse
	GetStatusCode() *int32
	SetBody(v *BatchAddDocumentResult) *BatchAddDocumentResponse
	GetBody() *BatchAddDocumentResult
}

type BatchAddDocumentResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchAddDocumentResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchAddDocumentResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchAddDocumentResponse) GoString() string {
	return s.String()
}

func (s *BatchAddDocumentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchAddDocumentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchAddDocumentResponse) GetBody() *BatchAddDocumentResult {
	return s.Body
}

func (s *BatchAddDocumentResponse) SetHeaders(v map[string]*string) *BatchAddDocumentResponse {
	s.Headers = v
	return s
}

func (s *BatchAddDocumentResponse) SetStatusCode(v int32) *BatchAddDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchAddDocumentResponse) SetBody(v *BatchAddDocumentResult) *BatchAddDocumentResponse {
	s.Body = v
	return s
}

func (s *BatchAddDocumentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
