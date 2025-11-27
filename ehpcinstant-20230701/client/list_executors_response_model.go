// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExecutorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListExecutorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListExecutorsResponse
	GetStatusCode() *int32
	SetBody(v *ListExecutorsResponseBody) *ListExecutorsResponse
	GetBody() *ListExecutorsResponseBody
}

type ListExecutorsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExecutorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExecutorsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListExecutorsResponse) GoString() string {
	return s.String()
}

func (s *ListExecutorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListExecutorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListExecutorsResponse) GetBody() *ListExecutorsResponseBody {
	return s.Body
}

func (s *ListExecutorsResponse) SetHeaders(v map[string]*string) *ListExecutorsResponse {
	s.Headers = v
	return s
}

func (s *ListExecutorsResponse) SetStatusCode(v int32) *ListExecutorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExecutorsResponse) SetBody(v *ListExecutorsResponseBody) *ListExecutorsResponse {
	s.Body = v
	return s
}

func (s *ListExecutorsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
