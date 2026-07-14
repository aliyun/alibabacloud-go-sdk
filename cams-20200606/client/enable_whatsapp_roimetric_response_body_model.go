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
  // The access denied details.
  // 
  // example:
  // 
  // None
  AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
  // The request status code. Valid values:
  // 
  // - OK: The request was successful.
  // 
  // - For other error codes, see [API error codes](https://help.aliyun.com/document_detail/196974.html).
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
  // 608F9CCA-B5EB-3D72-8047-B25D6D75BDEC
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

