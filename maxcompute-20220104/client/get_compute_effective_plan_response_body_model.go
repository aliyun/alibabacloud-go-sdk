// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetComputeEffectivePlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetComputeEffectivePlanResponseBodyData) *GetComputeEffectivePlanResponseBody
	GetData() *GetComputeEffectivePlanResponseBodyData
	SetErrorCode(v string) *GetComputeEffectivePlanResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GetComputeEffectivePlanResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *GetComputeEffectivePlanResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetComputeEffectivePlanResponseBody
	GetRequestId() *string
}

type GetComputeEffectivePlanResponseBody struct {
	Data      *GetComputeEffectivePlanResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrorCode *string                                  `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string                                  `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	HttpCode  *int32                                   `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	RequestId *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetComputeEffectivePlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetComputeEffectivePlanResponseBody) GoString() string {
	return s.String()
}

func (s *GetComputeEffectivePlanResponseBody) GetData() *GetComputeEffectivePlanResponseBodyData {
	return s.Data
}

func (s *GetComputeEffectivePlanResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetComputeEffectivePlanResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetComputeEffectivePlanResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetComputeEffectivePlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetComputeEffectivePlanResponseBody) SetData(v *GetComputeEffectivePlanResponseBodyData) *GetComputeEffectivePlanResponseBody {
	s.Data = v
	return s
}

func (s *GetComputeEffectivePlanResponseBody) SetErrorCode(v string) *GetComputeEffectivePlanResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBody) SetErrorMsg(v string) *GetComputeEffectivePlanResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBody) SetHttpCode(v int32) *GetComputeEffectivePlanResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBody) SetRequestId(v string) *GetComputeEffectivePlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetComputeEffectivePlanResponseBodyData struct {
	CreateTime  *string                                       `json:"createTime,omitempty" xml:"createTime,omitempty"`
	IsEffective *bool                                         `json:"isEffective,omitempty" xml:"isEffective,omitempty"`
	Name        *string                                       `json:"name,omitempty" xml:"name,omitempty"`
	Quota       *GetComputeEffectivePlanResponseBodyDataQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
}

func (s GetComputeEffectivePlanResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetComputeEffectivePlanResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetComputeEffectivePlanResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetComputeEffectivePlanResponseBodyData) GetIsEffective() *bool {
	return s.IsEffective
}

func (s *GetComputeEffectivePlanResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetComputeEffectivePlanResponseBodyData) GetQuota() *GetComputeEffectivePlanResponseBodyDataQuota {
	return s.Quota
}

func (s *GetComputeEffectivePlanResponseBodyData) SetCreateTime(v string) *GetComputeEffectivePlanResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyData) SetIsEffective(v bool) *GetComputeEffectivePlanResponseBodyData {
	s.IsEffective = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyData) SetName(v string) *GetComputeEffectivePlanResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyData) SetQuota(v *GetComputeEffectivePlanResponseBodyDataQuota) *GetComputeEffectivePlanResponseBodyData {
	s.Quota = v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyData) Validate() error {
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetComputeEffectivePlanResponseBodyDataQuota struct {
	Cluster          *string                                                         `json:"cluster,omitempty" xml:"cluster,omitempty"`
	CreateTime       *int64                                                          `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorId        *string                                                         `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Id               *string                                                         `json:"id,omitempty" xml:"id,omitempty"`
	Name             *string                                                         `json:"name,omitempty" xml:"name,omitempty"`
	NickName         *string                                                         `json:"nickName,omitempty" xml:"nickName,omitempty"`
	Parameter        map[string]interface{}                                          `json:"parameter,omitempty" xml:"parameter,omitempty"`
	RegionId         *string                                                         `json:"regionId,omitempty" xml:"regionId,omitempty"`
	Status           *string                                                         `json:"status,omitempty" xml:"status,omitempty"`
	SubQuotaInfoList []*GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList `json:"subQuotaInfoList,omitempty" xml:"subQuotaInfoList,omitempty" type:"Repeated"`
	TenantId         *string                                                         `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	Type             *string                                                         `json:"type,omitempty" xml:"type,omitempty"`
	Version          *string                                                         `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetComputeEffectivePlanResponseBodyDataQuota) String() string {
	return dara.Prettify(s)
}

func (s GetComputeEffectivePlanResponseBodyDataQuota) GoString() string {
	return s.String()
}

func (s *GetComputeEffectivePlanResponseBodyDataQuota) GetCluster() *string {
	return s.Cluster
}

func (s *GetComputeEffectivePlanResponseBodyDataQuota) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetComputeEffectivePlanResponseBodyDataQuota) GetCreatorId() *string {
	return s.CreatorId
}

func (s *GetComputeEffectivePlanResponseBodyDataQuota) GetId() *string {
	return s.Id
}

func (s *GetComputeEffectivePlanResponseBodyDataQuota) GetName() *string {
	return s.Name
}

func (s *GetComputeEffectivePlanResponseBodyDataQuota) GetNickName() *string {
	return s.NickName
}

func (s *GetComputeEffectivePlanResponseBodyDataQuota) GetParameter() map[string]interface{} {
	return s.Parameter
}

func (s *GetComputeEffectivePlanResponseBodyDataQuota) GetRegionId() *string {
	return s.RegionId
}

func (s *GetComputeEffectivePlanResponseBodyDataQuota) GetStatus() *string {
	return s.Status
}

func (s *GetComputeEffectivePlanResponseBodyDataQuota) GetSubQuotaInfoList() []*GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList {
	return s.SubQuotaInfoList
}

func (s *GetComputeEffectivePlanResponseBodyDataQuota) GetTenantId() *string {
	return s.TenantId
}

func (s *GetComputeEffectivePlanResponseBodyDataQuota) GetType() *string {
	return s.Type
}

func (s *GetComputeEffectivePlanResponseBodyDataQuota) GetVersion() *string {
	return s.Version
}

func (s *GetComputeEffectivePlanResponseBodyDataQuota) SetCluster(v string) *GetComputeEffectivePlanResponseBodyDataQuota {
	s.Cluster = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuota) SetCreateTime(v int64) *GetComputeEffectivePlanResponseBodyDataQuota {
	s.CreateTime = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuota) SetCreatorId(v string) *GetComputeEffectivePlanResponseBodyDataQuota {
	s.CreatorId = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuota) SetId(v string) *GetComputeEffectivePlanResponseBodyDataQuota {
	s.Id = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuota) SetName(v string) *GetComputeEffectivePlanResponseBodyDataQuota {
	s.Name = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuota) SetNickName(v string) *GetComputeEffectivePlanResponseBodyDataQuota {
	s.NickName = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuota) SetParameter(v map[string]interface{}) *GetComputeEffectivePlanResponseBodyDataQuota {
	s.Parameter = v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuota) SetRegionId(v string) *GetComputeEffectivePlanResponseBodyDataQuota {
	s.RegionId = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuota) SetStatus(v string) *GetComputeEffectivePlanResponseBodyDataQuota {
	s.Status = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuota) SetSubQuotaInfoList(v []*GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) *GetComputeEffectivePlanResponseBodyDataQuota {
	s.SubQuotaInfoList = v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuota) SetTenantId(v string) *GetComputeEffectivePlanResponseBodyDataQuota {
	s.TenantId = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuota) SetType(v string) *GetComputeEffectivePlanResponseBodyDataQuota {
	s.Type = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuota) SetVersion(v string) *GetComputeEffectivePlanResponseBodyDataQuota {
	s.Version = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuota) Validate() error {
	if s.SubQuotaInfoList != nil {
		for _, item := range s.SubQuotaInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList struct {
	Cluster    *string                `json:"cluster,omitempty" xml:"cluster,omitempty"`
	CreateTime *int64                 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorId  *string                `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Id         *string                `json:"id,omitempty" xml:"id,omitempty"`
	Name       *string                `json:"name,omitempty" xml:"name,omitempty"`
	NickName   *string                `json:"nickName,omitempty" xml:"nickName,omitempty"`
	Parameter  map[string]interface{} `json:"parameter,omitempty" xml:"parameter,omitempty"`
	RegionId   *string                `json:"regionId,omitempty" xml:"regionId,omitempty"`
	Status     *string                `json:"status,omitempty" xml:"status,omitempty"`
	TenantId   *string                `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	Type       *string                `json:"type,omitempty" xml:"type,omitempty"`
	Version    *string                `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) GoString() string {
	return s.String()
}

