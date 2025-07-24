// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackServiceInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RollbackServiceInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RollbackServiceInstanceResponse
	GetStatusCode() *int32
	SetBody(v *RollbackServiceInstanceResponseBody) *RollbackServiceInstanceResponse
	GetBody() *RollbackServiceInstanceResponseBody
}

type RollbackServiceInstanceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RollbackServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RollbackServiceInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RollbackServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *RollbackServiceInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RollbackServiceInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RollbackServiceInstanceResponse) GetBody() *RollbackServiceInstanceResponseBody {
	return s.Body
}

func (s *RollbackServiceInstanceResponse) SetHeaders(v map[string]*string) *RollbackServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *RollbackServiceInstanceResponse) SetStatusCode(v int32) *RollbackServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RollbackServiceInstanceResponse) SetBody(v *RollbackServiceInstanceResponseBody) *RollbackServiceInstanceResponse {
	s.Body = v
	return s
}

func (s *RollbackServiceInstanceResponse) Validate() error {
	return dara.Validate(s)
}
