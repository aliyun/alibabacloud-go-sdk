// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDocumentRetrieveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDocumentRetrieveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDocumentRetrieveResponse
	GetStatusCode() *int32
	SetBody(v *ListDocumentRetrieveResponseBody) *ListDocumentRetrieveResponse
	GetBody() *ListDocumentRetrieveResponseBody
}

type ListDocumentRetrieveResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDocumentRetrieveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDocumentRetrieveResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDocumentRetrieveResponse) GoString() string {
	return s.String()
}

func (s *ListDocumentRetrieveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDocumentRetrieveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDocumentRetrieveResponse) GetBody() *ListDocumentRetrieveResponseBody {
	return s.Body
}

func (s *ListDocumentRetrieveResponse) SetHeaders(v map[string]*string) *ListDocumentRetrieveResponse {
	s.Headers = v
	return s
}

func (s *ListDocumentRetrieveResponse) SetStatusCode(v int32) *ListDocumentRetrieveResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDocumentRetrieveResponse) SetBody(v *ListDocumentRetrieveResponseBody) *ListDocumentRetrieveResponse {
	s.Body = v
	return s
}

func (s *ListDocumentRetrieveResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
