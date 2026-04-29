// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudMonitorTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudMonitorTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudMonitorTaskResponse
	GetStatusCode() *int32
	SetBody(v *CloudMonitorTaskResponseBody) *CloudMonitorTaskResponse
	GetBody() *CloudMonitorTaskResponseBody
}

type CloudMonitorTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudMonitorTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudMonitorTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudMonitorTaskResponse) GoString() string {
	return s.String()
}

func (s *CloudMonitorTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudMonitorTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudMonitorTaskResponse) GetBody() *CloudMonitorTaskResponseBody {
	return s.Body
}

func (s *CloudMonitorTaskResponse) SetHeaders(v map[string]*string) *CloudMonitorTaskResponse {
	s.Headers = v
	return s
}

func (s *CloudMonitorTaskResponse) SetStatusCode(v int32) *CloudMonitorTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudMonitorTaskResponse) SetBody(v *CloudMonitorTaskResponseBody) *CloudMonitorTaskResponse {
	s.Body = v
	return s
}

func (s *CloudMonitorTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
