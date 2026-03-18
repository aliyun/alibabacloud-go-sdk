// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListProjectsResponseBodyData) *ListProjectsResponseBody
	GetData() *ListProjectsResponseBodyData
	SetRequestId(v string) *ListProjectsResponseBody
	GetRequestId() *string
}

type ListProjectsResponseBody struct {
	Data      *ListProjectsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	RequestId *string                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListProjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBody) GetData() *ListProjectsResponseBodyData {
	return s.Data
}

func (s *ListProjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProjectsResponseBody) SetData(v *ListProjectsResponseBodyData) *ListProjectsResponseBody {
	s.Data = v
	return s
}

func (s *ListProjectsResponseBody) SetRequestId(v string) *ListProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListProjectsResponseBodyData struct {
	NextToken *string                                 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Marker    *string                                 `json:"marker,omitempty" xml:"marker,omitempty"`
	MaxItem   *int32                                  `json:"maxItem,omitempty" xml:"maxItem,omitempty"`
	Projects  []*ListProjectsResponseBodyDataProjects `json:"projects,omitempty" xml:"projects,omitempty" type:"Repeated"`
}

func (s ListProjectsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProjectsResponseBodyData) GetMarker() *string {
	return s.Marker
}

func (s *ListProjectsResponseBodyData) GetMaxItem() *int32 {
	return s.MaxItem
}

func (s *ListProjectsResponseBodyData) GetProjects() []*ListProjectsResponseBodyDataProjects {
	return s.Projects
}

func (s *ListProjectsResponseBodyData) SetNextToken(v string) *ListProjectsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListProjectsResponseBodyData) SetMarker(v string) *ListProjectsResponseBodyData {
	s.Marker = &v
	return s
}

func (s *ListProjectsResponseBodyData) SetMaxItem(v int32) *ListProjectsResponseBodyData {
	s.MaxItem = &v
	return s
}

func (s *ListProjectsResponseBodyData) SetProjects(v []*ListProjectsResponseBodyDataProjects) *ListProjectsResponseBodyData {
	s.Projects = v
	return s
}

