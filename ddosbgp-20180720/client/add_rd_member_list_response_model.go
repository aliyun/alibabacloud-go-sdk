// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRdMemberListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddRdMemberListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddRdMemberListResponse
	GetStatusCode() *int32
	SetBody(v *AddRdMemberListResponseBody) *AddRdMemberListResponse
	GetBody() *AddRdMemberListResponseBody
}

type AddRdMemberListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddRdMemberListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddRdMemberListResponse) String() string {
	return dara.Prettify(s)
}

func (s AddRdMemberListResponse) GoString() string {
	return s.String()
}

func (s *AddRdMemberListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddRdMemberListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddRdMemberListResponse) GetBody() *AddRdMemberListResponseBody {
	return s.Body
}

func (s *AddRdMemberListResponse) SetHeaders(v map[string]*string) *AddRdMemberListResponse {
	s.Headers = v
	return s
}

func (s *AddRdMemberListResponse) SetStatusCode(v int32) *AddRdMemberListResponse {
	s.StatusCode = &v
	return s
}

func (s *AddRdMemberListResponse) SetBody(v *AddRdMemberListResponseBody) *AddRdMemberListResponse {
	s.Body = v
	return s
}

func (s *AddRdMemberListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
