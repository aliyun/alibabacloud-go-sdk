// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckControlPlaneLogEnableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckControlPlaneLogEnableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckControlPlaneLogEnableResponse
	GetStatusCode() *int32
	SetBody(v *CheckControlPlaneLogEnableResponseBody) *CheckControlPlaneLogEnableResponse
	GetBody() *CheckControlPlaneLogEnableResponseBody
}

type CheckControlPlaneLogEnableResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckControlPlaneLogEnableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckControlPlaneLogEnableResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckControlPlaneLogEnableResponse) GoString() string {
	return s.String()
}

func (s *CheckControlPlaneLogEnableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckControlPlaneLogEnableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckControlPlaneLogEnableResponse) GetBody() *CheckControlPlaneLogEnableResponseBody {
	return s.Body
}

func (s *CheckControlPlaneLogEnableResponse) SetHeaders(v map[string]*string) *CheckControlPlaneLogEnableResponse {
	s.Headers = v
	return s
}

func (s *CheckControlPlaneLogEnableResponse) SetStatusCode(v int32) *CheckControlPlaneLogEnableResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckControlPlaneLogEnableResponse) SetBody(v *CheckControlPlaneLogEnableResponseBody) *CheckControlPlaneLogEnableResponse {
	s.Body = v
	return s
}

func (s *CheckControlPlaneLogEnableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
