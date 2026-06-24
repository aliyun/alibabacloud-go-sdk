// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateMem0PublicConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AllocateMem0PublicConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AllocateMem0PublicConnectionResponse
	GetStatusCode() *int32
	SetBody(v *AllocateMem0PublicConnectionResponseBody) *AllocateMem0PublicConnectionResponse
	GetBody() *AllocateMem0PublicConnectionResponseBody
}

type AllocateMem0PublicConnectionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllocateMem0PublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AllocateMem0PublicConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s AllocateMem0PublicConnectionResponse) GoString() string {
	return s.String()
}

func (s *AllocateMem0PublicConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AllocateMem0PublicConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AllocateMem0PublicConnectionResponse) GetBody() *AllocateMem0PublicConnectionResponseBody {
	return s.Body
}

func (s *AllocateMem0PublicConnectionResponse) SetHeaders(v map[string]*string) *AllocateMem0PublicConnectionResponse {
	s.Headers = v
	return s
}

func (s *AllocateMem0PublicConnectionResponse) SetStatusCode(v int32) *AllocateMem0PublicConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *AllocateMem0PublicConnectionResponse) SetBody(v *AllocateMem0PublicConnectionResponseBody) *AllocateMem0PublicConnectionResponse {
	s.Body = v
	return s
}

func (s *AllocateMem0PublicConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
