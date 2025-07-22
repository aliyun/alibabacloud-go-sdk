// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMongoDBCurrentOpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterDoc(v string) *GetMongoDBCurrentOpRequest
	GetFilterDoc() *string
	SetInstanceId(v string) *GetMongoDBCurrentOpRequest
	GetInstanceId() *string
	SetNodeId(v string) *GetMongoDBCurrentOpRequest
	GetNodeId() *string
	SetRole(v string) *GetMongoDBCurrentOpRequest
	GetRole() *string
}

type GetMongoDBCurrentOpRequest struct {
	// The `db.currentOp()` command that is used to filter sessions. For more information, see [db.currentOp()](https://docs.mongodb.com/manual/reference/method/db.currentOp/) of MongoDB Documentation.
	//
	// example:
	//
	// { "active" : true }
	FilterDoc *string `json:"FilterDoc,omitempty" xml:"FilterDoc,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-uf608087********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The node ID.
	//
	// >  If you do not specify a node ID, the sessions of the primary node are queried by default.
	//
	// example:
	//
	// 23302531
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// A reserved parameter. You do not need to specify the parameter.
	//
	// example:
	//
	// None
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s GetMongoDBCurrentOpRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMongoDBCurrentOpRequest) GoString() string {
	return s.String()
}

func (s *GetMongoDBCurrentOpRequest) GetFilterDoc() *string {
	return s.FilterDoc
}

func (s *GetMongoDBCurrentOpRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetMongoDBCurrentOpRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *GetMongoDBCurrentOpRequest) GetRole() *string {
	return s.Role
}

func (s *GetMongoDBCurrentOpRequest) SetFilterDoc(v string) *GetMongoDBCurrentOpRequest {
	s.FilterDoc = &v
	return s
}

func (s *GetMongoDBCurrentOpRequest) SetInstanceId(v string) *GetMongoDBCurrentOpRequest {
	s.InstanceId = &v
	return s
}

func (s *GetMongoDBCurrentOpRequest) SetNodeId(v string) *GetMongoDBCurrentOpRequest {
	s.NodeId = &v
	return s
}

func (s *GetMongoDBCurrentOpRequest) SetRole(v string) *GetMongoDBCurrentOpRequest {
	s.Role = &v
	return s
}

func (s *GetMongoDBCurrentOpRequest) Validate() error {
	return dara.Validate(s)
}
