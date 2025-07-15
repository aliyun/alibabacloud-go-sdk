// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceResponseBody) *ModifyInstanceResponse
	GetBody() *ModifyInstanceResponseBody
}

type ModifyInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceResponse) GetBody() *ModifyInstanceResponseBody {
	return s.Body
}

func (s *ModifyInstanceResponse) SetHeaders(v map[string]*string) *ModifyInstanceResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceResponse) SetStatusCode(v int32) *ModifyInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceResponse) SetBody(v *ModifyInstanceResponseBody) *ModifyInstanceResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceResponse) Validate() error {
	return dara.Validate(s)
}
