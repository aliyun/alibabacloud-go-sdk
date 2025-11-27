// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateReadWriteSplittingConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AllocateReadWriteSplittingConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AllocateReadWriteSplittingConnectionResponse
	GetStatusCode() *int32
	SetBody(v *AllocateReadWriteSplittingConnectionResponseBody) *AllocateReadWriteSplittingConnectionResponse
	GetBody() *AllocateReadWriteSplittingConnectionResponseBody
}

type AllocateReadWriteSplittingConnectionResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllocateReadWriteSplittingConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AllocateReadWriteSplittingConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s AllocateReadWriteSplittingConnectionResponse) GoString() string {
	return s.String()
}

func (s *AllocateReadWriteSplittingConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AllocateReadWriteSplittingConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AllocateReadWriteSplittingConnectionResponse) GetBody() *AllocateReadWriteSplittingConnectionResponseBody {
	return s.Body
}

func (s *AllocateReadWriteSplittingConnectionResponse) SetHeaders(v map[string]*string) *AllocateReadWriteSplittingConnectionResponse {
	s.Headers = v
	return s
}

func (s *AllocateReadWriteSplittingConnectionResponse) SetStatusCode(v int32) *AllocateReadWriteSplittingConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *AllocateReadWriteSplittingConnectionResponse) SetBody(v *AllocateReadWriteSplittingConnectionResponseBody) *AllocateReadWriteSplittingConnectionResponse {
	s.Body = v
	return s
}

func (s *AllocateReadWriteSplittingConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
