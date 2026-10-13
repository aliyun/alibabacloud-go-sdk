// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMemoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddMemoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddMemoriesResponse
	GetStatusCode() *int32
	SetBody(v *AddMemoriesResponseBody) *AddMemoriesResponse
	GetBody() *AddMemoriesResponseBody
}

type AddMemoriesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddMemoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddMemoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s AddMemoriesResponse) GoString() string {
	return s.String()
}

func (s *AddMemoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddMemoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddMemoriesResponse) GetBody() *AddMemoriesResponseBody {
	return s.Body
}

func (s *AddMemoriesResponse) SetHeaders(v map[string]*string) *AddMemoriesResponse {
	s.Headers = v
	return s
}

func (s *AddMemoriesResponse) SetStatusCode(v int32) *AddMemoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *AddMemoriesResponse) SetBody(v *AddMemoriesResponseBody) *AddMemoriesResponse {
	s.Body = v
	return s
}

func (s *AddMemoriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