func (s *ListProjectsResponseBodyData) Validate() error {
	if s.Projects != nil {
		for _, item := range s.Projects {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProjectsResponseBodyDataProjects struct {
	Comment            *string                                                 `json:"comment,omitempty" xml:"comment,omitempty"`
	CostStorage        *string                                                 `json:"costStorage,omitempty" xml:"costStorage,omitempty"`
	CreatedTime        *int64                                                  `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	DefaultQuota       *string                                                 `json:"defaultQuota,omitempty" xml:"defaultQuota,omitempty"`
	IpWhiteList        *ListProjectsResponseBodyDataProjectsIpWhiteList        `json:"ipWhiteList,omitempty" xml:"ipWhiteList,omitempty" type:"Struct"`
	Name               *string                                                 `json:"name,omitempty" xml:"name,omitempty"`
	Owner              *string                                                 `json:"owner,omitempty" xml:"owner,omitempty"`
	Properties         *ListProjectsResponseBodyDataProjectsProperties         `json:"properties,omitempty" xml:"properties,omitempty" type:"Struct"`
	RegionId           *string                                                 `json:"regionId,omitempty" xml:"regionId,omitempty"`
	SaleTag            *ListProjectsResponseBodyDataProjectsSaleTag            `json:"saleTag,omitempty" xml:"saleTag,omitempty" type:"Struct"`
	SecurityProperties *ListProjectsResponseBodyDataProjectsSecurityProperties `json:"securityProperties,omitempty" xml:"securityProperties,omitempty" type:"Struct"`
	Status             *string                                                 `json:"status,omitempty" xml:"status,omitempty"`
	ThreeTierModel     *bool                                                   `json:"threeTierModel,omitempty" xml:"threeTierModel,omitempty"`
	Type               *string                                                 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListProjectsResponseBodyDataProjects) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBodyDataProjects) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyDataProjects) GetComment() *string {
	return s.Comment
}

func (s *ListProjectsResponseBodyDataProjects) GetCostStorage() *string {
	return s.CostStorage
}

func (s *ListProjectsResponseBodyDataProjects) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *ListProjectsResponseBodyDataProjects) GetDefaultQuota() *string {
	return s.DefaultQuota
}

func (s *ListProjectsResponseBodyDataProjects) GetIpWhiteList() *ListProjectsResponseBodyDataProjectsIpWhiteList {
	return s.IpWhiteList
}

func (s *ListProjectsResponseBodyDataProjects) GetName() *string {
	return s.Name
}

func (s *ListProjectsResponseBodyDataProjects) GetOwner() *string {
	return s.Owner
}

func (s *ListProjectsResponseBodyDataProjects) GetProperties() *ListProjectsResponseBodyDataProjectsProperties {
	return s.Properties
}

func (s *ListProjectsResponseBodyDataProjects) GetRegionId() *string {
	return s.RegionId
}

func (s *ListProjectsResponseBodyDataProjects) GetSaleTag() *ListProjectsResponseBodyDataProjectsSaleTag {
	return s.SaleTag
}

func (s *ListProjectsResponseBodyDataProjects) GetSecurityProperties() *ListProjectsResponseBodyDataProjectsSecurityProperties {
	return s.SecurityProperties
}

func (s *ListProjectsResponseBodyDataProjects) GetStatus() *string {
	return s.Status
}

func (s *ListProjectsResponseBodyDataProjects) GetThreeTierModel() *bool {
	return s.ThreeTierModel
}

func (s *ListProjectsResponseBodyDataProjects) GetType() *string {
	return s.Type
}

func (s *ListProjectsResponseBodyDataProjects) SetComment(v string) *ListProjectsResponseBodyDataProjects {
	s.Comment = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetCostStorage(v string) *ListProjectsResponseBodyDataProjects {
	s.CostStorage = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetCreatedTime(v int64) *ListProjectsResponseBodyDataProjects {
	s.CreatedTime = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetDefaultQuota(v string) *ListProjectsResponseBodyDataProjects {
	s.DefaultQuota = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetIpWhiteList(v *ListProjectsResponseBodyDataProjectsIpWhiteList) *ListProjectsResponseBodyDataProjects {
	s.IpWhiteList = v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetName(v string) *ListProjectsResponseBodyDataProjects {
	s.Name = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetOwner(v string) *ListProjectsResponseBodyDataProjects {
	s.Owner = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetProperties(v *ListProjectsResponseBodyDataProjectsProperties) *ListProjectsResponseBodyDataProjects {
	s.Properties = v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetRegionId(v string) *ListProjectsResponseBodyDataProjects {
	s.RegionId = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetSaleTag(v *ListProjectsResponseBodyDataProjectsSaleTag) *ListProjectsResponseBodyDataProjects {
	s.SaleTag = v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetSecurityProperties(v *ListProjectsResponseBodyDataProjectsSecurityProperties) *ListProjectsResponseBodyDataProjects {
	s.SecurityProperties = v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetStatus(v string) *ListProjectsResponseBodyDataProjects {
	s.Status = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetThreeTierModel(v bool) *ListProjectsResponseBodyDataProjects {
	s.ThreeTierModel = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetType(v string) *ListProjectsResponseBodyDataProjects {
	s.Type = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) Validate() error {
	if s.IpWhiteList != nil {
		if err := s.IpWhiteList.Validate(); err != nil {
			return err
		}
	}
	if s.Properties != nil {
		if err := s.Properties.Validate(); err != nil {
			return err
		}
	}
	if s.SaleTag != nil {
		if err := s.SaleTag.Validate(); err != nil {
			return err
		}
	}
	if s.SecurityProperties != nil {
		if err := s.SecurityProperties.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListProjectsResponseBodyDataProjectsIpWhiteList struct {
	IpList    *string `json:"ipList,omitempty" xml:"ipList,omitempty"`
	VpcIpList *string `json:"vpcIpList,omitempty" xml:"vpcIpList,omitempty"`
}

func (s ListProjectsResponseBodyDataProjectsIpWhiteList) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBodyDataProjectsIpWhiteList) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyDataProjectsIpWhiteList) GetIpList() *string {
	return s.IpList
}

func (s *ListProjectsResponseBodyDataProjectsIpWhiteList) GetVpcIpList() *string {
	return s.VpcIpList
}

func (s *ListProjectsResponseBodyDataProjectsIpWhiteList) SetIpList(v string) *ListProjectsResponseBodyDataProjectsIpWhiteList {
	s.IpList = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsIpWhiteList) SetVpcIpList(v string) *ListProjectsResponseBodyDataProjectsIpWhiteList {
	s.VpcIpList = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsIpWhiteList) Validate() error {
	return dara.Validate(s)
}

type ListProjectsResponseBodyDataProjectsProperties struct {
	AllowFullScan             *bool                                                                    `json:"allowFullScan,omitempty" xml:"allowFullScan,omitempty"`
	EnableDecimal2            *bool                                                                    `json:"enableDecimal2,omitempty" xml:"enableDecimal2,omitempty"`
	EnableTunnelQuotaRoute    *bool                                                                    `json:"enableTunnelQuotaRoute,omitempty" xml:"enableTunnelQuotaRoute,omitempty"`
	Encryption                *ListProjectsResponseBodyDataProjectsPropertiesEncryption                `json:"encryption,omitempty" xml:"encryption,omitempty" type:"Struct"`
	ExternalProjectProperties *ListProjectsResponseBodyDataProjectsPropertiesExternalProjectProperties `json:"externalProjectProperties,omitempty" xml:"externalProjectProperties,omitempty" type:"Struct"`
	RetentionDays             *int64                                                                   `json:"retentionDays,omitempty" xml:"retentionDays,omitempty"`
	SqlMeteringMax            *string                                                                  `json:"sqlMeteringMax,omitempty" xml:"sqlMeteringMax,omitempty"`
	TableLifecycle            *ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle            `json:"tableLifecycle,omitempty" xml:"tableLifecycle,omitempty" type:"Struct"`
	Timezone                  *string                                                                  `json:"timezone,omitempty" xml:"timezone,omitempty"`
	TunnelQuota               *string                                                                  `json:"tunnelQuota,omitempty" xml:"tunnelQuota,omitempty"`
	TypeSystem                *string                                                                  `json:"typeSystem,omitempty" xml:"typeSystem,omitempty"`
}

func (s ListProjectsResponseBodyDataProjectsProperties) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBodyDataProjectsProperties) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyDataProjectsProperties) GetAllowFullScan() *bool {
	return s.AllowFullScan
}

func (s *ListProjectsResponseBodyDataProjectsProperties) GetEnableDecimal2() *bool {
	return s.EnableDecimal2
}

func (s *ListProjectsResponseBodyDataProjectsProperties) GetEnableTunnelQuotaRoute() *bool {
	return s.EnableTunnelQuotaRoute
}

func (s *ListProjectsResponseBodyDataProjectsProperties) GetEncryption() *ListProjectsResponseBodyDataProjectsPropertiesEncryption {
	return s.Encryption
}

func (s *ListProjectsResponseBodyDataProjectsProperties) GetExternalProjectProperties() *ListProjectsResponseBodyDataProjectsPropertiesExternalProjectProperties {
	return s.ExternalProjectProperties
}

func (s *ListProjectsResponseBodyDataProjectsProperties) GetRetentionDays() *int64 {
	return s.RetentionDays
}

func (s *ListProjectsResponseBodyDataProjectsProperties) GetSqlMeteringMax() *string {
	return s.SqlMeteringMax
}

func (s *ListProjectsResponseBodyDataProjectsProperties) GetTableLifecycle() *ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle {
	return s.TableLifecycle
}

func (s *ListProjectsResponseBodyDataProjectsProperties) GetTimezone() *string {
	return s.Timezone
}

func (s *ListProjectsResponseBodyDataProjectsProperties) GetTunnelQuota() *string {
	return s.TunnelQuota
}

func (s *ListProjectsResponseBodyDataProjectsProperties) GetTypeSystem() *string {
	return s.TypeSystem
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetAllowFullScan(v bool) *ListProjectsResponseBodyDataProjectsProperties {
	s.AllowFullScan = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetEnableDecimal2(v bool) *ListProjectsResponseBodyDataProjectsProperties {
	s.EnableDecimal2 = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetEnableTunnelQuotaRoute(v bool) *ListProjectsResponseBodyDataProjectsProperties {
	s.EnableTunnelQuotaRoute = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetEncryption(v *ListProjectsResponseBodyDataProjectsPropertiesEncryption) *ListProjectsResponseBodyDataProjectsProperties {
	s.Encryption = v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetExternalProjectProperties(v *ListProjectsResponseBodyDataProjectsPropertiesExternalProjectProperties) *ListProjectsResponseBodyDataProjectsProperties {
	s.ExternalProjectProperties = v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetRetentionDays(v int64) *ListProjectsResponseBodyDataProjectsProperties {
	s.RetentionDays = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetSqlMeteringMax(v string) *ListProjectsResponseBodyDataProjectsProperties {
	s.SqlMeteringMax = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetTableLifecycle(v *ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle) *ListProjectsResponseBodyDataProjectsProperties {
	s.TableLifecycle = v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetTimezone(v string) *ListProjectsResponseBodyDataProjectsProperties {
	s.Timezone = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetTunnelQuota(v string) *ListProjectsResponseBodyDataProjectsProperties {
	s.TunnelQuota = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetTypeSystem(v string) *ListProjectsResponseBodyDataProjectsProperties {
	s.TypeSystem = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) Validate() error {
	if s.Encryption != nil {
		if err := s.Encryption.Validate(); err != nil {
			return err
		}
	}
	if s.ExternalProjectProperties != nil {
		if err := s.ExternalProjectProperties.Validate(); err != nil {
			return err
		}
	}
	if s.TableLifecycle != nil {
		if err := s.TableLifecycle.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListProjectsResponseBodyDataProjectsPropertiesEncryption struct {
	Algorithm *string `json:"algorithm,omitempty" xml:"algorithm,omitempty"`
	Enable    *bool   `json:"enable,omitempty" xml:"enable,omitempty"`
	Key       *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s ListProjectsResponseBodyDataProjectsPropertiesEncryption) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBodyDataProjectsPropertiesEncryption) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesEncryption) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesEncryption) GetEnable() *bool {
	return s.Enable
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesEncryption) GetKey() *string {
	return s.Key
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesEncryption) SetAlgorithm(v string) *ListProjectsResponseBodyDataProjectsPropertiesEncryption {
	s.Algorithm = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesEncryption) SetEnable(v bool) *ListProjectsResponseBodyDataProjectsPropertiesEncryption {
	s.Enable = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesEncryption) SetKey(v string) *ListProjectsResponseBodyDataProjectsPropertiesEncryption {
	s.Key = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesEncryption) Validate() error {
	return dara.Validate(s)
}

type ListProjectsResponseBodyDataProjectsPropertiesExternalProjectProperties struct {
	IsExternalCatalogBound *string `json:"isExternalCatalogBound,omitempty" xml:"isExternalCatalogBound,omitempty"`
}

func (s ListProjectsResponseBodyDataProjectsPropertiesExternalProjectProperties) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBodyDataProjectsPropertiesExternalProjectProperties) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesExternalProjectProperties) GetIsExternalCatalogBound() *string {
	return s.IsExternalCatalogBound
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesExternalProjectProperties) SetIsExternalCatalogBound(v string) *ListProjectsResponseBodyDataProjectsPropertiesExternalProjectProperties {
	s.IsExternalCatalogBound = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesExternalProjectProperties) Validate() error {
	return dara.Validate(s)
}

type ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle struct {
	Type  *string `json:"type,omitempty" xml:"type,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle) GetType() *string {
	return s.Type
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle) GetValue() *string {
	return s.Value
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle) SetType(v string) *ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle {
	s.Type = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle) SetValue(v string) *ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle {
	s.Value = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle) Validate() error {
	return dara.Validate(s)
}

type ListProjectsResponseBodyDataProjectsSaleTag struct {
	ResourceId   *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ListProjectsResponseBodyDataProjectsSaleTag) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBodyDataProjectsSaleTag) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyDataProjectsSaleTag) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListProjectsResponseBodyDataProjectsSaleTag) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListProjectsResponseBodyDataProjectsSaleTag) SetResourceId(v string) *ListProjectsResponseBodyDataProjectsSaleTag {
	s.ResourceId = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsSaleTag) SetResourceType(v string) *ListProjectsResponseBodyDataProjectsSaleTag {
	s.ResourceType = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsSaleTag) Validate() error {
	return dara.Validate(s)
}

