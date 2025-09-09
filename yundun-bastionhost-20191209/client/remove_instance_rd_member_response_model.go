// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveInstanceRdMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveInstanceRdMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveInstanceRdMemberResponse
	GetStatusCode() *int32
	SetBody(v *RemoveInstanceRdMemberResponseBody) *RemoveInstanceRdMemberResponse
	GetBody() *RemoveInstanceRdMemberResponseBody
}

type RemoveInstanceRdMemberResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveInstanceRdMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveInstanceRdMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveInstanceRdMemberResponse) GoString() string {
	return s.String()
}

func (s *RemoveInstanceRdMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveInstanceRdMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveInstanceRdMemberResponse) GetBody() *RemoveInstanceRdMemberResponseBody {
	return s.Body
}

func (s *RemoveInstanceRdMemberResponse) SetHeaders(v map[string]*string) *RemoveInstanceRdMemberResponse {
	s.Headers = v
	return s
}

func (s *RemoveInstanceRdMemberResponse) SetStatusCode(v int32) *RemoveInstanceRdMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveInstanceRdMemberResponse) SetBody(v *RemoveInstanceRdMemberResponseBody) *RemoveInstanceRdMemberResponse {
	s.Body = v
	return s
}

func (s *RemoveInstanceRdMemberResponse) Validate() error {
	return dara.Validate(s)
}
