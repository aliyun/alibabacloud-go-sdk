// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterInternalAccountForBucResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RegisterInternalAccountForBucResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RegisterInternalAccountForBucResponse
	GetStatusCode() *int32
	SetBody(v *RegisterInternalAccountForBucResponseBody) *RegisterInternalAccountForBucResponse
	GetBody() *RegisterInternalAccountForBucResponseBody
}

type RegisterInternalAccountForBucResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterInternalAccountForBucResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterInternalAccountForBucResponse) String() string {
	return dara.Prettify(s)
}

func (s RegisterInternalAccountForBucResponse) GoString() string {
	return s.String()
}

func (s *RegisterInternalAccountForBucResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RegisterInternalAccountForBucResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RegisterInternalAccountForBucResponse) GetBody() *RegisterInternalAccountForBucResponseBody {
	return s.Body
}

func (s *RegisterInternalAccountForBucResponse) SetHeaders(v map[string]*string) *RegisterInternalAccountForBucResponse {
	s.Headers = v
	return s
}

func (s *RegisterInternalAccountForBucResponse) SetStatusCode(v int32) *RegisterInternalAccountForBucResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterInternalAccountForBucResponse) SetBody(v *RegisterInternalAccountForBucResponseBody) *RegisterInternalAccountForBucResponse {
	s.Body = v
	return s
}

func (s *RegisterInternalAccountForBucResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
