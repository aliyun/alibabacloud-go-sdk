// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConcurrency(v int64) *CreateInstanceRequest
	GetConcurrency() *int64
	SetDescription(v string) *CreateInstanceRequest
	GetDescription() *string
	SetName(v string) *CreateInstanceRequest
	GetName() *string
	SetNluServiceParamsJson(v string) *CreateInstanceRequest
	GetNluServiceParamsJson() *string
	SetUnionInstanceId(v string) *CreateInstanceRequest
	GetUnionInstanceId() *string
	SetUnionSource(v string) *CreateInstanceRequest
	GetUnionSource() *string
}

type CreateInstanceRequest struct {
	// The maximum concurrency of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Concurrency *int64 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// 这是一个测试场景
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the Voice Navigator instance, which identifies the digital employee scenario.
	//
	// This parameter is required.
	//
	// example:
	//
	// 测试场景
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Configuration parameters for the large language model service, in JSON format.
	//
	// - Use this parameter to specify a Function Compute service.
	NluServiceParamsJson *string `json:"NluServiceParamsJson,omitempty" xml:"NluServiceParamsJson,omitempty"`
	// The ID of the source instance.
	//
	// > If you set UnionSource to CCC, set this parameter to the ID of the Cloud Contact Center instance.
	//
	// example:
	//
	// demo-lctms
	UnionInstanceId *string `json:"UnionInstanceId,omitempty" xml:"UnionInstanceId,omitempty"`
	// The source service.
	//
	// - CCC: Cloud Contact Center
	//
	// example:
	//
	// CCC
	UnionSource *string `json:"UnionSource,omitempty" xml:"UnionSource,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) GetConcurrency() *int64 {
	return s.Concurrency
}

func (s *CreateInstanceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateInstanceRequest) GetName() *string {
	return s.Name
}

func (s *CreateInstanceRequest) GetNluServiceParamsJson() *string {
	return s.NluServiceParamsJson
}

func (s *CreateInstanceRequest) GetUnionInstanceId() *string {
	return s.UnionInstanceId
}

func (s *CreateInstanceRequest) GetUnionSource() *string {
	return s.UnionSource
}

func (s *CreateInstanceRequest) SetConcurrency(v int64) *CreateInstanceRequest {
	s.Concurrency = &v
	return s
}

func (s *CreateInstanceRequest) SetDescription(v string) *CreateInstanceRequest {
	s.Description = &v
	return s
}

func (s *CreateInstanceRequest) SetName(v string) *CreateInstanceRequest {
	s.Name = &v
	return s
}

func (s *CreateInstanceRequest) SetNluServiceParamsJson(v string) *CreateInstanceRequest {
	s.NluServiceParamsJson = &v
	return s
}

func (s *CreateInstanceRequest) SetUnionInstanceId(v string) *CreateInstanceRequest {
	s.UnionInstanceId = &v
	return s
}

func (s *CreateInstanceRequest) SetUnionSource(v string) *CreateInstanceRequest {
	s.UnionSource = &v
	return s
}

func (s *CreateInstanceRequest) Validate() error {
	return dara.Validate(s)
}
