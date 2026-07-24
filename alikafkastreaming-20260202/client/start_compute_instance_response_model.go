// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartComputeInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartComputeInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartComputeInstanceResponse
	GetStatusCode() *int32
	SetBody(v *StartComputeInstanceResponseBody) *StartComputeInstanceResponse
	GetBody() *StartComputeInstanceResponseBody
}

type StartComputeInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartComputeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartComputeInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s StartComputeInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartComputeInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartComputeInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartComputeInstanceResponse) GetBody() *StartComputeInstanceResponseBody {
	return s.Body
}

func (s *StartComputeInstanceResponse) SetHeaders(v map[string]*string) *StartComputeInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartComputeInstanceResponse) SetStatusCode(v int32) *StartComputeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartComputeInstanceResponse) SetBody(v *StartComputeInstanceResponseBody) *StartComputeInstanceResponse {
	s.Body = v
	return s
}

func (s *StartComputeInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
