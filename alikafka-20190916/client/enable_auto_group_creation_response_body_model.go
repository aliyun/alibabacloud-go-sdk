// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableAutoGroupCreationResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v int32) *EnableAutoGroupCreationResponseBody
  GetCode() *int32 
  SetMessage(v string) *EnableAutoGroupCreationResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnableAutoGroupCreationResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnableAutoGroupCreationResponseBody
  GetSuccess() *bool 
}

type EnableAutoGroupCreationResponseBody struct {
  // The returned HTTP status code.
  // 
  // If the value **200*	- is returned, the request is successful.
  // 
  // example:
  // 
  // 200
  Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
  // The returned message.
  // 
  // example:
  // 
  // operation success
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // A421CCD7-5BC5-4B32-8DD8-64668A8FCB56
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the request is successful.
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableAutoGroupCreationResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableAutoGroupCreationResponseBody) GoString() string {
  return s.String()
}

func (s *EnableAutoGroupCreationResponseBody) GetCode() *int32  {
  return s.Code
}

func (s *EnableAutoGroupCreationResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnableAutoGroupCreationResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableAutoGroupCreationResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableAutoGroupCreationResponseBody) SetCode(v int32) *EnableAutoGroupCreationResponseBody {
  s.Code = &v
  return s
}

func (s *EnableAutoGroupCreationResponseBody) SetMessage(v string) *EnableAutoGroupCreationResponseBody {
  s.Message = &v
  return s
}

func (s *EnableAutoGroupCreationResponseBody) SetRequestId(v string) *EnableAutoGroupCreationResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableAutoGroupCreationResponseBody) SetSuccess(v bool) *EnableAutoGroupCreationResponseBody {
  s.Success = &v
  return s
}

func (s *EnableAutoGroupCreationResponseBody) Validate() error {
  return dara.Validate(s)
}

