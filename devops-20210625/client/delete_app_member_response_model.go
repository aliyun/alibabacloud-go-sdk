// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAppMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAppMemberResponse
	GetStatusCode() *int32
	SetBody(v string) *DeleteAppMemberResponse
	GetBody() *string
}

type DeleteAppMemberResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppMemberResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAppMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAppMemberResponse) GetBody() *string {
	return s.Body
}

func (s *DeleteAppMemberResponse) SetHeaders(v map[string]*string) *DeleteAppMemberResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppMemberResponse) SetStatusCode(v int32) *DeleteAppMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppMemberResponse) SetBody(v string) *DeleteAppMemberResponse {
	s.Body = &v
	return s
}

func (s *DeleteAppMemberResponse) Validate() error {
	return dara.Validate(s)
}
