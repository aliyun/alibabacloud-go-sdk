// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackendModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBackendModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBackendModelResponse
	GetStatusCode() *int32
	SetBody(v *CreateBackendModelResponseBody) *CreateBackendModelResponse
	GetBody() *CreateBackendModelResponseBody
}

type CreateBackendModelResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBackendModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBackendModelResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBackendModelResponse) GoString() string {
	return s.String()
}

func (s *CreateBackendModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBackendModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBackendModelResponse) GetBody() *CreateBackendModelResponseBody {
	return s.Body
}

func (s *CreateBackendModelResponse) SetHeaders(v map[string]*string) *CreateBackendModelResponse {
	s.Headers = v
	return s
}

func (s *CreateBackendModelResponse) SetStatusCode(v int32) *CreateBackendModelResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBackendModelResponse) SetBody(v *CreateBackendModelResponseBody) *CreateBackendModelResponse {
	s.Body = v
	return s
}

func (s *CreateBackendModelResponse) Validate() error {
	return dara.Validate(s)
}
