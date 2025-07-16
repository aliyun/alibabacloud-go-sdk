// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceLabelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateServiceLabelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateServiceLabelResponse
	GetStatusCode() *int32
	SetBody(v *UpdateServiceLabelResponseBody) *UpdateServiceLabelResponse
	GetBody() *UpdateServiceLabelResponseBody
}

type UpdateServiceLabelResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceLabelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceLabelResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceLabelResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceLabelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateServiceLabelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateServiceLabelResponse) GetBody() *UpdateServiceLabelResponseBody {
	return s.Body
}

func (s *UpdateServiceLabelResponse) SetHeaders(v map[string]*string) *UpdateServiceLabelResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceLabelResponse) SetStatusCode(v int32) *UpdateServiceLabelResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceLabelResponse) SetBody(v *UpdateServiceLabelResponseBody) *UpdateServiceLabelResponse {
	s.Body = v
	return s
}

func (s *UpdateServiceLabelResponse) Validate() error {
	return dara.Validate(s)
}
