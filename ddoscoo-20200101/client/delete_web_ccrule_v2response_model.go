// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWebCCRuleV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWebCCRuleV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWebCCRuleV2Response
	GetStatusCode() *int32
	SetBody(v *DeleteWebCCRuleV2ResponseBody) *DeleteWebCCRuleV2Response
	GetBody() *DeleteWebCCRuleV2ResponseBody
}

type DeleteWebCCRuleV2Response struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWebCCRuleV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWebCCRuleV2Response) String() string {
	return dara.Prettify(s)
}

func (s DeleteWebCCRuleV2Response) GoString() string {
	return s.String()
}

func (s *DeleteWebCCRuleV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWebCCRuleV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWebCCRuleV2Response) GetBody() *DeleteWebCCRuleV2ResponseBody {
	return s.Body
}

func (s *DeleteWebCCRuleV2Response) SetHeaders(v map[string]*string) *DeleteWebCCRuleV2Response {
	s.Headers = v
	return s
}

func (s *DeleteWebCCRuleV2Response) SetStatusCode(v int32) *DeleteWebCCRuleV2Response {
	s.StatusCode = &v
	return s
}

func (s *DeleteWebCCRuleV2Response) SetBody(v *DeleteWebCCRuleV2ResponseBody) *DeleteWebCCRuleV2Response {
	s.Body = v
	return s
}

func (s *DeleteWebCCRuleV2Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
