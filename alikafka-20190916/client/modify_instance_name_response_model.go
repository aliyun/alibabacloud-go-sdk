// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceNameResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceNameResponseBody) *ModifyInstanceNameResponse
	GetBody() *ModifyInstanceNameResponseBody
}

type ModifyInstanceNameResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceNameResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceNameResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceNameResponse) GetBody() *ModifyInstanceNameResponseBody {
	return s.Body
}

func (s *ModifyInstanceNameResponse) SetHeaders(v map[string]*string) *ModifyInstanceNameResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceNameResponse) SetStatusCode(v int32) *ModifyInstanceNameResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceNameResponse) SetBody(v *ModifyInstanceNameResponseBody) *ModifyInstanceNameResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceNameResponse) Validate() error {
	return dara.Validate(s)
}
