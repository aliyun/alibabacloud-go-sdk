// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDigitalEmployeeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDigitalEmployeeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDigitalEmployeeResponse
	GetStatusCode() *int32
	SetBody(v *GetDigitalEmployeeResponseBody) *GetDigitalEmployeeResponse
	GetBody() *GetDigitalEmployeeResponseBody
}

type GetDigitalEmployeeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDigitalEmployeeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDigitalEmployeeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDigitalEmployeeResponse) GoString() string {
	return s.String()
}

func (s *GetDigitalEmployeeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDigitalEmployeeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDigitalEmployeeResponse) GetBody() *GetDigitalEmployeeResponseBody {
	return s.Body
}

func (s *GetDigitalEmployeeResponse) SetHeaders(v map[string]*string) *GetDigitalEmployeeResponse {
	s.Headers = v
	return s
}

func (s *GetDigitalEmployeeResponse) SetStatusCode(v int32) *GetDigitalEmployeeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDigitalEmployeeResponse) SetBody(v *GetDigitalEmployeeResponseBody) *GetDigitalEmployeeResponse {
	s.Body = v
	return s
}

func (s *GetDigitalEmployeeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
