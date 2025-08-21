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
	// example:
	//
	// 20%
	BatchCount *float64 `json:"batchCount,omitempty" xml:"batchCount,omitempty"`
	// example:
	//
	// false
	BlueGreenDep *bool     `json:"blueGreenDep,omitempty" xml:"blueGreenDep,omitempty"`
	NodeTypes    []*string `json:"nodeTypes,omitempty" xml:"nodeTypes,omitempty" type:"Repeated"`
	Nodes        []*string `json:"nodes,omitempty" xml:"nodes,omitempty" type:"Repeated"`
	// example:
	//
	// instance
	RestartType *string `json:"restartType,omitempty" xml:"restartType,omitempty"`
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
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
