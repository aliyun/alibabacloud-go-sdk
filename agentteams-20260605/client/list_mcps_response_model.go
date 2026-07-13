// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMcpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMcpsResponse
	GetStatusCode() *int32
	SetBody(v *ListMcpsResponseBody) *ListMcpsResponse
	GetBody() *ListMcpsResponseBody
}

type ListMcpsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMcpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMcpsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMcpsResponse) GoString() string {
	return s.String()
}

func (s *ListMcpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMcpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMcpsResponse) GetBody() *ListMcpsResponseBody {
	return s.Body
}

func (s *ListMcpsResponse) SetHeaders(v map[string]*string) *ListMcpsResponse {
	s.Headers = v
	return s
}

func (s *ListMcpsResponse) SetStatusCode(v int32) *ListMcpsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMcpsResponse) SetBody(v *ListMcpsResponseBody) *ListMcpsResponse {
	s.Body = v
	return s
}

func (s *ListMcpsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
