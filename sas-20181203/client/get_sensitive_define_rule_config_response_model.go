// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSensitiveDefineRuleConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSensitiveDefineRuleConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSensitiveDefineRuleConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetSensitiveDefineRuleConfigResponseBody) *GetSensitiveDefineRuleConfigResponse
	GetBody() *GetSensitiveDefineRuleConfigResponseBody
}

type GetSensitiveDefineRuleConfigResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSensitiveDefineRuleConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSensitiveDefineRuleConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSensitiveDefineRuleConfigResponse) GoString() string {
	return s.String()
}

func (s *GetSensitiveDefineRuleConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSensitiveDefineRuleConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSensitiveDefineRuleConfigResponse) GetBody() *GetSensitiveDefineRuleConfigResponseBody {
	return s.Body
}

func (s *GetSensitiveDefineRuleConfigResponse) SetHeaders(v map[string]*string) *GetSensitiveDefineRuleConfigResponse {
	s.Headers = v
	return s
}

func (s *GetSensitiveDefineRuleConfigResponse) SetStatusCode(v int32) *GetSensitiveDefineRuleConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSensitiveDefineRuleConfigResponse) SetBody(v *GetSensitiveDefineRuleConfigResponseBody) *GetSensitiveDefineRuleConfigResponse {
	s.Body = v
	return s
}

func (s *GetSensitiveDefineRuleConfigResponse) Validate() error {
	return dara.Validate(s)
}
