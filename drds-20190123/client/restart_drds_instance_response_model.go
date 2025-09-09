// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartDrdsInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartDrdsInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartDrdsInstanceResponse
	GetStatusCode() *int32
	SetBody(v *RestartDrdsInstanceResponseBody) *RestartDrdsInstanceResponse
	GetBody() *RestartDrdsInstanceResponseBody
}

type RestartDrdsInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartDrdsInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartDrdsInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartDrdsInstanceResponse) GoString() string {
	return s.String()
}

func (s *RestartDrdsInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartDrdsInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartDrdsInstanceResponse) GetBody() *RestartDrdsInstanceResponseBody {
	return s.Body
}

func (s *RestartDrdsInstanceResponse) SetHeaders(v map[string]*string) *RestartDrdsInstanceResponse {
	s.Headers = v
	return s
}

func (s *RestartDrdsInstanceResponse) SetStatusCode(v int32) *RestartDrdsInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartDrdsInstanceResponse) SetBody(v *RestartDrdsInstanceResponseBody) *RestartDrdsInstanceResponse {
	s.Body = v
	return s
}

func (s *RestartDrdsInstanceResponse) Validate() error {
	return dara.Validate(s)
}
