// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationDescriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateApplicationDescriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateApplicationDescriptionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateApplicationDescriptionResponseBody) *UpdateApplicationDescriptionResponse
	GetBody() *UpdateApplicationDescriptionResponseBody
}

type UpdateApplicationDescriptionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApplicationDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApplicationDescriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationDescriptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplicationDescriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateApplicationDescriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateApplicationDescriptionResponse) GetBody() *UpdateApplicationDescriptionResponseBody {
	return s.Body
}

func (s *UpdateApplicationDescriptionResponse) SetHeaders(v map[string]*string) *UpdateApplicationDescriptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateApplicationDescriptionResponse) SetStatusCode(v int32) *UpdateApplicationDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApplicationDescriptionResponse) SetBody(v *UpdateApplicationDescriptionResponseBody) *UpdateApplicationDescriptionResponse {
	s.Body = v
	return s
}

func (s *UpdateApplicationDescriptionResponse) Validate() error {
	return dara.Validate(s)
}
