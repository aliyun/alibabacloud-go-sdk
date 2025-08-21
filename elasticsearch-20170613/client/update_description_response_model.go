// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDescriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDescriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDescriptionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDescriptionResponseBody) *UpdateDescriptionResponse
	GetBody() *UpdateDescriptionResponseBody
}

type UpdateDescriptionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDescriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDescriptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateDescriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDescriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDescriptionResponse) GetBody() *UpdateDescriptionResponseBody {
	return s.Body
}

func (s *UpdateDescriptionResponse) SetHeaders(v map[string]*string) *UpdateDescriptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateDescriptionResponse) SetStatusCode(v int32) *UpdateDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDescriptionResponse) SetBody(v *UpdateDescriptionResponseBody) *UpdateDescriptionResponse {
	s.Body = v
	return s
}

func (s *UpdateDescriptionResponse) Validate() error {
	return dara.Validate(s)
}
