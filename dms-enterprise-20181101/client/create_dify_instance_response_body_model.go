// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDifyInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateDifyInstanceResponseBody
	GetCode() *string
	SetData(v *CreateDifyInstanceResponseBodyData) *CreateDifyInstanceResponseBody
	GetData() *CreateDifyInstanceResponseBodyData
	SetErrorCode(v string) *CreateDifyInstanceResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *CreateDifyInstanceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateDifyInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateDifyInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDifyInstanceResponseBody
	GetSuccess() *bool
}

type CreateDifyInstanceResponseBody struct {
	// example:
	//
	// 200
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateDifyInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// UnknownError
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// ABCF54A2-4D74-5DE1-9F0F-5221DDEDD9AD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDifyInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDifyInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDifyInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateDifyInstanceResponseBody) GetData() *CreateDifyInstanceResponseBodyData {
	return s.Data
}

func (s *CreateDifyInstanceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateDifyInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateDifyInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDifyInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDifyInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDifyInstanceResponseBody) SetCode(v string) *CreateDifyInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDifyInstanceResponseBody) SetData(v *CreateDifyInstanceResponseBodyData) *CreateDifyInstanceResponseBody {
	s.Data = v
	return s
}

func (s *CreateDifyInstanceResponseBody) SetErrorCode(v string) *CreateDifyInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDifyInstanceResponseBody) SetHttpStatusCode(v int32) *CreateDifyInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateDifyInstanceResponseBody) SetMessage(v string) *CreateDifyInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDifyInstanceResponseBody) SetRequestId(v string) *CreateDifyInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDifyInstanceResponseBody) SetSuccess(v bool) *CreateDifyInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDifyInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDifyInstanceResponseBodyData struct {
	// example:
	//
	// abc
	AppUuid *string `json:"AppUuid,omitempty" xml:"AppUuid,omitempty"`
	// example:
	//
	// abc1-def2-ghi3-jkl4
	DifyInstanceId *string `json:"DifyInstanceId,omitempty" xml:"DifyInstanceId,omitempty"`
	// example:
	//
	// Dify Instance
	DifyInstanceName *string `json:"DifyInstanceName,omitempty" xml:"DifyInstanceName,omitempty"`
	// example:
	//
	// 291XXX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// trScore
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// 1
	Replicas *int32 `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	// example:
	//
	// 4CU
	ResourceQuota *string `json:"ResourceQuota,omitempty" xml:"ResourceQuota,omitempty"`
	// example:
	//
	// sg-uf6hs6f3m6j5gm6jj0we
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// vsw-bp1m5bwgv41nfoi5el6y1
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// vpc-xxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// 863020290155****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// example:
	//
	// cn-beijing-f
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateDifyInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateDifyInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDifyInstanceResponseBodyData) GetAppUuid() *string {
	return s.AppUuid
}

func (s *CreateDifyInstanceResponseBodyData) GetDifyInstanceId() *string {
	return s.DifyInstanceId
}

func (s *CreateDifyInstanceResponseBodyData) GetDifyInstanceName() *string {
	return s.DifyInstanceName
}

func (s *CreateDifyInstanceResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateDifyInstanceResponseBodyData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateDifyInstanceResponseBodyData) GetReplicas() *int32 {
	return s.Replicas
}

func (s *CreateDifyInstanceResponseBodyData) GetResourceQuota() *string {
	return s.ResourceQuota
}

func (s *CreateDifyInstanceResponseBodyData) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateDifyInstanceResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateDifyInstanceResponseBodyData) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateDifyInstanceResponseBodyData) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateDifyInstanceResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateDifyInstanceResponseBodyData) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateDifyInstanceResponseBodyData) SetAppUuid(v string) *CreateDifyInstanceResponseBodyData {
	s.AppUuid = &v
	return s
}

func (s *CreateDifyInstanceResponseBodyData) SetDifyInstanceId(v string) *CreateDifyInstanceResponseBodyData {
	s.DifyInstanceId = &v
	return s
}

func (s *CreateDifyInstanceResponseBodyData) SetDifyInstanceName(v string) *CreateDifyInstanceResponseBodyData {
	s.DifyInstanceName = &v
	return s
}

func (s *CreateDifyInstanceResponseBodyData) SetInstanceId(v string) *CreateDifyInstanceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *CreateDifyInstanceResponseBodyData) SetInstanceName(v string) *CreateDifyInstanceResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *CreateDifyInstanceResponseBodyData) SetReplicas(v int32) *CreateDifyInstanceResponseBodyData {
	s.Replicas = &v
	return s
}

func (s *CreateDifyInstanceResponseBodyData) SetResourceQuota(v string) *CreateDifyInstanceResponseBodyData {
	s.ResourceQuota = &v
	return s
}

func (s *CreateDifyInstanceResponseBodyData) SetSecurityGroupId(v string) *CreateDifyInstanceResponseBodyData {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateDifyInstanceResponseBodyData) SetStatus(v string) *CreateDifyInstanceResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateDifyInstanceResponseBodyData) SetVSwitchId(v string) *CreateDifyInstanceResponseBodyData {
	s.VSwitchId = &v
	return s
}

func (s *CreateDifyInstanceResponseBodyData) SetVpcId(v string) *CreateDifyInstanceResponseBodyData {
	s.VpcId = &v
	return s
}

func (s *CreateDifyInstanceResponseBodyData) SetWorkspaceId(v string) *CreateDifyInstanceResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *CreateDifyInstanceResponseBodyData) SetZoneId(v string) *CreateDifyInstanceResponseBodyData {
	s.ZoneId = &v
	return s
}

func (s *CreateDifyInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
