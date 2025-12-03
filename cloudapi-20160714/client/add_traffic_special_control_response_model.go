// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTrafficSpecialControlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddTrafficSpecialControlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddTrafficSpecialControlResponse
	GetStatusCode() *int32
	SetBody(v *AddTrafficSpecialControlResponseBody) *AddTrafficSpecialControlResponse
	GetBody() *AddTrafficSpecialControlResponseBody
}

type AddTrafficSpecialControlResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddTrafficSpecialControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddTrafficSpecialControlResponse) String() string {
	return dara.Prettify(s)
}

func (s AddTrafficSpecialControlResponse) GoString() string {
	return s.String()
}

func (s *AddTrafficSpecialControlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddTrafficSpecialControlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddTrafficSpecialControlResponse) GetBody() *AddTrafficSpecialControlResponseBody {
	return s.Body
}

func (s *AddTrafficSpecialControlResponse) SetHeaders(v map[string]*string) *AddTrafficSpecialControlResponse {
	s.Headers = v
	return s
}

func (s *AddTrafficSpecialControlResponse) SetStatusCode(v int32) *AddTrafficSpecialControlResponse {
	s.StatusCode = &v
	return s
}

func (s *AddTrafficSpecialControlResponse) SetBody(v *AddTrafficSpecialControlResponseBody) *AddTrafficSpecialControlResponse {
	s.Body = v
	return s
}

func (s *AddTrafficSpecialControlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
