// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateBlackNoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateBlackNoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateBlackNoResponse
	GetStatusCode() *int32
	SetBody(v *OperateBlackNoResponseBody) *OperateBlackNoResponse
	GetBody() *OperateBlackNoResponseBody
}

type OperateBlackNoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateBlackNoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateBlackNoResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateBlackNoResponse) GoString() string {
	return s.String()
}

func (s *OperateBlackNoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateBlackNoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateBlackNoResponse) GetBody() *OperateBlackNoResponseBody {
	return s.Body
}

func (s *OperateBlackNoResponse) SetHeaders(v map[string]*string) *OperateBlackNoResponse {
	s.Headers = v
	return s
}

func (s *OperateBlackNoResponse) SetStatusCode(v int32) *OperateBlackNoResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateBlackNoResponse) SetBody(v *OperateBlackNoResponseBody) *OperateBlackNoResponse {
	s.Body = v
	return s
}

func (s *OperateBlackNoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
