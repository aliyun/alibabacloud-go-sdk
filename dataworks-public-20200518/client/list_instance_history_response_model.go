// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceHistoryResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceHistoryResponseBody) *ListInstanceHistoryResponse
	GetBody() *ListInstanceHistoryResponseBody
}

type ListInstanceHistoryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceHistoryResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceHistoryResponse) GetBody() *ListInstanceHistoryResponseBody {
	return s.Body
}

func (s *ListInstanceHistoryResponse) SetHeaders(v map[string]*string) *ListInstanceHistoryResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceHistoryResponse) SetStatusCode(v int32) *ListInstanceHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceHistoryResponse) SetBody(v *ListInstanceHistoryResponseBody) *ListInstanceHistoryResponse {
	s.Body = v
	return s
}

func (s *ListInstanceHistoryResponse) Validate() error {
	return dara.Validate(s)
}
