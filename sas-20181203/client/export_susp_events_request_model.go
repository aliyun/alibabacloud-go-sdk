// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportSuspEventsRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAssetsTypeList(v []*string) *ExportSuspEventsRequest
  GetAssetsTypeList() []*string 
  SetClusterId(v string) *ExportSuspEventsRequest
  GetClusterId() *string 
  SetContainerFieldName(v string) *ExportSuspEventsRequest
  GetContainerFieldName() *string 
  SetContainerFieldValue(v string) *ExportSuspEventsRequest
  GetContainerFieldValue() *string 
  SetCurrentPage(v string) *ExportSuspEventsRequest
  GetCurrentPage() *string 
  SetDealed(v string) *ExportSuspEventsRequest
  GetDealed() *string 
  SetFrom(v string) *ExportSuspEventsRequest
  GetFrom() *string 
  SetGroupId(v int64) *ExportSuspEventsRequest
  GetGroupId() *int64 
  SetId(v int64) *ExportSuspEventsRequest
  GetId() *int64 
  SetLang(v string) *ExportSuspEventsRequest
  GetLang() *string 
  SetLevels(v string) *ExportSuspEventsRequest
  GetLevels() *string 
  SetName(v string) *ExportSuspEventsRequest
  GetName() *string 
  SetOperateErrorCodeList(v []*string) *ExportSuspEventsRequest
  GetOperateErrorCodeList() []*string 
  SetPageSize(v string) *ExportSuspEventsRequest
  GetPageSize() *string 
  SetParentEventTypes(v string) *ExportSuspEventsRequest
  GetParentEventTypes() *string 
  SetRemark(v string) *ExportSuspEventsRequest
  GetRemark() *string 
  SetSourceIp(v string) *ExportSuspEventsRequest
  GetSourceIp() *string 
  SetStatus(v string) *ExportSuspEventsRequest
  GetStatus() *string 
  SetTargetType(v string) *ExportSuspEventsRequest
  GetTargetType() *string 
  SetTimeEnd(v string) *ExportSuspEventsRequest
  GetTimeEnd() *string 
  SetTimeStart(v string) *ExportSuspEventsRequest
  GetTimeStart() *string 
  SetUniqueInfo(v string) *ExportSuspEventsRequest
  GetUniqueInfo() *string 
  SetUuid(v string) *ExportSuspEventsRequest
  GetUuid() *string 
}

