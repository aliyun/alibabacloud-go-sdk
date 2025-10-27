// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUnfinishedOnceTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUnfinishedOnceTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUnfinishedOnceTaskResponse
	GetStatusCode() *int32
	SetBody(v *ListUnfinishedOnceTaskResponseBody) *ListUnfinishedOnceTaskResponse
	GetBody() *ListUnfinishedOnceTaskResponseBody
}

type ListUnfinishedOnceTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUnfinishedOnceTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUnfinishedOnceTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUnfinishedOnceTaskResponse) GoString() string {
	return s.String()
}

func (s *ListUnfinishedOnceTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUnfinishedOnceTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUnfinishedOnceTaskResponse) GetBody() *ListUnfinishedOnceTaskResponseBody {
	return s.Body
}

func (s *ListUnfinishedOnceTaskResponse) SetHeaders(v map[string]*string) *ListUnfinishedOnceTaskResponse {
	s.Headers = v
	return s
}

func (s *ListUnfinishedOnceTaskResponse) SetStatusCode(v int32) *ListUnfinishedOnceTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUnfinishedOnceTaskResponse) SetBody(v *ListUnfinishedOnceTaskResponseBody) *ListUnfinishedOnceTaskResponse {
	s.Body = v
	return s
}

func (s *ListUnfinishedOnceTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
