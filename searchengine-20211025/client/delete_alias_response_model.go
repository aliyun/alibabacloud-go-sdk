// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAliasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAliasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAliasResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAliasResponseBody) *DeleteAliasResponse
	GetBody() *DeleteAliasResponseBody
}

type DeleteAliasResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAliasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAliasResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAliasResponse) GoString() string {
	return s.String()
}

func (s *DeleteAliasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAliasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAliasResponse) GetBody() *DeleteAliasResponseBody {
	return s.Body
}

func (s *DeleteAliasResponse) SetHeaders(v map[string]*string) *DeleteAliasResponse {
	s.Headers = v
	return s
}

func (s *DeleteAliasResponse) SetStatusCode(v int32) *DeleteAliasResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAliasResponse) SetBody(v *DeleteAliasResponseBody) *DeleteAliasResponse {
	s.Body = v
	return s
}

func (s *DeleteAliasResponse) Validate() error {
	return dara.Validate(s)
}
