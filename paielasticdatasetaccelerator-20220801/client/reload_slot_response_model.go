// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReloadSlotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReloadSlotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReloadSlotResponse
	GetStatusCode() *int32
	SetBody(v *ReloadSlotResponseBody) *ReloadSlotResponse
	GetBody() *ReloadSlotResponseBody
}

type ReloadSlotResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReloadSlotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReloadSlotResponse) String() string {
	return dara.Prettify(s)
}

func (s ReloadSlotResponse) GoString() string {
	return s.String()
}

func (s *ReloadSlotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReloadSlotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReloadSlotResponse) GetBody() *ReloadSlotResponseBody {
	return s.Body
}

func (s *ReloadSlotResponse) SetHeaders(v map[string]*string) *ReloadSlotResponse {
	s.Headers = v
	return s
}

func (s *ReloadSlotResponse) SetStatusCode(v int32) *ReloadSlotResponse {
	s.StatusCode = &v
	return s
}

func (s *ReloadSlotResponse) SetBody(v *ReloadSlotResponseBody) *ReloadSlotResponse {
	s.Body = v
	return s
}

func (s *ReloadSlotResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
