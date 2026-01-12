// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomAgentId(v string) *GetCustomAgentRequest
	GetCustomAgentId() *string
}

type GetCustomAgentRequest struct {
	// The operation that you want to perform. Set the value to **GetCustomAgent**.
	//
	// example:
	//
	// ebe44453-3b41-4c74-94d1-01d088d7xxxx
	CustomAgentId *string `json:"CustomAgentId,omitempty" xml:"CustomAgentId,omitempty"`
}

func (s GetCustomAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAgentRequest) GoString() string {
	return s.String()
}

func (s *GetCustomAgentRequest) GetCustomAgentId() *string {
	return s.CustomAgentId
}

func (s *GetCustomAgentRequest) SetCustomAgentId(v string) *GetCustomAgentRequest {
	s.CustomAgentId = &v
	return s
}

func (s *GetCustomAgentRequest) Validate() error {
	return dara.Validate(s)
}
