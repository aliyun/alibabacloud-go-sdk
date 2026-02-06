// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInflightTaskTimeoutResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InflightTaskTimeoutResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InflightTaskTimeoutResponse
	GetStatusCode() *int32
	SetBody(v *InflightTaskTimeoutResponseBody) *InflightTaskTimeoutResponse
	GetBody() *InflightTaskTimeoutResponseBody
}

type InflightTaskTimeoutResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InflightTaskTimeoutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InflightTaskTimeoutResponse) String() string {
	return dara.Prettify(s)
}

func (s InflightTaskTimeoutResponse) GoString() string {
	return s.String()
}

func (s *InflightTaskTimeoutResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InflightTaskTimeoutResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InflightTaskTimeoutResponse) GetBody() *InflightTaskTimeoutResponseBody {
	return s.Body
}

func (s *InflightTaskTimeoutResponse) SetHeaders(v map[string]*string) *InflightTaskTimeoutResponse {
	s.Headers = v
	return s
}

func (s *InflightTaskTimeoutResponse) SetStatusCode(v int32) *InflightTaskTimeoutResponse {
	s.StatusCode = &v
	return s
}

func (s *InflightTaskTimeoutResponse) SetBody(v *InflightTaskTimeoutResponseBody) *InflightTaskTimeoutResponse {
	s.Body = v
	return s
}

func (s *InflightTaskTimeoutResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
