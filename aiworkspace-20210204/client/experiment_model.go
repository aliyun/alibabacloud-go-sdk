// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExperiment interface {
  dara.Model
  String() string
  GoString() string
  SetAccessibility(v string) *Experiment
  GetAccessibility() *string 
  SetArtifactUri(v string) *Experiment
  GetArtifactUri() *string 
  SetExperimentId(v string) *Experiment
  GetExperimentId() *string 
  SetGmtCreateTime(v string) *Experiment
  GetGmtCreateTime() *string 
  SetGmtModifiedTime(v string) *Experiment
  GetGmtModifiedTime() *string 
  SetLabels(v []*ExperimentLabel) *Experiment
  GetLabels() []*ExperimentLabel 
  SetLatestRun(v *Run) *Experiment
  GetLatestRun() *Run 
  SetName(v string) *Experiment
  GetName() *string 
  SetOwnerId(v string) *Experiment
  GetOwnerId() *string 
  SetRequestId(v string) *Experiment
  GetRequestId() *string 
  SetTensorboardLogUri(v string) *Experiment
  GetTensorboardLogUri() *string 
  SetUserId(v string) *Experiment
  GetUserId() *string 
  SetWorkspaceId(v string) *Experiment
  GetWorkspaceId() *string 
}

type Experiment struct {
  Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
  ArtifactUri *string `json:"ArtifactUri,omitempty" xml:"ArtifactUri,omitempty"`
  ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
  GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
  GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
  Labels []*ExperimentLabel `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
  LatestRun *Run `json:"LatestRun,omitempty" xml:"LatestRun,omitempty"`
  Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
  OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  TensorboardLogUri *string `json:"TensorboardLogUri,omitempty" xml:"TensorboardLogUri,omitempty"`
  UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
  WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s Experiment) String() string {
  return dara.Prettify(s)
}

func (s Experiment) GoString() string {
  return s.String()
}

func (s *Experiment) GetAccessibility() *string  {
  return s.Accessibility
}

func (s *Experiment) GetArtifactUri() *string  {
  return s.ArtifactUri
}

func (s *Experiment) GetExperimentId() *string  {
  return s.ExperimentId
}

func (s *Experiment) GetGmtCreateTime() *string  {
  return s.GmtCreateTime
}

func (s *Experiment) GetGmtModifiedTime() *string  {
  return s.GmtModifiedTime
}

func (s *Experiment) GetLabels() []*ExperimentLabel  {
  return s.Labels
}

func (s *Experiment) GetLatestRun() *Run  {
  return s.LatestRun
}

func (s *Experiment) GetName() *string  {
  return s.Name
}

func (s *Experiment) GetOwnerId() *string  {
  return s.OwnerId
}

func (s *Experiment) GetRequestId() *string  {
  return s.RequestId
}

func (s *Experiment) GetTensorboardLogUri() *string  {
  return s.TensorboardLogUri
}

func (s *Experiment) GetUserId() *string  {
  return s.UserId
}

func (s *Experiment) GetWorkspaceId() *string  {
  return s.WorkspaceId
}

func (s *Experiment) SetAccessibility(v string) *Experiment {
  s.Accessibility = &v
  return s
}

func (s *Experiment) SetArtifactUri(v string) *Experiment {
  s.ArtifactUri = &v
  return s
}

func (s *Experiment) SetExperimentId(v string) *Experiment {
  s.ExperimentId = &v
  return s
}

func (s *Experiment) SetGmtCreateTime(v string) *Experiment {
  s.GmtCreateTime = &v
  return s
}

func (s *Experiment) SetGmtModifiedTime(v string) *Experiment {
  s.GmtModifiedTime = &v
  return s
}

func (s *Experiment) SetLabels(v []*ExperimentLabel) *Experiment {
  s.Labels = v
  return s
}

func (s *Experiment) SetLatestRun(v *Run) *Experiment {
  s.LatestRun = v
  return s
}

func (s *Experiment) SetName(v string) *Experiment {
  s.Name = &v
  return s
}

func (s *Experiment) SetOwnerId(v string) *Experiment {
  s.OwnerId = &v
  return s
}

func (s *Experiment) SetRequestId(v string) *Experiment {
  s.RequestId = &v
  return s
}

func (s *Experiment) SetTensorboardLogUri(v string) *Experiment {
  s.TensorboardLogUri = &v
  return s
}

func (s *Experiment) SetUserId(v string) *Experiment {
  s.UserId = &v
  return s
}

func (s *Experiment) SetWorkspaceId(v string) *Experiment {
  s.WorkspaceId = &v
  return s
}

func (s *Experiment) Validate() error {
  if s.Labels != nil {
    for _, item := range s.Labels {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  if s.LatestRun != nil {
    if err := s.LatestRun.Validate(); err != nil {
      return err
    }
  }
  return nil
}

