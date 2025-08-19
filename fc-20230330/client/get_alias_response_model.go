// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAliasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAliasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAliasResponse
	GetStatusCode() *int32
	SetBody(v *Alias) *GetAliasResponse
	GetBody() *Alias
}

type GetAliasResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Alias             `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAliasResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAliasResponse) GoString() string {
	return s.String()
}

func (s *GetAliasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAliasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAliasResponse) GetBody() *Alias {
	return s.Body
}

func (s *GetAliasResponse) SetHeaders(v map[string]*string) *GetAliasResponse {
	s.Headers = v
	return s
}

func (s *GetAliasResponse) SetStatusCode(v int32) *GetAliasResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAliasResponse) SetBody(v *Alias) *GetAliasResponse {
	s.Body = v
	return s
}

func (s *GetAliasResponse) Validate() error {
	return dara.Validate(s)
}
