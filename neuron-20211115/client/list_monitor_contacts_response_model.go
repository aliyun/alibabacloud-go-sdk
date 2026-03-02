// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMonitorContactsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMonitorContactsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMonitorContactsResponse
	GetStatusCode() *int32
	SetBody(v *MonitorContactPageResult) *ListMonitorContactsResponse
	GetBody() *MonitorContactPageResult
}

type ListMonitorContactsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MonitorContactPageResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMonitorContactsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMonitorContactsResponse) GoString() string {
	return s.String()
}

func (s *ListMonitorContactsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMonitorContactsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMonitorContactsResponse) GetBody() *MonitorContactPageResult {
	return s.Body
}

func (s *ListMonitorContactsResponse) SetHeaders(v map[string]*string) *ListMonitorContactsResponse {
	s.Headers = v
	return s
}

func (s *ListMonitorContactsResponse) SetStatusCode(v int32) *ListMonitorContactsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMonitorContactsResponse) SetBody(v *MonitorContactPageResult) *ListMonitorContactsResponse {
	s.Body = v
	return s
}

func (s *ListMonitorContactsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
