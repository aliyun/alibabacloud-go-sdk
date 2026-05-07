// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDigitalEmployeeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDigitalEmployeeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDigitalEmployeeResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDigitalEmployeeResponseBody) *DeleteDigitalEmployeeResponse
	GetBody() *DeleteDigitalEmployeeResponseBody
}

type DeleteDigitalEmployeeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDigitalEmployeeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDigitalEmployeeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDigitalEmployeeResponse) GoString() string {
	return s.String()
}

func (s *DeleteDigitalEmployeeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDigitalEmployeeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDigitalEmployeeResponse) GetBody() *DeleteDigitalEmployeeResponseBody {
	return s.Body
}

func (s *DeleteDigitalEmployeeResponse) SetHeaders(v map[string]*string) *DeleteDigitalEmployeeResponse {
	s.Headers = v
	return s
}

func (s *DeleteDigitalEmployeeResponse) SetStatusCode(v int32) *DeleteDigitalEmployeeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDigitalEmployeeResponse) SetBody(v *DeleteDigitalEmployeeResponseBody) *DeleteDigitalEmployeeResponse {
	s.Body = v
	return s
}

func (s *DeleteDigitalEmployeeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
