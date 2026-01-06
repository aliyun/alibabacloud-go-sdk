// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkMeetingMemberListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDingtalkMeetingMemberListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDingtalkMeetingMemberListResponse
	GetStatusCode() *int32
	SetBody(v *GetDingtalkMeetingMemberListResponseBody) *GetDingtalkMeetingMemberListResponse
	GetBody() *GetDingtalkMeetingMemberListResponseBody
}

type GetDingtalkMeetingMemberListResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDingtalkMeetingMemberListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDingtalkMeetingMemberListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingMemberListResponse) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingMemberListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDingtalkMeetingMemberListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDingtalkMeetingMemberListResponse) GetBody() *GetDingtalkMeetingMemberListResponseBody {
	return s.Body
}

func (s *GetDingtalkMeetingMemberListResponse) SetHeaders(v map[string]*string) *GetDingtalkMeetingMemberListResponse {
	s.Headers = v
	return s
}

func (s *GetDingtalkMeetingMemberListResponse) SetStatusCode(v int32) *GetDingtalkMeetingMemberListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDingtalkMeetingMemberListResponse) SetBody(v *GetDingtalkMeetingMemberListResponseBody) *GetDingtalkMeetingMemberListResponse {
	s.Body = v
	return s
}

func (s *GetDingtalkMeetingMemberListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
