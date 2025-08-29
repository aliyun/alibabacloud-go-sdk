// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateInstanceResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateInstanceResourceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateInstanceResourceResponseBody) *UpdateInstanceResourceResponse
	GetBody() *UpdateInstanceResourceResponseBody
}

type UpdateInstanceResourceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceResourceResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateInstanceResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateInstanceResourceResponse) GetBody() *UpdateInstanceResourceResponseBody {
	return s.Body
}

func (s *UpdateInstanceResourceResponse) SetHeaders(v map[string]*string) *UpdateInstanceResourceResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceResourceResponse) SetStatusCode(v int32) *UpdateInstanceResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceResourceResponse) SetBody(v *UpdateInstanceResourceResponseBody) *UpdateInstanceResourceResponse {
	s.Body = v
	return s
}

func (s *UpdateInstanceResourceResponse) Validate() error {
	return dara.Validate(s)
}
