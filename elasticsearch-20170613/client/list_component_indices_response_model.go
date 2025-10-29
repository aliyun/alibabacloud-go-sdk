// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComponentIndicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListComponentIndicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListComponentIndicesResponse
	GetStatusCode() *int32
	SetBody(v *ListComponentIndicesResponseBody) *ListComponentIndicesResponse
	GetBody() *ListComponentIndicesResponseBody
}

type ListComponentIndicesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListComponentIndicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListComponentIndicesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListComponentIndicesResponse) GoString() string {
	return s.String()
}

func (s *ListComponentIndicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListComponentIndicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListComponentIndicesResponse) GetBody() *ListComponentIndicesResponseBody {
	return s.Body
}

func (s *ListComponentIndicesResponse) SetHeaders(v map[string]*string) *ListComponentIndicesResponse {
	s.Headers = v
	return s
}

func (s *ListComponentIndicesResponse) SetStatusCode(v int32) *ListComponentIndicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListComponentIndicesResponse) SetBody(v *ListComponentIndicesResponseBody) *ListComponentIndicesResponse {
	s.Body = v
	return s
}

func (s *ListComponentIndicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
