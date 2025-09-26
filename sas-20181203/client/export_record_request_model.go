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
}

type ExportRecordRequest struct {
  // The type of the file to export. Valid values:
  // 
  // 	- **virusScanExport**: The details of the virus scan tasks are exported. This parameter is available and required when ExportType is set to virusScanExport.
  // 
  // example:
  // 
  // virusScanExport
  ExportFileType *string `json:"ExportFileType,omitempty" xml:"ExportFileType,omitempty"`
  // The type of detection result list to be exported. Values:
  // 
  // - **assetInstance**: List of servers in the asset center 
  // 
  // - **user**: List of asset fingerprints for accounts 
  // 
  // - **port**: List of asset fingerprints for ports 
  // 
  // - **process**: List of asset fingerprints for processes 
  // 
  // - **sca**: List of asset fingerprints for middleware 
  // 
  // - **database**: List of asset fingerprints for databases 
  // 
  // - **web**: List of asset fingerprints for web services 
  // 
  // - **software**: List of asset fingerprints for software 
  // 
  // - **cron**: List of asset fingerprints for scheduled tasks 
  // 
  // - **autorun**: List of asset fingerprints for startup items 
  // 
  // - **lkm**: List of asset fingerprints for kernel modules 
  // 
  // - **webserver**: List of asset fingerprints for web sites 
  // 
  // - **virusScanExport**: List of details for virus scan tasks 
  // 
  // - **imageVulExport**: List of system vulnerabilities in images 
  // 
  // - **imageBaseLineExport**: List of baseline check results in images 
  // 
  // - **imageAffectedMaliciousExport**: List of malicious sample check results in images 
  // 
  // - **baselineCspm**: List of detection results for cloud platform configuration checks 
  // 
  // - **attack**: List of alert events for attack analysis 
  // 
  // - **accessKey**: List of alert events for AK leak detection 
  // 
  // - **exportObjectScanEvents**: List of alert events for malicious file detection 
  // 
  // - **domainDetail**: Website assets 
  // 
  // - **assetsPropertyScaProcessDetail**: RASP protection process for application protection 
  // 
  // - **exportHcWarning**: List of system baseline risks 
  // 
  // - **raspAttackAlert**: List of attack alerts for Application Protection
  // 
  // - **raspApplicationConfiguration**: List of application configurations for Application Protection
  // 
  // - **raspWeaknessDetection**: List of weakness detections for Application Protection
  // 
  // - **raspInMemoryWebshellDetection**: List of in-memory webshell detection alerts for Application Protection
  // 
  // - **raspInMemoryWebshellInsertion**: List of in-memory webshell insertion alerts for Application Protection
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // database
  ExportType *string `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
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
  // The conditions that are used to filter check results.
  // 
  // > This operation is a common export operation for multiple features of Security Center. The available configuration fields of this parameter vary based on the features. We recommend that you do not specify this parameter when you call the operation. You can export an information list without specifying this parameter, and then filter data in the exported Excel file.
  // 
  // example:
  // 
  // {"extend":"1","currentPage":1,"pageSize":10}
  Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
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

func (s *ExportRecordRequest) Validate() error {
  return dara.Validate(s)
}

