// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOperationHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOperationHistoryResponse
	GetStatusCode() *int32
	SetBody(v *ListOperationHistoryResponseBody) *ListOperationHistoryResponse
	GetBody() *ListOperationHistoryResponseBody
}

type ListOperationHistoryResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOperationHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOperationHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOperationHistoryResponse) GoString() string {
	return s.String()
}

func (s *ListOperationHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOperationHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOperationHistoryResponse) GetBody() *ListOperationHistoryResponseBody {
	return s.Body
}

func (s *ListOperationHistoryResponse) SetHeaders(v map[string]*string) *ListOperationHistoryResponse {
	s.Headers = v
	return s
}

func (s *ListOperationHistoryResponse) SetStatusCode(v int32) *ListOperationHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOperationHistoryResponse) SetBody(v *ListOperationHistoryResponseBody) *ListOperationHistoryResponse {
	s.Body = v
	return s
}

func (s *ListOperationHistoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
