// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableWhatsappROIMetricResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetAccessDeniedDetail(v string) *EnableWhatsappROIMetricResponseBody
  GetAccessDeniedDetail() *string 
  SetCode(v string) *EnableWhatsappROIMetricResponseBody
  GetCode() *string 
  SetMessage(v string) *EnableWhatsappROIMetricResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnableWhatsappROIMetricResponseBody
  GetRequestId() *string 
}

type EnableWhatsappROIMetricResponseBody struct {
  // The details about the access denial.
  // 
  // example:
  // 
  // NONE
  AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
  // The value OK indicates that the request was successful.
  // 
  // example:
  // 
  // OK
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // The error message.
  // 
  // example:
  // 
  // None
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // 90E63D28-E31D-1EB2-8939-A9486641****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableWhatsappROIMetricResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableWhatsappROIMetricResponseBody) GoString() string {
  return s.String()
}

func (s *EnableWhatsappROIMetricResponseBody) GetAccessDeniedDetail() *string  {
  return s.AccessDeniedDetail
}

func (s *EnableWhatsappROIMetricResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnableWhatsappROIMetricResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnableWhatsappROIMetricResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableWhatsappROIMetricResponseBody) SetAccessDeniedDetail(v string) *EnableWhatsappROIMetricResponseBody {
  s.AccessDeniedDetail = &v
  return s
}

func (s *EnableWhatsappROIMetricResponseBody) SetCode(v string) *EnableWhatsappROIMetricResponseBody {
  s.Code = &v
  return s
}

func (s *EnableWhatsappROIMetricResponseBody) SetMessage(v string) *EnableWhatsappROIMetricResponseBody {
  s.Message = &v
  return s
}

func (s *EnableWhatsappROIMetricResponseBody) SetRequestId(v string) *EnableWhatsappROIMetricResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableWhatsappROIMetricResponseBody) Validate() error {
  return dara.Validate(s)
}

