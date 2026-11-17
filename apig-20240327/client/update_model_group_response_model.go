// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateModelGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateModelGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateModelGroupResponseBody) *UpdateModelGroupResponse
	GetBody() *UpdateModelGroupResponseBody
}

type UpdateModelGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateModelGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateModelGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateModelGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateModelGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateModelGroupResponse) GetBody() *UpdateModelGroupResponseBody {
	return s.Body
}

func (s *UpdateModelGroupResponse) SetHeaders(v map[string]*string) *UpdateModelGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateModelGroupResponse) SetStatusCode(v int32) *UpdateModelGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateModelGroupResponse) SetBody(v *UpdateModelGroupResponseBody) *UpdateModelGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateModelGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
