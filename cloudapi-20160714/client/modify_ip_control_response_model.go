// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIpControlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyIpControlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyIpControlResponse
	GetStatusCode() *int32
	SetBody(v *ModifyIpControlResponseBody) *ModifyIpControlResponse
	GetBody() *ModifyIpControlResponseBody
}

type ModifyIpControlResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyIpControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyIpControlResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyIpControlResponse) GoString() string {
	return s.String()
}

func (s *ModifyIpControlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyIpControlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyIpControlResponse) GetBody() *ModifyIpControlResponseBody {
	return s.Body
}

func (s *ModifyIpControlResponse) SetHeaders(v map[string]*string) *ModifyIpControlResponse {
	s.Headers = v
	return s
}

func (s *ModifyIpControlResponse) SetStatusCode(v int32) *ModifyIpControlResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyIpControlResponse) SetBody(v *ModifyIpControlResponseBody) *ModifyIpControlResponse {
	s.Body = v
	return s
}

func (s *ModifyIpControlResponse) Validate() error {
	return dara.Validate(s)
}
