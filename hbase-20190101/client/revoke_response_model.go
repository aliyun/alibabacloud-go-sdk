// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeResponse
	GetStatusCode() *int32
	SetBody(v *RevokeResponseBody) *RevokeResponse
	GetBody() *RevokeResponseBody
}

type RevokeResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeResponse) GoString() string {
	return s.String()
}

func (s *RevokeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeResponse) GetBody() *RevokeResponseBody {
	return s.Body
}

func (s *RevokeResponse) SetHeaders(v map[string]*string) *RevokeResponse {
	s.Headers = v
	return s
}

func (s *RevokeResponse) SetStatusCode(v int32) *RevokeResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeResponse) SetBody(v *RevokeResponseBody) *RevokeResponse {
	s.Body = v
	return s
}

func (s *RevokeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
