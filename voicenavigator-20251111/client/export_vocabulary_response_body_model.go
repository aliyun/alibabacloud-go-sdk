// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportVocabularyResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExportVocabularyResponseBody
  GetCode() *string 
  SetData(v string) *ExportVocabularyResponseBody
  GetData() *string 
  SetHttpStatusCode(v int32) *ExportVocabularyResponseBody
  GetHttpStatusCode() *int32 
  SetMessage(v string) *ExportVocabularyResponseBody
  GetMessage() *string 
  SetParams(v []*string) *ExportVocabularyResponseBody
  GetParams() []*string 
  SetRequestId(v string) *ExportVocabularyResponseBody
  GetRequestId() *string 
}

type ExportVocabularyResponseBody struct {
  // example:
  // 
  // OK
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  // example:
  // 
  // Instance af81a389-91f0-4157-8d82-720edd02b66a
  // 
  //  does not exist.
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
  // example:
  // 
  // 7401D698-0AAC-5909-B68E-88C68805FFB8
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportVocabularyResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportVocabularyResponseBody) GoString() string {
  return s.String()
}

func (s *ExportVocabularyResponseBody) GetCode() *string  {
  return s.Code
}

func (s *ExportVocabularyResponseBody) GetData() *string  {
  return s.Data
}

func (s *ExportVocabularyResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *ExportVocabularyResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExportVocabularyResponseBody) GetParams() []*string  {
  return s.Params
}

func (s *ExportVocabularyResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportVocabularyResponseBody) SetCode(v string) *ExportVocabularyResponseBody {
  s.Code = &v
  return s
}

func (s *ExportVocabularyResponseBody) SetData(v string) *ExportVocabularyResponseBody {
  s.Data = &v
  return s
}

func (s *ExportVocabularyResponseBody) SetHttpStatusCode(v int32) *ExportVocabularyResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExportVocabularyResponseBody) SetMessage(v string) *ExportVocabularyResponseBody {
  s.Message = &v
  return s
}

func (s *ExportVocabularyResponseBody) SetParams(v []*string) *ExportVocabularyResponseBody {
  s.Params = v
  return s
}

func (s *ExportVocabularyResponseBody) SetRequestId(v string) *ExportVocabularyResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportVocabularyResponseBody) Validate() error {
  return dara.Validate(s)
}

