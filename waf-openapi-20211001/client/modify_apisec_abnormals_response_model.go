// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApisecAbnormalsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyApisecAbnormalsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyApisecAbnormalsResponse
	GetStatusCode() *int32
	SetBody(v *ModifyApisecAbnormalsResponseBody) *ModifyApisecAbnormalsResponse
	GetBody() *ModifyApisecAbnormalsResponseBody
}

type ModifyApisecAbnormalsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyApisecAbnormalsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyApisecAbnormalsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyApisecAbnormalsResponse) GoString() string {
	return s.String()
}

func (s *ModifyApisecAbnormalsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyApisecAbnormalsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyApisecAbnormalsResponse) GetBody() *ModifyApisecAbnormalsResponseBody {
	return s.Body
}

func (s *ModifyApisecAbnormalsResponse) SetHeaders(v map[string]*string) *ModifyApisecAbnormalsResponse {
	s.Headers = v
	return s
}

func (s *ModifyApisecAbnormalsResponse) SetStatusCode(v int32) *ModifyApisecAbnormalsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyApisecAbnormalsResponse) SetBody(v *ModifyApisecAbnormalsResponseBody) *ModifyApisecAbnormalsResponse {
	s.Body = v
	return s
}

func (s *ModifyApisecAbnormalsResponse) Validate() error {
	return dara.Validate(s)
}
