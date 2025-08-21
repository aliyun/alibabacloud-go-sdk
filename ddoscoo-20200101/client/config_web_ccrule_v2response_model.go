// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigWebCCRuleV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigWebCCRuleV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigWebCCRuleV2Response
	GetStatusCode() *int32
	SetBody(v *ConfigWebCCRuleV2ResponseBody) *ConfigWebCCRuleV2Response
	GetBody() *ConfigWebCCRuleV2ResponseBody
}

type ConfigWebCCRuleV2Response struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigWebCCRuleV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigWebCCRuleV2Response) String() string {
	return dara.Prettify(s)
}

func (s ConfigWebCCRuleV2Response) GoString() string {
	return s.String()
}

func (s *ConfigWebCCRuleV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigWebCCRuleV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigWebCCRuleV2Response) GetBody() *ConfigWebCCRuleV2ResponseBody {
	return s.Body
}

func (s *ConfigWebCCRuleV2Response) SetHeaders(v map[string]*string) *ConfigWebCCRuleV2Response {
	s.Headers = v
	return s
}

func (s *ConfigWebCCRuleV2Response) SetStatusCode(v int32) *ConfigWebCCRuleV2Response {
	s.StatusCode = &v
	return s
}

func (s *ConfigWebCCRuleV2Response) SetBody(v *ConfigWebCCRuleV2ResponseBody) *ConfigWebCCRuleV2Response {
	s.Body = v
	return s
}

func (s *ConfigWebCCRuleV2Response) Validate() error {
	return dara.Validate(s)
}
