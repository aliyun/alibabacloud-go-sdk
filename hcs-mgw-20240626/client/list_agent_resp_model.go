// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentResp interface {
	dara.Model
	String() string
	GoString() string
	SetImportAgent(v []*GetAgentResp) *ListAgentResp
	GetImportAgent() []*GetAgentResp
	SetNextMarker(v string) *ListAgentResp
	GetNextMarker() *string
	SetTruncated(v bool) *ListAgentResp
	GetTruncated() *bool
}

type ListAgentResp struct {
	ImportAgent []*GetAgentResp `json:"ImportAgent,omitempty" xml:"ImportAgent,omitempty" type:"Repeated"`
	// example:
	//
	// <your-next-agent-name>
	NextMarker *string `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	// example:
	//
	// true
	Truncated *bool `json:"Truncated,omitempty" xml:"Truncated,omitempty"`
}

func (s ListAgentResp) String() string {
	return dara.Prettify(s)
}

func (s ListAgentResp) GoString() string {
	return s.String()
}

func (s *ListAgentResp) GetImportAgent() []*GetAgentResp {
	return s.ImportAgent
}

func (s *ListAgentResp) GetNextMarker() *string {
	return s.NextMarker
}

func (s *ListAgentResp) GetTruncated() *bool {
	return s.Truncated
}

func (s *ListAgentResp) SetImportAgent(v []*GetAgentResp) *ListAgentResp {
	s.ImportAgent = v
	return s
}

func (s *ListAgentResp) SetNextMarker(v string) *ListAgentResp {
	s.NextMarker = &v
	return s
}

func (s *ListAgentResp) SetTruncated(v bool) *ListAgentResp {
	s.Truncated = &v
	return s
}

func (s *ListAgentResp) Validate() error {
	if s.ImportAgent != nil {
		for _, item := range s.ImportAgent {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
