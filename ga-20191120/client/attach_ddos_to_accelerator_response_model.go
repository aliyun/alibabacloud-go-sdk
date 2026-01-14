// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachDdosToAcceleratorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachDdosToAcceleratorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachDdosToAcceleratorResponse
	GetStatusCode() *int32
	SetBody(v *AttachDdosToAcceleratorResponseBody) *AttachDdosToAcceleratorResponse
	GetBody() *AttachDdosToAcceleratorResponseBody
}

type AttachDdosToAcceleratorResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachDdosToAcceleratorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachDdosToAcceleratorResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachDdosToAcceleratorResponse) GoString() string {
	return s.String()
}

func (s *AttachDdosToAcceleratorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachDdosToAcceleratorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachDdosToAcceleratorResponse) GetBody() *AttachDdosToAcceleratorResponseBody {
	return s.Body
}

func (s *AttachDdosToAcceleratorResponse) SetHeaders(v map[string]*string) *AttachDdosToAcceleratorResponse {
	s.Headers = v
	return s
}

func (s *AttachDdosToAcceleratorResponse) SetStatusCode(v int32) *AttachDdosToAcceleratorResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachDdosToAcceleratorResponse) SetBody(v *AttachDdosToAcceleratorResponseBody) *AttachDdosToAcceleratorResponse {
	s.Body = v
	return s
}

func (s *AttachDdosToAcceleratorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
