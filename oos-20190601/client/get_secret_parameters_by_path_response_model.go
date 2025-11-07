// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecretParametersByPathResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSecretParametersByPathResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSecretParametersByPathResponse
	GetStatusCode() *int32
	SetBody(v *GetSecretParametersByPathResponseBody) *GetSecretParametersByPathResponse
	GetBody() *GetSecretParametersByPathResponseBody
}

type GetSecretParametersByPathResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSecretParametersByPathResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSecretParametersByPathResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSecretParametersByPathResponse) GoString() string {
	return s.String()
}

func (s *GetSecretParametersByPathResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSecretParametersByPathResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSecretParametersByPathResponse) GetBody() *GetSecretParametersByPathResponseBody {
	return s.Body
}

func (s *GetSecretParametersByPathResponse) SetHeaders(v map[string]*string) *GetSecretParametersByPathResponse {
	s.Headers = v
	return s
}

func (s *GetSecretParametersByPathResponse) SetStatusCode(v int32) *GetSecretParametersByPathResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSecretParametersByPathResponse) SetBody(v *GetSecretParametersByPathResponseBody) *GetSecretParametersByPathResponse {
	s.Body = v
	return s
}

func (s *GetSecretParametersByPathResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
