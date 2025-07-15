// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEndConferenceRequest interface {
  dara.Model
  String() string
  GoString() string
  SetInstanceId(v string) *EndConferenceRequest
  GetInstanceId() *string 
  SetJobId(v string) *EndConferenceRequest
  GetJobId() *string 
  SetUserId(v string) *EndConferenceRequest
  GetUserId() *string 
}

type EndConferenceRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // 9cfad875-6260-4a53-ab6e-b13e3fb31f7d
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // job-6538214103685****
  JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
  // example:
  // 
  // agent@ccc-test
  UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s EndConferenceRequest) String() string {
  return dara.Prettify(s)
}

func (s EndConferenceRequest) GoString() string {
  return s.String()
}

func (s *EndConferenceRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EndConferenceRequest) GetJobId() *string  {
  return s.JobId
}

func (s *EndConferenceRequest) GetUserId() *string  {
  return s.UserId
}

func (s *EndConferenceRequest) SetInstanceId(v string) *EndConferenceRequest {
  s.InstanceId = &v
  return s
}

func (s *EndConferenceRequest) SetJobId(v string) *EndConferenceRequest {
  s.JobId = &v
  return s
}

func (s *EndConferenceRequest) SetUserId(v string) *EndConferenceRequest {
  s.UserId = &v
  return s
}

func (s *EndConferenceRequest) Validate() error {
  return dara.Validate(s)
}

