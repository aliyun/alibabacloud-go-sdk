// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTaskDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTaskDetailResponse
	GetStatusCode() *int32
	SetBody(v *ListTaskDetailResponseBody) *ListTaskDetailResponse
	GetBody() *ListTaskDetailResponseBody
}

type ListTaskDetailResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTaskDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTaskDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTaskDetailResponse) GoString() string {
	return s.String()
}

func (s *ListTaskDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTaskDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTaskDetailResponse) GetBody() *ListTaskDetailResponseBody {
	return s.Body
}

func (s *ListTaskDetailResponse) SetHeaders(v map[string]*string) *ListTaskDetailResponse {
	s.Headers = v
	return s
}

func (s *ListTaskDetailResponse) SetStatusCode(v int32) *ListTaskDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTaskDetailResponse) SetBody(v *ListTaskDetailResponseBody) *ListTaskDetailResponse {
	s.Body = v
	return s
}

func (s *ListTaskDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
