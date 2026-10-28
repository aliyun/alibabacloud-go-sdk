// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContactGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateContactGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateContactGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateContactGroupResponseBody) *UpdateContactGroupResponse
	GetBody() *UpdateContactGroupResponseBody
}

type UpdateContactGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateContactGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateContactGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateContactGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateContactGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateContactGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateContactGroupResponse) GetBody() *UpdateContactGroupResponseBody {
	return s.Body
}

func (s *UpdateContactGroupResponse) SetHeaders(v map[string]*string) *UpdateContactGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateContactGroupResponse) SetStatusCode(v int32) *UpdateContactGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateContactGroupResponse) SetBody(v *UpdateContactGroupResponseBody) *UpdateContactGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateContactGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
