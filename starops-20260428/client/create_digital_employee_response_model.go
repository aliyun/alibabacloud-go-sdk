// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDigitalEmployeeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDigitalEmployeeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDigitalEmployeeResponse
	GetStatusCode() *int32
	SetBody(v *CreateDigitalEmployeeResponseBody) *CreateDigitalEmployeeResponse
	GetBody() *CreateDigitalEmployeeResponseBody
}

type CreateDigitalEmployeeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDigitalEmployeeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDigitalEmployeeResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDigitalEmployeeResponse) GoString() string {
	return s.String()
}

func (s *CreateDigitalEmployeeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDigitalEmployeeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDigitalEmployeeResponse) GetBody() *CreateDigitalEmployeeResponseBody {
	return s.Body
}

func (s *CreateDigitalEmployeeResponse) SetHeaders(v map[string]*string) *CreateDigitalEmployeeResponse {
	s.Headers = v
	return s
}

func (s *CreateDigitalEmployeeResponse) SetStatusCode(v int32) *CreateDigitalEmployeeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDigitalEmployeeResponse) SetBody(v *CreateDigitalEmployeeResponseBody) *CreateDigitalEmployeeResponse {
	s.Body = v
	return s
}

func (s *CreateDigitalEmployeeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
