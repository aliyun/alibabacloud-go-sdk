// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartProcessInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartProcessInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartProcessInstanceResponse
	GetStatusCode() *int32
	SetBody(v *StartProcessInstanceResponseBody) *StartProcessInstanceResponse
	GetBody() *StartProcessInstanceResponseBody
}

type StartProcessInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartProcessInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartProcessInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s StartProcessInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartProcessInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartProcessInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartProcessInstanceResponse) GetBody() *StartProcessInstanceResponseBody {
	return s.Body
}

func (s *StartProcessInstanceResponse) SetHeaders(v map[string]*string) *StartProcessInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartProcessInstanceResponse) SetStatusCode(v int32) *StartProcessInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartProcessInstanceResponse) SetBody(v *StartProcessInstanceResponseBody) *StartProcessInstanceResponse {
	s.Body = v
	return s
}

func (s *StartProcessInstanceResponse) Validate() error {
	return dara.Validate(s)
}
