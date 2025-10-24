// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAndroidInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopAndroidInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopAndroidInstanceResponse
	GetStatusCode() *int32
	SetBody(v *StopAndroidInstanceResponseBody) *StopAndroidInstanceResponse
	GetBody() *StopAndroidInstanceResponseBody
}

type StopAndroidInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopAndroidInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopAndroidInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s StopAndroidInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopAndroidInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopAndroidInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopAndroidInstanceResponse) GetBody() *StopAndroidInstanceResponseBody {
	return s.Body
}

func (s *StopAndroidInstanceResponse) SetHeaders(v map[string]*string) *StopAndroidInstanceResponse {
	s.Headers = v
	return s
}

func (s *StopAndroidInstanceResponse) SetStatusCode(v int32) *StopAndroidInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StopAndroidInstanceResponse) SetBody(v *StopAndroidInstanceResponseBody) *StopAndroidInstanceResponse {
	s.Body = v
	return s
}

func (s *StopAndroidInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
