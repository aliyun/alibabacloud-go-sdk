// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListForwardStrategiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListForwardStrategiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListForwardStrategiesResponse
	GetStatusCode() *int32
	SetBody(v *ListForwardStrategiesResponseBody) *ListForwardStrategiesResponse
	GetBody() *ListForwardStrategiesResponseBody
}

type ListForwardStrategiesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListForwardStrategiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListForwardStrategiesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListForwardStrategiesResponse) GoString() string {
	return s.String()
}

func (s *ListForwardStrategiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListForwardStrategiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListForwardStrategiesResponse) GetBody() *ListForwardStrategiesResponseBody {
	return s.Body
}

func (s *ListForwardStrategiesResponse) SetHeaders(v map[string]*string) *ListForwardStrategiesResponse {
	s.Headers = v
	return s
}

func (s *ListForwardStrategiesResponse) SetStatusCode(v int32) *ListForwardStrategiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListForwardStrategiesResponse) SetBody(v *ListForwardStrategiesResponseBody) *ListForwardStrategiesResponse {
	s.Body = v
	return s
}

func (s *ListForwardStrategiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
