// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivateTTSVoicesCustomResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPrivateTTSVoicesCustomResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPrivateTTSVoicesCustomResponse
	GetStatusCode() *int32
	SetBody(v *ListPrivateTTSVoicesCustomResponseBody) *ListPrivateTTSVoicesCustomResponse
	GetBody() *ListPrivateTTSVoicesCustomResponseBody
}

type ListPrivateTTSVoicesCustomResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPrivateTTSVoicesCustomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPrivateTTSVoicesCustomResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateTTSVoicesCustomResponse) GoString() string {
	return s.String()
}

func (s *ListPrivateTTSVoicesCustomResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPrivateTTSVoicesCustomResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPrivateTTSVoicesCustomResponse) GetBody() *ListPrivateTTSVoicesCustomResponseBody {
	return s.Body
}

func (s *ListPrivateTTSVoicesCustomResponse) SetHeaders(v map[string]*string) *ListPrivateTTSVoicesCustomResponse {
	s.Headers = v
	return s
}

func (s *ListPrivateTTSVoicesCustomResponse) SetStatusCode(v int32) *ListPrivateTTSVoicesCustomResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrivateTTSVoicesCustomResponse) SetBody(v *ListPrivateTTSVoicesCustomResponseBody) *ListPrivateTTSVoicesCustomResponse {
	s.Body = v
	return s
}

func (s *ListPrivateTTSVoicesCustomResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
