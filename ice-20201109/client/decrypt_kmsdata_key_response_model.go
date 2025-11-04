// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDecryptKMSDataKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DecryptKMSDataKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DecryptKMSDataKeyResponse
	GetStatusCode() *int32
	SetBody(v *DecryptKMSDataKeyResponseBody) *DecryptKMSDataKeyResponse
	GetBody() *DecryptKMSDataKeyResponseBody
}

type DecryptKMSDataKeyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DecryptKMSDataKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DecryptKMSDataKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s DecryptKMSDataKeyResponse) GoString() string {
	return s.String()
}

func (s *DecryptKMSDataKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DecryptKMSDataKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DecryptKMSDataKeyResponse) GetBody() *DecryptKMSDataKeyResponseBody {
	return s.Body
}

func (s *DecryptKMSDataKeyResponse) SetHeaders(v map[string]*string) *DecryptKMSDataKeyResponse {
	s.Headers = v
	return s
}

func (s *DecryptKMSDataKeyResponse) SetStatusCode(v int32) *DecryptKMSDataKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DecryptKMSDataKeyResponse) SetBody(v *DecryptKMSDataKeyResponseBody) *DecryptKMSDataKeyResponse {
	s.Body = v
	return s
}

func (s *DecryptKMSDataKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
