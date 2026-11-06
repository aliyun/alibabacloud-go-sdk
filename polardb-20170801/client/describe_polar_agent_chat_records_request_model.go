// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarAgentChatRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSessionId(v string) *DescribePolarAgentChatRecordsRequest
	GetSessionId() *string
	SetSource(v string) *DescribePolarAgentChatRecordsRequest
	GetSource() *string
}

type DescribePolarAgentChatRecordsRequest struct {
	// The session ID. To obtain the SessionId, use Signature Version 3.0 to call GetPolarAgent to initiate a conversation and extract the SessionId from the returned stream.
	//
	// This parameter is required.
	//
	// example:
	//
	// 44dcdf31-04cd-4a44-9bae-834dd6657e29
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The product type source. Valid value: polardb-console.
	//
	// example:
	//
	// polardb-console
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s DescribePolarAgentChatRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarAgentChatRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribePolarAgentChatRecordsRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *DescribePolarAgentChatRecordsRequest) GetSource() *string {
	return s.Source
}

func (s *DescribePolarAgentChatRecordsRequest) SetSessionId(v string) *DescribePolarAgentChatRecordsRequest {
	s.SessionId = &v
	return s
}

func (s *DescribePolarAgentChatRecordsRequest) SetSource(v string) *DescribePolarAgentChatRecordsRequest {
	s.Source = &v
	return s
}

func (s *DescribePolarAgentChatRecordsRequest) Validate() error {
	return dara.Validate(s)
}
