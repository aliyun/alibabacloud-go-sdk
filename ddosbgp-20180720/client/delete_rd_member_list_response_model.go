// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRdMemberListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRdMemberListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRdMemberListResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRdMemberListResponseBody) *DeleteRdMemberListResponse
	GetBody() *DeleteRdMemberListResponseBody
}

type DeleteRdMemberListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRdMemberListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRdMemberListResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRdMemberListResponse) GoString() string {
	return s.String()
}

func (s *DeleteRdMemberListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRdMemberListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRdMemberListResponse) GetBody() *DeleteRdMemberListResponseBody {
	return s.Body
}

func (s *DeleteRdMemberListResponse) SetHeaders(v map[string]*string) *DeleteRdMemberListResponse {
	s.Headers = v
	return s
}

func (s *DeleteRdMemberListResponse) SetStatusCode(v int32) *DeleteRdMemberListResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRdMemberListResponse) SetBody(v *DeleteRdMemberListResponseBody) *DeleteRdMemberListResponse {
	s.Body = v
	return s
}

func (s *DeleteRdMemberListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
