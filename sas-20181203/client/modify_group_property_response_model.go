// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGroupPropertyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyGroupPropertyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyGroupPropertyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyGroupPropertyResponseBody) *ModifyGroupPropertyResponse
	GetBody() *ModifyGroupPropertyResponseBody
}

type ModifyGroupPropertyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyGroupPropertyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyGroupPropertyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyGroupPropertyResponse) GoString() string {
	return s.String()
}

func (s *ModifyGroupPropertyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyGroupPropertyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyGroupPropertyResponse) GetBody() *ModifyGroupPropertyResponseBody {
	return s.Body
}

func (s *ModifyGroupPropertyResponse) SetHeaders(v map[string]*string) *ModifyGroupPropertyResponse {
	s.Headers = v
	return s
}

func (s *ModifyGroupPropertyResponse) SetStatusCode(v int32) *ModifyGroupPropertyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyGroupPropertyResponse) SetBody(v *ModifyGroupPropertyResponseBody) *ModifyGroupPropertyResponse {
	s.Body = v
	return s
}

func (s *ModifyGroupPropertyResponse) Validate() error {
	return dara.Validate(s)
}
