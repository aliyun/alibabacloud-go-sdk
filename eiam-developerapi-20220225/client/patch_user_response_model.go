// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPatchUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PatchUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PatchUserResponse
	GetStatusCode() *int32
}

type PatchUserResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PatchUserResponse) String() string {
	return dara.Prettify(s)
}

func (s PatchUserResponse) GoString() string {
	return s.String()
}

func (s *PatchUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PatchUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PatchUserResponse) SetHeaders(v map[string]*string) *PatchUserResponse {
	s.Headers = v
	return s
}

func (s *PatchUserResponse) SetStatusCode(v int32) *PatchUserResponse {
	s.StatusCode = &v
	return s
}

func (s *PatchUserResponse) Validate() error {
	return dara.Validate(s)
}
