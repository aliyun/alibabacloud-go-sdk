// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduledTaskExecutionRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetScheduledTaskExecutionRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetScheduledTaskExecutionRecordsResponse
	GetStatusCode() *int32
	SetBody(v *GetScheduledTaskExecutionRecordsResponseBody) *GetScheduledTaskExecutionRecordsResponse
	GetBody() *GetScheduledTaskExecutionRecordsResponseBody
}

type GetScheduledTaskExecutionRecordsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetScheduledTaskExecutionRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetScheduledTaskExecutionRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledTaskExecutionRecordsResponse) GoString() string {
	return s.String()
}

func (s *GetScheduledTaskExecutionRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetScheduledTaskExecutionRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetScheduledTaskExecutionRecordsResponse) GetBody() *GetScheduledTaskExecutionRecordsResponseBody {
	return s.Body
}

func (s *GetScheduledTaskExecutionRecordsResponse) SetHeaders(v map[string]*string) *GetScheduledTaskExecutionRecordsResponse {
	s.Headers = v
	return s
}

func (s *GetScheduledTaskExecutionRecordsResponse) SetStatusCode(v int32) *GetScheduledTaskExecutionRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScheduledTaskExecutionRecordsResponse) SetBody(v *GetScheduledTaskExecutionRecordsResponseBody) *GetScheduledTaskExecutionRecordsResponse {
	s.Body = v
	return s
}

func (s *GetScheduledTaskExecutionRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
