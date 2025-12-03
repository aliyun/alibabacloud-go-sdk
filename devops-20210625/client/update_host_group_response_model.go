// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHostGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateHostGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateHostGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateHostGroupResponseBody) *UpdateHostGroupResponse
	GetBody() *UpdateHostGroupResponseBody
}

type UpdateHostGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHostGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHostGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateHostGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateHostGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateHostGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateHostGroupResponse) GetBody() *UpdateHostGroupResponseBody {
	return s.Body
}

func (s *UpdateHostGroupResponse) SetHeaders(v map[string]*string) *UpdateHostGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateHostGroupResponse) SetStatusCode(v int32) *UpdateHostGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHostGroupResponse) SetBody(v *UpdateHostGroupResponseBody) *UpdateHostGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateHostGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
