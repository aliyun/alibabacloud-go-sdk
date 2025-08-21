// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBlockStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyBlockStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyBlockStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifyBlockStatusResponseBody) *ModifyBlockStatusResponse
	GetBody() *ModifyBlockStatusResponseBody
}

type ModifyBlockStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyBlockStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyBlockStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyBlockStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyBlockStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyBlockStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyBlockStatusResponse) GetBody() *ModifyBlockStatusResponseBody {
	return s.Body
}

func (s *ModifyBlockStatusResponse) SetHeaders(v map[string]*string) *ModifyBlockStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyBlockStatusResponse) SetStatusCode(v int32) *ModifyBlockStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyBlockStatusResponse) SetBody(v *ModifyBlockStatusResponseBody) *ModifyBlockStatusResponse {
	s.Body = v
	return s
}

func (s *ModifyBlockStatusResponse) Validate() error {
	return dara.Validate(s)
}
