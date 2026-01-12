// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendLocationServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SuspendLocationServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SuspendLocationServiceResponse
	GetStatusCode() *int32
	SetBody(v *SuspendLocationServiceResponseBody) *SuspendLocationServiceResponse
	GetBody() *SuspendLocationServiceResponseBody
}

type SuspendLocationServiceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SuspendLocationServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SuspendLocationServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s SuspendLocationServiceResponse) GoString() string {
	return s.String()
}

func (s *SuspendLocationServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SuspendLocationServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SuspendLocationServiceResponse) GetBody() *SuspendLocationServiceResponseBody {
	return s.Body
}

func (s *SuspendLocationServiceResponse) SetHeaders(v map[string]*string) *SuspendLocationServiceResponse {
	s.Headers = v
	return s
}

func (s *SuspendLocationServiceResponse) SetStatusCode(v int32) *SuspendLocationServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *SuspendLocationServiceResponse) SetBody(v *SuspendLocationServiceResponseBody) *SuspendLocationServiceResponse {
	s.Body = v
	return s
}

func (s *SuspendLocationServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
