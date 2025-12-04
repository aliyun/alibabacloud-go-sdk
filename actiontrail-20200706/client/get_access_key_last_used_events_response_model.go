// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessKeyLastUsedEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAccessKeyLastUsedEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAccessKeyLastUsedEventsResponse
	GetStatusCode() *int32
	SetBody(v *GetAccessKeyLastUsedEventsResponseBody) *GetAccessKeyLastUsedEventsResponse
	GetBody() *GetAccessKeyLastUsedEventsResponseBody
}

type GetAccessKeyLastUsedEventsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccessKeyLastUsedEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccessKeyLastUsedEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAccessKeyLastUsedEventsResponse) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAccessKeyLastUsedEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAccessKeyLastUsedEventsResponse) GetBody() *GetAccessKeyLastUsedEventsResponseBody {
	return s.Body
}

func (s *GetAccessKeyLastUsedEventsResponse) SetHeaders(v map[string]*string) *GetAccessKeyLastUsedEventsResponse {
	s.Headers = v
	return s
}

func (s *GetAccessKeyLastUsedEventsResponse) SetStatusCode(v int32) *GetAccessKeyLastUsedEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccessKeyLastUsedEventsResponse) SetBody(v *GetAccessKeyLastUsedEventsResponseBody) *GetAccessKeyLastUsedEventsResponse {
	s.Body = v
	return s
}

func (s *GetAccessKeyLastUsedEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
