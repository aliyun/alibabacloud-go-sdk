// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerNetworkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionType(v string) *TriggerNetworkRequest
	GetActionType() *string
	SetNetworkType(v string) *TriggerNetworkRequest
	GetNetworkType() *string
	SetNodeType(v string) *TriggerNetworkRequest
	GetNodeType() *string
	SetClientToken(v string) *TriggerNetworkRequest
	GetClientToken() *string
}

type TriggerNetworkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// OPEN
	ActionType *string `json:"actionType,omitempty" xml:"actionType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PUBLIC
	NetworkType *string `json:"networkType,omitempty" xml:"networkType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// KIBANA
	NodeType *string `json:"nodeType,omitempty" xml:"nodeType,omitempty"`
	// example:
	//
	// 407d02b74c49beb5bfdac7ec8bde2488
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s TriggerNetworkRequest) String() string {
	return dara.Prettify(s)
}

func (s TriggerNetworkRequest) GoString() string {
	return s.String()
}

func (s *TriggerNetworkRequest) GetActionType() *string {
	return s.ActionType
}

func (s *TriggerNetworkRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *TriggerNetworkRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *TriggerNetworkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *TriggerNetworkRequest) SetActionType(v string) *TriggerNetworkRequest {
	s.ActionType = &v
	return s
}

func (s *TriggerNetworkRequest) SetNetworkType(v string) *TriggerNetworkRequest {
	s.NetworkType = &v
	return s
}

func (s *TriggerNetworkRequest) SetNodeType(v string) *TriggerNetworkRequest {
	s.NodeType = &v
	return s
}

func (s *TriggerNetworkRequest) SetClientToken(v string) *TriggerNetworkRequest {
	s.ClientToken = &v
	return s
}

func (s *TriggerNetworkRequest) Validate() error {
	return dara.Validate(s)
}
