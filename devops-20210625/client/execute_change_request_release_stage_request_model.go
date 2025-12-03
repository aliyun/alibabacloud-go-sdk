// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteChangeRequestReleaseStageRequest interface {
  dara.Model
  String() string
  GoString() string
  SetParams(v map[string]interface{}) *ExecuteChangeRequestReleaseStageRequest
  GetParams() map[string]interface{} 
  SetOrganizationId(v string) *ExecuteChangeRequestReleaseStageRequest
  GetOrganizationId() *string 
}

type ExecuteChangeRequestReleaseStageRequest struct {
  Params map[string]interface{} `json:"params,omitempty" xml:"params,omitempty"`
  // example:
  // 
  // 66c0c9fffeb86b450c199fcd
  OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s ExecuteChangeRequestReleaseStageRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteChangeRequestReleaseStageRequest) GoString() string {
  return s.String()
}

func (s *ExecuteChangeRequestReleaseStageRequest) GetParams() map[string]interface{}  {
  return s.Params
}

func (s *ExecuteChangeRequestReleaseStageRequest) GetOrganizationId() *string  {
  return s.OrganizationId
}

func (s *ExecuteChangeRequestReleaseStageRequest) SetParams(v map[string]interface{}) *ExecuteChangeRequestReleaseStageRequest {
  s.Params = v
  return s
}

func (s *ExecuteChangeRequestReleaseStageRequest) SetOrganizationId(v string) *ExecuteChangeRequestReleaseStageRequest {
  s.OrganizationId = &v
  return s
}

func (s *ExecuteChangeRequestReleaseStageRequest) Validate() error {
  return dara.Validate(s)
}