type ListProjectsResponseBodyDataProjectsSecurityProperties struct {
	EnableDownloadPrivilege          *bool                                                                    `json:"enableDownloadPrivilege,omitempty" xml:"enableDownloadPrivilege,omitempty"`
	LabelSecurity                    *bool                                                                    `json:"labelSecurity,omitempty" xml:"labelSecurity,omitempty"`
	ObjectCreatorHasAccessPermission *bool                                                                    `json:"objectCreatorHasAccessPermission,omitempty" xml:"objectCreatorHasAccessPermission,omitempty"`
	ObjectCreatorHasGrantPermission  *bool                                                                    `json:"objectCreatorHasGrantPermission,omitempty" xml:"objectCreatorHasGrantPermission,omitempty"`
	ProjectProtection                *ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection `json:"projectProtection,omitempty" xml:"projectProtection,omitempty" type:"Struct"`
	UsingAcl                         *bool                                                                    `json:"usingAcl,omitempty" xml:"usingAcl,omitempty"`
	UsingPolicy                      *bool                                                                    `json:"usingPolicy,omitempty" xml:"usingPolicy,omitempty"`
}

func (s ListProjectsResponseBodyDataProjectsSecurityProperties) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBodyDataProjectsSecurityProperties) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) GetEnableDownloadPrivilege() *bool {
	return s.EnableDownloadPrivilege
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) GetLabelSecurity() *bool {
	return s.LabelSecurity
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) GetObjectCreatorHasAccessPermission() *bool {
	return s.ObjectCreatorHasAccessPermission
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) GetObjectCreatorHasGrantPermission() *bool {
	return s.ObjectCreatorHasGrantPermission
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) GetProjectProtection() *ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection {
	return s.ProjectProtection
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) GetUsingAcl() *bool {
	return s.UsingAcl
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) GetUsingPolicy() *bool {
	return s.UsingPolicy
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) SetEnableDownloadPrivilege(v bool) *ListProjectsResponseBodyDataProjectsSecurityProperties {
	s.EnableDownloadPrivilege = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) SetLabelSecurity(v bool) *ListProjectsResponseBodyDataProjectsSecurityProperties {
	s.LabelSecurity = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) SetObjectCreatorHasAccessPermission(v bool) *ListProjectsResponseBodyDataProjectsSecurityProperties {
	s.ObjectCreatorHasAccessPermission = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) SetObjectCreatorHasGrantPermission(v bool) *ListProjectsResponseBodyDataProjectsSecurityProperties {
	s.ObjectCreatorHasGrantPermission = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) SetProjectProtection(v *ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection) *ListProjectsResponseBodyDataProjectsSecurityProperties {
	s.ProjectProtection = v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) SetUsingAcl(v bool) *ListProjectsResponseBodyDataProjectsSecurityProperties {
	s.UsingAcl = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) SetUsingPolicy(v bool) *ListProjectsResponseBodyDataProjectsSecurityProperties {
	s.UsingPolicy = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) Validate() error {
	if s.ProjectProtection != nil {
		if err := s.ProjectProtection.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection struct {
	ExceptionPolicy *string `json:"exceptionPolicy,omitempty" xml:"exceptionPolicy,omitempty"`
	Protected       *bool   `json:"protected,omitempty" xml:"protected,omitempty"`
}

func (s ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection) GetExceptionPolicy() *string {
	return s.ExceptionPolicy
}

func (s *ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection) GetProtected() *bool {
	return s.Protected
}

func (s *ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection) SetExceptionPolicy(v string) *ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection {
	s.ExceptionPolicy = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection) SetProtected(v bool) *ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection {
	s.Protected = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection) Validate() error {
	return dara.Validate(s)
}
