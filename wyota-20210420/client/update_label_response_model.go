// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLabelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLabelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLabelResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLabelResponseBody) *UpdateLabelResponse
	GetBody() *UpdateLabelResponseBody
}

type UpdateLabelResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLabelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLabelResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLabelResponse) GoString() string {
	return s.String()
}

func (s *UpdateLabelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLabelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLabelResponse) GetBody() *UpdateLabelResponseBody {
	return s.Body
}

func (s *UpdateLabelResponse) SetHeaders(v map[string]*string) *UpdateLabelResponse {
	s.Headers = v
	return s
}

func (s *UpdateLabelResponse) SetStatusCode(v int32) *UpdateLabelResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLabelResponse) SetBody(v *UpdateLabelResponseBody) *UpdateLabelResponse {
	s.Body = v
	return s
}

func (s *UpdateLabelResponse) Validate() error {
	return dara.Validate(s)
}
