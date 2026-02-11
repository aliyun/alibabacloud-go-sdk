// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableAlertTemplateResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableAlertTemplateResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnableAlertTemplateResponseBody
  GetSuccess() *bool 
}

type EnableAlertTemplateResponseBody struct {
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableAlertTemplateResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableAlertTemplateResponseBody) GoString() string {
  return s.String()
}

func (s *EnableAlertTemplateResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableAlertTemplateResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableAlertTemplateResponseBody) SetRequestId(v string) *EnableAlertTemplateResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableAlertTemplateResponseBody) SetSuccess(v bool) *EnableAlertTemplateResponseBody {
  s.Success = &v
  return s
}

func (s *EnableAlertTemplateResponseBody) Validate() error {
  return dara.Validate(s)
}

