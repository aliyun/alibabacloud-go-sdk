// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddScenegroupMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddScenegroupMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddScenegroupMemberResponse
	GetStatusCode() *int32
	SetBody(v *AddScenegroupMemberResponseBody) *AddScenegroupMemberResponse
	GetBody() *AddScenegroupMemberResponseBody
}

type AddScenegroupMemberResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddScenegroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddScenegroupMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s AddScenegroupMemberResponse) GoString() string {
	return s.String()
}

func (s *AddScenegroupMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddScenegroupMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddScenegroupMemberResponse) GetBody() *AddScenegroupMemberResponseBody {
	return s.Body
}

func (s *AddScenegroupMemberResponse) SetHeaders(v map[string]*string) *AddScenegroupMemberResponse {
	s.Headers = v
	return s
}

func (s *AddScenegroupMemberResponse) SetStatusCode(v int32) *AddScenegroupMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *AddScenegroupMemberResponse) SetBody(v *AddScenegroupMemberResponseBody) *AddScenegroupMemberResponse {
	s.Body = v
	return s
}

func (s *AddScenegroupMemberResponse) Validate() error {
	return dara.Validate(s)
}
