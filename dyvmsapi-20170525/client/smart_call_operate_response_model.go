// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSmartCallOperateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SmartCallOperateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SmartCallOperateResponse
	GetStatusCode() *int32
	SetBody(v *SmartCallOperateResponseBody) *SmartCallOperateResponse
	GetBody() *SmartCallOperateResponseBody
}

type SmartCallOperateResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SmartCallOperateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SmartCallOperateResponse) String() string {
	return dara.Prettify(s)
}

func (s SmartCallOperateResponse) GoString() string {
	return s.String()
}

func (s *SmartCallOperateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SmartCallOperateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SmartCallOperateResponse) GetBody() *SmartCallOperateResponseBody {
	return s.Body
}

func (s *SmartCallOperateResponse) SetHeaders(v map[string]*string) *SmartCallOperateResponse {
	s.Headers = v
	return s
}

func (s *SmartCallOperateResponse) SetStatusCode(v int32) *SmartCallOperateResponse {
	s.StatusCode = &v
	return s
}

func (s *SmartCallOperateResponse) SetBody(v *SmartCallOperateResponseBody) *SmartCallOperateResponse {
	s.Body = v
	return s
}

func (s *SmartCallOperateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
