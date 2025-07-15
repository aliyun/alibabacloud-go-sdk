// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPropertiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *GetPropertiesRequest
	GetAgentKey() *string
}

type GetPropertiesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxx_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
}

func (s GetPropertiesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPropertiesRequest) GoString() string {
	return s.String()
}

func (s *GetPropertiesRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *GetPropertiesRequest) SetAgentKey(v string) *GetPropertiesRequest {
	s.AgentKey = &v
	return s
}

func (s *GetPropertiesRequest) Validate() error {
	return dara.Validate(s)
}
