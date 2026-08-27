// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduledTaskPushOptionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetScheduledTaskPushOptionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetScheduledTaskPushOptionsResponse
	GetStatusCode() *int32
	SetBody(v *GetScheduledTaskPushOptionsResponseBody) *GetScheduledTaskPushOptionsResponse
	GetBody() *GetScheduledTaskPushOptionsResponseBody
}

type GetScheduledTaskPushOptionsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetScheduledTaskPushOptionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetScheduledTaskPushOptionsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledTaskPushOptionsResponse) GoString() string {
	return s.String()
}

func (s *GetScheduledTaskPushOptionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetScheduledTaskPushOptionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetScheduledTaskPushOptionsResponse) GetBody() *GetScheduledTaskPushOptionsResponseBody {
	return s.Body
}

func (s *GetScheduledTaskPushOptionsResponse) SetHeaders(v map[string]*string) *GetScheduledTaskPushOptionsResponse {
	s.Headers = v
	return s
}

func (s *GetScheduledTaskPushOptionsResponse) SetStatusCode(v int32) *GetScheduledTaskPushOptionsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScheduledTaskPushOptionsResponse) SetBody(v *GetScheduledTaskPushOptionsResponseBody) *GetScheduledTaskPushOptionsResponse {
	s.Body = v
	return s
}

func (s *GetScheduledTaskPushOptionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
