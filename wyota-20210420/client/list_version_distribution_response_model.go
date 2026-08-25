// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVersionDistributionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVersionDistributionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVersionDistributionResponse
	GetStatusCode() *int32
	SetBody(v *ListVersionDistributionResponseBody) *ListVersionDistributionResponse
	GetBody() *ListVersionDistributionResponseBody
}

type ListVersionDistributionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVersionDistributionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVersionDistributionResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVersionDistributionResponse) GoString() string {
	return s.String()
}

func (s *ListVersionDistributionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVersionDistributionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVersionDistributionResponse) GetBody() *ListVersionDistributionResponseBody {
	return s.Body
}

func (s *ListVersionDistributionResponse) SetHeaders(v map[string]*string) *ListVersionDistributionResponse {
	s.Headers = v
	return s
}

func (s *ListVersionDistributionResponse) SetStatusCode(v int32) *ListVersionDistributionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVersionDistributionResponse) SetBody(v *ListVersionDistributionResponseBody) *ListVersionDistributionResponse {
	s.Body = v
	return s
}

func (s *ListVersionDistributionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
