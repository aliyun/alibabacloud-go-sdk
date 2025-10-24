// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMajorProtectionBlackIpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMajorProtectionBlackIpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMajorProtectionBlackIpResponse
	GetStatusCode() *int32
	SetBody(v *CreateMajorProtectionBlackIpResponseBody) *CreateMajorProtectionBlackIpResponse
	GetBody() *CreateMajorProtectionBlackIpResponseBody
}

type CreateMajorProtectionBlackIpResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMajorProtectionBlackIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMajorProtectionBlackIpResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMajorProtectionBlackIpResponse) GoString() string {
	return s.String()
}

func (s *CreateMajorProtectionBlackIpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMajorProtectionBlackIpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMajorProtectionBlackIpResponse) GetBody() *CreateMajorProtectionBlackIpResponseBody {
	return s.Body
}

func (s *CreateMajorProtectionBlackIpResponse) SetHeaders(v map[string]*string) *CreateMajorProtectionBlackIpResponse {
	s.Headers = v
	return s
}

func (s *CreateMajorProtectionBlackIpResponse) SetStatusCode(v int32) *CreateMajorProtectionBlackIpResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMajorProtectionBlackIpResponse) SetBody(v *CreateMajorProtectionBlackIpResponseBody) *CreateMajorProtectionBlackIpResponse {
	s.Body = v
	return s
}

func (s *CreateMajorProtectionBlackIpResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
