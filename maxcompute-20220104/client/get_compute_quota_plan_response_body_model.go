// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetComputeQuotaPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetComputeQuotaPlanResponseBodyData) *GetComputeQuotaPlanResponseBody
	GetData() *GetComputeQuotaPlanResponseBodyData
	SetErrorCode(v string) *GetComputeQuotaPlanResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GetComputeQuotaPlanResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *GetComputeQuotaPlanResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetComputeQuotaPlanResponseBody
	GetRequestId() *string
}

type GetComputeQuotaPlanResponseBody struct {
	Data      *GetComputeQuotaPlanResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrorCode *string                              `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string                              `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	HttpCode  *int32                               `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	RequestId *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetComputeQuotaPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetComputeQuotaPlanResponseBody) GoString() string {
	return s.String()
}

func (s *GetComputeQuotaPlanResponseBody) GetData() *GetComputeQuotaPlanResponseBodyData {
	return s.Data
}

func (s *GetComputeQuotaPlanResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetComputeQuotaPlanResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetComputeQuotaPlanResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetComputeQuotaPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetComputeQuotaPlanResponseBody) SetData(v *GetComputeQuotaPlanResponseBodyData) *GetComputeQuotaPlanResponseBody {
	s.Data = v
	return s
}

func (s *GetComputeQuotaPlanResponseBody) SetErrorCode(v string) *GetComputeQuotaPlanResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBody) SetErrorMsg(v string) *GetComputeQuotaPlanResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBody) SetHttpCode(v int32) *GetComputeQuotaPlanResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBody) SetRequestId(v string) *GetComputeQuotaPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetComputeQuotaPlanResponseBodyData struct {
	CreateTime  *string                                   `json:"createTime,omitempty" xml:"createTime,omitempty"`
	IsEffective *bool                                     `json:"isEffective,omitempty" xml:"isEffective,omitempty"`
	Name        *string                                   `json:"name,omitempty" xml:"name,omitempty"`
	Quota       *GetComputeQuotaPlanResponseBodyDataQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
}

func (s GetComputeQuotaPlanResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetComputeQuotaPlanResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetComputeQuotaPlanResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetComputeQuotaPlanResponseBodyData) GetIsEffective() *bool {
	return s.IsEffective
}

func (s *GetComputeQuotaPlanResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetComputeQuotaPlanResponseBodyData) GetQuota() *GetComputeQuotaPlanResponseBodyDataQuota {
	return s.Quota
}

func (s *GetComputeQuotaPlanResponseBodyData) SetCreateTime(v string) *GetComputeQuotaPlanResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyData) SetIsEffective(v bool) *GetComputeQuotaPlanResponseBodyData {
	s.IsEffective = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyData) SetName(v string) *GetComputeQuotaPlanResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyData) SetQuota(v *GetComputeQuotaPlanResponseBodyDataQuota) *GetComputeQuotaPlanResponseBodyData {
	s.Quota = v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyData) Validate() error {
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetComputeQuotaPlanResponseBodyDataQuota struct {
	Cluster          *string                                                     `json:"cluster,omitempty" xml:"cluster,omitempty"`
	CreateTime       *int64                                                      `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorId        *string                                                     `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Id               *string                                                     `json:"id,omitempty" xml:"id,omitempty"`
	Name             *string                                                     `json:"name,omitempty" xml:"name,omitempty"`
	NickName         *string                                                     `json:"nickName,omitempty" xml:"nickName,omitempty"`
	Parameter        *GetComputeQuotaPlanResponseBodyDataQuotaParameter          `json:"parameter,omitempty" xml:"parameter,omitempty" type:"Struct"`
	RegionId         *string                                                     `json:"regionId,omitempty" xml:"regionId,omitempty"`
	Status           *string                                                     `json:"status,omitempty" xml:"status,omitempty"`
	SubQuotaInfoList []*GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList `json:"subQuotaInfoList,omitempty" xml:"subQuotaInfoList,omitempty" type:"Repeated"`
	TenantId         *string                                                     `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	Type             *string                                                     `json:"type,omitempty" xml:"type,omitempty"`
	Version          *string                                                     `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetComputeQuotaPlanResponseBodyDataQuota) String() string {
	return dara.Prettify(s)
}

func (s GetComputeQuotaPlanResponseBodyDataQuota) GoString() string {
	return s.String()
}

func (s *GetComputeQuotaPlanResponseBodyDataQuota) GetCluster() *string {
	return s.Cluster
}

func (s *GetComputeQuotaPlanResponseBodyDataQuota) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetComputeQuotaPlanResponseBodyDataQuota) GetCreatorId() *string {
	return s.CreatorId
}

func (s *GetComputeQuotaPlanResponseBodyDataQuota) GetId() *string {
	return s.Id
}

func (s *GetComputeQuotaPlanResponseBodyDataQuota) GetName() *string {
	return s.Name
}

func (s *GetComputeQuotaPlanResponseBodyDataQuota) GetNickName() *string {
	return s.NickName
}

func (s *GetComputeQuotaPlanResponseBodyDataQuota) GetParameter() *GetComputeQuotaPlanResponseBodyDataQuotaParameter {
	return s.Parameter
}

func (s *GetComputeQuotaPlanResponseBodyDataQuota) GetRegionId() *string {
	return s.RegionId
}

func (s *GetComputeQuotaPlanResponseBodyDataQuota) GetStatus() *string {
	return s.Status
}

func (s *GetComputeQuotaPlanResponseBodyDataQuota) GetSubQuotaInfoList() []*GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	return s.SubQuotaInfoList
}

func (s *GetComputeQuotaPlanResponseBodyDataQuota) GetTenantId() *string {
	return s.TenantId
}

func (s *GetComputeQuotaPlanResponseBodyDataQuota) GetType() *string {
	return s.Type
}

func (s *GetComputeQuotaPlanResponseBodyDataQuota) GetVersion() *string {
	return s.Version
}

func (s *GetComputeQuotaPlanResponseBodyDataQuota) SetCluster(v string) *GetComputeQuotaPlanResponseBodyDataQuota {
	s.Cluster = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuota) SetCreateTime(v int64) *GetComputeQuotaPlanResponseBodyDataQuota {
	s.CreateTime = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuota) SetCreatorId(v string) *GetComputeQuotaPlanResponseBodyDataQuota {
	s.CreatorId = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuota) SetId(v string) *GetComputeQuotaPlanResponseBodyDataQuota {
	s.Id = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuota) SetName(v string) *GetComputeQuotaPlanResponseBodyDataQuota {
	s.Name = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuota) SetNickName(v string) *GetComputeQuotaPlanResponseBodyDataQuota {
	s.NickName = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuota) SetParameter(v *GetComputeQuotaPlanResponseBodyDataQuotaParameter) *GetComputeQuotaPlanResponseBodyDataQuota {
	s.Parameter = v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuota) SetRegionId(v string) *GetComputeQuotaPlanResponseBodyDataQuota {
	s.RegionId = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuota) SetStatus(v string) *GetComputeQuotaPlanResponseBodyDataQuota {
	s.Status = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuota) SetSubQuotaInfoList(v []*GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) *GetComputeQuotaPlanResponseBodyDataQuota {
	s.SubQuotaInfoList = v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuota) SetTenantId(v string) *GetComputeQuotaPlanResponseBodyDataQuota {
	s.TenantId = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuota) SetType(v string) *GetComputeQuotaPlanResponseBodyDataQuota {
	s.Type = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuota) SetVersion(v string) *GetComputeQuotaPlanResponseBodyDataQuota {
	s.Version = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuota) Validate() error {
	if s.Parameter != nil {
		if err := s.Parameter.Validate(); err != nil {
			return err
		}
	}
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

type GetComputeQuotaPlanResponseBodyDataQuotaParameter struct {
	ElasticReservedCU *int64 `json:"elasticReservedCU,omitempty" xml:"elasticReservedCU,omitempty"`
	MaxCU             *int64 `json:"maxCU,omitempty" xml:"maxCU,omitempty"`
	MinCU             *int64 `json:"minCU,omitempty" xml:"minCU,omitempty"`
}

func (s GetComputeQuotaPlanResponseBodyDataQuotaParameter) String() string {
	return dara.Prettify(s)
}

func (s GetComputeQuotaPlanResponseBodyDataQuotaParameter) GoString() string {
	return s.String()
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaParameter) GetElasticReservedCU() *int64 {
	return s.ElasticReservedCU
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaParameter) GetMaxCU() *int64 {
	return s.MaxCU
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaParameter) GetMinCU() *int64 {
	return s.MinCU
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaParameter) SetElasticReservedCU(v int64) *GetComputeQuotaPlanResponseBodyDataQuotaParameter {
	s.ElasticReservedCU = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaParameter) SetMaxCU(v int64) *GetComputeQuotaPlanResponseBodyDataQuotaParameter {
	s.MaxCU = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaParameter) SetMinCU(v int64) *GetComputeQuotaPlanResponseBodyDataQuotaParameter {
	s.MinCU = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaParameter) Validate() error {
	return dara.Validate(s)
}

type GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList struct {
	Cluster    *string                                                            `json:"cluster,omitempty" xml:"cluster,omitempty"`
	CreateTime *int64                                                             `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorId  *string                                                            `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Id         *string                                                            `json:"id,omitempty" xml:"id,omitempty"`
	Name       *string                                                            `json:"name,omitempty" xml:"name,omitempty"`
	NickName   *string                                                            `json:"nickName,omitempty" xml:"nickName,omitempty"`
	Parameter  *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter `json:"parameter,omitempty" xml:"parameter,omitempty" type:"Struct"`
	RegionId   *string                                                            `json:"regionId,omitempty" xml:"regionId,omitempty"`
	Status     *string                                                            `json:"status,omitempty" xml:"status,omitempty"`
	TenantId   *string                                                            `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	Type       *string                                                            `json:"type,omitempty" xml:"type,omitempty"`
	Version    *string                                                            `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) GoString() string {
	return s.String()
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) GetCluster() *string {
	return s.Cluster
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) GetCreatorId() *string {
	return s.CreatorId
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) GetId() *string {
	return s.Id
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) GetName() *string {
	return s.Name
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) GetNickName() *string {
	return s.NickName
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) GetParameter() *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter {
	return s.Parameter
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) GetRegionId() *string {
	return s.RegionId
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) GetStatus() *string {
	return s.Status
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) GetTenantId() *string {
	return s.TenantId
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) GetType() *string {
	return s.Type
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) GetVersion() *string {
	return s.Version
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetCluster(v string) *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Cluster = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetCreateTime(v int64) *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.CreateTime = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetCreatorId(v string) *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.CreatorId = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetId(v string) *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Id = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetName(v string) *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Name = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetNickName(v string) *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.NickName = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetParameter(v *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter) *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Parameter = v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetRegionId(v string) *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.RegionId = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetStatus(v string) *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Status = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetTenantId(v string) *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.TenantId = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetType(v string) *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Type = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetVersion(v string) *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Version = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) Validate() error {
	if s.Parameter != nil {
		if err := s.Parameter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter struct {
	ElasticReservedCU *int64  `json:"elasticReservedCU,omitempty" xml:"elasticReservedCU,omitempty"`
	EnablePriority    *bool   `json:"enablePriority,omitempty" xml:"enablePriority,omitempty"`
	ForceReservedMin  *bool   `json:"forceReservedMin,omitempty" xml:"forceReservedMin,omitempty"`
	MaxCU             *int64  `json:"maxCU,omitempty" xml:"maxCU,omitempty"`
	MinCU             *int64  `json:"minCU,omitempty" xml:"minCU,omitempty"`
	SchedulerType     *string `json:"schedulerType,omitempty" xml:"schedulerType,omitempty"`
	SingleJobCULimit  *int64  `json:"singleJobCULimit,omitempty" xml:"singleJobCULimit,omitempty"`
}

func (s GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter) String() string {
	return dara.Prettify(s)
}

func (s GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter) GoString() string {
	return s.String()
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter) GetElasticReservedCU() *int64 {
	return s.ElasticReservedCU
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter) GetEnablePriority() *bool {
	return s.EnablePriority
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter) GetForceReservedMin() *bool {
	return s.ForceReservedMin
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter) GetMaxCU() *int64 {
	return s.MaxCU
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter) GetMinCU() *int64 {
	return s.MinCU
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter) GetSchedulerType() *string {
	return s.SchedulerType
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter) GetSingleJobCULimit() *int64 {
	return s.SingleJobCULimit
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter) SetElasticReservedCU(v int64) *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter {
	s.ElasticReservedCU = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter) SetEnablePriority(v bool) *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter {
	s.EnablePriority = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter) SetForceReservedMin(v bool) *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter {
	s.ForceReservedMin = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter) SetMaxCU(v int64) *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter {
	s.MaxCU = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter) SetMinCU(v int64) *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter {
	s.MinCU = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter) SetSchedulerType(v string) *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter {
	s.SchedulerType = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter) SetSingleJobCULimit(v int64) *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter {
	s.SingleJobCULimit = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter) Validate() error {
	return dara.Validate(s)
}
