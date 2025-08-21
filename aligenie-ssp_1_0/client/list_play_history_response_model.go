// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPlayHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPlayHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPlayHistoryResponse
	GetStatusCode() *int32
	SetBody(v *ListPlayHistoryResponseBody) *ListPlayHistoryResponse
	GetBody() *ListPlayHistoryResponseBody
}

type ListPlayHistoryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPlayHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPlayHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPlayHistoryResponse) GoString() string {
	return s.String()
}

func (s *ListPlayHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPlayHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPlayHistoryResponse) GetBody() *ListPlayHistoryResponseBody {
	return s.Body
}

func (s *ListPlayHistoryResponse) SetHeaders(v map[string]*string) *ListPlayHistoryResponse {
	s.Headers = v
	return s
}

func (s *ListPlayHistoryResponse) SetStatusCode(v int32) *ListPlayHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPlayHistoryResponse) SetBody(v *ListPlayHistoryResponseBody) *ListPlayHistoryResponse {
	s.Body = v
	return s
}

func (s *ListPlayHistoryResponse) Validate() error {
	return dara.Validate(s)
}
