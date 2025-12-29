// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCartoonResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddCartoonResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddCartoonResponse
	GetStatusCode() *int32
	SetBody(v *AddCartoonResponseBody) *AddCartoonResponse
	GetBody() *AddCartoonResponseBody
}

type AddCartoonResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCartoonResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCartoonResponse) String() string {
	return dara.Prettify(s)
}

func (s AddCartoonResponse) GoString() string {
	return s.String()
}

func (s *AddCartoonResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddCartoonResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddCartoonResponse) GetBody() *AddCartoonResponseBody {
	return s.Body
}

func (s *AddCartoonResponse) SetHeaders(v map[string]*string) *AddCartoonResponse {
	s.Headers = v
	return s
}

func (s *AddCartoonResponse) SetStatusCode(v int32) *AddCartoonResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCartoonResponse) SetBody(v *AddCartoonResponseBody) *AddCartoonResponse {
	s.Body = v
	return s
}

func (s *AddCartoonResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
