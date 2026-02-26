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
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

