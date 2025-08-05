// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAliyunOfficialEventSourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAliyunOfficialEventSourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAliyunOfficialEventSourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListAliyunOfficialEventSourcesResponseBody) *ListAliyunOfficialEventSourcesResponse
	GetBody() *ListAliyunOfficialEventSourcesResponseBody
}

type ListAliyunOfficialEventSourcesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAliyunOfficialEventSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAliyunOfficialEventSourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAliyunOfficialEventSourcesResponse) GoString() string {
	return s.String()
}

func (s *ListAliyunOfficialEventSourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAliyunOfficialEventSourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAliyunOfficialEventSourcesResponse) GetBody() *ListAliyunOfficialEventSourcesResponseBody {
	return s.Body
}

func (s *ListAliyunOfficialEventSourcesResponse) SetHeaders(v map[string]*string) *ListAliyunOfficialEventSourcesResponse {
	s.Headers = v
	return s
}

func (s *ListAliyunOfficialEventSourcesResponse) SetStatusCode(v int32) *ListAliyunOfficialEventSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAliyunOfficialEventSourcesResponse) SetBody(v *ListAliyunOfficialEventSourcesResponseBody) *ListAliyunOfficialEventSourcesResponse {
	s.Body = v
	return s
}

func (s *ListAliyunOfficialEventSourcesResponse) Validate() error {
	return dara.Validate(s)
}
