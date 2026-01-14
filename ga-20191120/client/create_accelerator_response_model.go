// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAcceleratorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAcceleratorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAcceleratorResponse
	GetStatusCode() *int32
	SetBody(v *CreateAcceleratorResponseBody) *CreateAcceleratorResponse
	GetBody() *CreateAcceleratorResponseBody
}

type CreateAcceleratorResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAcceleratorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAcceleratorResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAcceleratorResponse) GoString() string {
	return s.String()
}

func (s *CreateAcceleratorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAcceleratorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAcceleratorResponse) GetBody() *CreateAcceleratorResponseBody {
	return s.Body
}

func (s *CreateAcceleratorResponse) SetHeaders(v map[string]*string) *CreateAcceleratorResponse {
	s.Headers = v
	return s
}

func (s *CreateAcceleratorResponse) SetStatusCode(v int32) *CreateAcceleratorResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAcceleratorResponse) SetBody(v *CreateAcceleratorResponseBody) *CreateAcceleratorResponse {
	s.Body = v
	return s
}

func (s *CreateAcceleratorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
