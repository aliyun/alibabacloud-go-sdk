// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetDataServiceAppSecretResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetDataServiceAppSecretResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetDataServiceAppSecretResponse
	GetStatusCode() *int32
	SetBody(v *ResetDataServiceAppSecretResponseBody) *ResetDataServiceAppSecretResponse
	GetBody() *ResetDataServiceAppSecretResponseBody
}

type ResetDataServiceAppSecretResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetDataServiceAppSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetDataServiceAppSecretResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetDataServiceAppSecretResponse) GoString() string {
	return s.String()
}

func (s *ResetDataServiceAppSecretResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetDataServiceAppSecretResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetDataServiceAppSecretResponse) GetBody() *ResetDataServiceAppSecretResponseBody {
	return s.Body
}

func (s *ResetDataServiceAppSecretResponse) SetHeaders(v map[string]*string) *ResetDataServiceAppSecretResponse {
	s.Headers = v
	return s
}

func (s *ResetDataServiceAppSecretResponse) SetStatusCode(v int32) *ResetDataServiceAppSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetDataServiceAppSecretResponse) SetBody(v *ResetDataServiceAppSecretResponseBody) *ResetDataServiceAppSecretResponse {
	s.Body = v
	return s
}

func (s *ResetDataServiceAppSecretResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
