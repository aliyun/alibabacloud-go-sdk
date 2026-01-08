// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDmTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDmTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDmTagResponse
	GetStatusCode() *int32
	SetBody(v *ListDmTagResponseBody) *ListDmTagResponse
	GetBody() *ListDmTagResponseBody
}

type ListDmTagResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDmTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDmTagResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDmTagResponse) GoString() string {
	return s.String()
}

func (s *ListDmTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDmTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDmTagResponse) GetBody() *ListDmTagResponseBody {
	return s.Body
}

func (s *ListDmTagResponse) SetHeaders(v map[string]*string) *ListDmTagResponse {
	s.Headers = v
	return s
}

func (s *ListDmTagResponse) SetStatusCode(v int32) *ListDmTagResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDmTagResponse) SetBody(v *ListDmTagResponseBody) *ListDmTagResponse {
	s.Body = v
	return s
}

func (s *ListDmTagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
