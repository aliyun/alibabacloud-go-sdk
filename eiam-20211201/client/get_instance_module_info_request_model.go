// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceModuleInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetInstanceModuleInfoRequest
	GetInstanceId() *string
	SetModuleKey(v string) *GetInstanceModuleInfoRequest
	GetModuleKey() *string
}

type GetInstanceModuleInfoRequest struct {
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 一级模块标识，必填
	//
	// This parameter is required.
	//
	// example:
	//
	// urn:alibaba:idaas:license:module:ud
	ModuleKey *string `json:"ModuleKey,omitempty" xml:"ModuleKey,omitempty"`
}

func (s GetInstanceModuleInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceModuleInfoRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceModuleInfoRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceModuleInfoRequest) GetModuleKey() *string {
	return s.ModuleKey
}

func (s *GetInstanceModuleInfoRequest) SetInstanceId(v string) *GetInstanceModuleInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceModuleInfoRequest) SetModuleKey(v string) *GetInstanceModuleInfoRequest {
	s.ModuleKey = &v
	return s
}

func (s *GetInstanceModuleInfoRequest) Validate() error {
	return dara.Validate(s)
}
