// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListJobHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListJobHistoryResponse
	GetStatusCode() *int32
	SetBody(v *ListJobHistoryResponseBody) *ListJobHistoryResponse
	GetBody() *ListJobHistoryResponseBody
}

type ListJobHistoryResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListJobHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListJobHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListJobHistoryResponse) GoString() string {
	return s.String()
}

func (s *ListJobHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListJobHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListJobHistoryResponse) GetBody() *ListJobHistoryResponseBody {
	return s.Body
}

func (s *ListJobHistoryResponse) SetHeaders(v map[string]*string) *ListJobHistoryResponse {
	s.Headers = v
	return s
}

func (s *ListJobHistoryResponse) SetStatusCode(v int32) *ListJobHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobHistoryResponse) SetBody(v *ListJobHistoryResponseBody) *ListJobHistoryResponse {
	s.Body = v
	return s
}

func (s *ListJobHistoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
