// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostRuleStatusChangeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PostRuleStatusChangeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PostRuleStatusChangeResponse
	GetStatusCode() *int32
	SetBody(v *PostRuleStatusChangeResponseBody) *PostRuleStatusChangeResponse
	GetBody() *PostRuleStatusChangeResponseBody
}

type PostRuleStatusChangeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PostRuleStatusChangeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PostRuleStatusChangeResponse) String() string {
	return dara.Prettify(s)
}

func (s PostRuleStatusChangeResponse) GoString() string {
	return s.String()
}

func (s *PostRuleStatusChangeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PostRuleStatusChangeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PostRuleStatusChangeResponse) GetBody() *PostRuleStatusChangeResponseBody {
	return s.Body
}

func (s *PostRuleStatusChangeResponse) SetHeaders(v map[string]*string) *PostRuleStatusChangeResponse {
	s.Headers = v
	return s
}

func (s *PostRuleStatusChangeResponse) SetStatusCode(v int32) *PostRuleStatusChangeResponse {
	s.StatusCode = &v
	return s
}

func (s *PostRuleStatusChangeResponse) SetBody(v *PostRuleStatusChangeResponseBody) *PostRuleStatusChangeResponse {
	s.Body = v
	return s
}

func (s *PostRuleStatusChangeResponse) Validate() error {
	return dara.Validate(s)
}
