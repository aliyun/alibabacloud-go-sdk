// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExceedApplySyncRequest interface {
  dara.Model
  String() string
  GoString() string
  SetApplyId(v int64) *ExceedApplySyncRequest
  GetApplyId() *int64 
  SetBizCategory(v int32) *ExceedApplySyncRequest
  GetBizCategory() *int32 
  SetRemark(v string) *ExceedApplySyncRequest
  GetRemark() *string 
  SetStatus(v int32) *ExceedApplySyncRequest
  GetStatus() *int32 
  SetThirdpartyFlowId(v string) *ExceedApplySyncRequest
  GetThirdpartyFlowId() *string 
  SetUserId(v string) *ExceedApplySyncRequest
  GetUserId() *string 
}

type ExceedApplySyncRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // 823744
  ApplyId *int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
  // example:
  // 
  // 3
  BizCategory *int32 `json:"biz_category,omitempty" xml:"biz_category,omitempty"`
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 1
  Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
  // example:
  // 
  // 123421
  ThirdpartyFlowId *string `json:"thirdparty_flow_id,omitempty" xml:"thirdparty_flow_id,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // open5145141
  UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s ExceedApplySyncRequest) String() string {
  return dara.Prettify(s)
}

func (s ExceedApplySyncRequest) GoString() string {
  return s.String()
}

func (s *ExceedApplySyncRequest) GetApplyId() *int64  {
  return s.ApplyId
}

func (s *ExceedApplySyncRequest) GetBizCategory() *int32  {
  return s.BizCategory
}

func (s *ExceedApplySyncRequest) GetRemark() *string  {
  return s.Remark
}

func (s *ExceedApplySyncRequest) GetStatus() *int32  {
  return s.Status
}

func (s *ExceedApplySyncRequest) GetThirdpartyFlowId() *string  {
  return s.ThirdpartyFlowId
}

func (s *ExceedApplySyncRequest) GetUserId() *string  {
  return s.UserId
}

func (s *ExceedApplySyncRequest) SetApplyId(v int64) *ExceedApplySyncRequest {
  s.ApplyId = &v
  return s
}

func (s *ExceedApplySyncRequest) SetBizCategory(v int32) *ExceedApplySyncRequest {
  s.BizCategory = &v
  return s
}

func (s *ExceedApplySyncRequest) SetRemark(v string) *ExceedApplySyncRequest {
  s.Remark = &v
  return s
}

func (s *ExceedApplySyncRequest) SetStatus(v int32) *ExceedApplySyncRequest {
  s.Status = &v
  return s
}

func (s *ExceedApplySyncRequest) SetThirdpartyFlowId(v string) *ExceedApplySyncRequest {
  s.ThirdpartyFlowId = &v
  return s
}

func (s *ExceedApplySyncRequest) SetUserId(v string) *ExceedApplySyncRequest {
  s.UserId = &v
  return s
}

func (s *ExceedApplySyncRequest) Validate() error {
  return dara.Validate(s)
}

