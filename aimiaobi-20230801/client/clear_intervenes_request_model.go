// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearIntervenesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ClearIntervenesRequest
	GetAgentKey() *string
}

type ClearIntervenesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
}

func (s ClearIntervenesRequest) String() string {
	return dara.Prettify(s)
}

func (s ClearIntervenesRequest) GoString() string {
	return s.String()
}

func (s *ClearIntervenesRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ClearIntervenesRequest) SetAgentKey(v string) *ClearIntervenesRequest {
	s.AgentKey = &v
	return s
}

func (s *ClearIntervenesRequest) Validate() error {
	return dara.Validate(s)
}
