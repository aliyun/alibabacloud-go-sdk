// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDocsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDocsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDocsResponse
	GetStatusCode() *int32
	SetBody(v *ListDocsResponseBody) *ListDocsResponse
	GetBody() *ListDocsResponseBody
}

type ListDocsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDocsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDocsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDocsResponse) GoString() string {
	return s.String()
}

func (s *ListDocsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDocsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDocsResponse) GetBody() *ListDocsResponseBody {
	return s.Body
}

func (s *ListDocsResponse) SetHeaders(v map[string]*string) *ListDocsResponse {
	s.Headers = v
	return s
}

func (s *ListDocsResponse) SetStatusCode(v int32) *ListDocsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDocsResponse) SetBody(v *ListDocsResponseBody) *ListDocsResponse {
	s.Body = v
	return s
}

func (s *ListDocsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
