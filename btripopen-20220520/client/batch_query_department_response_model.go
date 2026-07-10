// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchQueryDepartmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchQueryDepartmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchQueryDepartmentResponse
	GetStatusCode() *int32
	SetBody(v *BatchQueryDepartmentResponseBody) *BatchQueryDepartmentResponse
	GetBody() *BatchQueryDepartmentResponseBody
}

type BatchQueryDepartmentResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchQueryDepartmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchQueryDepartmentResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryDepartmentResponse) GoString() string {
	return s.String()
}

func (s *BatchQueryDepartmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchQueryDepartmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchQueryDepartmentResponse) GetBody() *BatchQueryDepartmentResponseBody {
	return s.Body
}

func (s *BatchQueryDepartmentResponse) SetHeaders(v map[string]*string) *BatchQueryDepartmentResponse {
	s.Headers = v
	return s
}

func (s *BatchQueryDepartmentResponse) SetStatusCode(v int32) *BatchQueryDepartmentResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchQueryDepartmentResponse) SetBody(v *BatchQueryDepartmentResponseBody) *BatchQueryDepartmentResponse {
	s.Body = v
	return s
}

func (s *BatchQueryDepartmentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
