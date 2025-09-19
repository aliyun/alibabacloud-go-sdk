// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSasContainerWebDefenseRuleCriteriaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSasContainerWebDefenseRuleCriteriaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSasContainerWebDefenseRuleCriteriaResponse
	GetStatusCode() *int32
	SetBody(v *GetSasContainerWebDefenseRuleCriteriaResponseBody) *GetSasContainerWebDefenseRuleCriteriaResponse
	GetBody() *GetSasContainerWebDefenseRuleCriteriaResponseBody
}

type GetSasContainerWebDefenseRuleCriteriaResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSasContainerWebDefenseRuleCriteriaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSasContainerWebDefenseRuleCriteriaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSasContainerWebDefenseRuleCriteriaResponse) GoString() string {
	return s.String()
}

func (s *GetSasContainerWebDefenseRuleCriteriaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSasContainerWebDefenseRuleCriteriaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSasContainerWebDefenseRuleCriteriaResponse) GetBody() *GetSasContainerWebDefenseRuleCriteriaResponseBody {
	return s.Body
}

func (s *GetSasContainerWebDefenseRuleCriteriaResponse) SetHeaders(v map[string]*string) *GetSasContainerWebDefenseRuleCriteriaResponse {
	s.Headers = v
	return s
}

func (s *GetSasContainerWebDefenseRuleCriteriaResponse) SetStatusCode(v int32) *GetSasContainerWebDefenseRuleCriteriaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSasContainerWebDefenseRuleCriteriaResponse) SetBody(v *GetSasContainerWebDefenseRuleCriteriaResponseBody) *GetSasContainerWebDefenseRuleCriteriaResponse {
	s.Body = v
	return s
}

func (s *GetSasContainerWebDefenseRuleCriteriaResponse) Validate() error {
	return dara.Validate(s)
}
