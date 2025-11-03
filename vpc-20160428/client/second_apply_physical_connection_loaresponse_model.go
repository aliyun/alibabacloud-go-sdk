// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSecondApplyPhysicalConnectionLOAResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SecondApplyPhysicalConnectionLOAResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SecondApplyPhysicalConnectionLOAResponse
	GetStatusCode() *int32
	SetBody(v *SecondApplyPhysicalConnectionLOAResponseBody) *SecondApplyPhysicalConnectionLOAResponse
	GetBody() *SecondApplyPhysicalConnectionLOAResponseBody
}

type SecondApplyPhysicalConnectionLOAResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SecondApplyPhysicalConnectionLOAResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SecondApplyPhysicalConnectionLOAResponse) String() string {
	return dara.Prettify(s)
}

func (s SecondApplyPhysicalConnectionLOAResponse) GoString() string {
	return s.String()
}

func (s *SecondApplyPhysicalConnectionLOAResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SecondApplyPhysicalConnectionLOAResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SecondApplyPhysicalConnectionLOAResponse) GetBody() *SecondApplyPhysicalConnectionLOAResponseBody {
	return s.Body
}

func (s *SecondApplyPhysicalConnectionLOAResponse) SetHeaders(v map[string]*string) *SecondApplyPhysicalConnectionLOAResponse {
	s.Headers = v
	return s
}

func (s *SecondApplyPhysicalConnectionLOAResponse) SetStatusCode(v int32) *SecondApplyPhysicalConnectionLOAResponse {
	s.StatusCode = &v
	return s
}

func (s *SecondApplyPhysicalConnectionLOAResponse) SetBody(v *SecondApplyPhysicalConnectionLOAResponseBody) *SecondApplyPhysicalConnectionLOAResponse {
	s.Body = v
	return s
}

func (s *SecondApplyPhysicalConnectionLOAResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
