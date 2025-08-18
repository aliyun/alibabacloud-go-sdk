// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCnameFlatteningResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCnameFlatteningResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCnameFlatteningResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCnameFlatteningResponseBody) *UpdateCnameFlatteningResponse
	GetBody() *UpdateCnameFlatteningResponseBody
}

type UpdateCnameFlatteningResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCnameFlatteningResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCnameFlatteningResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCnameFlatteningResponse) GoString() string {
	return s.String()
}

func (s *UpdateCnameFlatteningResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCnameFlatteningResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCnameFlatteningResponse) GetBody() *UpdateCnameFlatteningResponseBody {
	return s.Body
}

func (s *UpdateCnameFlatteningResponse) SetHeaders(v map[string]*string) *UpdateCnameFlatteningResponse {
	s.Headers = v
	return s
}

func (s *UpdateCnameFlatteningResponse) SetStatusCode(v int32) *UpdateCnameFlatteningResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCnameFlatteningResponse) SetBody(v *UpdateCnameFlatteningResponseBody) *UpdateCnameFlatteningResponse {
	s.Body = v
	return s
}

func (s *UpdateCnameFlatteningResponse) Validate() error {
	return dara.Validate(s)
}
