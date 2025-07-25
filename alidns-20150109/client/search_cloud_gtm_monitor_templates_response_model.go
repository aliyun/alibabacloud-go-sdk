// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchCloudGtmMonitorTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchCloudGtmMonitorTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchCloudGtmMonitorTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *SearchCloudGtmMonitorTemplatesResponseBody) *SearchCloudGtmMonitorTemplatesResponse
	GetBody() *SearchCloudGtmMonitorTemplatesResponseBody
}

type SearchCloudGtmMonitorTemplatesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchCloudGtmMonitorTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchCloudGtmMonitorTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmMonitorTemplatesResponse) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmMonitorTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchCloudGtmMonitorTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchCloudGtmMonitorTemplatesResponse) GetBody() *SearchCloudGtmMonitorTemplatesResponseBody {
	return s.Body
}

func (s *SearchCloudGtmMonitorTemplatesResponse) SetHeaders(v map[string]*string) *SearchCloudGtmMonitorTemplatesResponse {
	s.Headers = v
	return s
}

func (s *SearchCloudGtmMonitorTemplatesResponse) SetStatusCode(v int32) *SearchCloudGtmMonitorTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchCloudGtmMonitorTemplatesResponse) SetBody(v *SearchCloudGtmMonitorTemplatesResponseBody) *SearchCloudGtmMonitorTemplatesResponse {
	s.Body = v
	return s
}

func (s *SearchCloudGtmMonitorTemplatesResponse) Validate() error {
	return dara.Validate(s)
}
