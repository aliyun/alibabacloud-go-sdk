// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudGtmMonitorNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCloudGtmMonitorNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCloudGtmMonitorNodesResponse
	GetStatusCode() *int32
	SetBody(v *ListCloudGtmMonitorNodesResponseBody) *ListCloudGtmMonitorNodesResponse
	GetBody() *ListCloudGtmMonitorNodesResponseBody
}

type ListCloudGtmMonitorNodesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCloudGtmMonitorNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCloudGtmMonitorNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmMonitorNodesResponse) GoString() string {
	return s.String()
}

func (s *ListCloudGtmMonitorNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCloudGtmMonitorNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCloudGtmMonitorNodesResponse) GetBody() *ListCloudGtmMonitorNodesResponseBody {
	return s.Body
}

func (s *ListCloudGtmMonitorNodesResponse) SetHeaders(v map[string]*string) *ListCloudGtmMonitorNodesResponse {
	s.Headers = v
	return s
}

func (s *ListCloudGtmMonitorNodesResponse) SetStatusCode(v int32) *ListCloudGtmMonitorNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCloudGtmMonitorNodesResponse) SetBody(v *ListCloudGtmMonitorNodesResponseBody) *ListCloudGtmMonitorNodesResponse {
	s.Body = v
	return s
}

func (s *ListCloudGtmMonitorNodesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
