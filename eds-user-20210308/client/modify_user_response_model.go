// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyUserResponse
	GetStatusCode() *int32
	SetBody(v *ModifyUserResponseBody) *ModifyUserResponse
	GetBody() *ModifyUserResponseBody
}

type ModifyUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserResponse) GoString() string {
	return s.String()
}

func (s *ModifyUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyUserResponse) GetBody() *ModifyUserResponseBody {
	return s.Body
}

func (s *ModifyUserResponse) SetHeaders(v map[string]*string) *ModifyUserResponse {
	s.Headers = v
	return s
}

func (s *ModifyUserResponse) SetStatusCode(v int32) *ModifyUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyUserResponse) SetBody(v *ModifyUserResponseBody) *ModifyUserResponse {
	s.Body = v
	return s
}

func (s *ModifyUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
