// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindDBResourcePoolWithUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindDBResourcePoolWithUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindDBResourcePoolWithUserResponse
	GetStatusCode() *int32
	SetBody(v *UnbindDBResourcePoolWithUserResponseBody) *UnbindDBResourcePoolWithUserResponse
	GetBody() *UnbindDBResourcePoolWithUserResponseBody
}

type UnbindDBResourcePoolWithUserResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindDBResourcePoolWithUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindDBResourcePoolWithUserResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindDBResourcePoolWithUserResponse) GoString() string {
	return s.String()
}

func (s *UnbindDBResourcePoolWithUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindDBResourcePoolWithUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindDBResourcePoolWithUserResponse) GetBody() *UnbindDBResourcePoolWithUserResponseBody {
	return s.Body
}

func (s *UnbindDBResourcePoolWithUserResponse) SetHeaders(v map[string]*string) *UnbindDBResourcePoolWithUserResponse {
	s.Headers = v
	return s
}

func (s *UnbindDBResourcePoolWithUserResponse) SetStatusCode(v int32) *UnbindDBResourcePoolWithUserResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindDBResourcePoolWithUserResponse) SetBody(v *UnbindDBResourcePoolWithUserResponseBody) *UnbindDBResourcePoolWithUserResponse {
	s.Body = v
	return s
}

func (s *UnbindDBResourcePoolWithUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
