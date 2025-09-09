// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHostResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyHostResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyHostResponse
	GetStatusCode() *int32
	SetBody(v *ModifyHostResponseBody) *ModifyHostResponse
	GetBody() *ModifyHostResponseBody
}

type ModifyHostResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHostResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyHostResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyHostResponse) GoString() string {
	return s.String()
}

func (s *ModifyHostResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyHostResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyHostResponse) GetBody() *ModifyHostResponseBody {
	return s.Body
}

func (s *ModifyHostResponse) SetHeaders(v map[string]*string) *ModifyHostResponse {
	s.Headers = v
	return s
}

func (s *ModifyHostResponse) SetStatusCode(v int32) *ModifyHostResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHostResponse) SetBody(v *ModifyHostResponseBody) *ModifyHostResponse {
	s.Body = v
	return s
}

func (s *ModifyHostResponse) Validate() error {
	return dara.Validate(s)
}
