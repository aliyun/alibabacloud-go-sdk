// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHostAliasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyHostAliasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyHostAliasResponse
	GetStatusCode() *int32
	SetBody(v *ModifyHostAliasResponseBody) *ModifyHostAliasResponse
	GetBody() *ModifyHostAliasResponseBody
}

type ModifyHostAliasResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHostAliasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyHostAliasResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyHostAliasResponse) GoString() string {
	return s.String()
}

func (s *ModifyHostAliasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyHostAliasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyHostAliasResponse) GetBody() *ModifyHostAliasResponseBody {
	return s.Body
}

func (s *ModifyHostAliasResponse) SetHeaders(v map[string]*string) *ModifyHostAliasResponse {
	s.Headers = v
	return s
}

func (s *ModifyHostAliasResponse) SetStatusCode(v int32) *ModifyHostAliasResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHostAliasResponse) SetBody(v *ModifyHostAliasResponseBody) *ModifyHostAliasResponse {
	s.Body = v
	return s
}

func (s *ModifyHostAliasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
