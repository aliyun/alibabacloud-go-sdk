// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduledTaskExecutionDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetScheduledTaskExecutionDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetScheduledTaskExecutionDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetScheduledTaskExecutionDetailResponseBody) *GetScheduledTaskExecutionDetailResponse
	GetBody() *GetScheduledTaskExecutionDetailResponseBody
}

type GetScheduledTaskExecutionDetailResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetScheduledTaskExecutionDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetScheduledTaskExecutionDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledTaskExecutionDetailResponse) GoString() string {
	return s.String()
}

func (s *GetScheduledTaskExecutionDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetScheduledTaskExecutionDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetScheduledTaskExecutionDetailResponse) GetBody() *GetScheduledTaskExecutionDetailResponseBody {
	return s.Body
}

func (s *GetScheduledTaskExecutionDetailResponse) SetHeaders(v map[string]*string) *GetScheduledTaskExecutionDetailResponse {
	s.Headers = v
	return s
}

func (s *GetScheduledTaskExecutionDetailResponse) SetStatusCode(v int32) *GetScheduledTaskExecutionDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScheduledTaskExecutionDetailResponse) SetBody(v *GetScheduledTaskExecutionDetailResponseBody) *GetScheduledTaskExecutionDetailResponse {
	s.Body = v
	return s
}

func (s *GetScheduledTaskExecutionDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
