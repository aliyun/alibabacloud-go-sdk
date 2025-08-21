// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSummaryAppsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSummaryAppsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSummaryAppsResponse
	GetStatusCode() *int32
	SetBody(v *ListSummaryAppsResponseBody) *ListSummaryAppsResponse
	GetBody() *ListSummaryAppsResponseBody
}

type ListSummaryAppsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSummaryAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSummaryAppsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSummaryAppsResponse) GoString() string {
	return s.String()
}

func (s *ListSummaryAppsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSummaryAppsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSummaryAppsResponse) GetBody() *ListSummaryAppsResponseBody {
	return s.Body
}

func (s *ListSummaryAppsResponse) SetHeaders(v map[string]*string) *ListSummaryAppsResponse {
	s.Headers = v
	return s
}

func (s *ListSummaryAppsResponse) SetStatusCode(v int32) *ListSummaryAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSummaryAppsResponse) SetBody(v *ListSummaryAppsResponseBody) *ListSummaryAppsResponse {
	s.Body = v
	return s
}

func (s *ListSummaryAppsResponse) Validate() error {
	return dara.Validate(s)
}
