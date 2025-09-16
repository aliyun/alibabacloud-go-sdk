// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOnlineConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOnlineConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOnlineConfigsResponse
	GetStatusCode() *int32
	SetBody(v *ListOnlineConfigsResponseBody) *ListOnlineConfigsResponse
	GetBody() *ListOnlineConfigsResponseBody
}

type ListOnlineConfigsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOnlineConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOnlineConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOnlineConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListOnlineConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOnlineConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOnlineConfigsResponse) GetBody() *ListOnlineConfigsResponseBody {
	return s.Body
}

func (s *ListOnlineConfigsResponse) SetHeaders(v map[string]*string) *ListOnlineConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListOnlineConfigsResponse) SetStatusCode(v int32) *ListOnlineConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOnlineConfigsResponse) SetBody(v *ListOnlineConfigsResponseBody) *ListOnlineConfigsResponse {
	s.Body = v
	return s
}

func (s *ListOnlineConfigsResponse) Validate() error {
	return dara.Validate(s)
}
