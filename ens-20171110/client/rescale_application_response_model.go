// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRescaleApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RescaleApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RescaleApplicationResponse
	GetStatusCode() *int32
	SetBody(v *RescaleApplicationResponseBody) *RescaleApplicationResponse
	GetBody() *RescaleApplicationResponseBody
}

type RescaleApplicationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RescaleApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RescaleApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s RescaleApplicationResponse) GoString() string {
	return s.String()
}

func (s *RescaleApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RescaleApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RescaleApplicationResponse) GetBody() *RescaleApplicationResponseBody {
	return s.Body
}

func (s *RescaleApplicationResponse) SetHeaders(v map[string]*string) *RescaleApplicationResponse {
	s.Headers = v
	return s
}

func (s *RescaleApplicationResponse) SetStatusCode(v int32) *RescaleApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *RescaleApplicationResponse) SetBody(v *RescaleApplicationResponseBody) *RescaleApplicationResponse {
	s.Body = v
	return s
}

func (s *RescaleApplicationResponse) Validate() error {
	return dara.Validate(s)
}
