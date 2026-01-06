// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkMeetingListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDingtalkMeetingListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDingtalkMeetingListResponse
	GetStatusCode() *int32
	SetBody(v *GetDingtalkMeetingListResponseBody) *GetDingtalkMeetingListResponse
	GetBody() *GetDingtalkMeetingListResponseBody
}

type GetDingtalkMeetingListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDingtalkMeetingListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDingtalkMeetingListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingListResponse) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDingtalkMeetingListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDingtalkMeetingListResponse) GetBody() *GetDingtalkMeetingListResponseBody {
	return s.Body
}

func (s *GetDingtalkMeetingListResponse) SetHeaders(v map[string]*string) *GetDingtalkMeetingListResponse {
	s.Headers = v
	return s
}

func (s *GetDingtalkMeetingListResponse) SetStatusCode(v int32) *GetDingtalkMeetingListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDingtalkMeetingListResponse) SetBody(v *GetDingtalkMeetingListResponseBody) *GetDingtalkMeetingListResponse {
	s.Body = v
	return s
}

func (s *GetDingtalkMeetingListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
