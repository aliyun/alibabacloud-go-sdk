// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportVulRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAliasName(v string) *ExportVulRequest
  GetAliasName() *string 
  SetAssetType(v string) *ExportVulRequest
  GetAssetType() *string 
  SetAttachTypes(v string) *ExportVulRequest
  GetAttachTypes() *string 
  SetContainerName(v string) *ExportVulRequest
  GetContainerName() *string 
  SetCreateTsEnd(v int64) *ExportVulRequest
  GetCreateTsEnd() *int64 
  SetCreateTsStart(v int64) *ExportVulRequest
  GetCreateTsStart() *int64 
  SetCveId(v string) *ExportVulRequest
  GetCveId() *string 
  SetDealed(v string) *ExportVulRequest
  GetDealed() *string 
  SetGroupId(v string) *ExportVulRequest
  GetGroupId() *string 
  SetImageName(v string) *ExportVulRequest
  GetImageName() *string 
  SetLang(v string) *ExportVulRequest
  GetLang() *string 
  SetNecessity(v string) *ExportVulRequest
  GetNecessity() *string 
  SetPath(v string) *ExportVulRequest
  GetPath() *string 
  SetRaspDefend(v int32) *ExportVulRequest
  GetRaspDefend() *int32 
  SetResourceDirectoryAccountId(v int64) *ExportVulRequest
  GetResourceDirectoryAccountId() *int64 
  SetSearchTags(v string) *ExportVulRequest
  GetSearchTags() *string 
  SetType(v string) *ExportVulRequest
  GetType() *string 
  SetUuids(v string) *ExportVulRequest
  GetUuids() *string 
  SetVpcInstanceIds(v string) *ExportVulRequest
  GetVpcInstanceIds() *string 
  SetVulEntityList(v []*ExportVulRequestVulEntityList) *ExportVulRequest
  GetVulEntityList() []*ExportVulRequestVulEntityList 
}

