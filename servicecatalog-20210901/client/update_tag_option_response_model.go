// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTagOptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTagOptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTagOptionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTagOptionResponseBody) *UpdateTagOptionResponse
	GetBody() *UpdateTagOptionResponseBody
}

type UpdateTagOptionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTagOptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTagOptionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTagOptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateTagOptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTagOptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTagOptionResponse) GetBody() *UpdateTagOptionResponseBody {
	return s.Body
}

func (s *UpdateTagOptionResponse) SetHeaders(v map[string]*string) *UpdateTagOptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateTagOptionResponse) SetStatusCode(v int32) *UpdateTagOptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTagOptionResponse) SetBody(v *UpdateTagOptionResponseBody) *UpdateTagOptionResponse {
	s.Body = v
	return s
}

func (s *UpdateTagOptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
