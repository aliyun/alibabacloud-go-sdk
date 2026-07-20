// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryResourceControlEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryResourceControlEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryResourceControlEventsResponse
	GetStatusCode() *int32
	SetBody(v *QueryResourceControlEventsResponseBody) *QueryResourceControlEventsResponse
	GetBody() *QueryResourceControlEventsResponseBody
}

type QueryResourceControlEventsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryResourceControlEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryResourceControlEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryResourceControlEventsResponse) GoString() string {
	return s.String()
}

func (s *QueryResourceControlEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryResourceControlEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryResourceControlEventsResponse) GetBody() *QueryResourceControlEventsResponseBody {
	return s.Body
}

func (s *QueryResourceControlEventsResponse) SetHeaders(v map[string]*string) *QueryResourceControlEventsResponse {
	s.Headers = v
	return s
}

func (s *QueryResourceControlEventsResponse) SetStatusCode(v int32) *QueryResourceControlEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryResourceControlEventsResponse) SetBody(v *QueryResourceControlEventsResponseBody) *QueryResourceControlEventsResponse {
	s.Body = v
	return s
}

func (s *QueryResourceControlEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
