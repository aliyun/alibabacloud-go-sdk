// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsMessagePushResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnsMessagePushResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnsMessagePushResponse
	GetStatusCode() *int32
	SetBody(v *OnsMessagePushResponseBody) *OnsMessagePushResponse
	GetBody() *OnsMessagePushResponseBody
}

type OnsMessagePushResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsMessagePushResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsMessagePushResponse) String() string {
	return dara.Prettify(s)
}

func (s OnsMessagePushResponse) GoString() string {
	return s.String()
}

func (s *OnsMessagePushResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnsMessagePushResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnsMessagePushResponse) GetBody() *OnsMessagePushResponseBody {
	return s.Body
}

func (s *OnsMessagePushResponse) SetHeaders(v map[string]*string) *OnsMessagePushResponse {
	s.Headers = v
	return s
}

func (s *OnsMessagePushResponse) SetStatusCode(v int32) *OnsMessagePushResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsMessagePushResponse) SetBody(v *OnsMessagePushResponseBody) *OnsMessagePushResponse {
	s.Body = v
	return s
}

func (s *OnsMessagePushResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
