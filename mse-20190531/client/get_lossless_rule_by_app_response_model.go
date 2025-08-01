// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLosslessRuleByAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLosslessRuleByAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLosslessRuleByAppResponse
	GetStatusCode() *int32
	SetBody(v *GetLosslessRuleByAppResponseBody) *GetLosslessRuleByAppResponse
	GetBody() *GetLosslessRuleByAppResponseBody
}

type GetLosslessRuleByAppResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLosslessRuleByAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLosslessRuleByAppResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLosslessRuleByAppResponse) GoString() string {
	return s.String()
}

func (s *GetLosslessRuleByAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLosslessRuleByAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLosslessRuleByAppResponse) GetBody() *GetLosslessRuleByAppResponseBody {
	return s.Body
}

func (s *GetLosslessRuleByAppResponse) SetHeaders(v map[string]*string) *GetLosslessRuleByAppResponse {
	s.Headers = v
	return s
}

func (s *GetLosslessRuleByAppResponse) SetStatusCode(v int32) *GetLosslessRuleByAppResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLosslessRuleByAppResponse) SetBody(v *GetLosslessRuleByAppResponseBody) *GetLosslessRuleByAppResponse {
	s.Body = v
	return s
}

func (s *GetLosslessRuleByAppResponse) Validate() error {
	return dara.Validate(s)
}
