// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportScriptResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExportScriptResponseBody
  GetCode() *string 
  SetDownloadParams(v *ExportScriptResponseBodyDownloadParams) *ExportScriptResponseBody
  GetDownloadParams() *ExportScriptResponseBodyDownloadParams 
  SetHttpStatusCode(v int32) *ExportScriptResponseBody
  GetHttpStatusCode() *int32 
  SetMessage(v string) *ExportScriptResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExportScriptResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExportScriptResponseBody
  GetSuccess() *bool 
}

type ExportScriptResponseBody struct {
  // example:
  // 
  // OK
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  DownloadParams *ExportScriptResponseBodyDownloadParams `json:"DownloadParams,omitempty" xml:"DownloadParams,omitempty" type:"Struct"`
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  // example:
  // 
  // Success
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // example:
  // 
  // 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExportScriptResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportScriptResponseBody) GoString() string {
  return s.String()
}

func (s *ExportScriptResponseBody) GetCode() *string  {
  return s.Code
}

func (s *ExportScriptResponseBody) GetDownloadParams() *ExportScriptResponseBodyDownloadParams  {
  return s.DownloadParams
}

func (s *ExportScriptResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *ExportScriptResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExportScriptResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportScriptResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExportScriptResponseBody) SetCode(v string) *ExportScriptResponseBody {
  s.Code = &v
  return s
}

func (s *ExportScriptResponseBody) SetDownloadParams(v *ExportScriptResponseBodyDownloadParams) *ExportScriptResponseBody {
  s.DownloadParams = v
  return s
}

func (s *ExportScriptResponseBody) SetHttpStatusCode(v int32) *ExportScriptResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExportScriptResponseBody) SetMessage(v string) *ExportScriptResponseBody {
  s.Message = &v
  return s
}

func (s *ExportScriptResponseBody) SetRequestId(v string) *ExportScriptResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportScriptResponseBody) SetSuccess(v bool) *ExportScriptResponseBody {
  s.Success = &v
  return s
}

func (s *ExportScriptResponseBody) Validate() error {
  if s.DownloadParams != nil {
    if err := s.DownloadParams.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExportScriptResponseBodyDownloadParams struct {
  // example:
  // 
  // http://tiangong-staging.oss-cn-shanghai.aliyuncs.com/record/281eb174-3865-41c1-9274-7b6813edadab.wav?Expires=1578624046&OSSAccessKeyId=LTAI****cqw&Signature=dL2dxWS6VcdZrvG9xOMOBMSP3Fg%3D
  SignatureUrl *string `json:"SignatureUrl,omitempty" xml:"SignatureUrl,omitempty"`
}

func (s ExportScriptResponseBodyDownloadParams) String() string {
  return dara.Prettify(s)
}

func (s ExportScriptResponseBodyDownloadParams) GoString() string {
  return s.String()
}

func (s *ExportScriptResponseBodyDownloadParams) GetSignatureUrl() *string  {
  return s.SignatureUrl
}

func (s *ExportScriptResponseBodyDownloadParams) SetSignatureUrl(v string) *ExportScriptResponseBodyDownloadParams {
  s.SignatureUrl = &v
  return s
}

func (s *ExportScriptResponseBodyDownloadParams) Validate() error {
  return dara.Validate(s)
}

