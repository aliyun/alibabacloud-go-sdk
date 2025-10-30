// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelIsUsingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelIsUsingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelIsUsingResponse
	GetStatusCode() *int32
	SetBody(v *ModelIsUsingResponseBody) *ModelIsUsingResponse
	GetBody() *ModelIsUsingResponseBody
}

type ModelIsUsingResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelIsUsingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelIsUsingResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelIsUsingResponse) GoString() string {
	return s.String()
}

func (s *ModelIsUsingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelIsUsingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelIsUsingResponse) GetBody() *ModelIsUsingResponseBody {
	return s.Body
}

func (s *ModelIsUsingResponse) SetHeaders(v map[string]*string) *ModelIsUsingResponse {
	s.Headers = v
	return s
}

func (s *ModelIsUsingResponse) SetStatusCode(v int32) *ModelIsUsingResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelIsUsingResponse) SetBody(v *ModelIsUsingResponseBody) *ModelIsUsingResponse {
	s.Body = v
	return s
}

func (s *ModelIsUsingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
