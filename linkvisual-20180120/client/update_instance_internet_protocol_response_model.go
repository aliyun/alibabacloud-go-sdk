// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceInternetProtocolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateInstanceInternetProtocolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateInstanceInternetProtocolResponse
	GetStatusCode() *int32
	SetBody(v *UpdateInstanceInternetProtocolResponseBody) *UpdateInstanceInternetProtocolResponse
	GetBody() *UpdateInstanceInternetProtocolResponseBody
}

type UpdateInstanceInternetProtocolResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceInternetProtocolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceInternetProtocolResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceInternetProtocolResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInternetProtocolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateInstanceInternetProtocolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateInstanceInternetProtocolResponse) GetBody() *UpdateInstanceInternetProtocolResponseBody {
	return s.Body
}

func (s *UpdateInstanceInternetProtocolResponse) SetHeaders(v map[string]*string) *UpdateInstanceInternetProtocolResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceInternetProtocolResponse) SetStatusCode(v int32) *UpdateInstanceInternetProtocolResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceInternetProtocolResponse) SetBody(v *UpdateInstanceInternetProtocolResponseBody) *UpdateInstanceInternetProtocolResponse {
	s.Body = v
	return s
}

func (s *UpdateInstanceInternetProtocolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
