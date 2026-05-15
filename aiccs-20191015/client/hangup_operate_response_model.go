// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHangupOperateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HangupOperateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HangupOperateResponse
	GetStatusCode() *int32
	SetBody(v *HangupOperateResponseBody) *HangupOperateResponse
	GetBody() *HangupOperateResponseBody
}

type HangupOperateResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HangupOperateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HangupOperateResponse) String() string {
	return dara.Prettify(s)
}

func (s HangupOperateResponse) GoString() string {
	return s.String()
}

func (s *HangupOperateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HangupOperateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HangupOperateResponse) GetBody() *HangupOperateResponseBody {
	return s.Body
}

func (s *HangupOperateResponse) SetHeaders(v map[string]*string) *HangupOperateResponse {
	s.Headers = v
	return s
}

func (s *HangupOperateResponse) SetStatusCode(v int32) *HangupOperateResponse {
	s.StatusCode = &v
	return s
}

func (s *HangupOperateResponse) SetBody(v *HangupOperateResponseBody) *HangupOperateResponse {
	s.Body = v
	return s
}

func (s *HangupOperateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
