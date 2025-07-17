// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHttpApiOperationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateHttpApiOperationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateHttpApiOperationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateHttpApiOperationResponseBody) *UpdateHttpApiOperationResponse
	GetBody() *UpdateHttpApiOperationResponseBody
}

type UpdateHttpApiOperationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHttpApiOperationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHttpApiOperationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiOperationResponse) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiOperationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateHttpApiOperationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateHttpApiOperationResponse) GetBody() *UpdateHttpApiOperationResponseBody {
	return s.Body
}

func (s *UpdateHttpApiOperationResponse) SetHeaders(v map[string]*string) *UpdateHttpApiOperationResponse {
	s.Headers = v
	return s
}

func (s *UpdateHttpApiOperationResponse) SetStatusCode(v int32) *UpdateHttpApiOperationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHttpApiOperationResponse) SetBody(v *UpdateHttpApiOperationResponseBody) *UpdateHttpApiOperationResponse {
	s.Body = v
	return s
}

func (s *UpdateHttpApiOperationResponse) Validate() error {
	return dara.Validate(s)
}
