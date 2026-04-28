// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecyclebinResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRecyclebinResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRecyclebinResponse
	GetStatusCode() *int32
	SetBody(v *ListRecyclebinResponseBody) *ListRecyclebinResponse
	GetBody() *ListRecyclebinResponseBody
}

type ListRecyclebinResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRecyclebinResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRecyclebinResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRecyclebinResponse) GoString() string {
	return s.String()
}

func (s *ListRecyclebinResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRecyclebinResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRecyclebinResponse) GetBody() *ListRecyclebinResponseBody {
	return s.Body
}

func (s *ListRecyclebinResponse) SetHeaders(v map[string]*string) *ListRecyclebinResponse {
	s.Headers = v
	return s
}

func (s *ListRecyclebinResponse) SetStatusCode(v int32) *ListRecyclebinResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRecyclebinResponse) SetBody(v *ListRecyclebinResponseBody) *ListRecyclebinResponse {
	s.Body = v
	return s
}

func (s *ListRecyclebinResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
