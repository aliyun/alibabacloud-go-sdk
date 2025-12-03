// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVariableGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVariableGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVariableGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateVariableGroupResponseBody) *CreateVariableGroupResponse
	GetBody() *CreateVariableGroupResponseBody
}

type CreateVariableGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVariableGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVariableGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVariableGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateVariableGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVariableGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVariableGroupResponse) GetBody() *CreateVariableGroupResponseBody {
	return s.Body
}

func (s *CreateVariableGroupResponse) SetHeaders(v map[string]*string) *CreateVariableGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateVariableGroupResponse) SetStatusCode(v int32) *CreateVariableGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVariableGroupResponse) SetBody(v *CreateVariableGroupResponseBody) *CreateVariableGroupResponse {
	s.Body = v
	return s
}

func (s *CreateVariableGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
