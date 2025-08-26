// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAirflowsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAirflowsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAirflowsResponse
	GetStatusCode() *int32
	SetBody(v *ListAirflowsResponseBody) *ListAirflowsResponse
	GetBody() *ListAirflowsResponseBody
}

type ListAirflowsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAirflowsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAirflowsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAirflowsResponse) GoString() string {
	return s.String()
}

func (s *ListAirflowsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAirflowsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAirflowsResponse) GetBody() *ListAirflowsResponseBody {
	return s.Body
}

func (s *ListAirflowsResponse) SetHeaders(v map[string]*string) *ListAirflowsResponse {
	s.Headers = v
	return s
}

func (s *ListAirflowsResponse) SetStatusCode(v int32) *ListAirflowsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAirflowsResponse) SetBody(v *ListAirflowsResponseBody) *ListAirflowsResponse {
	s.Body = v
	return s
}

func (s *ListAirflowsResponse) Validate() error {
	return dara.Validate(s)
}
