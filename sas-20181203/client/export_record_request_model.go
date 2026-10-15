// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportRecordRequest interface {
  dara.Model
  String() string
  GoString() string
  SetExportFileType(v string) *ExportRecordRequest
  GetExportFileType() *string 
  SetExportType(v string) *ExportRecordRequest
  GetExportType() *string 
  SetLang(v string) *ExportRecordRequest
  GetLang() *string 
  SetParams(v string) *ExportRecordRequest
  GetParams() *string 
  SetResourceDirectoryAccountId(v int64) *ExportRecordRequest
  GetResourceDirectoryAccountId() *int64 
}

type ExportRecordRequest struct {
  // Specifies the type of the file to be exported. Valid values:
  // 
  // - **virusScanExport**: Exports the virus scan task details. This value is required and can only be used when ExportType is set to virusScanExport.
  // 
  // example:
  // 
  // virusScanExport
  ExportFileType *string `json:"ExportFileType,omitempty" xml:"ExportFileType,omitempty"`
  // Specifies the type of the detection result list to be exported. Valid values:
  // 
  // - **assetInstance**: the server list in the Asset Center
  // 
  // - **user**: the Asset Fingerprints list of accounts
  // 
  // - **port**: the Asset Fingerprints list of ports
  // 
  // - **process**: the Asset Fingerprints list of processes
  // 
  // - **sca**: the Asset Fingerprints list of middleware
  // 
  // - **database**: the Asset Fingerprints list of databases
  // 
  // - **web**: the Asset Fingerprints list of web services
  // 
  // - **software**: the Asset Fingerprints list of software
  // 
  // - **cron**: the Asset Fingerprints list of scheduled nodes
  // 
  // - **autorun**: the Asset Fingerprints list of startup items
  // 
  // - **lkm**: the Asset Fingerprints list of kernel modules
  // 
  // - **webserver**: the Asset Fingerprints list of websites
  // 
  // - **virusScanExport**: the virus scan node details list
  // 
  // - **imageVulExport**: the image system vulnerability list
  // 
  // - **imageBaseLineExport**: the image baseline check result list
  // 
  // - **imageAffectedMaliciousExport**: the image malicious sample check result list
  // 
  // - **baselineCspm**: the cloud platform configuration check result list
  // 
  // - **attack**: the attack analysis alerting event list
  // 
  // - **accessKey**: the AccessKey leak detection alerting event list
  // 
  // - **exportObjectScanEvents**: the malicious file detection alerting event list
  // 
  // - **domainDetail**: the website assets
  // 
  // - **assetsPropertyScaProcessDetail**: the RASP-protected processes in Application Protection
  // 
  // - **exportHcWarning**: the system baseline risk list
  // 
  // - **raspAttackAlert**: the Application Protection attack alerting list
  // 
  // - **raspApplicationConfiguration**: the Application Protection application configuration list
  // 
  // - **raspWeaknessDetection**: the Application Protection weakness detection list
  // 
  // - **raspInMemoryWebshellDetection**: the Application Protection in-memory webshell detection alerting list
  // 
  // - **raspInMemoryWebshellInsertion**: the Application Protection in-memory webshell injection alerting list
  // 
  // - **listAgentExport**: the agent list
  // 
  // - **listSkillExport**: the skill list
  // 
  // - **listModelExport**: the model list
  // 
  // - **listAiAgentEventExport**: the agent risk list
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // database
  ExportType *string `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
  // Specifies the language of the request and response. Default value: **zh**. Valid values:
  // 
  // - **zh**: Chinese
  // 
  // - **en**: English
  // 
  // example:
  // 
  // zh
  Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
  // Specifies the filter conditions for the exported content.
  // 
  // > This operation is a general export operation for exporting detection lists of multiple features in Security Center. The configuration fields of this parameter vary based on the feature. Do not specify this parameter when you call the operation. You can export the information list without specifying this parameter and then filter the target data in the exported Excel file.
  // 
  // example:
  // 
  // {"extend":"1","currentPage":1,"pageSize":10}
  Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
  // The Alibaba Cloud account ID of the member accounts in the resource folder.
  // 
  // > You can invoke the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain this parameter.
  // 
  // example:
  // 
  // 1232428423234****
  ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
}

func (s ExportRecordRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportRecordRequest) GoString() string {
  return s.String()
}

func (s *ExportRecordRequest) GetExportFileType() *string  {
  return s.ExportFileType
}

func (s *ExportRecordRequest) GetExportType() *string  {
  return s.ExportType
}

func (s *ExportRecordRequest) GetLang() *string  {
  return s.Lang
}

func (s *ExportRecordRequest) GetParams() *string  {
  return s.Params
}

func (s *ExportRecordRequest) GetResourceDirectoryAccountId() *int64  {
  return s.ResourceDirectoryAccountId
}

func (s *ExportRecordRequest) SetExportFileType(v string) *ExportRecordRequest {
  s.ExportFileType = &v
  return s
}

func (s *ExportRecordRequest) SetExportType(v string) *ExportRecordRequest {
  s.ExportType = &v
  return s
}

func (s *ExportRecordRequest) SetLang(v string) *ExportRecordRequest {
  s.Lang = &v
  return s
}

func (s *ExportRecordRequest) SetParams(v string) *ExportRecordRequest {
  s.Params = &v
  return s
}

func (s *ExportRecordRequest) SetResourceDirectoryAccountId(v int64) *ExportRecordRequest {
  s.ResourceDirectoryAccountId = &v
  return s
}

func (s *ExportRecordRequest) Validate() error {
  return dara.Validate(s)
}

