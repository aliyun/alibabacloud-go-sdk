// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportContactFlowResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExportContactFlowResponseBody
  GetCode() *string 
  SetFlowPackageData(v string) *ExportContactFlowResponseBody
  GetFlowPackageData() *string 
  SetHttpStatusCode(v int32) *ExportContactFlowResponseBody
  GetHttpStatusCode() *int32 
  SetMessage(v string) *ExportContactFlowResponseBody
  GetMessage() *string 
  SetParams(v []*string) *ExportContactFlowResponseBody
  GetParams() []*string 
  SetRequestId(v string) *ExportContactFlowResponseBody
  GetRequestId() *string 
}

type ExportContactFlowResponseBody struct {
  // example:
  // 
  // OK
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // example:
  // 
  // {}
  FlowPackageData *string `json:"FlowPackageData,omitempty" xml:"FlowPackageData,omitempty"`
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
  // example:
  // 
  // DE803553-8AA9-4B9D-9E4E-A82BC69EDCEE
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportContactFlowResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportContactFlowResponseBody) GoString() string {
  return s.String()
}

func (s *ExportContactFlowResponseBody) GetCode() *string  {
  return s.Code
}

func (s *ExportContactFlowResponseBody) GetFlowPackageData() *string  {
  return s.FlowPackageData
}

func (s *ExportContactFlowResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *ExportContactFlowResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExportContactFlowResponseBody) GetParams() []*string  {
  return s.Params
}

func (s *ExportContactFlowResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportContactFlowResponseBody) SetCode(v string) *ExportContactFlowResponseBody {
  s.Code = &v
  return s
}

func (s *ExportContactFlowResponseBody) SetFlowPackageData(v string) *ExportContactFlowResponseBody {
  s.FlowPackageData = &v
  return s
}

func (s *ExportContactFlowResponseBody) SetHttpStatusCode(v int32) *ExportContactFlowResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExportContactFlowResponseBody) SetMessage(v string) *ExportContactFlowResponseBody {
  s.Message = &v
  return s
}

func (s *ExportContactFlowResponseBody) SetParams(v []*string) *ExportContactFlowResponseBody {
  s.Params = v
  return s
}

func (s *ExportContactFlowResponseBody) SetRequestId(v string) *ExportContactFlowResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportContactFlowResponseBody) Validate() error {
  return dara.Validate(s)
}

