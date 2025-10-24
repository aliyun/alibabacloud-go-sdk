// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateAppResponse
	GetStatusCode() *int32
	SetBody(v *OperateAppResponseBody) *OperateAppResponse
	GetBody() *OperateAppResponseBody
}

type OperateAppResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateAppResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateAppResponse) GoString() string {
	return s.String()
}

func (s *OperateAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateAppResponse) GetBody() *OperateAppResponseBody {
	return s.Body
}

func (s *OperateAppResponse) SetHeaders(v map[string]*string) *OperateAppResponse {
	s.Headers = v
	return s
}

func (s *OperateAppResponse) SetStatusCode(v int32) *OperateAppResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateAppResponse) SetBody(v *OperateAppResponseBody) *OperateAppResponse {
	s.Body = v
	return s
}

func (s *OperateAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
