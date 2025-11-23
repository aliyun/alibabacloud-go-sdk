// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStandardGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateStandardGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateStandardGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateStandardGroupResponseBody) *UpdateStandardGroupResponse
	GetBody() *UpdateStandardGroupResponseBody
}

type UpdateStandardGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateStandardGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateStandardGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateStandardGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateStandardGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateStandardGroupResponse) GetBody() *UpdateStandardGroupResponseBody {
	return s.Body
}

func (s *UpdateStandardGroupResponse) SetHeaders(v map[string]*string) *UpdateStandardGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateStandardGroupResponse) SetStatusCode(v int32) *UpdateStandardGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateStandardGroupResponse) SetBody(v *UpdateStandardGroupResponseBody) *UpdateStandardGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateStandardGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
