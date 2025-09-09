// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceBuildLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListServiceBuildLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListServiceBuildLogsResponse
	GetStatusCode() *int32
	SetBody(v *ListServiceBuildLogsResponseBody) *ListServiceBuildLogsResponse
	GetBody() *ListServiceBuildLogsResponseBody
}

type ListServiceBuildLogsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceBuildLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceBuildLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListServiceBuildLogsResponse) GoString() string {
	return s.String()
}

func (s *ListServiceBuildLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListServiceBuildLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServiceBuildLogsResponse) GetBody() *ListServiceBuildLogsResponseBody {
	return s.Body
}

func (s *ListServiceBuildLogsResponse) SetHeaders(v map[string]*string) *ListServiceBuildLogsResponse {
	s.Headers = v
	return s
}

func (s *ListServiceBuildLogsResponse) SetStatusCode(v int32) *ListServiceBuildLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceBuildLogsResponse) SetBody(v *ListServiceBuildLogsResponseBody) *ListServiceBuildLogsResponse {
	s.Body = v
	return s
}

func (s *ListServiceBuildLogsResponse) Validate() error {
	return dara.Validate(s)
}
