// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeMemoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeMemoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeMemoryResponse
	GetStatusCode() *int32
	SetBody(v *RevokeMemoryResponseBody) *RevokeMemoryResponse
	GetBody() *RevokeMemoryResponseBody
}

type RevokeMemoryResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeMemoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeMemoryResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeMemoryResponse) GoString() string {
	return s.String()
}

func (s *RevokeMemoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeMemoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeMemoryResponse) GetBody() *RevokeMemoryResponseBody {
	return s.Body
}

func (s *RevokeMemoryResponse) SetHeaders(v map[string]*string) *RevokeMemoryResponse {
	s.Headers = v
	return s
}

func (s *RevokeMemoryResponse) SetStatusCode(v int32) *RevokeMemoryResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeMemoryResponse) SetBody(v *RevokeMemoryResponseBody) *RevokeMemoryResponse {
	s.Body = v
	return s
}

func (s *RevokeMemoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
