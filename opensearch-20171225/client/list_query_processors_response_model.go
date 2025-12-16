// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQueryProcessorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListQueryProcessorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListQueryProcessorsResponse
	GetStatusCode() *int32
	SetBody(v *ListQueryProcessorsResponseBody) *ListQueryProcessorsResponse
	GetBody() *ListQueryProcessorsResponseBody
}

type ListQueryProcessorsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQueryProcessorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQueryProcessorsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListQueryProcessorsResponse) GoString() string {
	return s.String()
}

func (s *ListQueryProcessorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListQueryProcessorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListQueryProcessorsResponse) GetBody() *ListQueryProcessorsResponseBody {
	return s.Body
}

func (s *ListQueryProcessorsResponse) SetHeaders(v map[string]*string) *ListQueryProcessorsResponse {
	s.Headers = v
	return s
}

func (s *ListQueryProcessorsResponse) SetStatusCode(v int32) *ListQueryProcessorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQueryProcessorsResponse) SetBody(v *ListQueryProcessorsResponseBody) *ListQueryProcessorsResponse {
	s.Body = v
	return s
}

func (s *ListQueryProcessorsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
