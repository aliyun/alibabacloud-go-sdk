// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFullLogTtlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyFullLogTtlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyFullLogTtlResponse
	GetStatusCode() *int32
	SetBody(v *ModifyFullLogTtlResponseBody) *ModifyFullLogTtlResponse
	GetBody() *ModifyFullLogTtlResponseBody
}

type ModifyFullLogTtlResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyFullLogTtlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyFullLogTtlResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyFullLogTtlResponse) GoString() string {
	return s.String()
}

func (s *ModifyFullLogTtlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyFullLogTtlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyFullLogTtlResponse) GetBody() *ModifyFullLogTtlResponseBody {
	return s.Body
}

func (s *ModifyFullLogTtlResponse) SetHeaders(v map[string]*string) *ModifyFullLogTtlResponse {
	s.Headers = v
	return s
}

func (s *ModifyFullLogTtlResponse) SetStatusCode(v int32) *ModifyFullLogTtlResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyFullLogTtlResponse) SetBody(v *ModifyFullLogTtlResponseBody) *ModifyFullLogTtlResponse {
	s.Body = v
	return s
}

func (s *ModifyFullLogTtlResponse) Validate() error {
	return dara.Validate(s)
}