type ExportVulRequest struct {
  // The name of the vulnerability.
  // 
  // example:
  // 
  // RHSA-2019:3197-Important: sudo security update
  AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
  // The asset type where the vulnerability is detected. Separate multiple types with commas (,). Valid values:
  // 
  // - **ECS**: host asset
  // 
  // - **CONTAINER**: container asset
  // 
  // example:
  // 
  // ECS
  AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
  // The additional vulnerability type when querying application vulnerabilities. This parameter is required when Type is set to app. The value is fixed as **sca**.
  // 
  // > If this parameter is set to **sca**, both application vulnerabilities (**app*	- type) and software composition analysis (**sca*	- type) vulnerabilities are queried. If this parameter is not set, only application vulnerabilities are queried.
  // 
  // example:
  // 
  // sca
  AttachTypes *string `json:"AttachTypes,omitempty" xml:"AttachTypes,omitempty"`
  // The name of the container affected by the vulnerability.
  // 
  // example:
  // 
  // xxljob-7b87597b99-mcskr
  ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
  // The end of the time range during which the first scan was performed.
  // 
  // > The value is a UNIX timestamp. Unit: milliseconds.
  // 
  // example:
  // 
  // 1696186800000
  CreateTsEnd *int64 `json:"CreateTsEnd,omitempty" xml:"CreateTsEnd,omitempty"`
  // The start of the time range during which the first scan was performed.
  // 
  // > The value is a UNIX timestamp. Unit: milliseconds.
  // 
  // example:
  // 
  // 1696128144000
  CreateTsStart *int64 `json:"CreateTsStart,omitempty" xml:"CreateTsStart,omitempty"`
  // The CVE ID.
  // 
  // example:
  // 
  // CVE-2022-44702
  CveId *string `json:"CveId,omitempty" xml:"CveId,omitempty"`
  // Specifies whether the vulnerability is fixed. Valid values:
  // 
  // - **y**: fixed
  // 
  // - **n**: not fixed
  // 
  // example:
  // 
  // n
  Dealed *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
  // The ID of the asset group to which the server with the vulnerability belongs.
  // 
  // > Call the [DescribeAllGroups](~~DescribeAllGroups~~) operation to obtain this parameter.
  // 
  // example:
  // 
  // 8834224
  GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
  // The name of the image affected by the vulnerability.
  // 
  // example:
  // 
  // container-***:****
  ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
  // The language of the content within the request and response. Default value: **zh**. Valid values:
  // 
  // - zh: Chinese
  // 
  // - en: English
  // 
  // example:
  // 
  // zh
  Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
  // The priority of the vulnerability to query. Separate multiple priorities with commas (,). Valid values:
  // 
  // - **asap**: high
  // 
  // - **later**: medium
  // 
  // - **nntf**: low
  // 
  // example:
  // 
  // asap
  Necessity *string `json:"Necessity,omitempty" xml:"Necessity,omitempty"`
  // The path of the process affected by the vulnerability.
  // 
  // example:
  // 
  // /etc/test
  Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
  // Specifies whether runtime application self-protection (RASP) supports real-time protection against the vulnerability. Valid values:
  // 
  // - **0**: Not supported.
  // 
  // - **1**: Supported.
  // 
  // example:
  // 
  // 0
  RaspDefend *int32 `json:"RaspDefend,omitempty" xml:"RaspDefend,omitempty"`
  // The ID of the resource directory account.
  // 
  // example:
  // 
  // 1
  ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
  // Filters results by label. Valid values:
  // 
  // <props="china">
  // 
  // - Restart required
  // 
  // - Remote utilization
  // 
  // - EXP exists
  // 
  // - Exploitable
  // 
  // - Privilege escalation
  // 
  // - Code execution
  // 
  // 
  // <props="intl">
  // 
  // - **Restart required**
  // 
  // - **Remote utilization**
  // 
  // - **EXP exists**
  // 
  // - **Available**
  // 
  // - **Elevation of Privilege**
  // 
  // - **Code Execution**
  // 
  // example:
  // 
  // Restart required
  SearchTags *string `json:"SearchTags,omitempty" xml:"SearchTags,omitempty"`
  // The type of vulnerabilities to export. Valid values:
  // 
  // - **cve**: Linux software vulnerability
  // 
  // - **sys**: Windows system vulnerability
  // 
  // - **cms**: Web-CMS vulnerability
  // 
  // - **app**: application vulnerability
  // 
  // - **emg**: emergency vulnerability
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // app
  Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
  // The UUIDs of the servers to query for vulnerabilities. Separate multiple UUIDs with commas (,).
  // 
  // example:
  // 
  // 1587bedb-fdb4-48c4-9330-****
  Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
  // The instance IDs of the VPC-connected instances to query for vulnerabilities. Separate multiple IDs with commas (,).
  // 
  // > Invoke the [DescribeVpcList](~~DescribeVpcList~~) operation to obtain this parameter.
  // 
  // example:
  // 
  // ins-133****,ins-5414****
  VpcInstanceIds *string `json:"VpcInstanceIds,omitempty" xml:"VpcInstanceIds,omitempty"`
  // The list of vulnerability component information.
  VulEntityList []*ExportVulRequestVulEntityList `json:"VulEntityList,omitempty" xml:"VulEntityList,omitempty" type:"Repeated"`
}

func (s ExportVulRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportVulRequest) GoString() string {
  return s.String()
}

func (s *ExportVulRequest) GetAliasName() *string  {
  return s.AliasName
}

func (s *ExportVulRequest) GetAssetType() *string  {
  return s.AssetType
}

func (s *ExportVulRequest) GetAttachTypes() *string  {
  return s.AttachTypes
}

func (s *ExportVulRequest) GetContainerName() *string  {
  return s.ContainerName
}

func (s *ExportVulRequest) GetCreateTsEnd() *int64  {
  return s.CreateTsEnd
}

func (s *ExportVulRequest) GetCreateTsStart() *int64  {
  return s.CreateTsStart
}

func (s *ExportVulRequest) GetCveId() *string  {
  return s.CveId
}

func (s *ExportVulRequest) GetDealed() *string  {
  return s.Dealed
}

func (s *ExportVulRequest) GetGroupId() *string  {
  return s.GroupId
}

func (s *ExportVulRequest) GetImageName() *string  {
  return s.ImageName
}

func (s *ExportVulRequest) GetLang() *string  {
  return s.Lang
}

func (s *ExportVulRequest) GetNecessity() *string  {
  return s.Necessity
}

func (s *ExportVulRequest) GetPath() *string  {
  return s.Path
}

func (s *ExportVulRequest) GetRaspDefend() *int32  {
  return s.RaspDefend
}

func (s *ExportVulRequest) GetResourceDirectoryAccountId() *int64  {
  return s.ResourceDirectoryAccountId
}

func (s *ExportVulRequest) GetSearchTags() *string  {
  return s.SearchTags
}

