// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkMeetingMemberEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDingtalkMeetingMemberEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDingtalkMeetingMemberEventResponse
	GetStatusCode() *int32
	SetBody(v *GetDingtalkMeetingMemberEventResponseBody) *GetDingtalkMeetingMemberEventResponse
	GetBody() *GetDingtalkMeetingMemberEventResponseBody
}

type GetDingtalkMeetingMemberEventResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDingtalkMeetingMemberEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDingtalkMeetingMemberEventResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingMemberEventResponse) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingMemberEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDingtalkMeetingMemberEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDingtalkMeetingMemberEventResponse) GetBody() *GetDingtalkMeetingMemberEventResponseBody {
	return s.Body
}

func (s *GetDingtalkMeetingMemberEventResponse) SetHeaders(v map[string]*string) *GetDingtalkMeetingMemberEventResponse {
	s.Headers = v
	return s
}

func (s *GetDingtalkMeetingMemberEventResponse) SetStatusCode(v int32) *GetDingtalkMeetingMemberEventResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDingtalkMeetingMemberEventResponse) SetBody(v *GetDingtalkMeetingMemberEventResponseBody) *GetDingtalkMeetingMemberEventResponse {
	s.Body = v
	return s
}

func (s *GetDingtalkMeetingMemberEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
