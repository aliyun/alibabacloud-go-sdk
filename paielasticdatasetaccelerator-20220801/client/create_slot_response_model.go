// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSlotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSlotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSlotResponse
	GetStatusCode() *int32
	SetBody(v *CreateSlotResponseBody) *CreateSlotResponse
	GetBody() *CreateSlotResponseBody
}

type CreateSlotResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSlotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSlotResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSlotResponse) GoString() string {
	return s.String()
}

func (s *CreateSlotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSlotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSlotResponse) GetBody() *CreateSlotResponseBody {
	return s.Body
}

func (s *CreateSlotResponse) SetHeaders(v map[string]*string) *CreateSlotResponse {
	s.Headers = v
	return s
}

func (s *CreateSlotResponse) SetStatusCode(v int32) *CreateSlotResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSlotResponse) SetBody(v *CreateSlotResponseBody) *CreateSlotResponse {
	s.Body = v
	return s
}

func (s *CreateSlotResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
