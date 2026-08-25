// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEncryptionConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEncryptionConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEncryptionConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetEncryptionConfigResponseBody) *GetEncryptionConfigResponse
	GetBody() *GetEncryptionConfigResponseBody
}

type GetEncryptionConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEncryptionConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEncryptionConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEncryptionConfigResponse) GoString() string {
	return s.String()
}

func (s *GetEncryptionConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEncryptionConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEncryptionConfigResponse) GetBody() *GetEncryptionConfigResponseBody {
	return s.Body
}

func (s *GetEncryptionConfigResponse) SetHeaders(v map[string]*string) *GetEncryptionConfigResponse {
	s.Headers = v
	return s
}

func (s *GetEncryptionConfigResponse) SetStatusCode(v int32) *GetEncryptionConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEncryptionConfigResponse) SetBody(v *GetEncryptionConfigResponseBody) *GetEncryptionConfigResponse {
	s.Body = v
	return s
}

func (s *GetEncryptionConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
