// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateContextDBPublicConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AllocateContextDBPublicConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AllocateContextDBPublicConnectionResponse
	GetStatusCode() *int32
	SetBody(v *AllocateContextDBPublicConnectionResponseBody) *AllocateContextDBPublicConnectionResponse
	GetBody() *AllocateContextDBPublicConnectionResponseBody
}

type AllocateContextDBPublicConnectionResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllocateContextDBPublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AllocateContextDBPublicConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s AllocateContextDBPublicConnectionResponse) GoString() string {
	return s.String()
}

func (s *AllocateContextDBPublicConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AllocateContextDBPublicConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AllocateContextDBPublicConnectionResponse) GetBody() *AllocateContextDBPublicConnectionResponseBody {
	return s.Body
}

func (s *AllocateContextDBPublicConnectionResponse) SetHeaders(v map[string]*string) *AllocateContextDBPublicConnectionResponse {
	s.Headers = v
	return s
}

func (s *AllocateContextDBPublicConnectionResponse) SetStatusCode(v int32) *AllocateContextDBPublicConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *AllocateContextDBPublicConnectionResponse) SetBody(v *AllocateContextDBPublicConnectionResponseBody) *AllocateContextDBPublicConnectionResponse {
	s.Body = v
	return s
}

func (s *AllocateContextDBPublicConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
