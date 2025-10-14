// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomLineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCustomLineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCustomLineResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCustomLineResponseBody) *UpdateCustomLineResponse
	GetBody() *UpdateCustomLineResponseBody
}

type UpdateCustomLineResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCustomLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCustomLineResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomLineResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomLineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCustomLineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCustomLineResponse) GetBody() *UpdateCustomLineResponseBody {
	return s.Body
}

func (s *UpdateCustomLineResponse) SetHeaders(v map[string]*string) *UpdateCustomLineResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomLineResponse) SetStatusCode(v int32) *UpdateCustomLineResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustomLineResponse) SetBody(v *UpdateCustomLineResponseBody) *UpdateCustomLineResponse {
	s.Body = v
	return s
}

func (s *UpdateCustomLineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
