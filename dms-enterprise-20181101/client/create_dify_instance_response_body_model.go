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
	Code           *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *CreateDifyInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode      *string                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpStatusCode *int32                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
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
	AppUuid         *string `json:"AppUuid,omitempty" xml:"AppUuid,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName    *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Replicas        *int32  `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	ResourceQuota   *string `json:"ResourceQuota,omitempty" xml:"ResourceQuota,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VSwitchId       *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId           *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	WorkspaceId     *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	ZoneId          *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
