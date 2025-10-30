// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAxgGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAxgGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAxgGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateAxgGroupResponseBody) *CreateAxgGroupResponse
	GetBody() *CreateAxgGroupResponseBody
}

type CreateAxgGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAxgGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAxgGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAxgGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateAxgGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAxgGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAxgGroupResponse) GetBody() *CreateAxgGroupResponseBody {
	return s.Body
}

func (s *CreateAxgGroupResponse) SetHeaders(v map[string]*string) *CreateAxgGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateAxgGroupResponse) SetStatusCode(v int32) *CreateAxgGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAxgGroupResponse) SetBody(v *CreateAxgGroupResponseBody) *CreateAxgGroupResponse {
	s.Body = v
	return s
}

func (s *CreateAxgGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
