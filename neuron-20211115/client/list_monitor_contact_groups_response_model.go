// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMonitorContactGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMonitorContactGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMonitorContactGroupsResponse
	GetStatusCode() *int32
	SetBody(v *MonitorContactPageResult) *ListMonitorContactGroupsResponse
	GetBody() *MonitorContactPageResult
}

type ListMonitorContactGroupsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MonitorContactPageResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMonitorContactGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMonitorContactGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListMonitorContactGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMonitorContactGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMonitorContactGroupsResponse) GetBody() *MonitorContactPageResult {
	return s.Body
}

func (s *ListMonitorContactGroupsResponse) SetHeaders(v map[string]*string) *ListMonitorContactGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListMonitorContactGroupsResponse) SetStatusCode(v int32) *ListMonitorContactGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMonitorContactGroupsResponse) SetBody(v *MonitorContactPageResult) *ListMonitorContactGroupsResponse {
	s.Body = v
	return s
}

func (s *ListMonitorContactGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
