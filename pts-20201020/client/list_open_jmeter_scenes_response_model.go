// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOpenJMeterScenesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOpenJMeterScenesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOpenJMeterScenesResponse
	GetStatusCode() *int32
	SetBody(v *ListOpenJMeterScenesResponseBody) *ListOpenJMeterScenesResponse
	GetBody() *ListOpenJMeterScenesResponseBody
}

type ListOpenJMeterScenesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOpenJMeterScenesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOpenJMeterScenesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOpenJMeterScenesResponse) GoString() string {
	return s.String()
}

func (s *ListOpenJMeterScenesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOpenJMeterScenesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOpenJMeterScenesResponse) GetBody() *ListOpenJMeterScenesResponseBody {
	return s.Body
}

func (s *ListOpenJMeterScenesResponse) SetHeaders(v map[string]*string) *ListOpenJMeterScenesResponse {
	s.Headers = v
	return s
}

func (s *ListOpenJMeterScenesResponse) SetStatusCode(v int32) *ListOpenJMeterScenesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOpenJMeterScenesResponse) SetBody(v *ListOpenJMeterScenesResponseBody) *ListOpenJMeterScenesResponse {
	s.Body = v
	return s
}

func (s *ListOpenJMeterScenesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
