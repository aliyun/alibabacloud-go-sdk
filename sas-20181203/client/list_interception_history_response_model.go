// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInterceptionHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInterceptionHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInterceptionHistoryResponse
	GetStatusCode() *int32
	SetBody(v *ListInterceptionHistoryResponseBody) *ListInterceptionHistoryResponse
	GetBody() *ListInterceptionHistoryResponseBody
}

type ListInterceptionHistoryResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInterceptionHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInterceptionHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInterceptionHistoryResponse) GoString() string {
	return s.String()
}

func (s *ListInterceptionHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInterceptionHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInterceptionHistoryResponse) GetBody() *ListInterceptionHistoryResponseBody {
	return s.Body
}

func (s *ListInterceptionHistoryResponse) SetHeaders(v map[string]*string) *ListInterceptionHistoryResponse {
	s.Headers = v
	return s
}

func (s *ListInterceptionHistoryResponse) SetStatusCode(v int32) *ListInterceptionHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInterceptionHistoryResponse) SetBody(v *ListInterceptionHistoryResponseBody) *ListInterceptionHistoryResponse {
	s.Body = v
	return s
}

func (s *ListInterceptionHistoryResponse) Validate() error {
	return dara.Validate(s)
}
