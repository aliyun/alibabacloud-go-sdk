// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTrafficControlTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTrafficControlTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTrafficControlTargetResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTrafficControlTargetResponseBody) *UpdateTrafficControlTargetResponse
	GetBody() *UpdateTrafficControlTargetResponseBody
}

type UpdateTrafficControlTargetResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTrafficControlTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTrafficControlTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTrafficControlTargetResponse) GoString() string {
	return s.String()
}

func (s *UpdateTrafficControlTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTrafficControlTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTrafficControlTargetResponse) GetBody() *UpdateTrafficControlTargetResponseBody {
	return s.Body
}

func (s *UpdateTrafficControlTargetResponse) SetHeaders(v map[string]*string) *UpdateTrafficControlTargetResponse {
	s.Headers = v
	return s
}

func (s *UpdateTrafficControlTargetResponse) SetStatusCode(v int32) *UpdateTrafficControlTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTrafficControlTargetResponse) SetBody(v *UpdateTrafficControlTargetResponseBody) *UpdateTrafficControlTargetResponse {
	s.Body = v
	return s
}

func (s *UpdateTrafficControlTargetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
