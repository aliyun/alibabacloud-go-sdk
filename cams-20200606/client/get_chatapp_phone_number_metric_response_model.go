// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatappPhoneNumberMetricResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetChatappPhoneNumberMetricResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetChatappPhoneNumberMetricResponse
	GetStatusCode() *int32
	SetBody(v *GetChatappPhoneNumberMetricResponseBody) *GetChatappPhoneNumberMetricResponse
	GetBody() *GetChatappPhoneNumberMetricResponseBody
}

type GetChatappPhoneNumberMetricResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChatappPhoneNumberMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChatappPhoneNumberMetricResponse) String() string {
	return dara.Prettify(s)
}

func (s GetChatappPhoneNumberMetricResponse) GoString() string {
	return s.String()
}

func (s *GetChatappPhoneNumberMetricResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetChatappPhoneNumberMetricResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetChatappPhoneNumberMetricResponse) GetBody() *GetChatappPhoneNumberMetricResponseBody {
	return s.Body
}

func (s *GetChatappPhoneNumberMetricResponse) SetHeaders(v map[string]*string) *GetChatappPhoneNumberMetricResponse {
	s.Headers = v
	return s
}

func (s *GetChatappPhoneNumberMetricResponse) SetStatusCode(v int32) *GetChatappPhoneNumberMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChatappPhoneNumberMetricResponse) SetBody(v *GetChatappPhoneNumberMetricResponseBody) *GetChatappPhoneNumberMetricResponse {
	s.Body = v
	return s
}

func (s *GetChatappPhoneNumberMetricResponse) Validate() error {
	return dara.Validate(s)
}
