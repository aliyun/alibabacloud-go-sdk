// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAdInsertionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAdInsertionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAdInsertionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAdInsertionResponseBody) *UpdateAdInsertionResponse
	GetBody() *UpdateAdInsertionResponseBody
}

type UpdateAdInsertionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAdInsertionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAdInsertionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAdInsertionResponse) GoString() string {
	return s.String()
}

func (s *UpdateAdInsertionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAdInsertionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAdInsertionResponse) GetBody() *UpdateAdInsertionResponseBody {
	return s.Body
}

func (s *UpdateAdInsertionResponse) SetHeaders(v map[string]*string) *UpdateAdInsertionResponse {
	s.Headers = v
	return s
}

func (s *UpdateAdInsertionResponse) SetStatusCode(v int32) *UpdateAdInsertionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAdInsertionResponse) SetBody(v *UpdateAdInsertionResponseBody) *UpdateAdInsertionResponse {
	s.Body = v
	return s
}

func (s *UpdateAdInsertionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
