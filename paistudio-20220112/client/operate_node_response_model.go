// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateNodeResponse
	GetStatusCode() *int32
	SetBody(v *OperateNodeResponseBody) *OperateNodeResponse
	GetBody() *OperateNodeResponseBody
}

type OperateNodeResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateNodeResponse) GoString() string {
	return s.String()
}

func (s *OperateNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateNodeResponse) GetBody() *OperateNodeResponseBody {
	return s.Body
}

func (s *OperateNodeResponse) SetHeaders(v map[string]*string) *OperateNodeResponse {
	s.Headers = v
	return s
}

func (s *OperateNodeResponse) SetStatusCode(v int32) *OperateNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateNodeResponse) SetBody(v *OperateNodeResponseBody) *OperateNodeResponse {
	s.Body = v
	return s
}

func (s *OperateNodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
