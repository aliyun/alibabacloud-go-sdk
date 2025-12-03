// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceAuthsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListServiceAuthsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListServiceAuthsResponse
	GetStatusCode() *int32
	SetBody(v *ListServiceAuthsResponseBody) *ListServiceAuthsResponse
	GetBody() *ListServiceAuthsResponseBody
}

type ListServiceAuthsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceAuthsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceAuthsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListServiceAuthsResponse) GoString() string {
	return s.String()
}

func (s *ListServiceAuthsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListServiceAuthsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServiceAuthsResponse) GetBody() *ListServiceAuthsResponseBody {
	return s.Body
}

func (s *ListServiceAuthsResponse) SetHeaders(v map[string]*string) *ListServiceAuthsResponse {
	s.Headers = v
	return s
}

func (s *ListServiceAuthsResponse) SetStatusCode(v int32) *ListServiceAuthsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceAuthsResponse) SetBody(v *ListServiceAuthsResponseBody) *ListServiceAuthsResponse {
	s.Body = v
	return s
}

func (s *ListServiceAuthsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
