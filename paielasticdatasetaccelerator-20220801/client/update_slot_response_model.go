// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSlotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSlotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSlotResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSlotResponseBody) *UpdateSlotResponse
	GetBody() *UpdateSlotResponseBody
}

type UpdateSlotResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSlotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSlotResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSlotResponse) GoString() string {
	return s.String()
}

func (s *UpdateSlotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSlotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSlotResponse) GetBody() *UpdateSlotResponseBody {
	return s.Body
}

func (s *UpdateSlotResponse) SetHeaders(v map[string]*string) *UpdateSlotResponse {
	s.Headers = v
	return s
}

func (s *UpdateSlotResponse) SetStatusCode(v int32) *UpdateSlotResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSlotResponse) SetBody(v *UpdateSlotResponseBody) *UpdateSlotResponse {
	s.Body = v
	return s
}

func (s *UpdateSlotResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
