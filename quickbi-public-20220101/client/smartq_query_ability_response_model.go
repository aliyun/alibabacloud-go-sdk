// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSmartqQueryAbilityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SmartqQueryAbilityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SmartqQueryAbilityResponse
	GetStatusCode() *int32
	SetBody(v *SmartqQueryAbilityResponseBody) *SmartqQueryAbilityResponse
	GetBody() *SmartqQueryAbilityResponseBody
}

type SmartqQueryAbilityResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SmartqQueryAbilityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SmartqQueryAbilityResponse) String() string {
	return dara.Prettify(s)
}

func (s SmartqQueryAbilityResponse) GoString() string {
	return s.String()
}

func (s *SmartqQueryAbilityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SmartqQueryAbilityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SmartqQueryAbilityResponse) GetBody() *SmartqQueryAbilityResponseBody {
	return s.Body
}

func (s *SmartqQueryAbilityResponse) SetHeaders(v map[string]*string) *SmartqQueryAbilityResponse {
	s.Headers = v
	return s
}

func (s *SmartqQueryAbilityResponse) SetStatusCode(v int32) *SmartqQueryAbilityResponse {
	s.StatusCode = &v
	return s
}

func (s *SmartqQueryAbilityResponse) SetBody(v *SmartqQueryAbilityResponseBody) *SmartqQueryAbilityResponse {
	s.Body = v
	return s
}

func (s *SmartqQueryAbilityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
