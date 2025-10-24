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
	// The data returned.
	Data *GetComputeQuotaPlanResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// QUOTA_PLAN_NOT_FOUND
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// plan \\"***\\" does not exist
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// The HTTP status code.
	//
	// - 1xx: informational response. The request is received and is being processed.
	//
	// - 2xx: success. The request is successfully received, understood, and accepted by the server.
	//
	// - 3xx: redirection. The request is redirected, and further actions are required to complete the request.
	//
	// - 4xx: client error. The request contains invalid request parameters or syntaxes, or specific request conditions cannot be met.
	//
	// - 5xx: server error. The server cannot meet requirements due to other reasons.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EA1320AB-7766-5EC7-B0F6-8B20E2298567
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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
	// The time when the quota plan was created.
	//
	// example:
	//
	// 1730946421757
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// Whether it is currently effective.
	//
	// >
	//
	// > - A Quota plan that has taken effect cannot be deleted, i.e., isEffective=true
	//
	// example:
	//
	// true/false
	IsEffective *bool `json:"isEffective,omitempty" xml:"isEffective,omitempty"`
	// The name of the quota plan.
	//
	// example:
	//
	// planA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The details of the quota.
	Quota *GetComputeQuotaPlanResponseBodyDataQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
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
	// Cluster ID.
	//
	// example:
	//
	// AT-120N
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// Creation time.
	//
	// example:
	//
	// 1719886322347
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// Creator\\"s cloud account UID.
	//
	// example:
	//
	// 672863518
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// The ID of the level-1 quota.
	//
	// example:
	//
	// 2413
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the level-1 quota.
	//
	// example:
	//
	// quota_a
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The nickname of the level-1 quota.
	//
	// example:
	//
	// quota_nickname
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// CU value parameters for the level-1 quota.
	Parameter *GetComputeQuotaPlanResponseBodyDataQuotaParameter `json:"parameter,omitempty" xml:"parameter,omitempty" type:"Struct"`
	// Region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// Resource status.
	//
	// example:
	//
	// ON
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The list of level-2 quotas.
	SubQuotaInfoList []*GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList `json:"subQuotaInfoList,omitempty" xml:"subQuotaInfoList,omitempty" type:"Repeated"`
	// Tenant ID.
	//
	// example:
	//
	// 478403690625249
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// Corresponds to the `resourceSystemType` field of the control cluster.
	//
	// example:
	//
	// FUXI_ONLINE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// Version number.
	//
	// example:
	//
	// 1964
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
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
	// The value of elastic Reserved CUs.
	//
	// example:
	//
	// 50
	ElasticReservedCU *int64 `json:"elasticReservedCU,omitempty" xml:"elasticReservedCU,omitempty"`
	// The value of maxCU in Reserved CUs.
	//
	// example:
	//
	// 50
	MaxCU *int64 `json:"maxCU,omitempty" xml:"maxCU,omitempty"`
	// The value of minCU in Reserved CUs.
	//
	// example:
	//
	// 50
	MinCU *int64 `json:"minCU,omitempty" xml:"minCU,omitempty"`
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
	// Cluster ID.
	//
	// example:
	//
	// AT-120N
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// Creation time.
	//
	// example:
	//
	// 1718155201628
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// Creator cloud account UID.
	//
	// example:
	//
	// 672863518
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// The ID of the level-2 quota.
	//
	// example:
	//
	// 10940
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the level-2 quota.
	//
	// example:
	//
	// dp_cn_shanghai_1696659792_p
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The nickname of the level-2 quota.
	//
	// example:
	//
	// subquotaA
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// The parameters of the level-2 quota.
	Parameter *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter `json:"parameter,omitempty" xml:"parameter,omitempty" type:"Struct"`
	// Region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// Resource status.
	//
	// example:
	//
	// ON
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Tenant ID.
	//
	// example:
	//
	// 478403690625249
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The type of quota.
	//
	// example:
	//
	// FUXI_ONLINE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// Version number.
	//
	// example:
	//
	// 1386
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
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
	// The value of elastic Reserved CUs.
	//
	// example:
	//
	// 50
	ElasticReservedCU *int64 `json:"elasticReservedCU,omitempty" xml:"elasticReservedCU,omitempty"`
	// whether to enable the priority feature.
	//
	// example:
	//
	// true/false
	EnablePriority *bool `json:"enablePriority,omitempty" xml:"enablePriority,omitempty"`
	// Whether it is exclusive.
	//
	// example:
	//
	// true/false
	ForceReservedMin *bool `json:"forceReservedMin,omitempty" xml:"forceReservedMin,omitempty"`
	// The value of maxCU in Reserved CUs.
	//
	// example:
	//
	// 50
	MaxCU *int64 `json:"maxCU,omitempty" xml:"maxCU,omitempty"`
	// The value of minCU in Reserved CUs.
	//
	// example:
	//
	// 50
	MinCU *int64 `json:"minCU,omitempty" xml:"minCU,omitempty"`
	// Scheduling policy.
	//
	// example:
	//
	// Fifo/Fair
	SchedulerType *string `json:"schedulerType,omitempty" xml:"schedulerType,omitempty"`
	// The upper limit for CUs that can be concurrently used by a job scheduled to the quota.
	//
	// example:
	//
	// 50
	SingleJobCULimit *int64 `json:"singleJobCULimit,omitempty" xml:"singleJobCULimit,omitempty"`
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
