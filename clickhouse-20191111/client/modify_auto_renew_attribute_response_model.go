// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAutoRenewAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAutoRenewAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAutoRenewAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAutoRenewAttributeResponseBody) *ModifyAutoRenewAttributeResponse
	GetBody() *ModifyAutoRenewAttributeResponseBody
}

type ModifyAutoRenewAttributeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAutoRenewAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAutoRenewAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoRenewAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyAutoRenewAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAutoRenewAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAutoRenewAttributeResponse) GetBody() *ModifyAutoRenewAttributeResponseBody {
	return s.Body
}

func (s *ModifyAutoRenewAttributeResponse) SetHeaders(v map[string]*string) *ModifyAutoRenewAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyAutoRenewAttributeResponse) SetStatusCode(v int32) *ModifyAutoRenewAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAutoRenewAttributeResponse) SetBody(v *ModifyAutoRenewAttributeResponseBody) *ModifyAutoRenewAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyAutoRenewAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
