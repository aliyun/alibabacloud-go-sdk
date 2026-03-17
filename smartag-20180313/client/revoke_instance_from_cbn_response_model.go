// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeInstanceFromCbnResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeInstanceFromCbnResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeInstanceFromCbnResponse
	GetStatusCode() *int32
	SetBody(v *RevokeInstanceFromCbnResponseBody) *RevokeInstanceFromCbnResponse
	GetBody() *RevokeInstanceFromCbnResponseBody
}

type RevokeInstanceFromCbnResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeInstanceFromCbnResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeInstanceFromCbnResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeInstanceFromCbnResponse) GoString() string {
	return s.String()
}

func (s *RevokeInstanceFromCbnResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeInstanceFromCbnResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeInstanceFromCbnResponse) GetBody() *RevokeInstanceFromCbnResponseBody {
	return s.Body
}

func (s *RevokeInstanceFromCbnResponse) SetHeaders(v map[string]*string) *RevokeInstanceFromCbnResponse {
	s.Headers = v
	return s
}

func (s *RevokeInstanceFromCbnResponse) SetStatusCode(v int32) *RevokeInstanceFromCbnResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeInstanceFromCbnResponse) SetBody(v *RevokeInstanceFromCbnResponseBody) *RevokeInstanceFromCbnResponse {
	s.Body = v
	return s
}

func (s *RevokeInstanceFromCbnResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
