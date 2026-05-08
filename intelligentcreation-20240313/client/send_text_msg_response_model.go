// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendTextMsgResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendTextMsgResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendTextMsgResponse
	GetStatusCode() *int32
	SetBody(v *SendTextMsgResponseBody) *SendTextMsgResponse
	GetBody() *SendTextMsgResponseBody
}

type SendTextMsgResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendTextMsgResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendTextMsgResponse) String() string {
	return dara.Prettify(s)
}

func (s SendTextMsgResponse) GoString() string {
	return s.String()
}

func (s *SendTextMsgResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendTextMsgResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendTextMsgResponse) GetBody() *SendTextMsgResponseBody {
	return s.Body
}

func (s *SendTextMsgResponse) SetHeaders(v map[string]*string) *SendTextMsgResponse {
	s.Headers = v
	return s
}

func (s *SendTextMsgResponse) SetStatusCode(v int32) *SendTextMsgResponse {
	s.StatusCode = &v
	return s
}

func (s *SendTextMsgResponse) SetBody(v *SendTextMsgResponseBody) *SendTextMsgResponse {
	s.Body = v
	return s
}

func (s *SendTextMsgResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
