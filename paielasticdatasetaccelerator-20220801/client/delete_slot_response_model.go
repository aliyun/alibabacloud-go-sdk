// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSlotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSlotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSlotResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSlotResponseBody) *DeleteSlotResponse
	GetBody() *DeleteSlotResponseBody
}

type DeleteSlotResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSlotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSlotResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSlotResponse) GoString() string {
	return s.String()
}

func (s *DeleteSlotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSlotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSlotResponse) GetBody() *DeleteSlotResponseBody {
	return s.Body
}

func (s *DeleteSlotResponse) SetHeaders(v map[string]*string) *DeleteSlotResponse {
	s.Headers = v
	return s
}

func (s *DeleteSlotResponse) SetStatusCode(v int32) *DeleteSlotResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSlotResponse) SetBody(v *DeleteSlotResponseBody) *DeleteSlotResponse {
	s.Body = v
	return s
}

func (s *DeleteSlotResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
