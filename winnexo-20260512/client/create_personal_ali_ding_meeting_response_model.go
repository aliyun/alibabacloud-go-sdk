// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalAliDingMeetingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePersonalAliDingMeetingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePersonalAliDingMeetingResponse
	GetStatusCode() *int32
	SetBody(v *CreatePersonalAliDingMeetingResponseBody) *CreatePersonalAliDingMeetingResponse
	GetBody() *CreatePersonalAliDingMeetingResponseBody
}

type CreatePersonalAliDingMeetingResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePersonalAliDingMeetingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePersonalAliDingMeetingResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalAliDingMeetingResponse) GoString() string {
	return s.String()
}

func (s *CreatePersonalAliDingMeetingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePersonalAliDingMeetingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePersonalAliDingMeetingResponse) GetBody() *CreatePersonalAliDingMeetingResponseBody {
	return s.Body
}

func (s *CreatePersonalAliDingMeetingResponse) SetHeaders(v map[string]*string) *CreatePersonalAliDingMeetingResponse {
	s.Headers = v
	return s
}

func (s *CreatePersonalAliDingMeetingResponse) SetStatusCode(v int32) *CreatePersonalAliDingMeetingResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePersonalAliDingMeetingResponse) SetBody(v *CreatePersonalAliDingMeetingResponseBody) *CreatePersonalAliDingMeetingResponse {
	s.Body = v
	return s
}

func (s *CreatePersonalAliDingMeetingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
