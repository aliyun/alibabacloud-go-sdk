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
	// The session ID, which is used to identify a visitor\\"s session and maintain context information. Usage notes: You can obtain this ID by calling the GetPolarAgent API and retrieving the SessionId field from the response. Note that GetPolarAgent requires Signature Version 3.0 and the SSE protocol. The ID must be in the standard UUID format, such as 1e28530a-0da2-c475-5f16-5b1b8b9a73c9.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1e28530a-0da2-c475-5f16-5b1b8b9a73c9
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The product source. Valid values: polardb-console.
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
