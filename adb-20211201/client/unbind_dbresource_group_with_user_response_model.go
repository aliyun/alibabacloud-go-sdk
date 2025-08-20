// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindDBResourceGroupWithUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindDBResourceGroupWithUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindDBResourceGroupWithUserResponse
	GetStatusCode() *int32
	SetBody(v *UnbindDBResourceGroupWithUserResponseBody) *UnbindDBResourceGroupWithUserResponse
	GetBody() *UnbindDBResourceGroupWithUserResponseBody
}

type UnbindDBResourceGroupWithUserResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindDBResourceGroupWithUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindDBResourceGroupWithUserResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindDBResourceGroupWithUserResponse) GoString() string {
	return s.String()
}

func (s *UnbindDBResourceGroupWithUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindDBResourceGroupWithUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindDBResourceGroupWithUserResponse) GetBody() *UnbindDBResourceGroupWithUserResponseBody {
	return s.Body
}

func (s *UnbindDBResourceGroupWithUserResponse) SetHeaders(v map[string]*string) *UnbindDBResourceGroupWithUserResponse {
	s.Headers = v
	return s
}

func (s *UnbindDBResourceGroupWithUserResponse) SetStatusCode(v int32) *UnbindDBResourceGroupWithUserResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindDBResourceGroupWithUserResponse) SetBody(v *UnbindDBResourceGroupWithUserResponseBody) *UnbindDBResourceGroupWithUserResponse {
	s.Body = v
	return s
}

func (s *UnbindDBResourceGroupWithUserResponse) Validate() error {
	return dara.Validate(s)
}
