// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddContextsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddContextsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddContextsResponse
	GetStatusCode() *int32
	SetBody(v *AddContextsResponseBody) *AddContextsResponse
	GetBody() *AddContextsResponseBody
}

type AddContextsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddContextsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddContextsResponse) String() string {
	return dara.Prettify(s)
}

func (s AddContextsResponse) GoString() string {
	return s.String()
}

func (s *AddContextsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddContextsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddContextsResponse) GetBody() *AddContextsResponseBody {
	return s.Body
}

func (s *AddContextsResponse) SetHeaders(v map[string]*string) *AddContextsResponse {
	s.Headers = v
	return s
}

func (s *AddContextsResponse) SetStatusCode(v int32) *AddContextsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddContextsResponse) SetBody(v *AddContextsResponseBody) *AddContextsResponse {
	s.Body = v
	return s
}

func (s *AddContextsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
