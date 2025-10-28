// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPackageStorageCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPackageStorageCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPackageStorageCredentialResponse
	GetStatusCode() *int32
	SetBody(v *GetPackageStorageCredentialResponseBody) *GetPackageStorageCredentialResponse
	GetBody() *GetPackageStorageCredentialResponseBody
}

type GetPackageStorageCredentialResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPackageStorageCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPackageStorageCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPackageStorageCredentialResponse) GoString() string {
	return s.String()
}

func (s *GetPackageStorageCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPackageStorageCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPackageStorageCredentialResponse) GetBody() *GetPackageStorageCredentialResponseBody {
	return s.Body
}

func (s *GetPackageStorageCredentialResponse) SetHeaders(v map[string]*string) *GetPackageStorageCredentialResponse {
	s.Headers = v
	return s
}

func (s *GetPackageStorageCredentialResponse) SetStatusCode(v int32) *GetPackageStorageCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPackageStorageCredentialResponse) SetBody(v *GetPackageStorageCredentialResponseBody) *GetPackageStorageCredentialResponse {
	s.Body = v
	return s
}

func (s *GetPackageStorageCredentialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
