// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutonomousNotifyEventsInRangeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAutonomousNotifyEventsInRangeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAutonomousNotifyEventsInRangeResponse
	GetStatusCode() *int32
	SetBody(v *GetAutonomousNotifyEventsInRangeResponseBody) *GetAutonomousNotifyEventsInRangeResponse
	GetBody() *GetAutonomousNotifyEventsInRangeResponseBody
}

type GetAutonomousNotifyEventsInRangeResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAutonomousNotifyEventsInRangeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAutonomousNotifyEventsInRangeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAutonomousNotifyEventsInRangeResponse) GoString() string {
	return s.String()
}

func (s *GetAutonomousNotifyEventsInRangeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAutonomousNotifyEventsInRangeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAutonomousNotifyEventsInRangeResponse) GetBody() *GetAutonomousNotifyEventsInRangeResponseBody {
	return s.Body
}

func (s *GetAutonomousNotifyEventsInRangeResponse) SetHeaders(v map[string]*string) *GetAutonomousNotifyEventsInRangeResponse {
	s.Headers = v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeResponse) SetStatusCode(v int32) *GetAutonomousNotifyEventsInRangeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeResponse) SetBody(v *GetAutonomousNotifyEventsInRangeResponseBody) *GetAutonomousNotifyEventsInRangeResponse {
	s.Body = v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
