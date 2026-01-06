// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkMeetingInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDingtalkMeetingInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDingtalkMeetingInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetDingtalkMeetingInfoResponseBody) *GetDingtalkMeetingInfoResponse
	GetBody() *GetDingtalkMeetingInfoResponseBody
}

type GetDingtalkMeetingInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDingtalkMeetingInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDingtalkMeetingInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingInfoResponse) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDingtalkMeetingInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDingtalkMeetingInfoResponse) GetBody() *GetDingtalkMeetingInfoResponseBody {
	return s.Body
}

func (s *GetDingtalkMeetingInfoResponse) SetHeaders(v map[string]*string) *GetDingtalkMeetingInfoResponse {
	s.Headers = v
	return s
}

func (s *GetDingtalkMeetingInfoResponse) SetStatusCode(v int32) *GetDingtalkMeetingInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDingtalkMeetingInfoResponse) SetBody(v *GetDingtalkMeetingInfoResponseBody) *GetDingtalkMeetingInfoResponse {
	s.Body = v
	return s
}

func (s *GetDingtalkMeetingInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
