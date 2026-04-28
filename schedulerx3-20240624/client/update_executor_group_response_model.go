// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExecutorGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateExecutorGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateExecutorGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateExecutorGroupResponseBody) *UpdateExecutorGroupResponse
	GetBody() *UpdateExecutorGroupResponseBody
}

type UpdateExecutorGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateExecutorGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateExecutorGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateExecutorGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateExecutorGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateExecutorGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateExecutorGroupResponse) GetBody() *UpdateExecutorGroupResponseBody {
	return s.Body
}

func (s *UpdateExecutorGroupResponse) SetHeaders(v map[string]*string) *UpdateExecutorGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateExecutorGroupResponse) SetStatusCode(v int32) *UpdateExecutorGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateExecutorGroupResponse) SetBody(v *UpdateExecutorGroupResponseBody) *UpdateExecutorGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateExecutorGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
