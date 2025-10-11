// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableResourceCenterResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableResourceCenterResponseBody
  GetRequestId() *string 
  SetStatus(v string) *EnableResourceCenterResponseBody
  GetStatus() *string 
}

type EnableResourceCenterResponseBody struct {
  // The ID of the request.
  // 
  // example:
  // 
  // 45357BEF-AB50-5E4D-B05D-5A882A4BE924
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // The activation status of the service. Valid values:
  // 
  // 	- Pending: The service is being activated.
  // 
  // 	- Enabled: The service is activated.
  // 
  // example:
  // 
  // Pending
  Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s EnableResourceCenterResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableResourceCenterResponseBody) GoString() string {
  return s.String()
}

func (s *EnableResourceCenterResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableResourceCenterResponseBody) GetStatus() *string  {
  return s.Status
}

func (s *EnableResourceCenterResponseBody) SetRequestId(v string) *EnableResourceCenterResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableResourceCenterResponseBody) SetStatus(v string) *EnableResourceCenterResponseBody {
  s.Status = &v
  return s
}

func (s *EnableResourceCenterResponseBody) Validate() error {
  return dara.Validate(s)
}

