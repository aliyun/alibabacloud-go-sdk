// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCallTaskDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCallTaskDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCallTaskDetailResponse
	GetStatusCode() *int32
	SetBody(v *ListCallTaskDetailResponseBody) *ListCallTaskDetailResponse
	GetBody() *ListCallTaskDetailResponseBody
}

type ListCallTaskDetailResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCallTaskDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCallTaskDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCallTaskDetailResponse) GoString() string {
	return s.String()
}

func (s *ListCallTaskDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCallTaskDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCallTaskDetailResponse) GetBody() *ListCallTaskDetailResponseBody {
	return s.Body
}

func (s *ListCallTaskDetailResponse) SetHeaders(v map[string]*string) *ListCallTaskDetailResponse {
	s.Headers = v
	return s
}

func (s *ListCallTaskDetailResponse) SetStatusCode(v int32) *ListCallTaskDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCallTaskDetailResponse) SetBody(v *ListCallTaskDetailResponseBody) *ListCallTaskDetailResponse {
	s.Body = v
	return s
}

func (s *ListCallTaskDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
