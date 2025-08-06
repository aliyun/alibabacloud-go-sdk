// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMessageCloudMonitorEventListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMessageCloudMonitorEventListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMessageCloudMonitorEventListResponse
	GetStatusCode() *int32
	SetBody(v *GetMessageCloudMonitorEventListResponseBody) *GetMessageCloudMonitorEventListResponse
	GetBody() *GetMessageCloudMonitorEventListResponseBody
}

type GetMessageCloudMonitorEventListResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMessageCloudMonitorEventListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMessageCloudMonitorEventListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMessageCloudMonitorEventListResponse) GoString() string {
	return s.String()
}

func (s *GetMessageCloudMonitorEventListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMessageCloudMonitorEventListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMessageCloudMonitorEventListResponse) GetBody() *GetMessageCloudMonitorEventListResponseBody {
	return s.Body
}

func (s *GetMessageCloudMonitorEventListResponse) SetHeaders(v map[string]*string) *GetMessageCloudMonitorEventListResponse {
	s.Headers = v
	return s
}

func (s *GetMessageCloudMonitorEventListResponse) SetStatusCode(v int32) *GetMessageCloudMonitorEventListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMessageCloudMonitorEventListResponse) SetBody(v *GetMessageCloudMonitorEventListResponseBody) *GetMessageCloudMonitorEventListResponse {
	s.Body = v
	return s
}

func (s *GetMessageCloudMonitorEventListResponse) Validate() error {
	return dara.Validate(s)
}
