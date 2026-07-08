// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVirtualThreeElementsVerificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VirtualThreeElementsVerificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VirtualThreeElementsVerificationResponse
	GetStatusCode() *int32
	SetBody(v *VirtualThreeElementsVerificationResponseBody) *VirtualThreeElementsVerificationResponse
	GetBody() *VirtualThreeElementsVerificationResponseBody
}

type VirtualThreeElementsVerificationResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VirtualThreeElementsVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VirtualThreeElementsVerificationResponse) String() string {
	return dara.Prettify(s)
}

func (s VirtualThreeElementsVerificationResponse) GoString() string {
	return s.String()
}

func (s *VirtualThreeElementsVerificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VirtualThreeElementsVerificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VirtualThreeElementsVerificationResponse) GetBody() *VirtualThreeElementsVerificationResponseBody {
	return s.Body
}

func (s *VirtualThreeElementsVerificationResponse) SetHeaders(v map[string]*string) *VirtualThreeElementsVerificationResponse {
	s.Headers = v
	return s
}

func (s *VirtualThreeElementsVerificationResponse) SetStatusCode(v int32) *VirtualThreeElementsVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *VirtualThreeElementsVerificationResponse) SetBody(v *VirtualThreeElementsVerificationResponseBody) *VirtualThreeElementsVerificationResponse {
	s.Body = v
	return s
}

func (s *VirtualThreeElementsVerificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
