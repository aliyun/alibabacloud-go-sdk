// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySensitiveSwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySensitiveSwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySensitiveSwitchResponse
	GetStatusCode() *int32
	SetBody(v *ModifySensitiveSwitchResponseBody) *ModifySensitiveSwitchResponse
	GetBody() *ModifySensitiveSwitchResponseBody
}

type ModifySensitiveSwitchResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySensitiveSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySensitiveSwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySensitiveSwitchResponse) GoString() string {
	return s.String()
}

func (s *ModifySensitiveSwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySensitiveSwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySensitiveSwitchResponse) GetBody() *ModifySensitiveSwitchResponseBody {
	return s.Body
}

func (s *ModifySensitiveSwitchResponse) SetHeaders(v map[string]*string) *ModifySensitiveSwitchResponse {
	s.Headers = v
	return s
}

func (s *ModifySensitiveSwitchResponse) SetStatusCode(v int32) *ModifySensitiveSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySensitiveSwitchResponse) SetBody(v *ModifySensitiveSwitchResponseBody) *ModifySensitiveSwitchResponse {
	s.Body = v
	return s
}

func (s *ModifySensitiveSwitchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
