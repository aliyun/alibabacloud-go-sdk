// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceInstanceLabelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateResourceInstanceLabelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateResourceInstanceLabelResponse
	GetStatusCode() *int32
	SetBody(v *UpdateResourceInstanceLabelResponseBody) *UpdateResourceInstanceLabelResponse
	GetBody() *UpdateResourceInstanceLabelResponseBody
}

type UpdateResourceInstanceLabelResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResourceInstanceLabelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResourceInstanceLabelResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceInstanceLabelResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceInstanceLabelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateResourceInstanceLabelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateResourceInstanceLabelResponse) GetBody() *UpdateResourceInstanceLabelResponseBody {
	return s.Body
}

func (s *UpdateResourceInstanceLabelResponse) SetHeaders(v map[string]*string) *UpdateResourceInstanceLabelResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourceInstanceLabelResponse) SetStatusCode(v int32) *UpdateResourceInstanceLabelResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResourceInstanceLabelResponse) SetBody(v *UpdateResourceInstanceLabelResponseBody) *UpdateResourceInstanceLabelResponse {
	s.Body = v
	return s
}

func (s *UpdateResourceInstanceLabelResponse) Validate() error {
	return dara.Validate(s)
}
