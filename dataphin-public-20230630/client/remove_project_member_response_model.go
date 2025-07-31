// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveProjectMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveProjectMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveProjectMemberResponse
	GetStatusCode() *int32
	SetBody(v *RemoveProjectMemberResponseBody) *RemoveProjectMemberResponse
	GetBody() *RemoveProjectMemberResponseBody
}

type RemoveProjectMemberResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveProjectMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveProjectMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveProjectMemberResponse) GoString() string {
	return s.String()
}

func (s *RemoveProjectMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveProjectMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveProjectMemberResponse) GetBody() *RemoveProjectMemberResponseBody {
	return s.Body
}

func (s *RemoveProjectMemberResponse) SetHeaders(v map[string]*string) *RemoveProjectMemberResponse {
	s.Headers = v
	return s
}

func (s *RemoveProjectMemberResponse) SetStatusCode(v int32) *RemoveProjectMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveProjectMemberResponse) SetBody(v *RemoveProjectMemberResponseBody) *RemoveProjectMemberResponse {
	s.Body = v
	return s
}

func (s *RemoveProjectMemberResponse) Validate() error {
	return dara.Validate(s)
}
