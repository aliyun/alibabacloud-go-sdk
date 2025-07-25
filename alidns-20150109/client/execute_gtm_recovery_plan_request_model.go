// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteGtmRecoveryPlanRequest interface {
  dara.Model
  String() string
  GoString() string
  SetLang(v string) *ExecuteGtmRecoveryPlanRequest
  GetLang() *string 
  SetRecoveryPlanId(v int64) *ExecuteGtmRecoveryPlanRequest
  GetRecoveryPlanId() *int64 
}

type ExecuteGtmRecoveryPlanRequest struct {
  // The language of the response. Valid values:
  // 
  // 	- zh: Chinese
  // 
  // 	- en: English
  // 
  // Default value: English.
  // 
  // example:
  // 
  // en
  Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
  // The ID of the disaster recovery plan.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 100
  RecoveryPlanId *int64 `json:"RecoveryPlanId,omitempty" xml:"RecoveryPlanId,omitempty"`
}

func (s ExecuteGtmRecoveryPlanRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteGtmRecoveryPlanRequest) GoString() string {
  return s.String()
}

func (s *ExecuteGtmRecoveryPlanRequest) GetLang() *string  {
  return s.Lang
}

func (s *ExecuteGtmRecoveryPlanRequest) GetRecoveryPlanId() *int64  {
  return s.RecoveryPlanId
}

func (s *ExecuteGtmRecoveryPlanRequest) SetLang(v string) *ExecuteGtmRecoveryPlanRequest {
  s.Lang = &v
  return s
}

func (s *ExecuteGtmRecoveryPlanRequest) SetRecoveryPlanId(v int64) *ExecuteGtmRecoveryPlanRequest {
  s.RecoveryPlanId = &v
  return s
}

func (s *ExecuteGtmRecoveryPlanRequest) Validate() error {
  return dara.Validate(s)
}

