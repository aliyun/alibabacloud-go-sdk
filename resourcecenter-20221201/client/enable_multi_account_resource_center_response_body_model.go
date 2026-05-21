// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableMultiAccountResourceCenterResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableMultiAccountResourceCenterResponseBody
  GetRequestId() *string 
  SetStatus(v string) *EnableMultiAccountResourceCenterResponseBody
  GetStatus() *string 
}

type EnableMultiAccountResourceCenterResponseBody struct {
  // The ID of the request.
  // 
  // example:
  // 
  // 767038B7-2027-5508-858B-E213232D57D5
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // The status of the feature. Valid values:
  // 
  // 	- Pending: The feature is being enabled.
  // 
  // 	- Enabled: The feature is enabled.
  // 
  // example:
  // 
  // Pending
  Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s EnableMultiAccountResourceCenterResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableMultiAccountResourceCenterResponseBody) GoString() string {
  return s.String()
}

func (s *EnableMultiAccountResourceCenterResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableMultiAccountResourceCenterResponseBody) GetStatus() *string  {
  return s.Status
}

func (s *EnableMultiAccountResourceCenterResponseBody) SetRequestId(v string) *EnableMultiAccountResourceCenterResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableMultiAccountResourceCenterResponseBody) SetStatus(v string) *EnableMultiAccountResourceCenterResponseBody {
  s.Status = &v
  return s
}

func (s *EnableMultiAccountResourceCenterResponseBody) Validate() error {
  return dara.Validate(s)
}

