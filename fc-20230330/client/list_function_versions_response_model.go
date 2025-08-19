// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFunctionVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFunctionVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFunctionVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListVersionsOutput) *ListFunctionVersionsResponse
	GetBody() *ListVersionsOutput
}

type ListFunctionVersionsResponse struct {
	Headers    map[string]*string  `json:"headers" xml:"headers"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVersionsOutput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFunctionVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListFunctionVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFunctionVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFunctionVersionsResponse) GetBody() *ListVersionsOutput {
	return s.Body
}

func (s *ListFunctionVersionsResponse) SetHeaders(v map[string]*string) *ListFunctionVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListFunctionVersionsResponse) SetStatusCode(v int32) *ListFunctionVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFunctionVersionsResponse) SetBody(v *ListVersionsOutput) *ListFunctionVersionsResponse {
	s.Body = v
	return s
}

func (s *ListFunctionVersionsResponse) Validate() error {
	return dara.Validate(s)
}
