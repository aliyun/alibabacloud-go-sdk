// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarAgentSessionStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSessionId(v string) *DescribePolarAgentSessionStatusRequest
	GetSessionId() *string
	SetSource(v string) *DescribePolarAgentSessionStatusRequest
	GetSource() *string
}

type DescribePolarAgentSessionStatusRequest struct {
	// The ID of the session. This ID is used to identify a visitor\\"s session and maintain context information.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1e28530a0da2c4755f165b1b8b9a73c9
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The source of the product. Set the value to polardb-console.
	//
	// example:
	//
	// polardb-console
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s DescribePolarAgentSessionStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarAgentSessionStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribePolarAgentSessionStatusRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *DescribePolarAgentSessionStatusRequest) GetSource() *string {
	return s.Source
}

func (s *DescribePolarAgentSessionStatusRequest) SetSessionId(v string) *DescribePolarAgentSessionStatusRequest {
	s.SessionId = &v
	return s
}

func (s *DescribePolarAgentSessionStatusRequest) SetSource(v string) *DescribePolarAgentSessionStatusRequest {
	s.Source = &v
	return s
}

func (s *DescribePolarAgentSessionStatusRequest) Validate() error {
	return dara.Validate(s)
}
