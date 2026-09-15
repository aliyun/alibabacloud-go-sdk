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
  SetResourceDirectoryAccountId(v int64) *ExportSuspEventsRequest
  GetResourceDirectoryAccountId() *int64 
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
  // The collection of asset types.
  AssetsTypeList []*string `json:"AssetsTypeList,omitempty" xml:"AssetsTypeList,omitempty" type:"Repeated"`
  // The ID of the cluster to query.
  // 
  // > You can call the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) operation to obtain this parameter.
  // 
  // example:
  // 
  // c4af4fdf38a98496a9b63c2be5dae****
  ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
  // The container search field. Valid values:
  // 
  // - **instanceId**: instance ID
  // 
  // - **appName**: application name
  // 
  // - **clusterId**: cluster ID
  // 
  // - **regionId**: region
  // 
  // - **nodeName**: node name
  // 
  // - **namespace**: namespace
  // 
  // - **clusterName**: cluster name
  // 
  // - **image**: image name
  // 
  // - **imageRepoName**: image repository name
  // 
  // - **imageRepoNamespace**: image repository namespace
  // 
  // - **imageRepoTag**: image tag
  // 
  // - **imageDigest**: image digest
  // 
  // example:
  // 
  // clusterId
  ContainerFieldName *string `json:"ContainerFieldName,omitempty" xml:"ContainerFieldName,omitempty"`
  // The value of the container search field.
  // 
  // example:
  // 
  // c819391d2d520485fa3e81e2dc2ea****
  ContainerFieldValue *string `json:"ContainerFieldValue,omitempty" xml:"ContainerFieldValue,omitempty"`
  // The page number of the current page in a paged query.
  // 
  // example:
  // 
  // 1
  CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
  // Specifies whether the alert event is handled. Valid values:
  // 
  // - **N**: Unhandled.
  // 
  // - **Y**: Handled.
  // 
  // example:
  // 
  // Y
  Dealed *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
  // The data source identifier of the anomaly event. Set the value to sas.
  // 
  // example:
  // 
  // sas
  From *string `json:"From,omitempty" xml:"From,omitempty"`
  // The ID of the asset group.
  // 
  // example:
  // 
  // 9454789
  GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
  // The unique ID of the alert event record.
  // 
  // example:
  // 
  // 17821
  Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
  // The language type for the request and response messages. Default value: **zh**. Valid values:
  // 
  // - **zh**: Chinese.
  // 
  // - **en**: English.
  // 
  // example:
  // 
  // zh
  Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
  // The severity levels of the alert events to query. Separate multiple severity levels with commas (,). The severity levels are listed in descending order. Valid values:
  // 
  // - **serious**: Urgent.
  // 
  // - **suspicious**: Suspicious.
  // 
  // - **remind**: Reminder.
  // 
  // example:
  // 
  // serious,suspicious,remind
  Levels *string `json:"Levels,omitempty" xml:"Levels,omitempty"`
  // The full name of the anomaly event.
  // 
  // example:
  // 
  // WEBSHELL
  Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
  // The collection of alert event handling result codes.
  OperateErrorCodeList []*string `json:"OperateErrorCodeList,omitempty" xml:"OperateErrorCodeList,omitempty" type:"Repeated"`
  // The maximum number of entries per page in a paged query. Default value: **20**.
  // 
  // example:
  // 
  // 20
  PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
  // The Alarm Metric of the alerting events to query. Valid values:
  // 
  // - **Abnormal process behavior**
  // 
  // - **Web shell**
  // 
  // - **Unusual logon**
  // 
  // - **Anomaly event**
  // 
  // - **Sensitive file tampering**
  // 
  // - **Malicious process (cloud scan)**
  // 
  // - **Suspicious network connection**
  // 
  // - **Abnormal account**
  // 
  // - **Application intrusion event**
  // 
  // - **Cloud service threat detection**
  // 
  // - **Precise defense**
  // 
  // - **Application whitelist**
  // 
  // - **Persistent backdoor**
  // 
  // - **Web application threat detection**
  // 
  // - **Malicious script**
  // 
  // - **Threat intelligence**
  // 
  // - **Malicious network connectivity behavior**
  // 
  // - **Container cluster exception**
  // 
  // - **Web shell (local scan)**
  // 
  // - **Vulnerability exploits**
  // 
  // - **Malicious process (local scan)**
  // 
  // - **Trusted exception**
  // 
  // - **Other**
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
  // The ID of the Alibaba Cloud account of the member accounts in the resource directory.
  // 
  // > You can call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain this parameter.
  ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
  // The source IP address of the request. You do not need to specify this parameter. The system automatically obtains this value.
  // 
  // example:
  // 
  // 127.0.XX.XX
  SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
  // The handling status of the anomaly event. Valid values:
  // 
  // - **0**: All.
  // 
  // - **1**: Unhandled.
  // 
  // - **2**: Ignored.
  // 
  // - **4**: Confirmed.
  // 
  // - **8**: Marked as false positive.
  // 
  // - **16**: Handling.
  // 
  // - **32**: Handled.
  // 
  // - **64**: Expired.
  // 
  // - **128**: Deleted.
  // 
  // example:
  // 
  // 0
  Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
  // The dimension of the target switch configuration. Valid values:
  // 
  // - **uuid**: asset UUID
  // 
  // - **image_repo**: image repository ID
  // 
  // - **Cluster**: cluster ID
  // 
  // example:
  // 
  // uuid
  TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
  // The end time of the anomaly event. Format: YYYY-MM-DD HH:mm:ss.
  // 
  // example:
  // 
  // 2022-12-05 00:00:00
  TimeEnd *string `json:"TimeEnd,omitempty" xml:"TimeEnd,omitempty"`
  // The start time of the anomaly event occurrence. Format: YYYY-MM-DD HH:mm:ss.
  // 
  // example:
  // 
  // 2022-10-01 00:00:00
  TimeStart *string `json:"TimeStart,omitempty" xml:"TimeStart,omitempty"`
  // The unique key of the security alert.
  // 
  // example:
  // 
  // 1fbe8d16727f61d1478a674d6fa0****
  UniqueInfo *string `json:"UniqueInfo,omitempty" xml:"UniqueInfo,omitempty"`
  // The unique identifier of the associated instance.
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

func (s *ExportSuspEventsRequest) GetResourceDirectoryAccountId() *int64  {
  return s.ResourceDirectoryAccountId
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

func (s *ExportSuspEventsRequest) SetResourceDirectoryAccountId(v int64) *ExportSuspEventsRequest {
  s.ResourceDirectoryAccountId = &v
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

