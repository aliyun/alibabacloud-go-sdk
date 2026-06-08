// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecurityStrategiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSecurityStrategiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSecurityStrategiesResponse
	GetStatusCode() *int32
	SetBody(v *ListSecurityStrategiesResponseBody) *ListSecurityStrategiesResponse
	GetBody() *ListSecurityStrategiesResponseBody
}

type ListSecurityStrategiesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSecurityStrategiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSecurityStrategiesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityStrategiesResponse) GoString() string {
	return s.String()
}

func (s *ListSecurityStrategiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSecurityStrategiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSecurityStrategiesResponse) GetBody() *ListSecurityStrategiesResponseBody {
	return s.Body
}

func (s *ListSecurityStrategiesResponse) SetHeaders(v map[string]*string) *ListSecurityStrategiesResponse {
	s.Headers = v
	return s
}

func (s *ListSecurityStrategiesResponse) SetStatusCode(v int32) *ListSecurityStrategiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSecurityStrategiesResponse) SetBody(v *ListSecurityStrategiesResponseBody) *ListSecurityStrategiesResponse {
	s.Body = v
	return s
}

func (s *ListSecurityStrategiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
