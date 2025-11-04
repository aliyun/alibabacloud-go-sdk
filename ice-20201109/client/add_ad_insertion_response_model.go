// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAdInsertionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddAdInsertionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddAdInsertionResponse
	GetStatusCode() *int32
	SetBody(v *AddAdInsertionResponseBody) *AddAdInsertionResponse
	GetBody() *AddAdInsertionResponseBody
}

type AddAdInsertionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddAdInsertionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddAdInsertionResponse) String() string {
	return dara.Prettify(s)
}

func (s AddAdInsertionResponse) GoString() string {
	return s.String()
}

func (s *AddAdInsertionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddAdInsertionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddAdInsertionResponse) GetBody() *AddAdInsertionResponseBody {
	return s.Body
}

func (s *AddAdInsertionResponse) SetHeaders(v map[string]*string) *AddAdInsertionResponse {
	s.Headers = v
	return s
}

func (s *AddAdInsertionResponse) SetStatusCode(v int32) *AddAdInsertionResponse {
	s.StatusCode = &v
	return s
}

func (s *AddAdInsertionResponse) SetBody(v *AddAdInsertionResponseBody) *AddAdInsertionResponse {
	s.Body = v
	return s
}

func (s *AddAdInsertionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
