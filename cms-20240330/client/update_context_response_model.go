// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContextResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateContextResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateContextResponse
	GetStatusCode() *int32
	SetBody(v map[string]interface{}) *UpdateContextResponse
	GetBody() map[string]interface{}
}

type UpdateContextResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateContextResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextResponse) GoString() string {
	return s.String()
}

func (s *UpdateContextResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateContextResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateContextResponse) GetBody() map[string]interface{} {
	return s.Body
}

func (s *UpdateContextResponse) SetHeaders(v map[string]*string) *UpdateContextResponse {
	s.Headers = v
	return s
}

func (s *UpdateContextResponse) SetStatusCode(v int32) *UpdateContextResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateContextResponse) SetBody(v map[string]interface{}) *UpdateContextResponse {
	s.Body = v
	return s
}

func (s *UpdateContextResponse) Validate() error {
	return dara.Validate(s)
}
