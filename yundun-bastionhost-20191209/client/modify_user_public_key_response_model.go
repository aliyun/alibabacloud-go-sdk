// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserPublicKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyUserPublicKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyUserPublicKeyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyUserPublicKeyResponseBody) *ModifyUserPublicKeyResponse
	GetBody() *ModifyUserPublicKeyResponseBody
}

type ModifyUserPublicKeyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyUserPublicKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyUserPublicKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserPublicKeyResponse) GoString() string {
	return s.String()
}

func (s *ModifyUserPublicKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyUserPublicKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyUserPublicKeyResponse) GetBody() *ModifyUserPublicKeyResponseBody {
	return s.Body
}

func (s *ModifyUserPublicKeyResponse) SetHeaders(v map[string]*string) *ModifyUserPublicKeyResponse {
	s.Headers = v
	return s
}

func (s *ModifyUserPublicKeyResponse) SetStatusCode(v int32) *ModifyUserPublicKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyUserPublicKeyResponse) SetBody(v *ModifyUserPublicKeyResponseBody) *ModifyUserPublicKeyResponse {
	s.Body = v
	return s
}

func (s *ModifyUserPublicKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
