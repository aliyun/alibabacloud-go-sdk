// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFiltersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFiltersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFiltersResponse
	GetStatusCode() *int32
	SetBody(v *ListFiltersResponseBody) *ListFiltersResponse
	GetBody() *ListFiltersResponseBody
}

type ListFiltersResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFiltersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFiltersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFiltersResponse) GoString() string {
	return s.String()
}

func (s *ListFiltersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFiltersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFiltersResponse) GetBody() *ListFiltersResponseBody {
	return s.Body
}

func (s *ListFiltersResponse) SetHeaders(v map[string]*string) *ListFiltersResponse {
	s.Headers = v
	return s
}

func (s *ListFiltersResponse) SetStatusCode(v int32) *ListFiltersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFiltersResponse) SetBody(v *ListFiltersResponseBody) *ListFiltersResponse {
	s.Body = v
	return s
}

func (s *ListFiltersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
