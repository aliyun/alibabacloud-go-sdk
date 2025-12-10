// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImportAgentList(v *ListAgentResp) *ListAgentResponseBody
	GetImportAgentList() *ListAgentResp
}

type ListAgentResponseBody struct {
	// The details of the agents.
	ImportAgentList *ListAgentResp `json:"ImportAgentList,omitempty" xml:"ImportAgentList,omitempty"`
}

func (s ListAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgentResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentResponseBody) GetImportAgentList() *ListAgentResp {
	return s.ImportAgentList
}

func (s *ListAgentResponseBody) SetImportAgentList(v *ListAgentResp) *ListAgentResponseBody {
	s.ImportAgentList = v
	return s
}

func (s *ListAgentResponseBody) Validate() error {
	if s.ImportAgentList != nil {
		if err := s.ImportAgentList.Validate(); err != nil {
			return err
		}
	}
	return nil
}
