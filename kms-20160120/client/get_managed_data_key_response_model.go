// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetManagedDataKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetManagedDataKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetManagedDataKeyResponse
	GetStatusCode() *int32
	SetBody(v *GetManagedDataKeyResponseBody) *GetManagedDataKeyResponse
	GetBody() *GetManagedDataKeyResponseBody
}

type GetManagedDataKeyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetManagedDataKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetManagedDataKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetManagedDataKeyResponse) GoString() string {
	return s.String()
}

func (s *GetManagedDataKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetManagedDataKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetManagedDataKeyResponse) GetBody() *GetManagedDataKeyResponseBody {
	return s.Body
}

func (s *GetManagedDataKeyResponse) SetHeaders(v map[string]*string) *GetManagedDataKeyResponse {
	s.Headers = v
	return s
}

func (s *GetManagedDataKeyResponse) SetStatusCode(v int32) *GetManagedDataKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetManagedDataKeyResponse) SetBody(v *GetManagedDataKeyResponseBody) *GetManagedDataKeyResponse {
	s.Body = v
	return s
}

func (s *GetManagedDataKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
