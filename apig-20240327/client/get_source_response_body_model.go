// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSourceResponseBody
	GetCode() *string
	SetData(v *GetSourceResponseBodyData) *GetSourceResponseBody
	GetData() *GetSourceResponseBodyData
	SetMessage(v string) *GetSourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSourceResponseBody
	GetRequestId() *string
}

type GetSourceResponseBody struct {
	// example:
	//
	// Ok
	Code *string                    `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetSourceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// DE10E3C0-A676-5169-812D-6610AACBFAFF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetSourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSourceResponseBody) GetData() *GetSourceResponseBodyData {
	return s.Data
}

func (s *GetSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSourceResponseBody) SetCode(v string) *GetSourceResponseBody {
	s.Code = &v
	return s
}

func (s *GetSourceResponseBody) SetData(v *GetSourceResponseBodyData) *GetSourceResponseBody {
	s.Data = v
	return s
}

func (s *GetSourceResponseBody) SetMessage(v string) *GetSourceResponseBody {
	s.Message = &v
	return s
}

func (s *GetSourceResponseBody) SetRequestId(v string) *GetSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSourceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSourceResponseBodyData struct {
	// example:
	//
	// 1719386834548
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// example:
	//
	// gw-cq2vundlhtg***
	GatewayId       *string                                   `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	K8SSourceInfo   *GetSourceResponseBodyDataK8SSourceInfo   `json:"k8SSourceInfo,omitempty" xml:"k8SSourceInfo,omitempty" type:"Struct"`
	NacosSourceInfo *GetSourceResponseBodyDataNacosSourceInfo `json:"nacosSourceInfo,omitempty" xml:"nacosSourceInfo,omitempty" type:"Struct"`
	// example:
	//
	// import-test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// rg-aekz4us4iruleja
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// example:
	//
	// src-crdddallhtgt***
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// example:
	//
	// K8S
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 1719386834548
	UpdateTimestamp *int64 `json:"updateTimestamp,omitempty" xml:"updateTimestamp,omitempty"`
}

func (s GetSourceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSourceResponseBodyData) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *GetSourceResponseBodyData) GetGatewayId() *string {
	return s.GatewayId
}

func (s *GetSourceResponseBodyData) GetK8SSourceInfo() *GetSourceResponseBodyDataK8SSourceInfo {
	return s.K8SSourceInfo
}

func (s *GetSourceResponseBodyData) GetNacosSourceInfo() *GetSourceResponseBodyDataNacosSourceInfo {
	return s.NacosSourceInfo
}

func (s *GetSourceResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetSourceResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetSourceResponseBodyData) GetSourceId() *string {
	return s.SourceId
}

func (s *GetSourceResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetSourceResponseBodyData) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *GetSourceResponseBodyData) SetCreateTimestamp(v int64) *GetSourceResponseBodyData {
	s.CreateTimestamp = &v
	return s
}

func (s *GetSourceResponseBodyData) SetGatewayId(v string) *GetSourceResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *GetSourceResponseBodyData) SetK8SSourceInfo(v *GetSourceResponseBodyDataK8SSourceInfo) *GetSourceResponseBodyData {
	s.K8SSourceInfo = v
	return s
}

func (s *GetSourceResponseBodyData) SetNacosSourceInfo(v *GetSourceResponseBodyDataNacosSourceInfo) *GetSourceResponseBodyData {
	s.NacosSourceInfo = v
	return s
}

func (s *GetSourceResponseBodyData) SetName(v string) *GetSourceResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetSourceResponseBodyData) SetResourceGroupId(v string) *GetSourceResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *GetSourceResponseBodyData) SetSourceId(v string) *GetSourceResponseBodyData {
	s.SourceId = &v
	return s
}

func (s *GetSourceResponseBodyData) SetType(v string) *GetSourceResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetSourceResponseBodyData) SetUpdateTimestamp(v int64) *GetSourceResponseBodyData {
	s.UpdateTimestamp = &v
	return s
}

func (s *GetSourceResponseBodyData) Validate() error {
	if s.K8SSourceInfo != nil {
		if err := s.K8SSourceInfo.Validate(); err != nil {
			return err
		}
	}
	if s.NacosSourceInfo != nil {
		if err := s.NacosSourceInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSourceResponseBodyDataK8SSourceInfo struct {
	// example:
	//
	// c2d290b2d8b5d4935864cace5f0173f31
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
}

func (s GetSourceResponseBodyDataK8SSourceInfo) String() string {
	return dara.Prettify(s)
}

func (s GetSourceResponseBodyDataK8SSourceInfo) GoString() string {
	return s.String()
}

func (s *GetSourceResponseBodyDataK8SSourceInfo) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetSourceResponseBodyDataK8SSourceInfo) SetClusterId(v string) *GetSourceResponseBodyDataK8SSourceInfo {
	s.ClusterId = &v
	return s
}

func (s *GetSourceResponseBodyDataK8SSourceInfo) Validate() error {
	return dara.Validate(s)
}

type GetSourceResponseBodyDataNacosSourceInfo struct {
	// example:
	//
	// mse
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// example:
	//
	// fluss-cn-w7k4hann601
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	// example:
	//
	// hgprecn-cn-cfn47q7oh001
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s GetSourceResponseBodyDataNacosSourceInfo) String() string {
	return dara.Prettify(s)
}

func (s GetSourceResponseBodyDataNacosSourceInfo) GoString() string {
	return s.String()
}

func (s *GetSourceResponseBodyDataNacosSourceInfo) GetAddress() *string {
	return s.Address
}

func (s *GetSourceResponseBodyDataNacosSourceInfo) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetSourceResponseBodyDataNacosSourceInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSourceResponseBodyDataNacosSourceInfo) SetAddress(v string) *GetSourceResponseBodyDataNacosSourceInfo {
	s.Address = &v
	return s
}

func (s *GetSourceResponseBodyDataNacosSourceInfo) SetClusterId(v string) *GetSourceResponseBodyDataNacosSourceInfo {
	s.ClusterId = &v
	return s
}

func (s *GetSourceResponseBodyDataNacosSourceInfo) SetInstanceId(v string) *GetSourceResponseBodyDataNacosSourceInfo {
	s.InstanceId = &v
	return s
}

func (s *GetSourceResponseBodyDataNacosSourceInfo) Validate() error {
	return dara.Validate(s)
}
