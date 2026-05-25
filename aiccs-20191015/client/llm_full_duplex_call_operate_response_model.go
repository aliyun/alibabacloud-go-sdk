// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLlmFullDuplexCallOperateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LlmFullDuplexCallOperateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LlmFullDuplexCallOperateResponse
	GetStatusCode() *int32
	SetBody(v *LlmFullDuplexCallOperateResponseBody) *LlmFullDuplexCallOperateResponse
	GetBody() *LlmFullDuplexCallOperateResponseBody
}

type LlmFullDuplexCallOperateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LlmFullDuplexCallOperateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LlmFullDuplexCallOperateResponse) String() string {
	return dara.Prettify(s)
}

func (s LlmFullDuplexCallOperateResponse) GoString() string {
	return s.String()
}

func (s *LlmFullDuplexCallOperateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LlmFullDuplexCallOperateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LlmFullDuplexCallOperateResponse) GetBody() *LlmFullDuplexCallOperateResponseBody {
	return s.Body
}

func (s *LlmFullDuplexCallOperateResponse) SetHeaders(v map[string]*string) *LlmFullDuplexCallOperateResponse {
	s.Headers = v
	return s
}

func (s *LlmFullDuplexCallOperateResponse) SetStatusCode(v int32) *LlmFullDuplexCallOperateResponse {
	s.StatusCode = &v
	return s
}

func (s *LlmFullDuplexCallOperateResponse) SetBody(v *LlmFullDuplexCallOperateResponseBody) *LlmFullDuplexCallOperateResponse {
	s.Body = v
	return s
}

func (s *LlmFullDuplexCallOperateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
