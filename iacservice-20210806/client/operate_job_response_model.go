// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateJobResponse
	GetStatusCode() *int32
	SetBody(v *OperateJobResponseBody) *OperateJobResponse
	GetBody() *OperateJobResponseBody
}

type OperateJobResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateJobResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateJobResponse) GoString() string {
	return s.String()
}

func (s *OperateJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateJobResponse) GetBody() *OperateJobResponseBody {
	return s.Body
}

func (s *OperateJobResponse) SetHeaders(v map[string]*string) *OperateJobResponse {
	s.Headers = v
	return s
}

func (s *OperateJobResponse) SetStatusCode(v int32) *OperateJobResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateJobResponse) SetBody(v *OperateJobResponseBody) *OperateJobResponse {
	s.Body = v
	return s
}

func (s *OperateJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
