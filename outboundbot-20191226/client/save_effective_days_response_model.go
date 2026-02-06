// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveEffectiveDaysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveEffectiveDaysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveEffectiveDaysResponse
	GetStatusCode() *int32
	SetBody(v *SaveEffectiveDaysResponseBody) *SaveEffectiveDaysResponse
	GetBody() *SaveEffectiveDaysResponseBody
}

type SaveEffectiveDaysResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveEffectiveDaysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveEffectiveDaysResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveEffectiveDaysResponse) GoString() string {
	return s.String()
}

func (s *SaveEffectiveDaysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveEffectiveDaysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveEffectiveDaysResponse) GetBody() *SaveEffectiveDaysResponseBody {
	return s.Body
}

func (s *SaveEffectiveDaysResponse) SetHeaders(v map[string]*string) *SaveEffectiveDaysResponse {
	s.Headers = v
	return s
}

func (s *SaveEffectiveDaysResponse) SetStatusCode(v int32) *SaveEffectiveDaysResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveEffectiveDaysResponse) SetBody(v *SaveEffectiveDaysResponseBody) *SaveEffectiveDaysResponse {
	s.Body = v
	return s
}

func (s *SaveEffectiveDaysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
