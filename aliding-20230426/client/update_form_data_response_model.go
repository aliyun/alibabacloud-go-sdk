// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFormDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFormDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFormDataResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFormDataResponseBody) *UpdateFormDataResponse
	GetBody() *UpdateFormDataResponseBody
}

type UpdateFormDataResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFormDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFormDataResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFormDataResponse) GoString() string {
	return s.String()
}

func (s *UpdateFormDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFormDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFormDataResponse) GetBody() *UpdateFormDataResponseBody {
	return s.Body
}

func (s *UpdateFormDataResponse) SetHeaders(v map[string]*string) *UpdateFormDataResponse {
	s.Headers = v
	return s
}

func (s *UpdateFormDataResponse) SetStatusCode(v int32) *UpdateFormDataResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFormDataResponse) SetBody(v *UpdateFormDataResponseBody) *UpdateFormDataResponse {
	s.Body = v
	return s
}

func (s *UpdateFormDataResponse) Validate() error {
	return dara.Validate(s)
}
