// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationAccessPointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApplicationAccessPointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApplicationAccessPointsResponse
	GetStatusCode() *int32
	SetBody(v *ListApplicationAccessPointsResponseBody) *ListApplicationAccessPointsResponse
	GetBody() *ListApplicationAccessPointsResponseBody
}

type ListApplicationAccessPointsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationAccessPointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationAccessPointsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationAccessPointsResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationAccessPointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApplicationAccessPointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApplicationAccessPointsResponse) GetBody() *ListApplicationAccessPointsResponseBody {
	return s.Body
}

func (s *ListApplicationAccessPointsResponse) SetHeaders(v map[string]*string) *ListApplicationAccessPointsResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationAccessPointsResponse) SetStatusCode(v int32) *ListApplicationAccessPointsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationAccessPointsResponse) SetBody(v *ListApplicationAccessPointsResponseBody) *ListApplicationAccessPointsResponse {
	s.Body = v
	return s
}

func (s *ListApplicationAccessPointsResponse) Validate() error {
	return dara.Validate(s)
}
