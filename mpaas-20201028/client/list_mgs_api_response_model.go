// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMgsApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMgsApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMgsApiResponse
	GetStatusCode() *int32
	SetBody(v *ListMgsApiResponseBody) *ListMgsApiResponse
	GetBody() *ListMgsApiResponseBody
}

type ListMgsApiResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMgsApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMgsApiResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMgsApiResponse) GoString() string {
	return s.String()
}

func (s *ListMgsApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMgsApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMgsApiResponse) GetBody() *ListMgsApiResponseBody {
	return s.Body
}

func (s *ListMgsApiResponse) SetHeaders(v map[string]*string) *ListMgsApiResponse {
	s.Headers = v
	return s
}

func (s *ListMgsApiResponse) SetStatusCode(v int32) *ListMgsApiResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMgsApiResponse) SetBody(v *ListMgsApiResponseBody) *ListMgsApiResponse {
	s.Body = v
	return s
}

func (s *ListMgsApiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
