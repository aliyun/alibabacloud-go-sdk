// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceVswitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceVswitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceVswitchResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceVswitchResponseBody) *ModifyInstanceVswitchResponse
	GetBody() *ModifyInstanceVswitchResponseBody
}

type ModifyInstanceVswitchResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceVswitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceVswitchResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceVswitchResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceVswitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceVswitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceVswitchResponse) GetBody() *ModifyInstanceVswitchResponseBody {
	return s.Body
}

func (s *ModifyInstanceVswitchResponse) SetHeaders(v map[string]*string) *ModifyInstanceVswitchResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceVswitchResponse) SetStatusCode(v int32) *ModifyInstanceVswitchResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceVswitchResponse) SetBody(v *ModifyInstanceVswitchResponseBody) *ModifyInstanceVswitchResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceVswitchResponse) Validate() error {
	return dara.Validate(s)
}
