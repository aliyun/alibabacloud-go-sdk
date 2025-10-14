// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudGtmMonitorTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCloudGtmMonitorTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCloudGtmMonitorTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListCloudGtmMonitorTemplatesResponseBody) *ListCloudGtmMonitorTemplatesResponse
	GetBody() *ListCloudGtmMonitorTemplatesResponseBody
}

type ListCloudGtmMonitorTemplatesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCloudGtmMonitorTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCloudGtmMonitorTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmMonitorTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListCloudGtmMonitorTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCloudGtmMonitorTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCloudGtmMonitorTemplatesResponse) GetBody() *ListCloudGtmMonitorTemplatesResponseBody {
	return s.Body
}

func (s *ListCloudGtmMonitorTemplatesResponse) SetHeaders(v map[string]*string) *ListCloudGtmMonitorTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponse) SetStatusCode(v int32) *ListCloudGtmMonitorTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponse) SetBody(v *ListCloudGtmMonitorTemplatesResponseBody) *ListCloudGtmMonitorTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
