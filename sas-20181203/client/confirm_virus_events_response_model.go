// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmVirusEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfirmVirusEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfirmVirusEventsResponse
	GetStatusCode() *int32
	SetBody(v *ConfirmVirusEventsResponseBody) *ConfirmVirusEventsResponse
	GetBody() *ConfirmVirusEventsResponseBody
}

type ConfirmVirusEventsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfirmVirusEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfirmVirusEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfirmVirusEventsResponse) GoString() string {
	return s.String()
}

func (s *ConfirmVirusEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfirmVirusEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfirmVirusEventsResponse) GetBody() *ConfirmVirusEventsResponseBody {
	return s.Body
}

func (s *ConfirmVirusEventsResponse) SetHeaders(v map[string]*string) *ConfirmVirusEventsResponse {
	s.Headers = v
	return s
}

func (s *ConfirmVirusEventsResponse) SetStatusCode(v int32) *ConfirmVirusEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfirmVirusEventsResponse) SetBody(v *ConfirmVirusEventsResponseBody) *ConfirmVirusEventsResponse {
	s.Body = v
	return s
}

func (s *ConfirmVirusEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
