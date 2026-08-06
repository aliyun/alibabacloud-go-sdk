// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDocumentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDocumentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDocumentsResponse
	GetStatusCode() *int32
	SetBody(v *AddDocumentsResponseBody) *AddDocumentsResponse
	GetBody() *AddDocumentsResponseBody
}

type AddDocumentsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDocumentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDocumentsResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDocumentsResponse) GoString() string {
	return s.String()
}

func (s *AddDocumentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDocumentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDocumentsResponse) GetBody() *AddDocumentsResponseBody {
	return s.Body
}

func (s *AddDocumentsResponse) SetHeaders(v map[string]*string) *AddDocumentsResponse {
	s.Headers = v
	return s
}

func (s *AddDocumentsResponse) SetStatusCode(v int32) *AddDocumentsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDocumentsResponse) SetBody(v *AddDocumentsResponseBody) *AddDocumentsResponse {
	s.Body = v
	return s
}

func (s *AddDocumentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
