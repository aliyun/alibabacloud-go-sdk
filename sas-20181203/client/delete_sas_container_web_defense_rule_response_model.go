// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSasContainerWebDefenseRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSasContainerWebDefenseRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSasContainerWebDefenseRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSasContainerWebDefenseRuleResponseBody) *DeleteSasContainerWebDefenseRuleResponse
	GetBody() *DeleteSasContainerWebDefenseRuleResponseBody
}

type DeleteSasContainerWebDefenseRuleResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSasContainerWebDefenseRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSasContainerWebDefenseRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSasContainerWebDefenseRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteSasContainerWebDefenseRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSasContainerWebDefenseRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSasContainerWebDefenseRuleResponse) GetBody() *DeleteSasContainerWebDefenseRuleResponseBody {
	return s.Body
}

func (s *DeleteSasContainerWebDefenseRuleResponse) SetHeaders(v map[string]*string) *DeleteSasContainerWebDefenseRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteSasContainerWebDefenseRuleResponse) SetStatusCode(v int32) *DeleteSasContainerWebDefenseRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSasContainerWebDefenseRuleResponse) SetBody(v *DeleteSasContainerWebDefenseRuleResponseBody) *DeleteSasContainerWebDefenseRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteSasContainerWebDefenseRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
