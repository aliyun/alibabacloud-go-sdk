// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetEncryptionConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetEncryptionConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetEncryptionConfigResponse
	GetStatusCode() *int32
	SetBody(v *SetEncryptionConfigResponseBody) *SetEncryptionConfigResponse
	GetBody() *SetEncryptionConfigResponseBody
}

type SetEncryptionConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetEncryptionConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetEncryptionConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SetEncryptionConfigResponse) GoString() string {
	return s.String()
}

func (s *SetEncryptionConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetEncryptionConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetEncryptionConfigResponse) GetBody() *SetEncryptionConfigResponseBody {
	return s.Body
}

func (s *SetEncryptionConfigResponse) SetHeaders(v map[string]*string) *SetEncryptionConfigResponse {
	s.Headers = v
	return s
}

func (s *SetEncryptionConfigResponse) SetStatusCode(v int32) *SetEncryptionConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetEncryptionConfigResponse) SetBody(v *SetEncryptionConfigResponseBody) *SetEncryptionConfigResponse {
	s.Body = v
	return s
}

func (s *SetEncryptionConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
