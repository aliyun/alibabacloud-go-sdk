// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsureOrderApplyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InsureOrderApplyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InsureOrderApplyResponse
	GetStatusCode() *int32
	SetBody(v *InsureOrderApplyResponseBody) *InsureOrderApplyResponse
	GetBody() *InsureOrderApplyResponseBody
}

type InsureOrderApplyResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsureOrderApplyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsureOrderApplyResponse) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderApplyResponse) GoString() string {
	return s.String()
}

func (s *InsureOrderApplyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InsureOrderApplyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InsureOrderApplyResponse) GetBody() *InsureOrderApplyResponseBody {
	return s.Body
}

func (s *InsureOrderApplyResponse) SetHeaders(v map[string]*string) *InsureOrderApplyResponse {
	s.Headers = v
	return s
}

func (s *InsureOrderApplyResponse) SetStatusCode(v int32) *InsureOrderApplyResponse {
	s.StatusCode = &v
	return s
}

func (s *InsureOrderApplyResponse) SetBody(v *InsureOrderApplyResponseBody) *InsureOrderApplyResponse {
	s.Body = v
	return s
}

func (s *InsureOrderApplyResponse) Validate() error {
	return dara.Validate(s)
}
