// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRescaleApplicationVerticallyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RescaleApplicationVerticallyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RescaleApplicationVerticallyResponse
	GetStatusCode() *int32
	SetBody(v *RescaleApplicationVerticallyResponseBody) *RescaleApplicationVerticallyResponse
	GetBody() *RescaleApplicationVerticallyResponseBody
}

type RescaleApplicationVerticallyResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RescaleApplicationVerticallyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RescaleApplicationVerticallyResponse) String() string {
	return dara.Prettify(s)
}

func (s RescaleApplicationVerticallyResponse) GoString() string {
	return s.String()
}

func (s *RescaleApplicationVerticallyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RescaleApplicationVerticallyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RescaleApplicationVerticallyResponse) GetBody() *RescaleApplicationVerticallyResponseBody {
	return s.Body
}

func (s *RescaleApplicationVerticallyResponse) SetHeaders(v map[string]*string) *RescaleApplicationVerticallyResponse {
	s.Headers = v
	return s
}

func (s *RescaleApplicationVerticallyResponse) SetStatusCode(v int32) *RescaleApplicationVerticallyResponse {
	s.StatusCode = &v
	return s
}

func (s *RescaleApplicationVerticallyResponse) SetBody(v *RescaleApplicationVerticallyResponseBody) *RescaleApplicationVerticallyResponse {
	s.Body = v
	return s
}

func (s *RescaleApplicationVerticallyResponse) Validate() error {
	return dara.Validate(s)
}
