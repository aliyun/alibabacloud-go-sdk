// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMemberResponse
	GetStatusCode() *int32
	SetBody(v *GetMemberResponseBody) *GetMemberResponse
	GetBody() *GetMemberResponseBody
}

type GetMemberResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMemberResponse) GoString() string {
	return s.String()
}

func (s *GetMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMemberResponse) GetBody() *GetMemberResponseBody {
	return s.Body
}

func (s *GetMemberResponse) SetHeaders(v map[string]*string) *GetMemberResponse {
	s.Headers = v
	return s
}

func (s *GetMemberResponse) SetStatusCode(v int32) *GetMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMemberResponse) SetBody(v *GetMemberResponseBody) *GetMemberResponse {
	s.Body = v
	return s
}

func (s *GetMemberResponse) Validate() error {
	return dara.Validate(s)
}
