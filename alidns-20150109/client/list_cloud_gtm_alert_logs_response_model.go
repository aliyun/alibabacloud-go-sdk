// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudGtmAlertLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCloudGtmAlertLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCloudGtmAlertLogsResponse
	GetStatusCode() *int32
	SetBody(v *ListCloudGtmAlertLogsResponseBody) *ListCloudGtmAlertLogsResponse
	GetBody() *ListCloudGtmAlertLogsResponseBody
}

type ListCloudGtmAlertLogsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCloudGtmAlertLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCloudGtmAlertLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmAlertLogsResponse) GoString() string {
	return s.String()
}

func (s *ListCloudGtmAlertLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCloudGtmAlertLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCloudGtmAlertLogsResponse) GetBody() *ListCloudGtmAlertLogsResponseBody {
	return s.Body
}

func (s *ListCloudGtmAlertLogsResponse) SetHeaders(v map[string]*string) *ListCloudGtmAlertLogsResponse {
	s.Headers = v
	return s
}

func (s *ListCloudGtmAlertLogsResponse) SetStatusCode(v int32) *ListCloudGtmAlertLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCloudGtmAlertLogsResponse) SetBody(v *ListCloudGtmAlertLogsResponseBody) *ListCloudGtmAlertLogsResponse {
	s.Body = v
	return s
}

func (s *ListCloudGtmAlertLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
