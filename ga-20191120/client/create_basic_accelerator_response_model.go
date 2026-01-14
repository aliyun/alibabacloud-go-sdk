// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBasicAcceleratorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBasicAcceleratorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBasicAcceleratorResponse
	GetStatusCode() *int32
	SetBody(v *CreateBasicAcceleratorResponseBody) *CreateBasicAcceleratorResponse
	GetBody() *CreateBasicAcceleratorResponseBody
}

type CreateBasicAcceleratorResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBasicAcceleratorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBasicAcceleratorResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBasicAcceleratorResponse) GoString() string {
	return s.String()
}

func (s *CreateBasicAcceleratorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBasicAcceleratorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBasicAcceleratorResponse) GetBody() *CreateBasicAcceleratorResponseBody {
	return s.Body
}

func (s *CreateBasicAcceleratorResponse) SetHeaders(v map[string]*string) *CreateBasicAcceleratorResponse {
	s.Headers = v
	return s
}

func (s *CreateBasicAcceleratorResponse) SetStatusCode(v int32) *CreateBasicAcceleratorResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBasicAcceleratorResponse) SetBody(v *CreateBasicAcceleratorResponseBody) *CreateBasicAcceleratorResponse {
	s.Body = v
	return s
}

func (s *CreateBasicAcceleratorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
