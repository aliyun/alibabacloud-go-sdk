// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRoutesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRoutesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRoutesResponse
	GetStatusCode() *int32
	SetBody(v *ListRoutesResponseBody) *ListRoutesResponse
	GetBody() *ListRoutesResponseBody
}

type ListRoutesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRoutesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRoutesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRoutesResponse) GoString() string {
	return s.String()
}

func (s *ListRoutesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRoutesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRoutesResponse) GetBody() *ListRoutesResponseBody {
	return s.Body
}

func (s *ListRoutesResponse) SetHeaders(v map[string]*string) *ListRoutesResponse {
	s.Headers = v
	return s
}

func (s *ListRoutesResponse) SetStatusCode(v int32) *ListRoutesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRoutesResponse) SetBody(v *ListRoutesResponseBody) *ListRoutesResponse {
	s.Body = v
	return s
}

func (s *ListRoutesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
