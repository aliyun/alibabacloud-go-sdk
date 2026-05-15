// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendCcoSmartCallOperateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendCcoSmartCallOperateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendCcoSmartCallOperateResponse
	GetStatusCode() *int32
	SetBody(v *SendCcoSmartCallOperateResponseBody) *SendCcoSmartCallOperateResponse
	GetBody() *SendCcoSmartCallOperateResponseBody
}

type SendCcoSmartCallOperateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendCcoSmartCallOperateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendCcoSmartCallOperateResponse) String() string {
	return dara.Prettify(s)
}

func (s SendCcoSmartCallOperateResponse) GoString() string {
	return s.String()
}

func (s *SendCcoSmartCallOperateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendCcoSmartCallOperateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendCcoSmartCallOperateResponse) GetBody() *SendCcoSmartCallOperateResponseBody {
	return s.Body
}

func (s *SendCcoSmartCallOperateResponse) SetHeaders(v map[string]*string) *SendCcoSmartCallOperateResponse {
	s.Headers = v
	return s
}

func (s *SendCcoSmartCallOperateResponse) SetStatusCode(v int32) *SendCcoSmartCallOperateResponse {
	s.StatusCode = &v
	return s
}

func (s *SendCcoSmartCallOperateResponse) SetBody(v *SendCcoSmartCallOperateResponseBody) *SendCcoSmartCallOperateResponse {
	s.Body = v
	return s
}

func (s *SendCcoSmartCallOperateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
