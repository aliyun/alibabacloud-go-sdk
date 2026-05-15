// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendCcoSmartCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendCcoSmartCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendCcoSmartCallResponse
	GetStatusCode() *int32
	SetBody(v *SendCcoSmartCallResponseBody) *SendCcoSmartCallResponse
	GetBody() *SendCcoSmartCallResponseBody
}

type SendCcoSmartCallResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendCcoSmartCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendCcoSmartCallResponse) String() string {
	return dara.Prettify(s)
}

func (s SendCcoSmartCallResponse) GoString() string {
	return s.String()
}

func (s *SendCcoSmartCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendCcoSmartCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendCcoSmartCallResponse) GetBody() *SendCcoSmartCallResponseBody {
	return s.Body
}

func (s *SendCcoSmartCallResponse) SetHeaders(v map[string]*string) *SendCcoSmartCallResponse {
	s.Headers = v
	return s
}

func (s *SendCcoSmartCallResponse) SetStatusCode(v int32) *SendCcoSmartCallResponse {
	s.StatusCode = &v
	return s
}

func (s *SendCcoSmartCallResponse) SetBody(v *SendCcoSmartCallResponseBody) *SendCcoSmartCallResponse {
	s.Body = v
	return s
}

func (s *SendCcoSmartCallResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
