// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBaselineConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBaselineConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBaselineConfigsResponse
	GetStatusCode() *int32
	SetBody(v *ListBaselineConfigsResponseBody) *ListBaselineConfigsResponse
	GetBody() *ListBaselineConfigsResponseBody
}

type ListBaselineConfigsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBaselineConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBaselineConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBaselineConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListBaselineConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBaselineConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBaselineConfigsResponse) GetBody() *ListBaselineConfigsResponseBody {
	return s.Body
}

func (s *ListBaselineConfigsResponse) SetHeaders(v map[string]*string) *ListBaselineConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListBaselineConfigsResponse) SetStatusCode(v int32) *ListBaselineConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBaselineConfigsResponse) SetBody(v *ListBaselineConfigsResponseBody) *ListBaselineConfigsResponse {
	s.Body = v
	return s
}

func (s *ListBaselineConfigsResponse) Validate() error {
	return dara.Validate(s)
}
