// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecretParameterVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSecretParameterVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSecretParameterVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListSecretParameterVersionsResponseBody) *ListSecretParameterVersionsResponse
	GetBody() *ListSecretParameterVersionsResponseBody
}

type ListSecretParameterVersionsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSecretParameterVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSecretParameterVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSecretParameterVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListSecretParameterVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSecretParameterVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSecretParameterVersionsResponse) GetBody() *ListSecretParameterVersionsResponseBody {
	return s.Body
}

func (s *ListSecretParameterVersionsResponse) SetHeaders(v map[string]*string) *ListSecretParameterVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListSecretParameterVersionsResponse) SetStatusCode(v int32) *ListSecretParameterVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSecretParameterVersionsResponse) SetBody(v *ListSecretParameterVersionsResponseBody) *ListSecretParameterVersionsResponse {
	s.Body = v
	return s
}

func (s *ListSecretParameterVersionsResponse) Validate() error {
	return dara.Validate(s)
}
