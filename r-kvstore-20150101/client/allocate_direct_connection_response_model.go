// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateDirectConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AllocateDirectConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AllocateDirectConnectionResponse
	GetStatusCode() *int32
	SetBody(v *AllocateDirectConnectionResponseBody) *AllocateDirectConnectionResponse
	GetBody() *AllocateDirectConnectionResponseBody
}

type AllocateDirectConnectionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllocateDirectConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AllocateDirectConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s AllocateDirectConnectionResponse) GoString() string {
	return s.String()
}

func (s *AllocateDirectConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AllocateDirectConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AllocateDirectConnectionResponse) GetBody() *AllocateDirectConnectionResponseBody {
	return s.Body
}

func (s *AllocateDirectConnectionResponse) SetHeaders(v map[string]*string) *AllocateDirectConnectionResponse {
	s.Headers = v
	return s
}

func (s *AllocateDirectConnectionResponse) SetStatusCode(v int32) *AllocateDirectConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *AllocateDirectConnectionResponse) SetBody(v *AllocateDirectConnectionResponseBody) *AllocateDirectConnectionResponse {
	s.Body = v
	return s
}

func (s *AllocateDirectConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
