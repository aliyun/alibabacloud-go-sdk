// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnablePolarClawCronJobResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetApplicationId(v string) *EnablePolarClawCronJobResponseBody
  GetApplicationId() *string 
  SetCode(v int32) *EnablePolarClawCronJobResponseBody
  GetCode() *int32 
  SetJobId(v string) *EnablePolarClawCronJobResponseBody
  GetJobId() *string 
  SetMessage(v string) *EnablePolarClawCronJobResponseBody
  GetMessage() *string 
  SetOk(v bool) *EnablePolarClawCronJobResponseBody
  GetOk() *bool 
  SetRequestId(v string) *EnablePolarClawCronJobResponseBody
  GetRequestId() *string 
}

type EnablePolarClawCronJobResponseBody struct {
  // example:
  // 
  // pa-**************
  ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
  // example:
  // 
  // 200
  Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
  // example:
  // 
  // 8006e51c-dab3-4602-bc69-4f728002c6ce
  JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
  // example:
  // 
  // successful
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // example:
  // 
  // true
  Ok *bool `json:"Ok,omitempty" xml:"Ok,omitempty"`
  // example:
  // 
  // 2281C6C9-CBAB-1AFD-8400-670750CF6025_2212
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnablePolarClawCronJobResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnablePolarClawCronJobResponseBody) GoString() string {
  return s.String()
}

func (s *EnablePolarClawCronJobResponseBody) GetApplicationId() *string  {
  return s.ApplicationId
}

func (s *EnablePolarClawCronJobResponseBody) GetCode() *int32  {
  return s.Code
}

func (s *EnablePolarClawCronJobResponseBody) GetJobId() *string  {
  return s.JobId
}

func (s *EnablePolarClawCronJobResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnablePolarClawCronJobResponseBody) GetOk() *bool  {
  return s.Ok
}

func (s *EnablePolarClawCronJobResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnablePolarClawCronJobResponseBody) SetApplicationId(v string) *EnablePolarClawCronJobResponseBody {
  s.ApplicationId = &v
  return s
}

func (s *EnablePolarClawCronJobResponseBody) SetCode(v int32) *EnablePolarClawCronJobResponseBody {
  s.Code = &v
  return s
}

func (s *EnablePolarClawCronJobResponseBody) SetJobId(v string) *EnablePolarClawCronJobResponseBody {
  s.JobId = &v
  return s
}

func (s *EnablePolarClawCronJobResponseBody) SetMessage(v string) *EnablePolarClawCronJobResponseBody {
  s.Message = &v
  return s
}

func (s *EnablePolarClawCronJobResponseBody) SetOk(v bool) *EnablePolarClawCronJobResponseBody {
  s.Ok = &v
  return s
}

func (s *EnablePolarClawCronJobResponseBody) SetRequestId(v string) *EnablePolarClawCronJobResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnablePolarClawCronJobResponseBody) Validate() error {
  return dara.Validate(s)
}

