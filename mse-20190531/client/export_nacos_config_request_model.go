// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportNacosConfigRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAcceptLanguage(v string) *ExportNacosConfigRequest
  GetAcceptLanguage() *string 
  SetAppName(v string) *ExportNacosConfigRequest
  GetAppName() *string 
  SetDataId(v string) *ExportNacosConfigRequest
  GetDataId() *string 
  SetDataIds(v string) *ExportNacosConfigRequest
  GetDataIds() *string 
  SetGroup(v string) *ExportNacosConfigRequest
  GetGroup() *string 
  SetIds(v string) *ExportNacosConfigRequest
  GetIds() *string 
  SetInstanceId(v string) *ExportNacosConfigRequest
  GetInstanceId() *string 
  SetNamespaceId(v string) *ExportNacosConfigRequest
  GetNamespaceId() *string 
}

type ExportNacosConfigRequest struct {
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
  // Deprecated
  // 
  // The application tag.
  // 
  // example:
  // 
  // qjl-gateway-openapi
  AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
  // Deprecated
  // 
  // The ID of the data that you want to export.
  // 
  // > 
  // 
  // 	- Multiple export methods are supported.
  // 
  // 	- If you want to export a single configuration, you must leave the Ids parameter empty and specify the DataID and Group parameters.
  // 
  // example:
  // 
  // sms-mes-develop.prop****
  DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
  // The configuration group name and the ID of the configuration that you want to export. Separate multiple configurations with comma (,).
  // 
  // example:
  // 
  // testGroup+testDataId1,testGroup+testDataId2
  DataIds *string `json:"DataIds,omitempty" xml:"DataIds,omitempty"`
  // Deprecated
  // 
  // The name of the configuration group.
  // 
  // example:
  // 
  // TIMEDTASK_COMMON_GROUP
  Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
  // Deprecated
  // 
  // The ID of the primary key of a configuration item.
  // 
  // >  - Multiple export methods are supported. You must specify this parameter if you want to export multiple configurations. - You can obtain the value of this parameter by calling the ListNacosConfigs operation. - If you specify this parameter, multiple configurations are exported. The DataId and Group parameters are invalid.
  // 
  // example:
  // 
  // 1709,1710
  Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
  // The ID of the instance.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // mse-cn-2r42ddc****
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  // The ID of the namespace.
  // 
  // example:
  // 
  // ae77c258-4d4f-478f-baaa-084aee0****
  NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s ExportNacosConfigRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportNacosConfigRequest) GoString() string {
  return s.String()
}

func (s *ExportNacosConfigRequest) GetAcceptLanguage() *string  {
  return s.AcceptLanguage
}

func (s *ExportNacosConfigRequest) GetAppName() *string  {
  return s.AppName
}

func (s *ExportNacosConfigRequest) GetDataId() *string  {
  return s.DataId
}

func (s *ExportNacosConfigRequest) GetDataIds() *string  {
  return s.DataIds
}

func (s *ExportNacosConfigRequest) GetGroup() *string  {
  return s.Group
}

func (s *ExportNacosConfigRequest) GetIds() *string  {
  return s.Ids
}

func (s *ExportNacosConfigRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *ExportNacosConfigRequest) GetNamespaceId() *string  {
  return s.NamespaceId
}

func (s *ExportNacosConfigRequest) SetAcceptLanguage(v string) *ExportNacosConfigRequest {
  s.AcceptLanguage = &v
  return s
}

func (s *ExportNacosConfigRequest) SetAppName(v string) *ExportNacosConfigRequest {
  s.AppName = &v
  return s
}

func (s *ExportNacosConfigRequest) SetDataId(v string) *ExportNacosConfigRequest {
  s.DataId = &v
  return s
}

func (s *ExportNacosConfigRequest) SetDataIds(v string) *ExportNacosConfigRequest {
  s.DataIds = &v
  return s
}

func (s *ExportNacosConfigRequest) SetGroup(v string) *ExportNacosConfigRequest {
  s.Group = &v
  return s
}

func (s *ExportNacosConfigRequest) SetIds(v string) *ExportNacosConfigRequest {
  s.Ids = &v
  return s
}

func (s *ExportNacosConfigRequest) SetInstanceId(v string) *ExportNacosConfigRequest {
  s.InstanceId = &v
  return s
}

func (s *ExportNacosConfigRequest) SetNamespaceId(v string) *ExportNacosConfigRequest {
  s.NamespaceId = &v
  return s
}

func (s *ExportNacosConfigRequest) Validate() error {
  return dara.Validate(s)
}

