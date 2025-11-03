// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDiscoverEventSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DiscoverEventSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DiscoverEventSourceResponse
	GetStatusCode() *int32
	SetBody(v *DiscoverEventSourceResponseBody) *DiscoverEventSourceResponse
	GetBody() *DiscoverEventSourceResponseBody
}

type DiscoverEventSourceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DiscoverEventSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DiscoverEventSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DiscoverEventSourceResponse) GoString() string {
	return s.String()
}

func (s *DiscoverEventSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DiscoverEventSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DiscoverEventSourceResponse) GetBody() *DiscoverEventSourceResponseBody {
	return s.Body
}

func (s *DiscoverEventSourceResponse) SetHeaders(v map[string]*string) *DiscoverEventSourceResponse {
	s.Headers = v
	return s
}

func (s *DiscoverEventSourceResponse) SetStatusCode(v int32) *DiscoverEventSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DiscoverEventSourceResponse) SetBody(v *DiscoverEventSourceResponseBody) *DiscoverEventSourceResponse {
	s.Body = v
	return s
}

func (s *DiscoverEventSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
