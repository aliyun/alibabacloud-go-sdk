// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPortResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPortResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPortResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPortResponseBody) *ModifyPortResponse
	GetBody() *ModifyPortResponseBody
}

type ModifyPortResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPortResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPortResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPortResponse) GoString() string {
	return s.String()
}

func (s *ModifyPortResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPortResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPortResponse) GetBody() *ModifyPortResponseBody {
	return s.Body
}

func (s *ModifyPortResponse) SetHeaders(v map[string]*string) *ModifyPortResponse {
	s.Headers = v
	return s
}

func (s *ModifyPortResponse) SetStatusCode(v int32) *ModifyPortResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPortResponse) SetBody(v *ModifyPortResponseBody) *ModifyPortResponse {
	s.Body = v
	return s
}

func (s *ModifyPortResponse) Validate() error {
	return dara.Validate(s)
}
