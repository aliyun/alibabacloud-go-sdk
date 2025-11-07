// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobileDetectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MobileDetectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MobileDetectResponse
	GetStatusCode() *int32
	SetBody(v *MobileDetectResponseBody) *MobileDetectResponse
	GetBody() *MobileDetectResponseBody
}

type MobileDetectResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MobileDetectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MobileDetectResponse) String() string {
	return dara.Prettify(s)
}

func (s MobileDetectResponse) GoString() string {
	return s.String()
}

func (s *MobileDetectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MobileDetectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MobileDetectResponse) GetBody() *MobileDetectResponseBody {
	return s.Body
}

func (s *MobileDetectResponse) SetHeaders(v map[string]*string) *MobileDetectResponse {
	s.Headers = v
	return s
}

func (s *MobileDetectResponse) SetStatusCode(v int32) *MobileDetectResponse {
	s.StatusCode = &v
	return s
}

func (s *MobileDetectResponse) SetBody(v *MobileDetectResponseBody) *MobileDetectResponse {
	s.Body = v
	return s
}

func (s *MobileDetectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
