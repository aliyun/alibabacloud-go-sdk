// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachDdosFromAcceleratorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachDdosFromAcceleratorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachDdosFromAcceleratorResponse
	GetStatusCode() *int32
	SetBody(v *DetachDdosFromAcceleratorResponseBody) *DetachDdosFromAcceleratorResponse
	GetBody() *DetachDdosFromAcceleratorResponseBody
}

type DetachDdosFromAcceleratorResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachDdosFromAcceleratorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachDdosFromAcceleratorResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachDdosFromAcceleratorResponse) GoString() string {
	return s.String()
}

func (s *DetachDdosFromAcceleratorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachDdosFromAcceleratorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachDdosFromAcceleratorResponse) GetBody() *DetachDdosFromAcceleratorResponseBody {
	return s.Body
}

func (s *DetachDdosFromAcceleratorResponse) SetHeaders(v map[string]*string) *DetachDdosFromAcceleratorResponse {
	s.Headers = v
	return s
}

func (s *DetachDdosFromAcceleratorResponse) SetStatusCode(v int32) *DetachDdosFromAcceleratorResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachDdosFromAcceleratorResponse) SetBody(v *DetachDdosFromAcceleratorResponseBody) *DetachDdosFromAcceleratorResponse {
	s.Body = v
	return s
}

func (s *DetachDdosFromAcceleratorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
