// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserDescriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateUserDescriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateUserDescriptionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateUserDescriptionResponseBody) *UpdateUserDescriptionResponse
	GetBody() *UpdateUserDescriptionResponseBody
}

type UpdateUserDescriptionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserDescriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserDescriptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserDescriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateUserDescriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateUserDescriptionResponse) GetBody() *UpdateUserDescriptionResponseBody {
	return s.Body
}

func (s *UpdateUserDescriptionResponse) SetHeaders(v map[string]*string) *UpdateUserDescriptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserDescriptionResponse) SetStatusCode(v int32) *UpdateUserDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserDescriptionResponse) SetBody(v *UpdateUserDescriptionResponseBody) *UpdateUserDescriptionResponse {
	s.Body = v
	return s
}

func (s *UpdateUserDescriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
