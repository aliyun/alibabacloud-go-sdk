// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStoragePoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateStoragePoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateStoragePoolResponse
	GetStatusCode() *int32
	SetBody(v *CreateStoragePoolResponseBody) *CreateStoragePoolResponse
	GetBody() *CreateStoragePoolResponseBody
}

type CreateStoragePoolResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateStoragePoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateStoragePoolResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateStoragePoolResponse) GoString() string {
	return s.String()
}

func (s *CreateStoragePoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateStoragePoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateStoragePoolResponse) GetBody() *CreateStoragePoolResponseBody {
	return s.Body
}

func (s *CreateStoragePoolResponse) SetHeaders(v map[string]*string) *CreateStoragePoolResponse {
	s.Headers = v
	return s
}

func (s *CreateStoragePoolResponse) SetStatusCode(v int32) *CreateStoragePoolResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStoragePoolResponse) SetBody(v *CreateStoragePoolResponseBody) *CreateStoragePoolResponse {
	s.Body = v
	return s
}

func (s *CreateStoragePoolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
