// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSasContainerWebDefenseRuleApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSasContainerWebDefenseRuleApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSasContainerWebDefenseRuleApplicationResponse
	GetStatusCode() *int32
	SetBody(v *GetSasContainerWebDefenseRuleApplicationResponseBody) *GetSasContainerWebDefenseRuleApplicationResponse
	GetBody() *GetSasContainerWebDefenseRuleApplicationResponseBody
}

type GetSasContainerWebDefenseRuleApplicationResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSasContainerWebDefenseRuleApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSasContainerWebDefenseRuleApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSasContainerWebDefenseRuleApplicationResponse) GoString() string {
	return s.String()
}

func (s *GetSasContainerWebDefenseRuleApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSasContainerWebDefenseRuleApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSasContainerWebDefenseRuleApplicationResponse) GetBody() *GetSasContainerWebDefenseRuleApplicationResponseBody {
	return s.Body
}

func (s *GetSasContainerWebDefenseRuleApplicationResponse) SetHeaders(v map[string]*string) *GetSasContainerWebDefenseRuleApplicationResponse {
	s.Headers = v
	return s
}

func (s *GetSasContainerWebDefenseRuleApplicationResponse) SetStatusCode(v int32) *GetSasContainerWebDefenseRuleApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSasContainerWebDefenseRuleApplicationResponse) SetBody(v *GetSasContainerWebDefenseRuleApplicationResponseBody) *GetSasContainerWebDefenseRuleApplicationResponse {
	s.Body = v
	return s
}

func (s *GetSasContainerWebDefenseRuleApplicationResponse) Validate() error {
	return dara.Validate(s)
}
