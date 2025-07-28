// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApisecApiResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyApisecApiResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyApisecApiResourceResponse
	GetStatusCode() *int32
	SetBody(v *ModifyApisecApiResourceResponseBody) *ModifyApisecApiResourceResponse
	GetBody() *ModifyApisecApiResourceResponseBody
}

type ModifyApisecApiResourceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyApisecApiResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyApisecApiResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyApisecApiResourceResponse) GoString() string {
	return s.String()
}

func (s *ModifyApisecApiResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyApisecApiResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyApisecApiResourceResponse) GetBody() *ModifyApisecApiResourceResponseBody {
	return s.Body
}

func (s *ModifyApisecApiResourceResponse) SetHeaders(v map[string]*string) *ModifyApisecApiResourceResponse {
	s.Headers = v
	return s
}

func (s *ModifyApisecApiResourceResponse) SetStatusCode(v int32) *ModifyApisecApiResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyApisecApiResourceResponse) SetBody(v *ModifyApisecApiResourceResponseBody) *ModifyApisecApiResourceResponse {
	s.Body = v
	return s
}

func (s *ModifyApisecApiResourceResponse) Validate() error {
	return dara.Validate(s)
}
