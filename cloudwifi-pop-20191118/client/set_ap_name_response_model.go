// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetApNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetApNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetApNameResponse
	GetStatusCode() *int32
	SetBody(v *SetApNameResponseBody) *SetApNameResponse
	GetBody() *SetApNameResponseBody
}

type SetApNameResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetApNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetApNameResponse) String() string {
	return dara.Prettify(s)
}

func (s SetApNameResponse) GoString() string {
	return s.String()
}

func (s *SetApNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetApNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetApNameResponse) GetBody() *SetApNameResponseBody {
	return s.Body
}

func (s *SetApNameResponse) SetHeaders(v map[string]*string) *SetApNameResponse {
	s.Headers = v
	return s
}

func (s *SetApNameResponse) SetStatusCode(v int32) *SetApNameResponse {
	s.StatusCode = &v
	return s
}

func (s *SetApNameResponse) SetBody(v *SetApNameResponseBody) *SetApNameResponse {
	s.Body = v
	return s
}

func (s *SetApNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
