// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAliasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAliasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAliasResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAliasResponseBody) *ModifyAliasResponse
	GetBody() *ModifyAliasResponseBody
}

type ModifyAliasResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAliasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAliasResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAliasResponse) GoString() string {
	return s.String()
}

func (s *ModifyAliasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAliasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAliasResponse) GetBody() *ModifyAliasResponseBody {
	return s.Body
}

func (s *ModifyAliasResponse) SetHeaders(v map[string]*string) *ModifyAliasResponse {
	s.Headers = v
	return s
}

func (s *ModifyAliasResponse) SetStatusCode(v int32) *ModifyAliasResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAliasResponse) SetBody(v *ModifyAliasResponseBody) *ModifyAliasResponse {
	s.Body = v
	return s
}

func (s *ModifyAliasResponse) Validate() error {
	return dara.Validate(s)
}
