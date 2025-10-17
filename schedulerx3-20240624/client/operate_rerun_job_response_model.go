// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateRerunJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateRerunJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateRerunJobResponse
	GetStatusCode() *int32
	SetBody(v *OperateRerunJobResponseBody) *OperateRerunJobResponse
	GetBody() *OperateRerunJobResponseBody
}

type OperateRerunJobResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateRerunJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateRerunJobResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateRerunJobResponse) GoString() string {
	return s.String()
}

func (s *OperateRerunJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateRerunJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateRerunJobResponse) GetBody() *OperateRerunJobResponseBody {
	return s.Body
}

func (s *OperateRerunJobResponse) SetHeaders(v map[string]*string) *OperateRerunJobResponse {
	s.Headers = v
	return s
}

func (s *OperateRerunJobResponse) SetStatusCode(v int32) *OperateRerunJobResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateRerunJobResponse) SetBody(v *OperateRerunJobResponseBody) *OperateRerunJobResponse {
	s.Body = v
	return s
}

func (s *OperateRerunJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
