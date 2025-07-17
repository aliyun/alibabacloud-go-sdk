// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetProjectMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetProjectMemberResponse
	GetStatusCode() *int32
	SetBody(v *GetProjectMemberResponseBody) *GetProjectMemberResponse
	GetBody() *GetProjectMemberResponseBody
}

type GetProjectMemberResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProjectMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProjectMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s GetProjectMemberResponse) GoString() string {
	return s.String()
}

func (s *GetProjectMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetProjectMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetProjectMemberResponse) GetBody() *GetProjectMemberResponseBody {
	return s.Body
}

func (s *GetProjectMemberResponse) SetHeaders(v map[string]*string) *GetProjectMemberResponse {
	s.Headers = v
	return s
}

func (s *GetProjectMemberResponse) SetStatusCode(v int32) *GetProjectMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectMemberResponse) SetBody(v *GetProjectMemberResponseBody) *GetProjectMemberResponse {
	s.Body = v
	return s
}

func (s *GetProjectMemberResponse) Validate() error {
	return dara.Validate(s)
}
