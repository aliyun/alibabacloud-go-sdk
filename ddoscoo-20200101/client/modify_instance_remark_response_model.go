// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceRemarkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceRemarkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceRemarkResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceRemarkResponseBody) *ModifyInstanceRemarkResponse
	GetBody() *ModifyInstanceRemarkResponseBody
}

type ModifyInstanceRemarkResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceRemarkResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceRemarkResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceRemarkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceRemarkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceRemarkResponse) GetBody() *ModifyInstanceRemarkResponseBody {
	return s.Body
}

func (s *ModifyInstanceRemarkResponse) SetHeaders(v map[string]*string) *ModifyInstanceRemarkResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceRemarkResponse) SetStatusCode(v int32) *ModifyInstanceRemarkResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceRemarkResponse) SetBody(v *ModifyInstanceRemarkResponseBody) *ModifyInstanceRemarkResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceRemarkResponse) Validate() error {
	return dara.Validate(s)
}
