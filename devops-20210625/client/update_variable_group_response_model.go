// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVariableGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateVariableGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateVariableGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateVariableGroupResponseBody) *UpdateVariableGroupResponse
	GetBody() *UpdateVariableGroupResponseBody
}

type UpdateVariableGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVariableGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVariableGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateVariableGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateVariableGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateVariableGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateVariableGroupResponse) GetBody() *UpdateVariableGroupResponseBody {
	return s.Body
}

func (s *UpdateVariableGroupResponse) SetHeaders(v map[string]*string) *UpdateVariableGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateVariableGroupResponse) SetStatusCode(v int32) *UpdateVariableGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVariableGroupResponse) SetBody(v *UpdateVariableGroupResponseBody) *UpdateVariableGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateVariableGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
