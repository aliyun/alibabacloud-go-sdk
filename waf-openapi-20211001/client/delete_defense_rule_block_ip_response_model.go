// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDefenseRuleBlockIpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDefenseRuleBlockIpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDefenseRuleBlockIpResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDefenseRuleBlockIpResponseBody) *DeleteDefenseRuleBlockIpResponse
	GetBody() *DeleteDefenseRuleBlockIpResponseBody
}

type DeleteDefenseRuleBlockIpResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDefenseRuleBlockIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDefenseRuleBlockIpResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDefenseRuleBlockIpResponse) GoString() string {
	return s.String()
}

func (s *DeleteDefenseRuleBlockIpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDefenseRuleBlockIpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDefenseRuleBlockIpResponse) GetBody() *DeleteDefenseRuleBlockIpResponseBody {
	return s.Body
}

func (s *DeleteDefenseRuleBlockIpResponse) SetHeaders(v map[string]*string) *DeleteDefenseRuleBlockIpResponse {
	s.Headers = v
	return s
}

func (s *DeleteDefenseRuleBlockIpResponse) SetStatusCode(v int32) *DeleteDefenseRuleBlockIpResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDefenseRuleBlockIpResponse) SetBody(v *DeleteDefenseRuleBlockIpResponseBody) *DeleteDefenseRuleBlockIpResponse {
	s.Body = v
	return s
}

func (s *DeleteDefenseRuleBlockIpResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
