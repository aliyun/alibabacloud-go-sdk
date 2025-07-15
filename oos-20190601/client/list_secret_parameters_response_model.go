// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecretParametersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSecretParametersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSecretParametersResponse
	GetStatusCode() *int32
	SetBody(v *ListSecretParametersResponseBody) *ListSecretParametersResponse
	GetBody() *ListSecretParametersResponseBody
}

type ListSecretParametersResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSecretParametersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSecretParametersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSecretParametersResponse) GoString() string {
	return s.String()
}

func (s *ListSecretParametersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSecretParametersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSecretParametersResponse) GetBody() *ListSecretParametersResponseBody {
	return s.Body
}

func (s *ListSecretParametersResponse) SetHeaders(v map[string]*string) *ListSecretParametersResponse {
	s.Headers = v
	return s
}

func (s *ListSecretParametersResponse) SetStatusCode(v int32) *ListSecretParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSecretParametersResponse) SetBody(v *ListSecretParametersResponseBody) *ListSecretParametersResponse {
	s.Body = v
	return s
}

func (s *ListSecretParametersResponse) Validate() error {
	return dara.Validate(s)
}
