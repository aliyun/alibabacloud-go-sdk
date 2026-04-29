// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnablePolarClawCronJobRequest interface {
  dara.Model
  String() string
  GoString() string
  SetApplicationId(v string) *EnablePolarClawCronJobRequest
  GetApplicationId() *string 
  SetJobId(v string) *EnablePolarClawCronJobRequest
  GetJobId() *string 
  SetRestart(v bool) *EnablePolarClawCronJobRequest
  GetRestart() *bool 
}

type EnablePolarClawCronJobRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // pa-**************
  ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 0ee00f56-f467-4d41-858c-ca4ede2c770e
  JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
  // example:
  // 
  // true
  Restart *bool `json:"Restart,omitempty" xml:"Restart,omitempty"`
}

func (s EnablePolarClawCronJobRequest) String() string {
  return dara.Prettify(s)
}

func (s EnablePolarClawCronJobRequest) GoString() string {
  return s.String()
}

func (s *EnablePolarClawCronJobRequest) GetApplicationId() *string  {
  return s.ApplicationId
}

func (s *EnablePolarClawCronJobRequest) GetJobId() *string  {
  return s.JobId
}

func (s *EnablePolarClawCronJobRequest) GetRestart() *bool  {
  return s.Restart
}

func (s *EnablePolarClawCronJobRequest) SetApplicationId(v string) *EnablePolarClawCronJobRequest {
  s.ApplicationId = &v
  return s
}

func (s *EnablePolarClawCronJobRequest) SetJobId(v string) *EnablePolarClawCronJobRequest {
  s.JobId = &v
  return s
}

func (s *EnablePolarClawCronJobRequest) SetRestart(v bool) *EnablePolarClawCronJobRequest {
  s.Restart = &v
  return s
}

func (s *EnablePolarClawCronJobRequest) Validate() error {
  return dara.Validate(s)
}

