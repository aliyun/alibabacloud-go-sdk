// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWordGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWordGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWordGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWordGroupResponseBody) *UpdateWordGroupResponse
	GetBody() *UpdateWordGroupResponseBody
}

type UpdateWordGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWordGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWordGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWordGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateWordGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWordGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWordGroupResponse) GetBody() *UpdateWordGroupResponseBody {
	return s.Body
}

func (s *UpdateWordGroupResponse) SetHeaders(v map[string]*string) *UpdateWordGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateWordGroupResponse) SetStatusCode(v int32) *UpdateWordGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWordGroupResponse) SetBody(v *UpdateWordGroupResponseBody) *UpdateWordGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateWordGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
