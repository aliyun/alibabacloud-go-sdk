// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExecutorGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListExecutorGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListExecutorGroupResponse
	GetStatusCode() *int32
	SetBody(v *ListExecutorGroupResponseBody) *ListExecutorGroupResponse
	GetBody() *ListExecutorGroupResponseBody
}

type ListExecutorGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExecutorGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExecutorGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ListExecutorGroupResponse) GoString() string {
	return s.String()
}

func (s *ListExecutorGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListExecutorGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListExecutorGroupResponse) GetBody() *ListExecutorGroupResponseBody {
	return s.Body
}

func (s *ListExecutorGroupResponse) SetHeaders(v map[string]*string) *ListExecutorGroupResponse {
	s.Headers = v
	return s
}

func (s *ListExecutorGroupResponse) SetStatusCode(v int32) *ListExecutorGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExecutorGroupResponse) SetBody(v *ListExecutorGroupResponseBody) *ListExecutorGroupResponse {
	s.Body = v
	return s
}

func (s *ListExecutorGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
