// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportApplicationConfigsResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetApplicationConfigs(v []*ExportApplicationConfigsResponseBodyApplicationConfigs) *ExportApplicationConfigsResponseBody
  GetApplicationConfigs() []*ExportApplicationConfigsResponseBodyApplicationConfigs 
  SetRequestId(v string) *ExportApplicationConfigsResponseBody
  GetRequestId() *string 
}

type ExportApplicationConfigsResponseBody struct {
  ApplicationConfigs []*ExportApplicationConfigsResponseBodyApplicationConfigs `json:"ApplicationConfigs,omitempty" xml:"ApplicationConfigs,omitempty" type:"Repeated"`
  // 请求ID。
  // 
  // example:
  // 
  // DD6B1B2A-5837-5237-ABE4-FF0C8944****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportApplicationConfigsResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportApplicationConfigsResponseBody) GoString() string {
  return s.String()
}

func (s *ExportApplicationConfigsResponseBody) GetApplicationConfigs() []*ExportApplicationConfigsResponseBodyApplicationConfigs  {
  return s.ApplicationConfigs
}

func (s *ExportApplicationConfigsResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportApplicationConfigsResponseBody) SetApplicationConfigs(v []*ExportApplicationConfigsResponseBodyApplicationConfigs) *ExportApplicationConfigsResponseBody {
  s.ApplicationConfigs = v
  return s
}

func (s *ExportApplicationConfigsResponseBody) SetRequestId(v string) *ExportApplicationConfigsResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportApplicationConfigsResponseBody) Validate() error {
  if s.ApplicationConfigs != nil {
    for _, item := range s.ApplicationConfigs {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type ExportApplicationConfigsResponseBodyApplicationConfigs struct {
  // 应用名称。
  // 
  // example:
  // 
  // YARN
  ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
  // 文件名称。
  // 
  // example:
  // 
  // yarn-site.xml
  ConfigFileName *string `json:"ConfigFileName,omitempty" xml:"ConfigFileName,omitempty"`
  // 文件内容，base64加密。
  // 
  // example:
  // 
  // export key=value
  Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s ExportApplicationConfigsResponseBodyApplicationConfigs) String() string {
  return dara.Prettify(s)
}

func (s ExportApplicationConfigsResponseBodyApplicationConfigs) GoString() string {
  return s.String()
}

func (s *ExportApplicationConfigsResponseBodyApplicationConfigs) GetApplicationName() *string  {
  return s.ApplicationName
}

func (s *ExportApplicationConfigsResponseBodyApplicationConfigs) GetConfigFileName() *string  {
  return s.ConfigFileName
}

func (s *ExportApplicationConfigsResponseBodyApplicationConfigs) GetContent() *string  {
  return s.Content
}

func (s *ExportApplicationConfigsResponseBodyApplicationConfigs) SetApplicationName(v string) *ExportApplicationConfigsResponseBodyApplicationConfigs {
  s.ApplicationName = &v
  return s
}

func (s *ExportApplicationConfigsResponseBodyApplicationConfigs) SetConfigFileName(v string) *ExportApplicationConfigsResponseBodyApplicationConfigs {
  s.ConfigFileName = &v
  return s
}

func (s *ExportApplicationConfigsResponseBodyApplicationConfigs) SetContent(v string) *ExportApplicationConfigsResponseBodyApplicationConfigs {
  s.Content = &v
  return s
}

func (s *ExportApplicationConfigsResponseBodyApplicationConfigs) Validate() error {
  return dara.Validate(s)
}

