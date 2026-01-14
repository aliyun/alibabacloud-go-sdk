// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBasicAcceleratorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBasicAcceleratorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBasicAcceleratorResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBasicAcceleratorResponseBody) *DeleteBasicAcceleratorResponse
	GetBody() *DeleteBasicAcceleratorResponseBody
}

type DeleteBasicAcceleratorResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBasicAcceleratorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBasicAcceleratorResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBasicAcceleratorResponse) GoString() string {
	return s.String()
}

func (s *DeleteBasicAcceleratorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBasicAcceleratorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBasicAcceleratorResponse) GetBody() *DeleteBasicAcceleratorResponseBody {
	return s.Body
}

func (s *DeleteBasicAcceleratorResponse) SetHeaders(v map[string]*string) *DeleteBasicAcceleratorResponse {
	s.Headers = v
	return s
}

func (s *DeleteBasicAcceleratorResponse) SetStatusCode(v int32) *DeleteBasicAcceleratorResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBasicAcceleratorResponse) SetBody(v *DeleteBasicAcceleratorResponseBody) *DeleteBasicAcceleratorResponse {
	s.Body = v
	return s
}

func (s *DeleteBasicAcceleratorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
