// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDBTaskSQLJobDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDBTaskSQLJobDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDBTaskSQLJobDetailResponse
	GetStatusCode() *int32
	SetBody(v *ListDBTaskSQLJobDetailResponseBody) *ListDBTaskSQLJobDetailResponse
	GetBody() *ListDBTaskSQLJobDetailResponseBody
}

type ListDBTaskSQLJobDetailResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDBTaskSQLJobDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDBTaskSQLJobDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDBTaskSQLJobDetailResponse) GoString() string {
	return s.String()
}

func (s *ListDBTaskSQLJobDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDBTaskSQLJobDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDBTaskSQLJobDetailResponse) GetBody() *ListDBTaskSQLJobDetailResponseBody {
	return s.Body
}

func (s *ListDBTaskSQLJobDetailResponse) SetHeaders(v map[string]*string) *ListDBTaskSQLJobDetailResponse {
	s.Headers = v
	return s
}

func (s *ListDBTaskSQLJobDetailResponse) SetStatusCode(v int32) *ListDBTaskSQLJobDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDBTaskSQLJobDetailResponse) SetBody(v *ListDBTaskSQLJobDetailResponseBody) *ListDBTaskSQLJobDetailResponse {
	s.Body = v
	return s
}

func (s *ListDBTaskSQLJobDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
