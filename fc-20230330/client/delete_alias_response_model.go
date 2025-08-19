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
}

type DeleteAliasResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

func (s *DeleteAliasResponse) SetHeaders(v map[string]*string) *DeleteAliasResponse {
	s.Headers = v
	return s
}

func (s *DeleteAliasResponse) SetStatusCode(v int32) *DeleteAliasResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAliasResponse) Validate() error {
	return dara.Validate(s)
}
