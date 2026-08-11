// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iE2BTeam interface {
  dara.Model
  String() string
  GoString() string
  SetAllowUpdateTeamName(v bool) *E2BTeam
  GetAllowUpdateTeamName() *bool 
  SetCreatedTime(v string) *E2BTeam
  GetCreatedTime() *string 
  SetDescription(v string) *E2BTeam
  GetDescription() *string 
  SetPlan(v string) *E2BTeam
  GetPlan() *string 
  SetResourceGroupID(v string) *E2BTeam
  GetResourceGroupID() *string 
  SetStatus(v string) *E2BTeam
  GetStatus() *string 
  SetTeamID(v string) *E2BTeam
  GetTeamID() *string 
  SetTeamName(v string) *E2BTeam
  GetTeamName() *string 
  SetUserID(v string) *E2BTeam
  GetUserID() *string 
}

type E2BTeam struct {
  // Indicates whether the team name can be modified.
  // 
  // example:
  // 
  // false
  AllowUpdateTeamName *bool `json:"allowUpdateTeamName,omitempty" xml:"allowUpdateTeamName,omitempty"`
  // The time when the team was created.
  // 
  // example:
  // 
  // 2025-05-23T11:26:25Z
  CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
  // The description.
  // 
  // example:
  // 
  // Open team
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  Plan *string `json:"plan,omitempty" xml:"plan,omitempty"`
  // The resource group ID.
  // 
  // example:
  // 
  // rg-acfmwxqyrgwabcd
  ResourceGroupID *string `json:"resourceGroupID,omitempty" xml:"resourceGroupID,omitempty"`
  // The status of the team.
  // 
  // example:
  // 
  // active
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  // The unique identifier of the team.
  // 
  // example:
  // 
  // 70d1c834-0383-58d8-97ac-5336eb91abcd
  TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
  // The name of the team.
  // 
  // example:
  // 
  // dev
  TeamName *string `json:"teamName,omitempty" xml:"teamName,omitempty"`
  // The UID of the creator.
  // 
  // example:
  // 
  // 12345
  UserID *string `json:"userID,omitempty" xml:"userID,omitempty"`
}

func (s E2BTeam) String() string {
  return dara.Prettify(s)
}

func (s E2BTeam) GoString() string {
  return s.String()
}

func (s *E2BTeam) GetAllowUpdateTeamName() *bool  {
  return s.AllowUpdateTeamName
}

func (s *E2BTeam) GetCreatedTime() *string  {
  return s.CreatedTime
}

func (s *E2BTeam) GetDescription() *string  {
  return s.Description
}

func (s *E2BTeam) GetPlan() *string  {
  return s.Plan
}

func (s *E2BTeam) GetResourceGroupID() *string  {
  return s.ResourceGroupID
}

func (s *E2BTeam) GetStatus() *string  {
  return s.Status
}

func (s *E2BTeam) GetTeamID() *string  {
  return s.TeamID
}

func (s *E2BTeam) GetTeamName() *string  {
  return s.TeamName
}

func (s *E2BTeam) GetUserID() *string  {
  return s.UserID
}

func (s *E2BTeam) SetAllowUpdateTeamName(v bool) *E2BTeam {
  s.AllowUpdateTeamName = &v
  return s
}

func (s *E2BTeam) SetCreatedTime(v string) *E2BTeam {
  s.CreatedTime = &v
  return s
}

func (s *E2BTeam) SetDescription(v string) *E2BTeam {
  s.Description = &v
  return s
}

func (s *E2BTeam) SetPlan(v string) *E2BTeam {
  s.Plan = &v
  return s
}

func (s *E2BTeam) SetResourceGroupID(v string) *E2BTeam {
  s.ResourceGroupID = &v
  return s
}

func (s *E2BTeam) SetStatus(v string) *E2BTeam {
  s.Status = &v
  return s
}

func (s *E2BTeam) SetTeamID(v string) *E2BTeam {
  s.TeamID = &v
  return s
}

func (s *E2BTeam) SetTeamName(v string) *E2BTeam {
  s.TeamName = &v
  return s
}

func (s *E2BTeam) SetUserID(v string) *E2BTeam {
  s.UserID = &v
  return s
}

func (s *E2BTeam) Validate() error {
  return dara.Validate(s)
}

