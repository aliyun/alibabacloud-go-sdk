// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iThreeElementsVerificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ThreeElementsVerificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ThreeElementsVerificationResponse
	GetStatusCode() *int32
	SetBody(v *ThreeElementsVerificationResponseBody) *ThreeElementsVerificationResponse
	GetBody() *ThreeElementsVerificationResponseBody
}

type ThreeElementsVerificationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ThreeElementsVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ThreeElementsVerificationResponse) String() string {
	return dara.Prettify(s)
}

func (s ThreeElementsVerificationResponse) GoString() string {
	return s.String()
}

func (s *ThreeElementsVerificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ThreeElementsVerificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ThreeElementsVerificationResponse) GetBody() *ThreeElementsVerificationResponseBody {
	return s.Body
}

func (s *ThreeElementsVerificationResponse) SetHeaders(v map[string]*string) *ThreeElementsVerificationResponse {
	s.Headers = v
	return s
}

func (s *ThreeElementsVerificationResponse) SetStatusCode(v int32) *ThreeElementsVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *ThreeElementsVerificationResponse) SetBody(v *ThreeElementsVerificationResponseBody) *ThreeElementsVerificationResponse {
	s.Body = v
	return s
}

func (s *ThreeElementsVerificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
