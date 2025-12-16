// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateABTestGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateABTestGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateABTestGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateABTestGroupResponseBody) *UpdateABTestGroupResponse
	GetBody() *UpdateABTestGroupResponseBody
}

type UpdateABTestGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateABTestGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateABTestGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateABTestGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateABTestGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateABTestGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateABTestGroupResponse) GetBody() *UpdateABTestGroupResponseBody {
	return s.Body
}

func (s *UpdateABTestGroupResponse) SetHeaders(v map[string]*string) *UpdateABTestGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateABTestGroupResponse) SetStatusCode(v int32) *UpdateABTestGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateABTestGroupResponse) SetBody(v *UpdateABTestGroupResponseBody) *UpdateABTestGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateABTestGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
