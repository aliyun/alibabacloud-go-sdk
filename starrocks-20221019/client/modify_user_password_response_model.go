// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserPasswordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyUserPasswordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyUserPasswordResponse
	GetStatusCode() *int32
	SetBody(v *ModifyUserPasswordResponseBody) *ModifyUserPasswordResponse
	GetBody() *ModifyUserPasswordResponseBody
}

type ModifyUserPasswordResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyUserPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyUserPasswordResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserPasswordResponse) GoString() string {
	return s.String()
}

func (s *ModifyUserPasswordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyUserPasswordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyUserPasswordResponse) GetBody() *ModifyUserPasswordResponseBody {
	return s.Body
}

func (s *ModifyUserPasswordResponse) SetHeaders(v map[string]*string) *ModifyUserPasswordResponse {
	s.Headers = v
	return s
}

func (s *ModifyUserPasswordResponse) SetStatusCode(v int32) *ModifyUserPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyUserPasswordResponse) SetBody(v *ModifyUserPasswordResponseBody) *ModifyUserPasswordResponse {
	s.Body = v
	return s
}

func (s *ModifyUserPasswordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
