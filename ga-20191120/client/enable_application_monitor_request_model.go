// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableApplicationMonitorRequest interface {
  dara.Model
  String() string
  GoString() string
  SetClientToken(v string) *EnableApplicationMonitorRequest
  GetClientToken() *string 
  SetRegionId(v string) *EnableApplicationMonitorRequest
  GetRegionId() *string 
  SetTaskId(v string) *EnableApplicationMonitorRequest
  GetTaskId() *string 
}

type EnableApplicationMonitorRequest struct {
  // The client token that is used to ensure the idempotence of the request.
  // 
  // You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
  // 
  // >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
  // 
  // example:
  // 
  // 123e4567-e89b-12d3-a456-426655440000
  ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
  // The region ID of the Global Accelerator (GA) instance. Set the value to **cn-hangzhou**.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The ID of the origin probing task that you want to enable.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // sm-bp1fpdjfju9k8yr1y****
  TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s EnableApplicationMonitorRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableApplicationMonitorRequest) GoString() string {
  return s.String()
}

func (s *EnableApplicationMonitorRequest) GetClientToken() *string  {
  return s.ClientToken
}

func (s *EnableApplicationMonitorRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableApplicationMonitorRequest) GetTaskId() *string  {
  return s.TaskId
}

func (s *EnableApplicationMonitorRequest) SetClientToken(v string) *EnableApplicationMonitorRequest {
  s.ClientToken = &v
  return s
}

func (s *EnableApplicationMonitorRequest) SetRegionId(v string) *EnableApplicationMonitorRequest {
  s.RegionId = &v
  return s
}

func (s *EnableApplicationMonitorRequest) SetTaskId(v string) *EnableApplicationMonitorRequest {
  s.TaskId = &v
  return s
}

func (s *EnableApplicationMonitorRequest) Validate() error {
  return dara.Validate(s)
}

