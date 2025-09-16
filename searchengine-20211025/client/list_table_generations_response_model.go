// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTableGenerationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTableGenerationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTableGenerationsResponse
	GetStatusCode() *int32
	SetBody(v *ListTableGenerationsResponseBody) *ListTableGenerationsResponse
	GetBody() *ListTableGenerationsResponseBody
}

type ListTableGenerationsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTableGenerationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTableGenerationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTableGenerationsResponse) GoString() string {
	return s.String()
}

func (s *ListTableGenerationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTableGenerationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTableGenerationsResponse) GetBody() *ListTableGenerationsResponseBody {
	return s.Body
}

func (s *ListTableGenerationsResponse) SetHeaders(v map[string]*string) *ListTableGenerationsResponse {
	s.Headers = v
	return s
}

func (s *ListTableGenerationsResponse) SetStatusCode(v int32) *ListTableGenerationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTableGenerationsResponse) SetBody(v *ListTableGenerationsResponseBody) *ListTableGenerationsResponse {
	s.Body = v
	return s
}

func (s *ListTableGenerationsResponse) Validate() error {
	return dara.Validate(s)
}
