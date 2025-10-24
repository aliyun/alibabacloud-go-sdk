// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTagGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTagGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTagGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTagGroupResponseBody) *UpdateTagGroupResponse
	GetBody() *UpdateTagGroupResponseBody
}

type UpdateTagGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTagGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTagGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTagGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateTagGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTagGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTagGroupResponse) GetBody() *UpdateTagGroupResponseBody {
	return s.Body
}

func (s *UpdateTagGroupResponse) SetHeaders(v map[string]*string) *UpdateTagGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateTagGroupResponse) SetStatusCode(v int32) *UpdateTagGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTagGroupResponse) SetBody(v *UpdateTagGroupResponseBody) *UpdateTagGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateTagGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
