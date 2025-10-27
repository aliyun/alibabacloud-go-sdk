// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSasContainerWebDefenseRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSasContainerWebDefenseRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSasContainerWebDefenseRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetSasContainerWebDefenseRuleResponseBody) *GetSasContainerWebDefenseRuleResponse
	GetBody() *GetSasContainerWebDefenseRuleResponseBody
}

type GetSasContainerWebDefenseRuleResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSasContainerWebDefenseRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSasContainerWebDefenseRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSasContainerWebDefenseRuleResponse) GoString() string {
	return s.String()
}

func (s *GetSasContainerWebDefenseRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSasContainerWebDefenseRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSasContainerWebDefenseRuleResponse) GetBody() *GetSasContainerWebDefenseRuleResponseBody {
	return s.Body
}

func (s *GetSasContainerWebDefenseRuleResponse) SetHeaders(v map[string]*string) *GetSasContainerWebDefenseRuleResponse {
	s.Headers = v
	return s
}

func (s *GetSasContainerWebDefenseRuleResponse) SetStatusCode(v int32) *GetSasContainerWebDefenseRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSasContainerWebDefenseRuleResponse) SetBody(v *GetSasContainerWebDefenseRuleResponseBody) *GetSasContainerWebDefenseRuleResponse {
	s.Body = v
	return s
}

func (s *GetSasContainerWebDefenseRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
