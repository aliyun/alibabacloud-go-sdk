// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateClusterPublicConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AllocateClusterPublicConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AllocateClusterPublicConnectionResponse
	GetStatusCode() *int32
	SetBody(v *AllocateClusterPublicConnectionResponseBody) *AllocateClusterPublicConnectionResponse
	GetBody() *AllocateClusterPublicConnectionResponseBody
}

type AllocateClusterPublicConnectionResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllocateClusterPublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AllocateClusterPublicConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s AllocateClusterPublicConnectionResponse) GoString() string {
	return s.String()
}

func (s *AllocateClusterPublicConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AllocateClusterPublicConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AllocateClusterPublicConnectionResponse) GetBody() *AllocateClusterPublicConnectionResponseBody {
	return s.Body
}

func (s *AllocateClusterPublicConnectionResponse) SetHeaders(v map[string]*string) *AllocateClusterPublicConnectionResponse {
	s.Headers = v
	return s
}

func (s *AllocateClusterPublicConnectionResponse) SetStatusCode(v int32) *AllocateClusterPublicConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *AllocateClusterPublicConnectionResponse) SetBody(v *AllocateClusterPublicConnectionResponseBody) *AllocateClusterPublicConnectionResponse {
	s.Body = v
	return s
}

func (s *AllocateClusterPublicConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
