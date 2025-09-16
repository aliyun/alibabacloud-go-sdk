// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAdvanceConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAdvanceConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAdvanceConfigsResponse
	GetStatusCode() *int32
	SetBody(v *ListAdvanceConfigsResponseBody) *ListAdvanceConfigsResponse
	GetBody() *ListAdvanceConfigsResponseBody
}

type ListAdvanceConfigsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAdvanceConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAdvanceConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAdvanceConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListAdvanceConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAdvanceConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAdvanceConfigsResponse) GetBody() *ListAdvanceConfigsResponseBody {
	return s.Body
}

func (s *ListAdvanceConfigsResponse) SetHeaders(v map[string]*string) *ListAdvanceConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListAdvanceConfigsResponse) SetStatusCode(v int32) *ListAdvanceConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAdvanceConfigsResponse) SetBody(v *ListAdvanceConfigsResponseBody) *ListAdvanceConfigsResponse {
	s.Body = v
	return s
}

func (s *ListAdvanceConfigsResponse) Validate() error {
	return dara.Validate(s)
}
