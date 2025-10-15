// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePredefinedDocumentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePredefinedDocumentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePredefinedDocumentResponse
	GetStatusCode() *int32
	SetBody(v *CreatePredefinedDocumentResponseBody) *CreatePredefinedDocumentResponse
	GetBody() *CreatePredefinedDocumentResponseBody
}

type CreatePredefinedDocumentResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePredefinedDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePredefinedDocumentResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePredefinedDocumentResponse) GoString() string {
	return s.String()
}

func (s *CreatePredefinedDocumentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePredefinedDocumentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePredefinedDocumentResponse) GetBody() *CreatePredefinedDocumentResponseBody {
	return s.Body
}

func (s *CreatePredefinedDocumentResponse) SetHeaders(v map[string]*string) *CreatePredefinedDocumentResponse {
	s.Headers = v
	return s
}

func (s *CreatePredefinedDocumentResponse) SetStatusCode(v int32) *CreatePredefinedDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePredefinedDocumentResponse) SetBody(v *CreatePredefinedDocumentResponseBody) *CreatePredefinedDocumentResponse {
	s.Body = v
	return s
}

func (s *CreatePredefinedDocumentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
