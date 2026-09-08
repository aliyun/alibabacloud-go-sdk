// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataMaskingEncryptionAlgorithmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDataMaskingEncryptionAlgorithmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDataMaskingEncryptionAlgorithmResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDataMaskingEncryptionAlgorithmResponseBody) *UpdateDataMaskingEncryptionAlgorithmResponse
	GetBody() *UpdateDataMaskingEncryptionAlgorithmResponseBody
}

type UpdateDataMaskingEncryptionAlgorithmResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataMaskingEncryptionAlgorithmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataMaskingEncryptionAlgorithmResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataMaskingEncryptionAlgorithmResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataMaskingEncryptionAlgorithmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDataMaskingEncryptionAlgorithmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDataMaskingEncryptionAlgorithmResponse) GetBody() *UpdateDataMaskingEncryptionAlgorithmResponseBody {
	return s.Body
}

func (s *UpdateDataMaskingEncryptionAlgorithmResponse) SetHeaders(v map[string]*string) *UpdateDataMaskingEncryptionAlgorithmResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataMaskingEncryptionAlgorithmResponse) SetStatusCode(v int32) *UpdateDataMaskingEncryptionAlgorithmResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataMaskingEncryptionAlgorithmResponse) SetBody(v *UpdateDataMaskingEncryptionAlgorithmResponseBody) *UpdateDataMaskingEncryptionAlgorithmResponse {
	s.Body = v
	return s
}

func (s *UpdateDataMaskingEncryptionAlgorithmResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