type ExportSuspEventsRequest struct {
  // The types of assets.
  AssetsTypeList []*string `json:"AssetsTypeList,omitempty" xml:"AssetsTypeList,omitempty" type:"Repeated"`
  // The ID of the cluster that you want to query.
  // 
  // > You can call the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) operation to query the IDs of clusters.
  // 
  // example:
  // 
  // c4af4fdf38a98496a9b63c2be5dae****
  ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
  // The key of the condition that is used to query alert events on containers. Valid values:
  // 
  // 	- **instanceId**: the ID of the asset
  // 
  // 	- **appName**: the name of the application
  // 
  // 	- **clusterId**: the ID of the cluster
  // 
  // 	- **regionId**: the ID of the region
  // 
  // 	- **nodeName**: the name of the node
  // 
  // 	- **namespace**: the namespace
  // 
  // 	- **clusterName**: the name of the cluster
  // 
  // 	- **image**: the name of the image
  // 
  // 	- **imageRepoName**: the name of the image repository
  // 
  // 	- **imageRepoNamespace**: the namespace to which the image repository belongs
  // 
  // 	- **imageRepoTag**: the tag that is added to the image
  // 
  // 	- **imageDigest**: the digest of the image
  // 
  // example:
  // 
  // clusterId
  ContainerFieldName *string `json:"ContainerFieldName,omitempty" xml:"ContainerFieldName,omitempty"`
  // The value of the condition that is used to query alert events on containers.
  // 
  // example:
  // 
  // c819391d2d520485fa3e81e2dc2ea****
  ContainerFieldValue *string `json:"ContainerFieldValue,omitempty" xml:"ContainerFieldValue,omitempty"`
  // The number of the page to return.
  // 
  // example:
  // 
  // 1
  CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
  // The status of the alert event. Valid values:
  // 
  // 	- **N**: unhandled
  // 
  // 	- **Y**: handled
  // 
  // example:
  // 
  // Y
  Dealed *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
  // The data source of the exception. Set the value to sas.
  // 
  // example:
  // 
  // sas
  From *string `json:"From,omitempty" xml:"From,omitempty"`
  // The ID of the asset group.
  // 
  // example:
  // 
  // 8076980
  GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
  // The unique ID of the alert event.
  // 
  // example:
  // 
  // 17821
  Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
  // The language of the content within the request and response. Default value: **zh**. Valid values:
  // 
  // 	- **zh**: Chinese
  // 
  // 	- **en**: English
  // 
  // example:
  // 
  // zh
  Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
  // The severity of the alert event. Separate multiple severities with commas (,). Valid values:
  // 
  // 	- **serious**
  // 
  // 	- **suspicious**
  // 
  // 	- **remind**
  // 
  // example:
  // 
  // serious,suspicious,remind
  Levels *string `json:"Levels,omitempty" xml:"Levels,omitempty"`
  // The complete name of the exception.
  // 
  // example:
  // 
  // WEBSHELL
  Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
  // The status codes of alert events.
  OperateErrorCodeList []*string `json:"OperateErrorCodeList,omitempty" xml:"OperateErrorCodeList,omitempty" type:"Repeated"`
  // The number of entries to return on each page. Default value: **20**.
  // 
  // example:
  // 
  // 20
  PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
  // The alert type of the alert event. Valid values:
  // 
  // 	- **Suspicious process**
  // 
  // 	- **Webshell**
  // 
  // 	- **Unusual logon**
  // 
  // 	- **Exception**
  // 
  // 	- **Sensitive file tampering**
  // 
  // 	- **Malicious process (cloud threat detection)**
  // 
  // 	- **Suspicious network connection**
  // 
  // 	- **Suspicious account**
  // 
  // 	- **Application intrusion event**
  // 
  // 	- **Cloud threat detection**
  // 
  // 	- **Precise defense**
  // 
  // 	- **Application whitelist**
  // 
  // 	- **Persistent webshell**
  // 
  // 	- **Web application threat detection**
  // 
  // 	- **Malicious script**
  // 
  // 	- **Threat intelligence**
  // 
  // 	- **Malicious network activity**
  // 
  // 	- **Cluster exception**
  // 
  // 	- **Webshell (on-premises threat detection)**
  // 
  // 	- **Vulnerability exploitation**
  // 
  // 	- **Malicious process (on-premises threat detection)**
  // 
  // 	- **Trusted exception**
  // 
  // 	- **Others**
  // 
  // example:
  // 
  // WEBSHELL
  ParentEventTypes *string `json:"ParentEventTypes,omitempty" xml:"ParentEventTypes,omitempty"`
  // The remarks.
  // 
  // example:
  // 
  // remark
  Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
  // The source IP address of the request. The value of this parameter is specified by the system.
  // 
  // example:
  // 
  // 127.0.XX.XX
  SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
  // The handling status of the exception. Valid values:
  // 
  // 	- **0**: all status
  // 
  // 	- **1**: pending handling
  // 
  // 	- **2**: ignored
  // 
  // 	- **4**: confirmed
  // 
  // 	- **8**: marked as false positive
  // 
  // 	- **16**: handling
  // 
  // 	- **32**: handled
  // 
  // 	- **64**: expired
  // 
  // 	- **128**: deleted
  // 
  // example:
  // 
  // 0
  Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
  // The dimension from which you want to configure the feature. Valid values:
  // 
  // 	- **uuid**: the UUID of the asset
  // 
  // 	- **image_repo**: the ID of the image repository
  // 
  // 	- **Cluster**: the ID of the cluster
  // 
  // example:
  // 
  // uuid
  TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
  // The end of the time range during which the exception is detected.
  // 
  // example:
  // 
  // 2022-12-05 00:00:00
  TimeEnd *string `json:"TimeEnd,omitempty" xml:"TimeEnd,omitempty"`
  // The beginning of the time range during which the exception is detected.
  // 
  // example:
  // 
  // 2022-10-01 00:00:00
  TimeStart *string `json:"TimeStart,omitempty" xml:"TimeStart,omitempty"`
  // The unique key of the alert event.
  // 
  // example:
  // 
  // 1fbe8d16727f61d1478a674d6fa0****
  UniqueInfo *string `json:"UniqueInfo,omitempty" xml:"UniqueInfo,omitempty"`
  // The unique ID of the associated instance.
  // 
  // example:
  // 
  // 18b7336e-d469-473b-af83-8e5420f9****
  Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ExportSuspEventsRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportSuspEventsRequest) GoString() string {
  return s.String()
}

func (s *ExportSuspEventsRequest) GetAssetsTypeList() []*string  {
  return s.AssetsTypeList
}

func (s *ExportSuspEventsRequest) GetClusterId() *string  {
  return s.ClusterId
}

func (s *ExportSuspEventsRequest) GetContainerFieldName() *string  {
  return s.ContainerFieldName
}

func (s *ExportSuspEventsRequest) GetContainerFieldValue() *string  {
  return s.ContainerFieldValue
}

func (s *ExportSuspEventsRequest) GetCurrentPage() *string  {
  return s.CurrentPage
}

func (s *ExportSuspEventsRequest) GetDealed() *string  {
  return s.Dealed
}

func (s *ExportSuspEventsRequest) GetFrom() *string  {
  return s.From
}

func (s *ExportSuspEventsRequest) GetGroupId() *int64  {
  return s.GroupId
}

func (s *ExportSuspEventsRequest) GetId() *int64  {
  return s.Id
}

func (s *ExportSuspEventsRequest) GetLang() *string  {
  return s.Lang
}

