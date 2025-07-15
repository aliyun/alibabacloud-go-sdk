// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLegacyAppraiseLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLegacyAppraiseLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLegacyAppraiseLogsResponse
	GetStatusCode() *int32
	SetBody(v *ListLegacyAppraiseLogsResponseBody) *ListLegacyAppraiseLogsResponse
	GetBody() *ListLegacyAppraiseLogsResponseBody
}

type ListLegacyAppraiseLogsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLegacyAppraiseLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLegacyAppraiseLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLegacyAppraiseLogsResponse) GoString() string {
	return s.String()
}

func (s *ListLegacyAppraiseLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLegacyAppraiseLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLegacyAppraiseLogsResponse) GetBody() *ListLegacyAppraiseLogsResponseBody {
	return s.Body
}

func (s *ListLegacyAppraiseLogsResponse) SetHeaders(v map[string]*string) *ListLegacyAppraiseLogsResponse {
	s.Headers = v
	return s
}

func (s *ListLegacyAppraiseLogsResponse) SetStatusCode(v int32) *ListLegacyAppraiseLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLegacyAppraiseLogsResponse) SetBody(v *ListLegacyAppraiseLogsResponseBody) *ListLegacyAppraiseLogsResponse {
	s.Body = v
	return s
}

func (s *ListLegacyAppraiseLogsResponse) Validate() error {
	return dara.Validate(s)
}
