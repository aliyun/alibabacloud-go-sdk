// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateHcWarningsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ValidateHcWarningsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ValidateHcWarningsResponse
	GetStatusCode() *int32
	SetBody(v *ValidateHcWarningsResponseBody) *ValidateHcWarningsResponse
	GetBody() *ValidateHcWarningsResponseBody
}

type ValidateHcWarningsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValidateHcWarningsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateHcWarningsResponse) String() string {
	return dara.Prettify(s)
}

func (s ValidateHcWarningsResponse) GoString() string {
	return s.String()
}

func (s *ValidateHcWarningsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ValidateHcWarningsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ValidateHcWarningsResponse) GetBody() *ValidateHcWarningsResponseBody {
	return s.Body
}

func (s *ValidateHcWarningsResponse) SetHeaders(v map[string]*string) *ValidateHcWarningsResponse {
	s.Headers = v
	return s
}

func (s *ValidateHcWarningsResponse) SetStatusCode(v int32) *ValidateHcWarningsResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateHcWarningsResponse) SetBody(v *ValidateHcWarningsResponseBody) *ValidateHcWarningsResponse {
	s.Body = v
	return s
}

func (s *ValidateHcWarningsResponse) Validate() error {
	return dara.Validate(s)
}
