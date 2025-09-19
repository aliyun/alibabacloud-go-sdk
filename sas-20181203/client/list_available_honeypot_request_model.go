// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvailableHoneypotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeId(v string) *ListAvailableHoneypotRequest
	GetNodeId() *string
}

type ListAvailableHoneypotRequest struct {
	// The ID of the management node to which the honeypot is deployed.
	//
	// example:
	//
	// 4341018b-8e01-43f6-b1d2-af29a2a4****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s ListAvailableHoneypotRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableHoneypotRequest) GoString() string {
	return s.String()
}

func (s *ListAvailableHoneypotRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *ListAvailableHoneypotRequest) SetNodeId(v string) *ListAvailableHoneypotRequest {
	s.NodeId = &v
	return s
}

func (s *ListAvailableHoneypotRequest) Validate() error {
	return dara.Validate(s)
}
