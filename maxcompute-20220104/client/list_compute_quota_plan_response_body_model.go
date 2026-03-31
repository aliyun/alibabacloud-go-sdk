// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeQuotaPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListComputeQuotaPlanResponseBodyData) *ListComputeQuotaPlanResponseBody
	GetData() *ListComputeQuotaPlanResponseBodyData
	SetErrorCode(v string) *ListComputeQuotaPlanResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *ListComputeQuotaPlanResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *ListComputeQuotaPlanResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *ListComputeQuotaPlanResponseBody
	GetRequestId() *string
}

type ListComputeQuotaPlanResponseBody struct {
	// The data returned.
	Data *ListComputeQuotaPlanResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// OBJECT_NOT_EXIST
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// This object does not exist.
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
	// 0bc3b4ae16685836687916212e7850
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListComputeQuotaPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListComputeQuotaPlanResponseBody) GoString() string {
	return s.String()
}

func (s *ListComputeQuotaPlanResponseBody) GetData() *ListComputeQuotaPlanResponseBodyData {
	return s.Data
}

func (s *ListComputeQuotaPlanResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListComputeQuotaPlanResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ListComputeQuotaPlanResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ListComputeQuotaPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListComputeQuotaPlanResponseBody) SetData(v *ListComputeQuotaPlanResponseBodyData) *ListComputeQuotaPlanResponseBody {
	s.Data = v
	return s
}

func (s *ListComputeQuotaPlanResponseBody) SetErrorCode(v string) *ListComputeQuotaPlanResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBody) SetErrorMsg(v string) *ListComputeQuotaPlanResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBody) SetHttpCode(v int32) *ListComputeQuotaPlanResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBody) SetRequestId(v string) *ListComputeQuotaPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListComputeQuotaPlanResponseBodyData struct {
	// The list of quota plan.
	PlanList []*ListComputeQuotaPlanResponseBodyDataPlanList `json:"planList,omitempty" xml:"planList,omitempty" type:"Repeated"`
}

func (s ListComputeQuotaPlanResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListComputeQuotaPlanResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListComputeQuotaPlanResponseBodyData) GetPlanList() []*ListComputeQuotaPlanResponseBodyDataPlanList {
	return s.PlanList
}

