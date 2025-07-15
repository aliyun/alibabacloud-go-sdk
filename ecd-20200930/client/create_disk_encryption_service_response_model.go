// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDiskEncryptionServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDiskEncryptionServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDiskEncryptionServiceResponse
	GetStatusCode() *int32
	SetBody(v *CreateDiskEncryptionServiceResponseBody) *CreateDiskEncryptionServiceResponse
	GetBody() *CreateDiskEncryptionServiceResponseBody
}

type CreateDiskEncryptionServiceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDiskEncryptionServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDiskEncryptionServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDiskEncryptionServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateDiskEncryptionServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDiskEncryptionServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDiskEncryptionServiceResponse) GetBody() *CreateDiskEncryptionServiceResponseBody {
	return s.Body
}

func (s *CreateDiskEncryptionServiceResponse) SetHeaders(v map[string]*string) *CreateDiskEncryptionServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateDiskEncryptionServiceResponse) SetStatusCode(v int32) *CreateDiskEncryptionServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDiskEncryptionServiceResponse) SetBody(v *CreateDiskEncryptionServiceResponseBody) *CreateDiskEncryptionServiceResponse {
	s.Body = v
	return s
}

func (s *CreateDiskEncryptionServiceResponse) Validate() error {
	return dara.Validate(s)
}
