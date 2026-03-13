// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserAliUidByInstanceIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *GetUserAliUidByInstanceIdResponseBody
	GetApplicationId() *string
	SetApplicationName(v string) *GetUserAliUidByInstanceIdResponseBody
	GetApplicationName() *string
	SetCode(v string) *GetUserAliUidByInstanceIdResponseBody
	GetCode() *string
	SetCpus(v int64) *GetUserAliUidByInstanceIdResponseBody
	GetCpus() *int64
	SetCreateTime(v int64) *GetUserAliUidByInstanceIdResponseBody
	GetCreateTime() *int64
	SetIntranetIp(v string) *GetUserAliUidByInstanceIdResponseBody
	GetIntranetIp() *string
	SetMemory(v int64) *GetUserAliUidByInstanceIdResponseBody
	GetMemory() *int64
	SetMessage(v string) *GetUserAliUidByInstanceIdResponseBody
	GetMessage() *string
	SetNamespaceId(v string) *GetUserAliUidByInstanceIdResponseBody
	GetNamespaceId() *string
	SetPodName(v string) *GetUserAliUidByInstanceIdResponseBody
	GetPodName() *string
	SetPodUid(v string) *GetUserAliUidByInstanceIdResponseBody
	GetPodUid() *string
	SetRequestId(v string) *GetUserAliUidByInstanceIdResponseBody
	GetRequestId() *string
	SetResouceType(v string) *GetUserAliUidByInstanceIdResponseBody
	GetResouceType() *string
	SetResourceInstanceId(v string) *GetUserAliUidByInstanceIdResponseBody
	GetResourceInstanceId() *string
	SetStatus(v string) *GetUserAliUidByInstanceIdResponseBody
	GetStatus() *string
	SetUpdateTime(v int64) *GetUserAliUidByInstanceIdResponseBody
	GetUpdateTime() *int64
	SetUserUid(v string) *GetUserAliUidByInstanceIdResponseBody
	GetUserUid() *string
}

type GetUserAliUidByInstanceIdResponseBody struct {
	ApplicationId      *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	ApplicationName    *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	Code               *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Cpus               *int64  `json:"Cpus,omitempty" xml:"Cpus,omitempty"`
	CreateTime         *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	IntranetIp         *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	Memory             *int64  `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Message            *string `json:"Message,omitempty" xml:"Message,omitempty"`
	NamespaceId        *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	PodName            *string `json:"PodName,omitempty" xml:"PodName,omitempty"`
	PodUid             *string `json:"PodUid,omitempty" xml:"PodUid,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResouceType        *string `json:"ResouceType,omitempty" xml:"ResouceType,omitempty"`
	ResourceInstanceId *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime         *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserUid            *string `json:"UserUid,omitempty" xml:"UserUid,omitempty"`
}

func (s GetUserAliUidByInstanceIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserAliUidByInstanceIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserAliUidByInstanceIdResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *GetUserAliUidByInstanceIdResponseBody) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *GetUserAliUidByInstanceIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetUserAliUidByInstanceIdResponseBody) GetCpus() *int64 {
	return s.Cpus
}

func (s *GetUserAliUidByInstanceIdResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetUserAliUidByInstanceIdResponseBody) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *GetUserAliUidByInstanceIdResponseBody) GetMemory() *int64 {
	return s.Memory
}

func (s *GetUserAliUidByInstanceIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUserAliUidByInstanceIdResponseBody) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *GetUserAliUidByInstanceIdResponseBody) GetPodName() *string {
	return s.PodName
}

func (s *GetUserAliUidByInstanceIdResponseBody) GetPodUid() *string {
	return s.PodUid
}

func (s *GetUserAliUidByInstanceIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserAliUidByInstanceIdResponseBody) GetResouceType() *string {
	return s.ResouceType
}

func (s *GetUserAliUidByInstanceIdResponseBody) GetResourceInstanceId() *string {
	return s.ResourceInstanceId
}

func (s *GetUserAliUidByInstanceIdResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetUserAliUidByInstanceIdResponseBody) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetUserAliUidByInstanceIdResponseBody) GetUserUid() *string {
	return s.UserUid
}

func (s *GetUserAliUidByInstanceIdResponseBody) SetApplicationId(v string) *GetUserAliUidByInstanceIdResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *GetUserAliUidByInstanceIdResponseBody) SetApplicationName(v string) *GetUserAliUidByInstanceIdResponseBody {
	s.ApplicationName = &v
	return s
}

func (s *GetUserAliUidByInstanceIdResponseBody) SetCode(v string) *GetUserAliUidByInstanceIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserAliUidByInstanceIdResponseBody) SetCpus(v int64) *GetUserAliUidByInstanceIdResponseBody {
	s.Cpus = &v
	return s
}

func (s *GetUserAliUidByInstanceIdResponseBody) SetCreateTime(v int64) *GetUserAliUidByInstanceIdResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetUserAliUidByInstanceIdResponseBody) SetIntranetIp(v string) *GetUserAliUidByInstanceIdResponseBody {
	s.IntranetIp = &v
	return s
}

func (s *GetUserAliUidByInstanceIdResponseBody) SetMemory(v int64) *GetUserAliUidByInstanceIdResponseBody {
	s.Memory = &v
	return s
}

func (s *GetUserAliUidByInstanceIdResponseBody) SetMessage(v string) *GetUserAliUidByInstanceIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserAliUidByInstanceIdResponseBody) SetNamespaceId(v string) *GetUserAliUidByInstanceIdResponseBody {
	s.NamespaceId = &v
	return s
}

func (s *GetUserAliUidByInstanceIdResponseBody) SetPodName(v string) *GetUserAliUidByInstanceIdResponseBody {
	s.PodName = &v
	return s
}

func (s *GetUserAliUidByInstanceIdResponseBody) SetPodUid(v string) *GetUserAliUidByInstanceIdResponseBody {
	s.PodUid = &v
	return s
}

func (s *GetUserAliUidByInstanceIdResponseBody) SetRequestId(v string) *GetUserAliUidByInstanceIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserAliUidByInstanceIdResponseBody) SetResouceType(v string) *GetUserAliUidByInstanceIdResponseBody {
	s.ResouceType = &v
	return s
}

func (s *GetUserAliUidByInstanceIdResponseBody) SetResourceInstanceId(v string) *GetUserAliUidByInstanceIdResponseBody {
	s.ResourceInstanceId = &v
	return s
}

func (s *GetUserAliUidByInstanceIdResponseBody) SetStatus(v string) *GetUserAliUidByInstanceIdResponseBody {
	s.Status = &v
	return s
}

func (s *GetUserAliUidByInstanceIdResponseBody) SetUpdateTime(v int64) *GetUserAliUidByInstanceIdResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetUserAliUidByInstanceIdResponseBody) SetUserUid(v string) *GetUserAliUidByInstanceIdResponseBody {
	s.UserUid = &v
	return s
}

func (s *GetUserAliUidByInstanceIdResponseBody) Validate() error {
	return dara.Validate(s)
}
