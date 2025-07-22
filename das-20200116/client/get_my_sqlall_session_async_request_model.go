// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMySQLAllSessionAsyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetMySQLAllSessionAsyncRequest
	GetInstanceId() *string
	SetNodeId(v string) *GetMySQLAllSessionAsyncRequest
	GetNodeId() *string
	SetResultId(v string) *GetMySQLAllSessionAsyncRequest
	GetResultId() *string
}

type GetMySQLAllSessionAsyncRequest struct {
	// The instance ID.
	//
	// >  Only ApsaraDB RDS for MySQL, PolarDB for MySQL, and PolarDB-X 2.0 instances are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze8g2am97624****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The node ID.
	//
	// >  You must specify this parameter for PolarDB for MySQL clusters. If you do not specify a node ID, the session data of the primary node is returned by default.
	//
	// example:
	//
	// pi-wz954ryd8f893****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The asynchronous request ID.
	//
	// >  GetMySQLAllSessionAsync is an asynchronous operation. After a request is sent, the system does not return complete results but returns a **request ID**. You need to use the **request ID*	- to initiate requests until the value of the **isFinish*	- field in the returned results is **true**, the complete results are returned. This indicates that to obtain complete data, you must call this operation at least twice.
	//
	// example:
	//
	// async__507044db6c4eadfa2dab9b084e80****
	ResultId *string `json:"ResultId,omitempty" xml:"ResultId,omitempty"`
}

func (s GetMySQLAllSessionAsyncRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMySQLAllSessionAsyncRequest) GoString() string {
	return s.String()
}

func (s *GetMySQLAllSessionAsyncRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetMySQLAllSessionAsyncRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *GetMySQLAllSessionAsyncRequest) GetResultId() *string {
	return s.ResultId
}

func (s *GetMySQLAllSessionAsyncRequest) SetInstanceId(v string) *GetMySQLAllSessionAsyncRequest {
	s.InstanceId = &v
	return s
}

func (s *GetMySQLAllSessionAsyncRequest) SetNodeId(v string) *GetMySQLAllSessionAsyncRequest {
	s.NodeId = &v
	return s
}

func (s *GetMySQLAllSessionAsyncRequest) SetResultId(v string) *GetMySQLAllSessionAsyncRequest {
	s.ResultId = &v
	return s
}

func (s *GetMySQLAllSessionAsyncRequest) Validate() error {
	return dara.Validate(s)
}
