// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateContext0PublicConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AllocateContext0PublicConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AllocateContext0PublicConnectionResponse
	GetStatusCode() *int32
	SetBody(v *AllocateContext0PublicConnectionResponseBody) *AllocateContext0PublicConnectionResponse
	GetBody() *AllocateContext0PublicConnectionResponseBody
}

type AllocateContext0PublicConnectionResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllocateContext0PublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AllocateContext0PublicConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s AllocateContext0PublicConnectionResponse) GoString() string {
	return s.String()
}

func (s *AllocateContext0PublicConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AllocateContext0PublicConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AllocateContext0PublicConnectionResponse) GetBody() *AllocateContext0PublicConnectionResponseBody {
	return s.Body
}

func (s *AllocateContext0PublicConnectionResponse) SetHeaders(v map[string]*string) *AllocateContext0PublicConnectionResponse {
	s.Headers = v
	return s
}

func (s *AllocateContext0PublicConnectionResponse) SetStatusCode(v int32) *AllocateContext0PublicConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *AllocateContext0PublicConnectionResponse) SetBody(v *AllocateContext0PublicConnectionResponseBody) *AllocateContext0PublicConnectionResponse {
	s.Body = v
	return s
}

func (s *AllocateContext0PublicConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
