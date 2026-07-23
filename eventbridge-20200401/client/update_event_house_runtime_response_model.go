// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEventHouseRuntimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEventHouseRuntimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEventHouseRuntimeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEventHouseRuntimeResponseBody) *UpdateEventHouseRuntimeResponse
	GetBody() *UpdateEventHouseRuntimeResponseBody
}

type UpdateEventHouseRuntimeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEventHouseRuntimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEventHouseRuntimeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventHouseRuntimeResponse) GoString() string {
	return s.String()
}

func (s *UpdateEventHouseRuntimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEventHouseRuntimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEventHouseRuntimeResponse) GetBody() *UpdateEventHouseRuntimeResponseBody {
	return s.Body
}

func (s *UpdateEventHouseRuntimeResponse) SetHeaders(v map[string]*string) *UpdateEventHouseRuntimeResponse {
	s.Headers = v
	return s
}

func (s *UpdateEventHouseRuntimeResponse) SetStatusCode(v int32) *UpdateEventHouseRuntimeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEventHouseRuntimeResponse) SetBody(v *UpdateEventHouseRuntimeResponseBody) *UpdateEventHouseRuntimeResponse {
	s.Body = v
	return s
}

func (s *UpdateEventHouseRuntimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
