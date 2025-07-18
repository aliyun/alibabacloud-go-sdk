// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateInstancePublicConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AllocateInstancePublicConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AllocateInstancePublicConnectionResponse
	GetStatusCode() *int32
	SetBody(v *AllocateInstancePublicConnectionResponseBody) *AllocateInstancePublicConnectionResponse
	GetBody() *AllocateInstancePublicConnectionResponseBody
}

type AllocateInstancePublicConnectionResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllocateInstancePublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AllocateInstancePublicConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s AllocateInstancePublicConnectionResponse) GoString() string {
	return s.String()
}

func (s *AllocateInstancePublicConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AllocateInstancePublicConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AllocateInstancePublicConnectionResponse) GetBody() *AllocateInstancePublicConnectionResponseBody {
	return s.Body
}

func (s *AllocateInstancePublicConnectionResponse) SetHeaders(v map[string]*string) *AllocateInstancePublicConnectionResponse {
	s.Headers = v
	return s
}

func (s *AllocateInstancePublicConnectionResponse) SetStatusCode(v int32) *AllocateInstancePublicConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *AllocateInstancePublicConnectionResponse) SetBody(v *AllocateInstancePublicConnectionResponseBody) *AllocateInstancePublicConnectionResponse {
	s.Body = v
	return s
}

func (s *AllocateInstancePublicConnectionResponse) Validate() error {
	return dara.Validate(s)
}
