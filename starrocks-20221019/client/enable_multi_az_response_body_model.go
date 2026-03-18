// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableMultiAzResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetAccessDeniedDetail(v string) *EnableMultiAzResponseBody
  GetAccessDeniedDetail() *string 
  SetData(v *EnableMultiAzResponseBodyData) *EnableMultiAzResponseBody
  GetData() *EnableMultiAzResponseBodyData 
  SetErrCode(v string) *EnableMultiAzResponseBody
  GetErrCode() *string 
  SetErrMessage(v string) *EnableMultiAzResponseBody
  GetErrMessage() *string 
  SetHttpStatusCode(v int32) *EnableMultiAzResponseBody
  GetHttpStatusCode() *int32 
  SetRequestId(v string) *EnableMultiAzResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnableMultiAzResponseBody
  GetSuccess() *bool 
}

type EnableMultiAzResponseBody struct {
  // AccessDeniedDetail
  // 
  // example:
  // 
  // {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
  AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
  Data *EnableMultiAzResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  // example:
  // 
  // InvalidParams
  ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
  // example:
  // 
  // Invalid params: [instance not exists].
  ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  // example:
  // 
  // 32A44F0D-BFF6-5664-999A-218BBDE74XXX
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // example:
  // 
  // True
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableMultiAzResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableMultiAzResponseBody) GoString() string {
  return s.String()
}

func (s *EnableMultiAzResponseBody) GetAccessDeniedDetail() *string  {
  return s.AccessDeniedDetail
}

func (s *EnableMultiAzResponseBody) GetData() *EnableMultiAzResponseBodyData  {
  return s.Data
}

func (s *EnableMultiAzResponseBody) GetErrCode() *string  {
  return s.ErrCode
}

func (s *EnableMultiAzResponseBody) GetErrMessage() *string  {
  return s.ErrMessage
}

func (s *EnableMultiAzResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *EnableMultiAzResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableMultiAzResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableMultiAzResponseBody) SetAccessDeniedDetail(v string) *EnableMultiAzResponseBody {
  s.AccessDeniedDetail = &v
  return s
}

func (s *EnableMultiAzResponseBody) SetData(v *EnableMultiAzResponseBodyData) *EnableMultiAzResponseBody {
  s.Data = v
  return s
}

func (s *EnableMultiAzResponseBody) SetErrCode(v string) *EnableMultiAzResponseBody {
  s.ErrCode = &v
  return s
}

func (s *EnableMultiAzResponseBody) SetErrMessage(v string) *EnableMultiAzResponseBody {
  s.ErrMessage = &v
  return s
}

func (s *EnableMultiAzResponseBody) SetHttpStatusCode(v int32) *EnableMultiAzResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *EnableMultiAzResponseBody) SetRequestId(v string) *EnableMultiAzResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableMultiAzResponseBody) SetSuccess(v bool) *EnableMultiAzResponseBody {
  s.Success = &v
  return s
}

func (s *EnableMultiAzResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EnableMultiAzResponseBodyData struct {
  // example:
  // 
  // ng-a9b2e9148196****
  NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
  // example:
  // 
  // 24782976697****
  OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s EnableMultiAzResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s EnableMultiAzResponseBodyData) GoString() string {
  return s.String()
}

func (s *EnableMultiAzResponseBodyData) GetNodeGroupId() *string  {
  return s.NodeGroupId
}

func (s *EnableMultiAzResponseBodyData) GetOrderId() *int64  {
  return s.OrderId
}

func (s *EnableMultiAzResponseBodyData) SetNodeGroupId(v string) *EnableMultiAzResponseBodyData {
  s.NodeGroupId = &v
  return s
}

func (s *EnableMultiAzResponseBodyData) SetOrderId(v int64) *EnableMultiAzResponseBodyData {
  s.OrderId = &v
  return s
}

func (s *EnableMultiAzResponseBodyData) Validate() error {
  return dara.Validate(s)
}

