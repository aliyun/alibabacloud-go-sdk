// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataCheckTaskHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataCheckTaskHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataCheckTaskHistoryResponse
	GetStatusCode() *int32
	SetBody(v *ListDataCheckTaskHistoryResponseBody) *ListDataCheckTaskHistoryResponse
	GetBody() *ListDataCheckTaskHistoryResponseBody
}

type ListDataCheckTaskHistoryResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataCheckTaskHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataCheckTaskHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataCheckTaskHistoryResponse) GoString() string {
	return s.String()
}

func (s *ListDataCheckTaskHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataCheckTaskHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataCheckTaskHistoryResponse) GetBody() *ListDataCheckTaskHistoryResponseBody {
	return s.Body
}

func (s *ListDataCheckTaskHistoryResponse) SetHeaders(v map[string]*string) *ListDataCheckTaskHistoryResponse {
	s.Headers = v
	return s
}

func (s *ListDataCheckTaskHistoryResponse) SetStatusCode(v int32) *ListDataCheckTaskHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataCheckTaskHistoryResponse) SetBody(v *ListDataCheckTaskHistoryResponseBody) *ListDataCheckTaskHistoryResponse {
	s.Body = v
	return s
}

func (s *ListDataCheckTaskHistoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
