// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyTagResponse
	GetStatusCode() *int32
	SetBody(v *ModifyTagResponseBody) *ModifyTagResponse
	GetBody() *ModifyTagResponseBody
}

type ModifyTagResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyTagResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyTagResponse) GoString() string {
	return s.String()
}

func (s *ModifyTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyTagResponse) GetBody() *ModifyTagResponseBody {
	return s.Body
}

func (s *ModifyTagResponse) SetHeaders(v map[string]*string) *ModifyTagResponse {
	s.Headers = v
	return s
}

func (s *ModifyTagResponse) SetStatusCode(v int32) *ModifyTagResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTagResponse) SetBody(v *ModifyTagResponseBody) *ModifyTagResponse {
	s.Body = v
	return s
}

func (s *ModifyTagResponse) Validate() error {
	return dara.Validate(s)
}
