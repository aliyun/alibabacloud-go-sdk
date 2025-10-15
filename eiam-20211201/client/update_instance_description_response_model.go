// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceDescriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateInstanceDescriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateInstanceDescriptionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateInstanceDescriptionResponseBody) *UpdateInstanceDescriptionResponse
	GetBody() *UpdateInstanceDescriptionResponseBody
}

type UpdateInstanceDescriptionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceDescriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceDescriptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceDescriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateInstanceDescriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateInstanceDescriptionResponse) GetBody() *UpdateInstanceDescriptionResponseBody {
	return s.Body
}

func (s *UpdateInstanceDescriptionResponse) SetHeaders(v map[string]*string) *UpdateInstanceDescriptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceDescriptionResponse) SetStatusCode(v int32) *UpdateInstanceDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceDescriptionResponse) SetBody(v *UpdateInstanceDescriptionResponseBody) *UpdateInstanceDescriptionResponse {
	s.Body = v
	return s
}

func (s *UpdateInstanceDescriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