func (s *ListComputeQuotaPlanResponseBodyData) SetPlanList(v []*ListComputeQuotaPlanResponseBodyDataPlanList) *ListComputeQuotaPlanResponseBodyData {
	s.PlanList = v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyData) Validate() error {
	if s.PlanList != nil {
		for _, item := range s.PlanList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListComputeQuotaPlanResponseBodyDataPlanList struct {
	// The time when the quota plan was created.
	//
	// example:
	//
	// 1731394621890
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The name of the quota plan.
	//
	// example:
	//
	// planA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The details of the quota.
	Quota *ListComputeQuotaPlanResponseBodyDataPlanListQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
}

func (s ListComputeQuotaPlanResponseBodyDataPlanList) String() string {
	return dara.Prettify(s)
}

func (s ListComputeQuotaPlanResponseBodyDataPlanList) GoString() string {
	return s.String()
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanList) GetName() *string {
	return s.Name
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanList) GetQuota() *ListComputeQuotaPlanResponseBodyDataPlanListQuota {
	return s.Quota
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanList) SetCreateTime(v string) *ListComputeQuotaPlanResponseBodyDataPlanList {
	s.CreateTime = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanList) SetName(v string) *ListComputeQuotaPlanResponseBodyDataPlanList {
	s.Name = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanList) SetQuota(v *ListComputeQuotaPlanResponseBodyDataPlanListQuota) *ListComputeQuotaPlanResponseBodyDataPlanList {
	s.Quota = v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanList) Validate() error {
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListComputeQuotaPlanResponseBodyDataPlanListQuota struct {
	// Cluster ID.
	//
	// example:
	//
	// AT-120N
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The time when the level-1 quota was created.
	//
	// example:
	//
	// 1730247361356
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the resource.
	//
	// example:
	//
	// 672863518
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// The ID of the level-1 quota.
	//
	// example:
	//
	// 186
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the level-1 quota.
	//
	// example:
	//
	// dp_cn_hangzhou_1717465943_p
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The nickname of the level-1 quota.
	//
	// example:
	//
	// os_MyQuota_p
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// The description of the level-1 quota.
	//
	// example:
	//
	// {
	//
	//   "enablePriority": false,
	//
	//   "minCU": 25,
	//
	//   "adhocCU": 0,
	//
	//   "elasticReservedCU": 0,
	//
	//   "forceReservedMin": false,
	//
	//   "maxCU": 50,
	//
	//   "schedulerType": "Fifo"
	//
	// }
	Parameter *ListComputeQuotaPlanResponseBodyDataPlanListQuotaParameter `json:"parameter,omitempty" xml:"parameter,omitempty" type:"Struct"`
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
	// The list of subquotas.
	SubQuotaInfoList []*ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList `json:"subQuotaInfoList,omitempty" xml:"subQuotaInfoList,omitempty" type:"Repeated"`
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
	// The version number.
	//
	// example:
	//
	// 2056
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListComputeQuotaPlanResponseBodyDataPlanListQuota) String() string {
	return dara.Prettify(s)
}

func (s ListComputeQuotaPlanResponseBodyDataPlanListQuota) GoString() string {
	return s.String()
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuota) GetCluster() *string {
	return s.Cluster
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuota) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuota) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuota) GetId() *string {
	return s.Id
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuota) GetName() *string {
	return s.Name
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuota) GetNickName() *string {
	return s.NickName
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuota) GetParameter() *ListComputeQuotaPlanResponseBodyDataPlanListQuotaParameter {
	return s.Parameter
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuota) GetRegionId() *string {
	return s.RegionId
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuota) GetStatus() *string {
	return s.Status
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuota) GetSubQuotaInfoList() []*ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList {
	return s.SubQuotaInfoList
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuota) GetTenantId() *string {
	return s.TenantId
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuota) GetType() *string {
	return s.Type
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuota) GetVersion() *string {
	return s.Version
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuota) SetCluster(v string) *ListComputeQuotaPlanResponseBodyDataPlanListQuota {
	s.Cluster = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuota) SetCreateTime(v int64) *ListComputeQuotaPlanResponseBodyDataPlanListQuota {
	s.CreateTime = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuota) SetCreatorId(v string) *ListComputeQuotaPlanResponseBodyDataPlanListQuota {
	s.CreatorId = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuota) SetId(v string) *ListComputeQuotaPlanResponseBodyDataPlanListQuota {
	s.Id = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuota) SetName(v string) *ListComputeQuotaPlanResponseBodyDataPlanListQuota {
	s.Name = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuota) SetNickName(v string) *ListComputeQuotaPlanResponseBodyDataPlanListQuota {
	s.NickName = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuota) SetParameter(v *ListComputeQuotaPlanResponseBodyDataPlanListQuotaParameter) *ListComputeQuotaPlanResponseBodyDataPlanListQuota {
	s.Parameter = v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuota) SetRegionId(v string) *ListComputeQuotaPlanResponseBodyDataPlanListQuota {
	s.RegionId = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuota) SetStatus(v string) *ListComputeQuotaPlanResponseBodyDataPlanListQuota {
	s.Status = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuota) SetSubQuotaInfoList(v []*ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) *ListComputeQuotaPlanResponseBodyDataPlanListQuota {
	s.SubQuotaInfoList = v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuota) SetTenantId(v string) *ListComputeQuotaPlanResponseBodyDataPlanListQuota {
	s.TenantId = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuota) SetType(v string) *ListComputeQuotaPlanResponseBodyDataPlanListQuota {
	s.Type = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuota) SetVersion(v string) *ListComputeQuotaPlanResponseBodyDataPlanListQuota {
	s.Version = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuota) Validate() error {
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

type ListComputeQuotaPlanResponseBodyDataPlanListQuotaParameter struct {
	ElasticReservedCU *int64 `json:"elasticReservedCU,omitempty" xml:"elasticReservedCU,omitempty"`
	MaxCU             *int64 `json:"maxCU,omitempty" xml:"maxCU,omitempty"`
	MinCU             *int64 `json:"minCU,omitempty" xml:"minCU,omitempty"`
}

func (s ListComputeQuotaPlanResponseBodyDataPlanListQuotaParameter) String() string {
	return dara.Prettify(s)
}

func (s ListComputeQuotaPlanResponseBodyDataPlanListQuotaParameter) GoString() string {
	return s.String()
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaParameter) GetElasticReservedCU() *int64 {
	return s.ElasticReservedCU
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaParameter) GetMaxCU() *int64 {
	return s.MaxCU
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaParameter) GetMinCU() *int64 {
	return s.MinCU
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaParameter) SetElasticReservedCU(v int64) *ListComputeQuotaPlanResponseBodyDataPlanListQuotaParameter {
	s.ElasticReservedCU = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaParameter) SetMaxCU(v int64) *ListComputeQuotaPlanResponseBodyDataPlanListQuotaParameter {
	s.MaxCU = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaParameter) SetMinCU(v int64) *ListComputeQuotaPlanResponseBodyDataPlanListQuotaParameter {
	s.MinCU = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaParameter) Validate() error {
	return dara.Validate(s)
}

type ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList struct {
	// Cluster ID.
	//
	// example:
	//
	// AT-120N
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1730946421757
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the resource.
	//
	// example:
	//
	// 672863518
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// The ID of the level-2 quota.
	//
	// example:
	//
	// 6790
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the level-2 quota.
	//
	// example:
	//
	// dp_cn_shanghai_1702627945_p
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The nickname of the level-2 quota.
	//
	// example:
	//
	// os_MyQuota
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// The description of the level-2 quota.
	//
	// example:
	//
	// {
	//
	//   "enablePriority": false,
	//
	//   "minCU": 25,
	//
	//   "adhocCU": 0,
	//
	//   "elasticReservedCU": 0,
	//
	//   "forceReservedMin": false,
	//
	//   "maxCU": 50,
	//
	//   "schedulerType": "Fifo"
	//
	// }
	Parameter *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoListParameter `json:"parameter,omitempty" xml:"parameter,omitempty" type:"Struct"`
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
	// The version number.
	//
	// example:
	//
	// 2056
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) GoString() string {
	return s.String()
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) GetCluster() *string {
	return s.Cluster
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) GetId() *string {
	return s.Id
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) GetName() *string {
	return s.Name
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) GetNickName() *string {
	return s.NickName
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) GetParameter() *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoListParameter {
	return s.Parameter
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) GetRegionId() *string {
	return s.RegionId
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) GetStatus() *string {
	return s.Status
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) GetTenantId() *string {
	return s.TenantId
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) GetType() *string {
	return s.Type
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) GetVersion() *string {
	return s.Version
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) SetCluster(v string) *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.Cluster = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) SetCreateTime(v int64) *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.CreateTime = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) SetCreatorId(v string) *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.CreatorId = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) SetId(v string) *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.Id = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) SetName(v string) *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.Name = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) SetNickName(v string) *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.NickName = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) SetParameter(v *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoListParameter) *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.Parameter = v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) SetRegionId(v string) *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.RegionId = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) SetStatus(v string) *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.Status = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) SetTenantId(v string) *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.TenantId = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) SetType(v string) *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.Type = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) SetVersion(v string) *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.Version = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) Validate() error {
	if s.Parameter != nil {
		if err := s.Parameter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoListParameter struct {
	ElasticReservedCU *int64 `json:"elasticReservedCU,omitempty" xml:"elasticReservedCU,omitempty"`
	MaxCU             *int64 `json:"maxCU,omitempty" xml:"maxCU,omitempty"`
	MinCU             *int64 `json:"minCU,omitempty" xml:"minCU,omitempty"`
}

func (s ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoListParameter) String() string {
	return dara.Prettify(s)
}

func (s ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoListParameter) GoString() string {
	return s.String()
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoListParameter) GetElasticReservedCU() *int64 {
	return s.ElasticReservedCU
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoListParameter) GetMaxCU() *int64 {
	return s.MaxCU
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoListParameter) GetMinCU() *int64 {
	return s.MinCU
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoListParameter) SetElasticReservedCU(v int64) *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoListParameter {
	s.ElasticReservedCU = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoListParameter) SetMaxCU(v int64) *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoListParameter {
	s.MaxCU = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoListParameter) SetMinCU(v int64) *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoListParameter {
	s.MinCU = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoListParameter) Validate() error {
	return dara.Validate(s)
}
