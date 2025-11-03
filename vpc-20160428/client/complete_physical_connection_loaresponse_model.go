// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompletePhysicalConnectionLOAResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CompletePhysicalConnectionLOAResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CompletePhysicalConnectionLOAResponse
	GetStatusCode() *int32
	SetBody(v *CompletePhysicalConnectionLOAResponseBody) *CompletePhysicalConnectionLOAResponse
	GetBody() *CompletePhysicalConnectionLOAResponseBody
}

type CompletePhysicalConnectionLOAResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CompletePhysicalConnectionLOAResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CompletePhysicalConnectionLOAResponse) String() string {
	return dara.Prettify(s)
}

func (s CompletePhysicalConnectionLOAResponse) GoString() string {
	return s.String()
}

func (s *CompletePhysicalConnectionLOAResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CompletePhysicalConnectionLOAResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CompletePhysicalConnectionLOAResponse) GetBody() *CompletePhysicalConnectionLOAResponseBody {
	return s.Body
}

func (s *CompletePhysicalConnectionLOAResponse) SetHeaders(v map[string]*string) *CompletePhysicalConnectionLOAResponse {
	s.Headers = v
	return s
}

func (s *CompletePhysicalConnectionLOAResponse) SetStatusCode(v int32) *CompletePhysicalConnectionLOAResponse {
	s.StatusCode = &v
	return s
}

func (s *CompletePhysicalConnectionLOAResponse) SetBody(v *CompletePhysicalConnectionLOAResponseBody) *CompletePhysicalConnectionLOAResponse {
	s.Body = v
	return s
}

func (s *CompletePhysicalConnectionLOAResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
