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
  // The type of the check result list that you want to export. Valid values:
  // 
  // 	- **assetInstance**: the list of servers displayed on the Host page
  // 
  // 	- **user**: the list of fingerprints for accounts
  // 
  // 	- **port**: the list of fingerprints for ports
  // 
  // 	- **process**: the list of fingerprints for processes
  // 
  // 	- **sca**: the list of fingerprints for middleware assets
  // 
  // 	- **database**: the list of fingerprints for databases
  // 
  // 	- **web**: the list of fingerprints for web services
  // 
  // 	- **software**: the list of fingerprints for software assets
  // 
  // 	- **cron**: the list of fingerprints for scheduled tasks
  // 
  // 	- **autorun**: the list of fingerprints for startup items
  // 
  // 	- **lkm**: the list of fingerprints for kernel modules
  // 
  // 	- **webserver**: the list of fingerprints for websites
  // 
  // 	- **virusScanExport**: the details of the virus scan tasks
  // 
  // 	- **imageVulExport**: the list of image system vulnerabilities
  // 
  // 	- **imageBaseLineExport**: the list of check results for image baselines
  // 
  // 	- **imageAffectedMaliciousExport**: the list of check results for malicious image samples
  // 
  // 	- **baselineCspm**: the list of check results for configuration assessment
  // 
  // 	- **attack**: the list of alert events for attack analysis
  // 
  // 	- **accessKey**: the list of alerts that are generated for AccessKey pair leaks
  // 
  // 	- **exportObjectScanEvents**: the list of alerts generated for malicious files
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

