// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAcceptRCInquiredSystemEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AcceptRCInquiredSystemEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AcceptRCInquiredSystemEventResponse
	GetStatusCode() *int32
	SetBody(v *AcceptRCInquiredSystemEventResponseBody) *AcceptRCInquiredSystemEventResponse
	GetBody() *AcceptRCInquiredSystemEventResponseBody
}

type AcceptRCInquiredSystemEventResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AcceptRCInquiredSystemEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AcceptRCInquiredSystemEventResponse) String() string {
	return dara.Prettify(s)
}

func (s AcceptRCInquiredSystemEventResponse) GoString() string {
	return s.String()
}

func (s *AcceptRCInquiredSystemEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AcceptRCInquiredSystemEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AcceptRCInquiredSystemEventResponse) GetBody() *AcceptRCInquiredSystemEventResponseBody {
	return s.Body
}

func (s *AcceptRCInquiredSystemEventResponse) SetHeaders(v map[string]*string) *AcceptRCInquiredSystemEventResponse {
	s.Headers = v
	return s
}

func (s *AcceptRCInquiredSystemEventResponse) SetStatusCode(v int32) *AcceptRCInquiredSystemEventResponse {
	s.StatusCode = &v
	return s
}

func (s *AcceptRCInquiredSystemEventResponse) SetBody(v *AcceptRCInquiredSystemEventResponseBody) *AcceptRCInquiredSystemEventResponse {
	s.Body = v
	return s
}

func (s *AcceptRCInquiredSystemEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
