// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocatePublicConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AllocatePublicConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AllocatePublicConnectionResponse
	GetStatusCode() *int32
	SetBody(v *AllocatePublicConnectionResponseBody) *AllocatePublicConnectionResponse
	GetBody() *AllocatePublicConnectionResponseBody
}

type AllocatePublicConnectionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllocatePublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AllocatePublicConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s AllocatePublicConnectionResponse) GoString() string {
	return s.String()
}

func (s *AllocatePublicConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AllocatePublicConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AllocatePublicConnectionResponse) GetBody() *AllocatePublicConnectionResponseBody {
	return s.Body
}

func (s *AllocatePublicConnectionResponse) SetHeaders(v map[string]*string) *AllocatePublicConnectionResponse {
	s.Headers = v
	return s
}

func (s *AllocatePublicConnectionResponse) SetStatusCode(v int32) *AllocatePublicConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *AllocatePublicConnectionResponse) SetBody(v *AllocatePublicConnectionResponseBody) *AllocatePublicConnectionResponse {
	s.Body = v
	return s
}

func (s *AllocatePublicConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
