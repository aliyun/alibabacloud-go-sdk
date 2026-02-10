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
  // The additional type of the vulnerabilities. You need to specify this parameter when you query application vulnerabilities. If you set the Type parameter to app, you must specify this parameter. Set the value to **sca**.
  // 
  // > If this parameter is set to **sca**, **application vulnerabilities*	- and the **vulnerabilities that are detected based on software component analysis*	- are queried. If you do not specify this parameter, only application vulnerabilities are queried.
  // 
  // example:
  // 
  // sca
  AttachTypes *string `json:"AttachTypes,omitempty" xml:"AttachTypes,omitempty"`
  // The name of the container that is affected by the vulnerability.
  // 
  // example:
  // 
  // xxljob-7b87597b99-mcskr
  ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
  // The end time of the first scan.
  // 
  // >  This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
  // 
  // example:
  // 
  // 1696186800000
  CreateTsEnd *int64 `json:"CreateTsEnd,omitempty" xml:"CreateTsEnd,omitempty"`
  // The start time of the first scan.
  // 
  // >  This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
  // 
  // example:
  // 
  // 1696128144000
  CreateTsStart *int64 `json:"CreateTsStart,omitempty" xml:"CreateTsStart,omitempty"`
  // The Common Vulnerabilities and Exposures (CVE) ID of the vulnerability.
  // 
  // example:
  // 
  // CVE-2022-44702
  CveId *string `json:"CveId,omitempty" xml:"CveId,omitempty"`
  // Specifies whether the vulnerability is fixed. Valid values:
  // 
  // 	- **y**: The vulnerability is fixed.
  // 
  // 	- **n**: The vulnerability is not fixed.
  // 
  // example:
  // 
  // n
  Dealed *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
  // The server group ID of the server on which the vulnerabilities are detected.
  // 
  // > You can call the [DescribeAllGroups](~~DescribeAllGroups~~) operation to query the IDs of server groups.
  // 
  // example:
  // 
  // 8834224
  GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
  // The name of the image that is affected by the vulnerability.
  // 
  // example:
  // 
  // container-***:****
  ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
  // The language of the content within the request and response. Default value: **zh**. Valid values:
  // 
  // 	- zh: Chinese
  // 
  // 	- en: English
  // 
  // example:
  // 
  // zh
  Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
  // The priority to fix the vulnerability. Separate multiple priorities with commas (,). Valid values:
  // 
  // 	- **asap**: high
  // 
  // 	- **later**: medium
  // 
  // 	- **nntf**: low
  // 
  // example:
  // 
  // asap
  Necessity *string `json:"Necessity,omitempty" xml:"Necessity,omitempty"`
  // The path of the process that is affected by the vulnerability.
  // 
  // example:
  // 
  // /etc/test
  Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
  // Indicates whether the application protection feature is supported. Valid values:
  // 
  // - **0**: no.
  // 
  // - **1**: yes.
  // 
  // example:
  // 
  // 0
  RaspDefend *int32 `json:"RaspDefend,omitempty" xml:"RaspDefend,omitempty"`
  // The tag that is used to search for the vulnerabilities. Valid values:
  // 
  // 	- Restart required
  // 
  // 	- Remote exploitation
  // 
  // 	- Exploit exists
  // 
  // 	- Exploitable
  // 
  // 	- Privilege escalation
  // 
  // 	- Code execution
  // 
  // example:
  // 
  // Restart required
  SearchTags *string `json:"SearchTags,omitempty" xml:"SearchTags,omitempty"`
  // The type of the vulnerability that you want to export. Valid values:
  // 
  // 	- **cve**: Linux software vulnerability
  // 
  // 	- **sys**: Windows system vulnerability
  // 
  // 	- **cms**: Web-CMS vulnerability
  // 
  // 	- **app**: application vulnerability
  // 
  // 	- **emg**: urgent vulnerability
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // app
  Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
  // The UUID of the server on which the vulnerabilities are detected. Separate multiple UUIDs with commas (,).
  // 
  // example:
  // 
  // 1587bedb-fdb4-48c4-9330-****
  Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
  // The ID of the virtual private cloud (VPC) in which the vulnerabilities are detected. Separate multiple IDs with commas (,).
  // 
  // > You can call the [DescribeVpcList](~~DescribeVpcList~~) operation to query the IDs of VPCs.
  // 
  // example:
  // 
  // ins-133****,ins-5414****
  VpcInstanceIds *string `json:"VpcInstanceIds,omitempty" xml:"VpcInstanceIds,omitempty"`
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
  // example:
  // 
  // Ollama
  EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
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