func (s *ExportVulRequest) GetType() *string  {
  return s.Type
}

func (s *ExportVulRequest) GetUuids() *string  {
  return s.Uuids
}

func (s *ExportVulRequest) GetVpcInstanceIds() *string  {
  return s.VpcInstanceIds
}

func (s *ExportVulRequest) GetVulEntityList() []*ExportVulRequestVulEntityList  {
  return s.VulEntityList
}

func (s *ExportVulRequest) SetAliasName(v string) *ExportVulRequest {
  s.AliasName = &v
  return s
}

func (s *ExportVulRequest) SetAssetType(v string) *ExportVulRequest {
  s.AssetType = &v
  return s
}

func (s *ExportVulRequest) SetAttachTypes(v string) *ExportVulRequest {
  s.AttachTypes = &v
  return s
}

func (s *ExportVulRequest) SetContainerName(v string) *ExportVulRequest {
  s.ContainerName = &v
  return s
}

func (s *ExportVulRequest) SetCreateTsEnd(v int64) *ExportVulRequest {
  s.CreateTsEnd = &v
  return s
}

func (s *ExportVulRequest) SetCreateTsStart(v int64) *ExportVulRequest {
  s.CreateTsStart = &v
  return s
}

func (s *ExportVulRequest) SetCveId(v string) *ExportVulRequest {
  s.CveId = &v
  return s
}

func (s *ExportVulRequest) SetDealed(v string) *ExportVulRequest {
  s.Dealed = &v
  return s
}

func (s *ExportVulRequest) SetGroupId(v string) *ExportVulRequest {
  s.GroupId = &v
  return s
}

func (s *ExportVulRequest) SetImageName(v string) *ExportVulRequest {
  s.ImageName = &v
  return s
}

func (s *ExportVulRequest) SetLang(v string) *ExportVulRequest {
  s.Lang = &v
  return s
}

func (s *ExportVulRequest) SetNecessity(v string) *ExportVulRequest {
  s.Necessity = &v
  return s
}

func (s *ExportVulRequest) SetPath(v string) *ExportVulRequest {
  s.Path = &v
  return s
}

func (s *ExportVulRequest) SetRaspDefend(v int32) *ExportVulRequest {
  s.RaspDefend = &v
  return s
}

func (s *ExportVulRequest) SetResourceDirectoryAccountId(v int64) *ExportVulRequest {
  s.ResourceDirectoryAccountId = &v
  return s
}

func (s *ExportVulRequest) SetSearchTags(v string) *ExportVulRequest {
  s.SearchTags = &v
  return s
}

func (s *ExportVulRequest) SetType(v string) *ExportVulRequest {
  s.Type = &v
  return s
}

func (s *ExportVulRequest) SetUuids(v string) *ExportVulRequest {
  s.Uuids = &v
  return s
}

func (s *ExportVulRequest) SetVpcInstanceIds(v string) *ExportVulRequest {
  s.VpcInstanceIds = &v
  return s
}

func (s *ExportVulRequest) SetVulEntityList(v []*ExportVulRequestVulEntityList) *ExportVulRequest {
  s.VulEntityList = v
  return s
}

func (s *ExportVulRequest) Validate() error {
  if s.VulEntityList != nil {
    for _, item := range s.VulEntityList {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type ExportVulRequestVulEntityList struct {
  // The name of the component.
  // 
  // example:
  // 
  // Ollama
  EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
  // The version of the component.
  // 
  // example:
  // 
  // 1.0.0
  EntityVersion *string `json:"EntityVersion,omitempty" xml:"EntityVersion,omitempty"`
}

func (s ExportVulRequestVulEntityList) String() string {
  return dara.Prettify(s)
}

func (s ExportVulRequestVulEntityList) GoString() string {
  return s.String()
}

func (s *ExportVulRequestVulEntityList) GetEntityName() *string  {
  return s.EntityName
}

func (s *ExportVulRequestVulEntityList) GetEntityVersion() *string  {
  return s.EntityVersion
}

func (s *ExportVulRequestVulEntityList) SetEntityName(v string) *ExportVulRequestVulEntityList {
  s.EntityName = &v
  return s
}

func (s *ExportVulRequestVulEntityList) SetEntityVersion(v string) *ExportVulRequestVulEntityList {
  s.EntityVersion = &v
  return s
}

func (s *ExportVulRequestVulEntityList) Validate() error {
  return dara.Validate(s)
}

