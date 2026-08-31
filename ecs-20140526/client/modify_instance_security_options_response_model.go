// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceSecurityOptionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceSecurityOptionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceSecurityOptionsResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceSecurityOptionsResponseBody) *ModifyInstanceSecurityOptionsResponse
	GetBody() *ModifyInstanceSecurityOptionsResponseBody
}

type ModifyInstanceSecurityOptionsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceSecurityOptionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceSecurityOptionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceSecurityOptionsResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceSecurityOptionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceSecurityOptionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceSecurityOptionsResponse) GetBody() *ModifyInstanceSecurityOptionsResponseBody {
	return s.Body
}

func (s *ModifyInstanceSecurityOptionsResponse) SetHeaders(v map[string]*string) *ModifyInstanceSecurityOptionsResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceSecurityOptionsResponse) SetStatusCode(v int32) *ModifyInstanceSecurityOptionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceSecurityOptionsResponse) SetBody(v *ModifyInstanceSecurityOptionsResponseBody) *ModifyInstanceSecurityOptionsResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceSecurityOptionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
