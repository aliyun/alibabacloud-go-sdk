// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAutoClipsTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAutoClipsTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAutoClipsTaskResponse
	GetStatusCode() *int32
	SetBody(v *ListAutoClipsTaskResponseBody) *ListAutoClipsTaskResponse
	GetBody() *ListAutoClipsTaskResponseBody
}

type ListAutoClipsTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAutoClipsTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAutoClipsTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAutoClipsTaskResponse) GoString() string {
	return s.String()
}

func (s *ListAutoClipsTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAutoClipsTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAutoClipsTaskResponse) GetBody() *ListAutoClipsTaskResponseBody {
	return s.Body
}

func (s *ListAutoClipsTaskResponse) SetHeaders(v map[string]*string) *ListAutoClipsTaskResponse {
	s.Headers = v
	return s
}

func (s *ListAutoClipsTaskResponse) SetStatusCode(v int32) *ListAutoClipsTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAutoClipsTaskResponse) SetBody(v *ListAutoClipsTaskResponseBody) *ListAutoClipsTaskResponse {
	s.Body = v
	return s
}

func (s *ListAutoClipsTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
