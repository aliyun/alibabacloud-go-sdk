// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAssistantCapabilityHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetAssistantCapabilityHeaders
	GetCommonHeaders() map[string]*string
	SetAccountId(v string) *GetAssistantCapabilityHeaders
	GetAccountId() *string
}

type GetAssistantCapabilityHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// 123456
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetAssistantCapabilityHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetAssistantCapabilityHeaders) GoString() string {
	return s.String()
}

func (s *GetAssistantCapabilityHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetAssistantCapabilityHeaders) GetAccountId() *string {
	return s.AccountId
}

func (s *GetAssistantCapabilityHeaders) SetCommonHeaders(v map[string]*string) *GetAssistantCapabilityHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAssistantCapabilityHeaders) SetAccountId(v string) *GetAssistantCapabilityHeaders {
	s.AccountId = &v
	return s
}

func (s *GetAssistantCapabilityHeaders) Validate() error {
	return dara.Validate(s)
}
