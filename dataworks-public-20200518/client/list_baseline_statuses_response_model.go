// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBaselineStatusesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBaselineStatusesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBaselineStatusesResponse
	GetStatusCode() *int32
	SetBody(v *ListBaselineStatusesResponseBody) *ListBaselineStatusesResponse
	GetBody() *ListBaselineStatusesResponseBody
}

type ListBaselineStatusesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBaselineStatusesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBaselineStatusesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBaselineStatusesResponse) GoString() string {
	return s.String()
}

func (s *ListBaselineStatusesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBaselineStatusesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBaselineStatusesResponse) GetBody() *ListBaselineStatusesResponseBody {
	return s.Body
}

func (s *ListBaselineStatusesResponse) SetHeaders(v map[string]*string) *ListBaselineStatusesResponse {
	s.Headers = v
	return s
}

func (s *ListBaselineStatusesResponse) SetStatusCode(v int32) *ListBaselineStatusesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBaselineStatusesResponse) SetBody(v *ListBaselineStatusesResponseBody) *ListBaselineStatusesResponse {
	s.Body = v
	return s
}

func (s *ListBaselineStatusesResponse) Validate() error {
	return dara.Validate(s)
}
