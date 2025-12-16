// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSearchStrategiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSearchStrategiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSearchStrategiesResponse
	GetStatusCode() *int32
	SetBody(v *ListSearchStrategiesResponseBody) *ListSearchStrategiesResponse
	GetBody() *ListSearchStrategiesResponseBody
}

type ListSearchStrategiesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSearchStrategiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSearchStrategiesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSearchStrategiesResponse) GoString() string {
	return s.String()
}

func (s *ListSearchStrategiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSearchStrategiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSearchStrategiesResponse) GetBody() *ListSearchStrategiesResponseBody {
	return s.Body
}

func (s *ListSearchStrategiesResponse) SetHeaders(v map[string]*string) *ListSearchStrategiesResponse {
	s.Headers = v
	return s
}

func (s *ListSearchStrategiesResponse) SetStatusCode(v int32) *ListSearchStrategiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSearchStrategiesResponse) SetBody(v *ListSearchStrategiesResponseBody) *ListSearchStrategiesResponse {
	s.Body = v
	return s
}

func (s *ListSearchStrategiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
