// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAcceleratorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAcceleratorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAcceleratorResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAcceleratorResponseBody) *DeleteAcceleratorResponse
	GetBody() *DeleteAcceleratorResponseBody
}

type DeleteAcceleratorResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAcceleratorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAcceleratorResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAcceleratorResponse) GoString() string {
	return s.String()
}

func (s *DeleteAcceleratorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAcceleratorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAcceleratorResponse) GetBody() *DeleteAcceleratorResponseBody {
	return s.Body
}

func (s *DeleteAcceleratorResponse) SetHeaders(v map[string]*string) *DeleteAcceleratorResponse {
	s.Headers = v
	return s
}

func (s *DeleteAcceleratorResponse) SetStatusCode(v int32) *DeleteAcceleratorResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAcceleratorResponse) SetBody(v *DeleteAcceleratorResponseBody) *DeleteAcceleratorResponse {
	s.Body = v
	return s
}

func (s *DeleteAcceleratorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
