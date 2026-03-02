// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStackDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetStackDetailRequest
	GetEndTime() *string
	SetEnv(v string) *GetStackDetailRequest
	GetEnv() *string
	SetRpcId(v string) *GetStackDetailRequest
	GetRpcId() *string
	SetServiceGroupId(v int64) *GetStackDetailRequest
	GetServiceGroupId() *int64
	SetServiceName(v string) *GetStackDetailRequest
	GetServiceName() *string
	SetStartTime(v string) *GetStackDetailRequest
	GetStartTime() *string
}

type GetStackDetailRequest struct {
	// example:
	//
	// 2022-11-08 15:03:22
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// TEST
	Env *string `json:"env,omitempty" xml:"env,omitempty"`
	// example:
	//
	// 12
	RpcId          *string `json:"rpcId,omitempty" xml:"rpcId,omitempty"`
	ServiceGroupId *int64  `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
	// example:
	//
	// feishu-attendance-app
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	// example:
	//
	// 2022-12-06 10:24:44
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetStackDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStackDetailRequest) GoString() string {
	return s.String()
}

func (s *GetStackDetailRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetStackDetailRequest) GetEnv() *string {
	return s.Env
}

func (s *GetStackDetailRequest) GetRpcId() *string {
	return s.RpcId
}

func (s *GetStackDetailRequest) GetServiceGroupId() *int64 {
	return s.ServiceGroupId
}

func (s *GetStackDetailRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetStackDetailRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetStackDetailRequest) SetEndTime(v string) *GetStackDetailRequest {
	s.EndTime = &v
	return s
}

func (s *GetStackDetailRequest) SetEnv(v string) *GetStackDetailRequest {
	s.Env = &v
	return s
}

func (s *GetStackDetailRequest) SetRpcId(v string) *GetStackDetailRequest {
	s.RpcId = &v
	return s
}

func (s *GetStackDetailRequest) SetServiceGroupId(v int64) *GetStackDetailRequest {
	s.ServiceGroupId = &v
	return s
}

func (s *GetStackDetailRequest) SetServiceName(v string) *GetStackDetailRequest {
	s.ServiceName = &v
	return s
}

func (s *GetStackDetailRequest) SetStartTime(v string) *GetStackDetailRequest {
	s.StartTime = &v
	return s
}

func (s *GetStackDetailRequest) Validate() error {
	return dara.Validate(s)
}
