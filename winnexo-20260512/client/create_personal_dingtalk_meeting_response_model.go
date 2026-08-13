// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalDingtalkMeetingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePersonalDingtalkMeetingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePersonalDingtalkMeetingResponse
	GetStatusCode() *int32
	SetBody(v *CreatePersonalDingtalkMeetingResponseBody) *CreatePersonalDingtalkMeetingResponse
	GetBody() *CreatePersonalDingtalkMeetingResponseBody
}

type CreatePersonalDingtalkMeetingResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePersonalDingtalkMeetingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePersonalDingtalkMeetingResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalDingtalkMeetingResponse) GoString() string {
	return s.String()
}

func (s *CreatePersonalDingtalkMeetingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePersonalDingtalkMeetingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePersonalDingtalkMeetingResponse) GetBody() *CreatePersonalDingtalkMeetingResponseBody {
	return s.Body
}

func (s *CreatePersonalDingtalkMeetingResponse) SetHeaders(v map[string]*string) *CreatePersonalDingtalkMeetingResponse {
	s.Headers = v
	return s
}

func (s *CreatePersonalDingtalkMeetingResponse) SetStatusCode(v int32) *CreatePersonalDingtalkMeetingResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePersonalDingtalkMeetingResponse) SetBody(v *CreatePersonalDingtalkMeetingResponseBody) *CreatePersonalDingtalkMeetingResponse {
	s.Body = v
	return s
}

func (s *CreatePersonalDingtalkMeetingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
