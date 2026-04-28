// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExecutorGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateExecutorGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateExecutorGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateExecutorGroupResponseBody) *CreateExecutorGroupResponse
	GetBody() *CreateExecutorGroupResponseBody
}

type CreateExecutorGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateExecutorGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateExecutorGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateExecutorGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateExecutorGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateExecutorGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateExecutorGroupResponse) GetBody() *CreateExecutorGroupResponseBody {
	return s.Body
}

func (s *CreateExecutorGroupResponse) SetHeaders(v map[string]*string) *CreateExecutorGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateExecutorGroupResponse) SetStatusCode(v int32) *CreateExecutorGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExecutorGroupResponse) SetBody(v *CreateExecutorGroupResponseBody) *CreateExecutorGroupResponse {
	s.Body = v
	return s
}

func (s *CreateExecutorGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