func (s *ExportSuspEventsRequest) GetLevels() *string  {
  return s.Levels
}

func (s *ExportSuspEventsRequest) GetName() *string  {
  return s.Name
}

func (s *ExportSuspEventsRequest) GetOperateErrorCodeList() []*string  {
  return s.OperateErrorCodeList
}

func (s *ExportSuspEventsRequest) GetPageSize() *string  {
  return s.PageSize
}

func (s *ExportSuspEventsRequest) GetParentEventTypes() *string  {
  return s.ParentEventTypes
}

func (s *ExportSuspEventsRequest) GetRemark() *string  {
  return s.Remark
}

func (s *ExportSuspEventsRequest) GetSourceIp() *string  {
  return s.SourceIp
}

func (s *ExportSuspEventsRequest) GetStatus() *string  {
  return s.Status
}

func (s *ExportSuspEventsRequest) GetTargetType() *string  {
  return s.TargetType
}

func (s *ExportSuspEventsRequest) GetTimeEnd() *string  {
  return s.TimeEnd
}

func (s *ExportSuspEventsRequest) GetTimeStart() *string  {
  return s.TimeStart
}

func (s *ExportSuspEventsRequest) GetUniqueInfo() *string  {
  return s.UniqueInfo
}

func (s *ExportSuspEventsRequest) GetUuid() *string  {
  return s.Uuid
}

func (s *ExportSuspEventsRequest) SetAssetsTypeList(v []*string) *ExportSuspEventsRequest {
  s.AssetsTypeList = v
  return s
}

func (s *ExportSuspEventsRequest) SetClusterId(v string) *ExportSuspEventsRequest {
  s.ClusterId = &v
  return s
}

func (s *ExportSuspEventsRequest) SetContainerFieldName(v string) *ExportSuspEventsRequest {
  s.ContainerFieldName = &v
  return s
}

func (s *ExportSuspEventsRequest) SetContainerFieldValue(v string) *ExportSuspEventsRequest {
  s.ContainerFieldValue = &v
  return s
}

func (s *ExportSuspEventsRequest) SetCurrentPage(v string) *ExportSuspEventsRequest {
  s.CurrentPage = &v
  return s
}

func (s *ExportSuspEventsRequest) SetDealed(v string) *ExportSuspEventsRequest {
  s.Dealed = &v
  return s
}

func (s *ExportSuspEventsRequest) SetFrom(v string) *ExportSuspEventsRequest {
  s.From = &v
  return s
}

func (s *ExportSuspEventsRequest) SetGroupId(v int64) *ExportSuspEventsRequest {
  s.GroupId = &v
  return s
}

func (s *ExportSuspEventsRequest) SetId(v int64) *ExportSuspEventsRequest {
  s.Id = &v
  return s
}

func (s *ExportSuspEventsRequest) SetLang(v string) *ExportSuspEventsRequest {
  s.Lang = &v
  return s
}

func (s *ExportSuspEventsRequest) SetLevels(v string) *ExportSuspEventsRequest {
  s.Levels = &v
  return s
}

func (s *ExportSuspEventsRequest) SetName(v string) *ExportSuspEventsRequest {
  s.Name = &v
  return s
}

func (s *ExportSuspEventsRequest) SetOperateErrorCodeList(v []*string) *ExportSuspEventsRequest {
  s.OperateErrorCodeList = v
  return s
}

func (s *ExportSuspEventsRequest) SetPageSize(v string) *ExportSuspEventsRequest {
  s.PageSize = &v
  return s
}

func (s *ExportSuspEventsRequest) SetParentEventTypes(v string) *ExportSuspEventsRequest {
  s.ParentEventTypes = &v
  return s
}

func (s *ExportSuspEventsRequest) SetRemark(v string) *ExportSuspEventsRequest {
  s.Remark = &v
  return s
}

func (s *ExportSuspEventsRequest) SetSourceIp(v string) *ExportSuspEventsRequest {
  s.SourceIp = &v
  return s
}

func (s *ExportSuspEventsRequest) SetStatus(v string) *ExportSuspEventsRequest {
  s.Status = &v
  return s
}

func (s *ExportSuspEventsRequest) SetTargetType(v string) *ExportSuspEventsRequest {
  s.TargetType = &v
  return s
}

func (s *ExportSuspEventsRequest) SetTimeEnd(v string) *ExportSuspEventsRequest {
  s.TimeEnd = &v
  return s
}

func (s *ExportSuspEventsRequest) SetTimeStart(v string) *ExportSuspEventsRequest {
  s.TimeStart = &v
  return s
}

func (s *ExportSuspEventsRequest) SetUniqueInfo(v string) *ExportSuspEventsRequest {
  s.UniqueInfo = &v
  return s
}

func (s *ExportSuspEventsRequest) SetUuid(v string) *ExportSuspEventsRequest {
  s.Uuid = &v
  return s
}

func (s *ExportSuspEventsRequest) Validate() error {
  return dara.Validate(s)
}

