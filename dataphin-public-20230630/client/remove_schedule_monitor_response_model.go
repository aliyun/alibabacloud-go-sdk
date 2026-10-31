// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveScheduleMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveScheduleMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveScheduleMonitorResponse
	GetStatusCode() *int32
	SetBody(v *RemoveScheduleMonitorResponseBody) *RemoveScheduleMonitorResponse
	GetBody() *RemoveScheduleMonitorResponseBody
}

type RemoveScheduleMonitorResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveScheduleMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveScheduleMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveScheduleMonitorResponse) GoString() string {
	return s.String()
}

func (s *RemoveScheduleMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveScheduleMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveScheduleMonitorResponse) GetBody() *RemoveScheduleMonitorResponseBody {
	return s.Body
}

func (s *RemoveScheduleMonitorResponse) SetHeaders(v map[string]*string) *RemoveScheduleMonitorResponse {
	s.Headers = v
	return s
}

func (s *RemoveScheduleMonitorResponse) SetStatusCode(v int32) *RemoveScheduleMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveScheduleMonitorResponse) SetBody(v *RemoveScheduleMonitorResponseBody) *RemoveScheduleMonitorResponse {
	s.Body = v
	return s
}

func (s *RemoveScheduleMonitorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