func (s *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) GetCluster() *string {
	return s.Cluster
}

func (s *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) GetCreatorId() *string {
	return s.CreatorId
}

func (s *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) GetId() *string {
	return s.Id
}

func (s *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) GetName() *string {
	return s.Name
}

func (s *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) GetNickName() *string {
	return s.NickName
}

func (s *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) GetParameter() map[string]interface{} {
	return s.Parameter
}

func (s *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) GetRegionId() *string {
	return s.RegionId
}

func (s *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) GetStatus() *string {
	return s.Status
}

func (s *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) GetTenantId() *string {
	return s.TenantId
}

func (s *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) GetType() *string {
	return s.Type
}

func (s *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) GetVersion() *string {
	return s.Version
}

func (s *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) SetCluster(v string) *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Cluster = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) SetCreateTime(v int64) *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList {
	s.CreateTime = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) SetCreatorId(v string) *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList {
	s.CreatorId = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) SetId(v string) *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Id = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) SetName(v string) *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Name = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) SetNickName(v string) *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList {
	s.NickName = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) SetParameter(v map[string]interface{}) *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Parameter = v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) SetRegionId(v string) *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList {
	s.RegionId = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) SetStatus(v string) *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Status = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) SetTenantId(v string) *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList {
	s.TenantId = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) SetType(v string) *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Type = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) SetVersion(v string) *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Version = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) Validate() error {
	return dara.Validate(s)
}
