// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeInstanceFromCenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeInstanceFromCenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeInstanceFromCenResponse
	GetStatusCode() *int32
	SetBody(v *RevokeInstanceFromCenResponseBody) *RevokeInstanceFromCenResponse
	GetBody() *RevokeInstanceFromCenResponseBody
}

type RevokeInstanceFromCenResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeInstanceFromCenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeInstanceFromCenResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeInstanceFromCenResponse) GoString() string {
	return s.String()
}

func (s *RevokeInstanceFromCenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeInstanceFromCenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeInstanceFromCenResponse) GetBody() *RevokeInstanceFromCenResponseBody {
	return s.Body
}

func (s *RevokeInstanceFromCenResponse) SetHeaders(v map[string]*string) *RevokeInstanceFromCenResponse {
	s.Headers = v
	return s
}

func (s *RevokeInstanceFromCenResponse) SetStatusCode(v int32) *RevokeInstanceFromCenResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeInstanceFromCenResponse) SetBody(v *RevokeInstanceFromCenResponseBody) *RevokeInstanceFromCenResponse {
	s.Body = v
	return s
}

func (s *RevokeInstanceFromCenResponse) Validate() error {
	return dara.Validate(s)
}
