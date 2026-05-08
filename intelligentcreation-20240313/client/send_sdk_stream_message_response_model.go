// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendSdkStreamMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendSdkStreamMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendSdkStreamMessageResponse
	GetStatusCode() *int32
	SetId(v string) *SendSdkStreamMessageResponse
	GetId() *string
	SetEvent(v string) *SendSdkStreamMessageResponse
	GetEvent() *string
	SetBody(v *SendSdkStreamMessageResponseBody) *SendSdkStreamMessageResponse
	GetBody() *SendSdkStreamMessageResponseBody
}

type SendSdkStreamMessageResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                           `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                           `json:"event,omitempty" xml:"event,omitempty"`
	Body       *SendSdkStreamMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendSdkStreamMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s SendSdkStreamMessageResponse) GoString() string {
	return s.String()
}

func (s *SendSdkStreamMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendSdkStreamMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendSdkStreamMessageResponse) GetId() *string {
	return s.Id
}

func (s *SendSdkStreamMessageResponse) GetEvent() *string {
	return s.Event
}

func (s *SendSdkStreamMessageResponse) GetBody() *SendSdkStreamMessageResponseBody {
	return s.Body
}

func (s *SendSdkStreamMessageResponse) SetHeaders(v map[string]*string) *SendSdkStreamMessageResponse {
	s.Headers = v
	return s
}

func (s *SendSdkStreamMessageResponse) SetStatusCode(v int32) *SendSdkStreamMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendSdkStreamMessageResponse) SetId(v string) *SendSdkStreamMessageResponse {
	s.Id = &v
	return s
}

func (s *SendSdkStreamMessageResponse) SetEvent(v string) *SendSdkStreamMessageResponse {
	s.Event = &v
	return s
}

func (s *SendSdkStreamMessageResponse) SetBody(v *SendSdkStreamMessageResponseBody) *SendSdkStreamMessageResponse {
	s.Body = v
	return s
}

func (s *SendSdkStreamMessageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
