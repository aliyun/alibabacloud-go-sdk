// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocatePxfsPublicConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AllocatePxfsPublicConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AllocatePxfsPublicConnectionResponse
	GetStatusCode() *int32
	SetBody(v *AllocatePxfsPublicConnectionResponseBody) *AllocatePxfsPublicConnectionResponse
	GetBody() *AllocatePxfsPublicConnectionResponseBody
}

type AllocatePxfsPublicConnectionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllocatePxfsPublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AllocatePxfsPublicConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s AllocatePxfsPublicConnectionResponse) GoString() string {
	return s.String()
}

func (s *AllocatePxfsPublicConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AllocatePxfsPublicConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AllocatePxfsPublicConnectionResponse) GetBody() *AllocatePxfsPublicConnectionResponseBody {
	return s.Body
}

func (s *AllocatePxfsPublicConnectionResponse) SetHeaders(v map[string]*string) *AllocatePxfsPublicConnectionResponse {
	s.Headers = v
	return s
}

func (s *AllocatePxfsPublicConnectionResponse) SetStatusCode(v int32) *AllocatePxfsPublicConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *AllocatePxfsPublicConnectionResponse) SetBody(v *AllocatePxfsPublicConnectionResponseBody) *AllocatePxfsPublicConnectionResponse {
	s.Body = v
	return s
}

func (s *AllocatePxfsPublicConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
