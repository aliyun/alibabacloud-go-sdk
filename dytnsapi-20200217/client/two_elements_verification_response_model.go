// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTwoElementsVerificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TwoElementsVerificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TwoElementsVerificationResponse
	GetStatusCode() *int32
	SetBody(v *TwoElementsVerificationResponseBody) *TwoElementsVerificationResponse
	GetBody() *TwoElementsVerificationResponseBody
}

type TwoElementsVerificationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TwoElementsVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TwoElementsVerificationResponse) String() string {
	return dara.Prettify(s)
}

func (s TwoElementsVerificationResponse) GoString() string {
	return s.String()
}

func (s *TwoElementsVerificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TwoElementsVerificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TwoElementsVerificationResponse) GetBody() *TwoElementsVerificationResponseBody {
	return s.Body
}

func (s *TwoElementsVerificationResponse) SetHeaders(v map[string]*string) *TwoElementsVerificationResponse {
	s.Headers = v
	return s
}

func (s *TwoElementsVerificationResponse) SetStatusCode(v int32) *TwoElementsVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *TwoElementsVerificationResponse) SetBody(v *TwoElementsVerificationResponseBody) *TwoElementsVerificationResponse {
	s.Body = v
	return s
}

func (s *TwoElementsVerificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
