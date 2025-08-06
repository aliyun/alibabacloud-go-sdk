// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditLicenseResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EditLicenseResponseBody
  GetCode() *string 
  SetData(v *EditLicenseResponseBodyData) *EditLicenseResponseBody
  GetData() *EditLicenseResponseBodyData 
  SetHttpStatusCode(v int32) *EditLicenseResponseBody
  GetHttpStatusCode() *int32 
  SetMessage(v string) *EditLicenseResponseBody
  GetMessage() *string 
  SetPageNo(v int64) *EditLicenseResponseBody
  GetPageNo() *int64 
  SetPageSize(v int64) *EditLicenseResponseBody
  GetPageSize() *int64 
  SetRequestId(v string) *EditLicenseResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EditLicenseResponseBody
  GetSuccess() *bool 
  SetTotal(v int64) *EditLicenseResponseBody
  GetTotal() *int64 
}

type EditLicenseResponseBody struct {
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Data *EditLicenseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
  PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
  Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s EditLicenseResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EditLicenseResponseBody) GoString() string {
  return s.String()
}

func (s *EditLicenseResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EditLicenseResponseBody) GetData() *EditLicenseResponseBodyData  {
  return s.Data
}

func (s *EditLicenseResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *EditLicenseResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EditLicenseResponseBody) GetPageNo() *int64  {
  return s.PageNo
}

func (s *EditLicenseResponseBody) GetPageSize() *int64  {
  return s.PageSize
}

func (s *EditLicenseResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EditLicenseResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EditLicenseResponseBody) GetTotal() *int64  {
  return s.Total
}

func (s *EditLicenseResponseBody) SetCode(v string) *EditLicenseResponseBody {
  s.Code = &v
  return s
}

func (s *EditLicenseResponseBody) SetData(v *EditLicenseResponseBodyData) *EditLicenseResponseBody {
  s.Data = v
  return s
}

func (s *EditLicenseResponseBody) SetHttpStatusCode(v int32) *EditLicenseResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *EditLicenseResponseBody) SetMessage(v string) *EditLicenseResponseBody {
  s.Message = &v
  return s
}

func (s *EditLicenseResponseBody) SetPageNo(v int64) *EditLicenseResponseBody {
  s.PageNo = &v
  return s
}

func (s *EditLicenseResponseBody) SetPageSize(v int64) *EditLicenseResponseBody {
  s.PageSize = &v
  return s
}

func (s *EditLicenseResponseBody) SetRequestId(v string) *EditLicenseResponseBody {
  s.RequestId = &v
  return s
}

func (s *EditLicenseResponseBody) SetSuccess(v bool) *EditLicenseResponseBody {
  s.Success = &v
  return s
}

func (s *EditLicenseResponseBody) SetTotal(v int64) *EditLicenseResponseBody {
  s.Total = &v
  return s
}

func (s *EditLicenseResponseBody) Validate() error {
  return dara.Validate(s)
}

type EditLicenseResponseBodyData struct {
  Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s EditLicenseResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s EditLicenseResponseBodyData) GoString() string {
  return s.String()
}

func (s *EditLicenseResponseBodyData) GetResult() *bool  {
  return s.Result
}

func (s *EditLicenseResponseBodyData) SetResult(v bool) *EditLicenseResponseBodyData {
  s.Result = &v
  return s
}

func (s *EditLicenseResponseBodyData) Validate() error {
  return dara.Validate(s)
}

