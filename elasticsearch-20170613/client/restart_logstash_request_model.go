// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartLogstashRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchCount(v float64) *RestartLogstashRequest
	GetBatchCount() *float64
	SetBlueGreenDep(v bool) *RestartLogstashRequest
	GetBlueGreenDep() *bool
	SetNodeTypes(v []*string) *RestartLogstashRequest
	GetNodeTypes() []*string
	SetNodes(v []*string) *RestartLogstashRequest
	GetNodes() []*string
	SetRestartType(v string) *RestartLogstashRequest
	GetRestartType() *string
	SetClientToken(v string) *RestartLogstashRequest
	GetClientToken() *string
	SetForce(v bool) *RestartLogstashRequest
	GetForce() *bool
}

type RestartLogstashRequest struct {
	// The concurrency for force restart. This parameter does not need to be set during a blue-green restart because force restart is not supported in that scenario.
	//
	// example:
	//
	// 20%
	BatchCount *float64 `json:"batchCount,omitempty" xml:"batchCount,omitempty"`
	// Specifies whether to perform a blue-green restart. Default value: false.
	//
	// example:
	//
	// false
	BlueGreenDep *bool `json:"blueGreenDep,omitempty" xml:"blueGreenDep,omitempty"`
	// The type of role node to restart. This parameter is not supported.
	NodeTypes []*string `json:"nodeTypes,omitempty" xml:"nodeTypes,omitempty" type:"Repeated"`
	// The node information selected when restarting nodes.
	Nodes []*string `json:"nodes,omitempty" xml:"nodes,omitempty" type:"Repeated"`
	// The restart type. Valid values:
	//
	// - instance: restart the instance
	//
	// - nodeIp: restart a node.
	//
	// example:
	//
	// instance
	RestartType *string `json:"restartType,omitempty" xml:"restartType,omitempty"`
	// A client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// Specifies whether to force restart the instance. Valid values:
	//
	// - true: force restart
	//
	// - false (default): do not force restart.
	//
	// example:
	//
	// true
	Force *bool `json:"force,omitempty" xml:"force,omitempty"`
}

func (s RestartLogstashRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartLogstashRequest) GoString() string {
	return s.String()
}

func (s *RestartLogstashRequest) GetBatchCount() *float64 {
	return s.BatchCount
}

func (s *RestartLogstashRequest) GetBlueGreenDep() *bool {
	return s.BlueGreenDep
}

func (s *RestartLogstashRequest) GetNodeTypes() []*string {
	return s.NodeTypes
}

func (s *RestartLogstashRequest) GetNodes() []*string {
	return s.Nodes
}

func (s *RestartLogstashRequest) GetRestartType() *string {
	return s.RestartType
}

func (s *RestartLogstashRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RestartLogstashRequest) GetForce() *bool {
	return s.Force
}

func (s *RestartLogstashRequest) SetBatchCount(v float64) *RestartLogstashRequest {
	s.BatchCount = &v
	return s
}

func (s *RestartLogstashRequest) SetBlueGreenDep(v bool) *RestartLogstashRequest {
	s.BlueGreenDep = &v
	return s
}

func (s *RestartLogstashRequest) SetNodeTypes(v []*string) *RestartLogstashRequest {
	s.NodeTypes = v
	return s
}

func (s *RestartLogstashRequest) SetNodes(v []*string) *RestartLogstashRequest {
	s.Nodes = v
	return s
}

func (s *RestartLogstashRequest) SetRestartType(v string) *RestartLogstashRequest {
	s.RestartType = &v
	return s
}

func (s *RestartLogstashRequest) SetClientToken(v string) *RestartLogstashRequest {
	s.ClientToken = &v
	return s
}

func (s *RestartLogstashRequest) SetForce(v bool) *RestartLogstashRequest {
	s.Force = &v
	return s
}

func (s *RestartLogstashRequest) Validate() error {
	return dara.Validate(s)
}
