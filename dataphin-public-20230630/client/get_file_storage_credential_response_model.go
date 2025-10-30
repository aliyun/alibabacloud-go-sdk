// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileStorageCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFileStorageCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFileStorageCredentialResponse
	GetStatusCode() *int32
	SetBody(v *GetFileStorageCredentialResponseBody) *GetFileStorageCredentialResponse
	GetBody() *GetFileStorageCredentialResponseBody
}

type GetFileStorageCredentialResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFileStorageCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFileStorageCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFileStorageCredentialResponse) GoString() string {
	return s.String()
}

func (s *GetFileStorageCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFileStorageCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFileStorageCredentialResponse) GetBody() *GetFileStorageCredentialResponseBody {
	return s.Body
}

func (s *GetFileStorageCredentialResponse) SetHeaders(v map[string]*string) *GetFileStorageCredentialResponse {
	s.Headers = v
	return s
}

func (s *GetFileStorageCredentialResponse) SetStatusCode(v int32) *GetFileStorageCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileStorageCredentialResponse) SetBody(v *GetFileStorageCredentialResponseBody) *GetFileStorageCredentialResponse {
	s.Body = v
	return s
}

func (s *GetFileStorageCredentialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
