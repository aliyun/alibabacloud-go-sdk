// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceSSLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceSSLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceSSLResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceSSLResponseBody) *ModifyInstanceSSLResponse
	GetBody() *ModifyInstanceSSLResponseBody
}

type ModifyInstanceSSLResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceSSLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceSSLResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceSSLResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceSSLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceSSLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceSSLResponse) GetBody() *ModifyInstanceSSLResponseBody {
	return s.Body
}

func (s *ModifyInstanceSSLResponse) SetHeaders(v map[string]*string) *ModifyInstanceSSLResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceSSLResponse) SetStatusCode(v int32) *ModifyInstanceSSLResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceSSLResponse) SetBody(v *ModifyInstanceSSLResponseBody) *ModifyInstanceSSLResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceSSLResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
