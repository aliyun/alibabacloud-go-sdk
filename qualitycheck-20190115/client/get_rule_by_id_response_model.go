// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRuleByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRuleByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRuleByIdResponse
	GetStatusCode() *int32
	SetBody(v *GetRuleByIdResponseBody) *GetRuleByIdResponse
	GetBody() *GetRuleByIdResponseBody
}

type GetRuleByIdResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRuleByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRuleByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRuleByIdResponse) GoString() string {
	return s.String()
}

func (s *GetRuleByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRuleByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRuleByIdResponse) GetBody() *GetRuleByIdResponseBody {
	return s.Body
}

func (s *GetRuleByIdResponse) SetHeaders(v map[string]*string) *GetRuleByIdResponse {
	s.Headers = v
	return s
}

func (s *GetRuleByIdResponse) SetStatusCode(v int32) *GetRuleByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRuleByIdResponse) SetBody(v *GetRuleByIdResponseBody) *GetRuleByIdResponse {
	s.Body = v
	return s
}

func (s *GetRuleByIdResponse) Validate() error {
	return dara.Validate(s)
}
