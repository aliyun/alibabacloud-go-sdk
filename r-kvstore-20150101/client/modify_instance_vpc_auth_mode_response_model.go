// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceVpcAuthModeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceVpcAuthModeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceVpcAuthModeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceVpcAuthModeResponseBody) *ModifyInstanceVpcAuthModeResponse
	GetBody() *ModifyInstanceVpcAuthModeResponseBody
}

type ModifyInstanceVpcAuthModeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceVpcAuthModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceVpcAuthModeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceVpcAuthModeResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceVpcAuthModeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceVpcAuthModeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceVpcAuthModeResponse) GetBody() *ModifyInstanceVpcAuthModeResponseBody {
	return s.Body
}

func (s *ModifyInstanceVpcAuthModeResponse) SetHeaders(v map[string]*string) *ModifyInstanceVpcAuthModeResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceVpcAuthModeResponse) SetStatusCode(v int32) *ModifyInstanceVpcAuthModeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceVpcAuthModeResponse) SetBody(v *ModifyInstanceVpcAuthModeResponseBody) *ModifyInstanceVpcAuthModeResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceVpcAuthModeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
