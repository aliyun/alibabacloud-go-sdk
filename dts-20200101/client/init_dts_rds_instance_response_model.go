// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitDtsRdsInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InitDtsRdsInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InitDtsRdsInstanceResponse
	GetStatusCode() *int32
	SetBody(v *InitDtsRdsInstanceResponseBody) *InitDtsRdsInstanceResponse
	GetBody() *InitDtsRdsInstanceResponseBody
}

type InitDtsRdsInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitDtsRdsInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitDtsRdsInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s InitDtsRdsInstanceResponse) GoString() string {
	return s.String()
}

func (s *InitDtsRdsInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InitDtsRdsInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InitDtsRdsInstanceResponse) GetBody() *InitDtsRdsInstanceResponseBody {
	return s.Body
}

func (s *InitDtsRdsInstanceResponse) SetHeaders(v map[string]*string) *InitDtsRdsInstanceResponse {
	s.Headers = v
	return s
}

func (s *InitDtsRdsInstanceResponse) SetStatusCode(v int32) *InitDtsRdsInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *InitDtsRdsInstanceResponse) SetBody(v *InitDtsRdsInstanceResponseBody) *InitDtsRdsInstanceResponse {
	s.Body = v
	return s
}

func (s *InitDtsRdsInstanceResponse) Validate() error {
	return dara.Validate(s)
}
