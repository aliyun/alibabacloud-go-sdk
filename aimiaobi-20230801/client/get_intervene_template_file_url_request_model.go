// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInterveneTemplateFileUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *GetInterveneTemplateFileUrlRequest
	GetAgentKey() *string
}

type GetInterveneTemplateFileUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c160c841c8e54295bf2f441432785944_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
}

func (s GetInterveneTemplateFileUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInterveneTemplateFileUrlRequest) GoString() string {
	return s.String()
}

func (s *GetInterveneTemplateFileUrlRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *GetInterveneTemplateFileUrlRequest) SetAgentKey(v string) *GetInterveneTemplateFileUrlRequest {
	s.AgentKey = &v
	return s
}

func (s *GetInterveneTemplateFileUrlRequest) Validate() error {
	return dara.Validate(s)
}
