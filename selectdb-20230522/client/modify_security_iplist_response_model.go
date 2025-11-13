// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityIPListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySecurityIPListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySecurityIPListResponse
	GetStatusCode() *int32
	SetBody(v *ModifySecurityIPListResponseBody) *ModifySecurityIPListResponse
	GetBody() *ModifySecurityIPListResponseBody
}

type ModifySecurityIPListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySecurityIPListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySecurityIPListResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityIPListResponse) GoString() string {
	return s.String()
}

func (s *ModifySecurityIPListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySecurityIPListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySecurityIPListResponse) GetBody() *ModifySecurityIPListResponseBody {
	return s.Body
}

func (s *ModifySecurityIPListResponse) SetHeaders(v map[string]*string) *ModifySecurityIPListResponse {
	s.Headers = v
	return s
}

func (s *ModifySecurityIPListResponse) SetStatusCode(v int32) *ModifySecurityIPListResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySecurityIPListResponse) SetBody(v *ModifySecurityIPListResponseBody) *ModifySecurityIPListResponse {
	s.Body = v
	return s
}

func (s *ModifySecurityIPListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
