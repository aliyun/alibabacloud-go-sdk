// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityIPGroupRelationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySecurityIPGroupRelationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySecurityIPGroupRelationResponse
	GetStatusCode() *int32
	SetBody(v *ModifySecurityIPGroupRelationResponseBody) *ModifySecurityIPGroupRelationResponse
	GetBody() *ModifySecurityIPGroupRelationResponseBody
}

type ModifySecurityIPGroupRelationResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySecurityIPGroupRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySecurityIPGroupRelationResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityIPGroupRelationResponse) GoString() string {
	return s.String()
}

func (s *ModifySecurityIPGroupRelationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySecurityIPGroupRelationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySecurityIPGroupRelationResponse) GetBody() *ModifySecurityIPGroupRelationResponseBody {
	return s.Body
}

func (s *ModifySecurityIPGroupRelationResponse) SetHeaders(v map[string]*string) *ModifySecurityIPGroupRelationResponse {
	s.Headers = v
	return s
}

func (s *ModifySecurityIPGroupRelationResponse) SetStatusCode(v int32) *ModifySecurityIPGroupRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySecurityIPGroupRelationResponse) SetBody(v *ModifySecurityIPGroupRelationResponseBody) *ModifySecurityIPGroupRelationResponse {
	s.Body = v
	return s
}

func (s *ModifySecurityIPGroupRelationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
