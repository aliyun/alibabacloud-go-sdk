// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportZookeeperDataRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAcceptLanguage(v string) *ExportZookeeperDataRequest
  GetAcceptLanguage() *string 
  SetExportType(v string) *ExportZookeeperDataRequest
  GetExportType() *string 
  SetInstanceId(v string) *ExportZookeeperDataRequest
  GetInstanceId() *string 
  SetRegionId(v string) *ExportZookeeperDataRequest
  GetRegionId() *string 
  SetRequestPars(v string) *ExportZookeeperDataRequest
  GetRequestPars() *string 
}

type ExportZookeeperDataRequest struct {
  // The language of the response. Valid values:
  // 
  // 	- zh: Chinese
  // 
  // 	- en: English
  // 
  // example:
  // 
  // zh
  AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
  // The type of the object that is exported. Valid values:
  // 
  // 	- transactionLog: transaction logs
  // 
  // 	- snapshot: snapshots
  // 
  // example:
  // 
  // snapshot
  ExportType *string `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
  // The instance ID.
  // 
  // example:
  // 
  // mse-cn-78v1l83****
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  // The ID of the region in which the instance resides. The region is supported by Microservices Engine (MSE).
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The extended request parameters in the JSON format.
  // 
  // example:
  // 
  // {}
  RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
}

func (s ExportZookeeperDataRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportZookeeperDataRequest) GoString() string {
  return s.String()
}

func (s *ExportZookeeperDataRequest) GetAcceptLanguage() *string  {
  return s.AcceptLanguage
}

func (s *ExportZookeeperDataRequest) GetExportType() *string  {
  return s.ExportType
}

func (s *ExportZookeeperDataRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *ExportZookeeperDataRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *ExportZookeeperDataRequest) GetRequestPars() *string  {
  return s.RequestPars
}

func (s *ExportZookeeperDataRequest) SetAcceptLanguage(v string) *ExportZookeeperDataRequest {
  s.AcceptLanguage = &v
  return s
}

func (s *ExportZookeeperDataRequest) SetExportType(v string) *ExportZookeeperDataRequest {
  s.ExportType = &v
  return s
}

func (s *ExportZookeeperDataRequest) SetInstanceId(v string) *ExportZookeeperDataRequest {
  s.InstanceId = &v
  return s
}

func (s *ExportZookeeperDataRequest) SetRegionId(v string) *ExportZookeeperDataRequest {
  s.RegionId = &v
  return s
}

func (s *ExportZookeeperDataRequest) SetRequestPars(v string) *ExportZookeeperDataRequest {
  s.RequestPars = &v
  return s
}

func (s *ExportZookeeperDataRequest) Validate() error {
  return dara.Validate(s)
}

