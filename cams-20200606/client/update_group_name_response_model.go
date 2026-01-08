// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGroupNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGroupNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGroupNameResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGroupNameResponseBody) *UpdateGroupNameResponse
	GetBody() *UpdateGroupNameResponseBody
}

type UpdateGroupNameResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGroupNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGroupNameResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGroupNameResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGroupNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGroupNameResponse) GetBody() *UpdateGroupNameResponseBody {
	return s.Body
}

func (s *UpdateGroupNameResponse) SetHeaders(v map[string]*string) *UpdateGroupNameResponse {
	s.Headers = v
	return s
}

func (s *UpdateGroupNameResponse) SetStatusCode(v int32) *UpdateGroupNameResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGroupNameResponse) SetBody(v *UpdateGroupNameResponseBody) *UpdateGroupNameResponse {
	s.Body = v
	return s
}

func (s *UpdateGroupNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
