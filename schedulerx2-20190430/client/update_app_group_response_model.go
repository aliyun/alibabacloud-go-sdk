// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAppGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAppGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAppGroupResponseBody) *UpdateAppGroupResponse
	GetBody() *UpdateAppGroupResponseBody
}

type UpdateAppGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAppGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAppGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAppGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAppGroupResponse) GetBody() *UpdateAppGroupResponseBody {
	return s.Body
}

func (s *UpdateAppGroupResponse) SetHeaders(v map[string]*string) *UpdateAppGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppGroupResponse) SetStatusCode(v int32) *UpdateAppGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAppGroupResponse) SetBody(v *UpdateAppGroupResponseBody) *UpdateAppGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateAppGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
