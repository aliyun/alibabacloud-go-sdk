// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListImageVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListImageVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListImageVersionsResponseBody) *ListImageVersionsResponse
	GetBody() *ListImageVersionsResponseBody
}

type ListImageVersionsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListImageVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListImageVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListImageVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListImageVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListImageVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListImageVersionsResponse) GetBody() *ListImageVersionsResponseBody {
	return s.Body
}

func (s *ListImageVersionsResponse) SetHeaders(v map[string]*string) *ListImageVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListImageVersionsResponse) SetStatusCode(v int32) *ListImageVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListImageVersionsResponse) SetBody(v *ListImageVersionsResponseBody) *ListImageVersionsResponse {
	s.Body = v
	return s
}

func (s *ListImageVersionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
